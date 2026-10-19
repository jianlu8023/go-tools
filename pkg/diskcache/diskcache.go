package diskcache

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Type 磁盘缓存类型
type Type string

const (
	TypeBody Type = "body" // 请求体缓存
	TypeFile Type = "file" // 文件数据缓存
)

// DefaultDirName 默认缓存目录名
const DefaultDirName = "go-tools-cache"

// Config 磁盘缓存配置
type Config struct {
	Enabled        bool   // 是否启用
	ThresholdBytes int64  // 触发磁盘缓存的数据大小阈值（字节）
	MaxSizeBytes   int64  // 缓存总大小上限（字节）
	Path           string // 缓存目录，空则用 os.TempDir()
	DirName        string // 目录名，空则用 DefaultDirName
}

var (
	mu     sync.RWMutex
	config = Config{Enabled: false, ThresholdBytes: 10 * 1024 * 1024, MaxSizeBytes: 1024 * 1024 * 1024}
)

// SetConfig 设置全局缓存配置
func SetConfig(c Config) {
	mu.Lock()
	defer mu.Unlock()
	config = c
}

// GetConfig 获取全局缓存配置
func GetConfig() Config {
	mu.RLock()
	defer mu.RUnlock()
	return config
}

// Dir 获取缓存目录（响应配置变化，每次调用重新计算）
func Dir() string {
	c := GetConfig()
	cachePath := c.Path
	if cachePath == "" {
		cachePath = os.TempDir()
	}
	dirName := c.DirName
	if dirName == "" {
		dirName = DefaultDirName
	}
	return filepath.Join(cachePath, dirName)
}

// EnsureDir 确保缓存目录存在
func EnsureDir() error {
	return os.MkdirAll(Dir(), 0755)
}

// CreateFile 创建磁盘缓存文件，返回文件路径和句柄（O_EXCL + 0600，文件名含uuid短前缀）
func CreateFile(cacheType Type) (string, *os.File, error) {
	if err := EnsureDir(); err != nil {
		return "", nil, fmt.Errorf("failed to create cache dir: %w", err)
	}
	filename := fmt.Sprintf("%s-%s-%d.tmp", cacheType, uuid.New().String()[:8], time.Now().UnixNano())
	filePath := filepath.Join(Dir(), filename)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_EXCL, 0600)
	if err != nil {
		return "", nil, fmt.Errorf("failed to create cache file: %w", err)
	}
	return filePath, file, nil
}

// WriteFile 写入数据到磁盘缓存文件，返回文件路径
func WriteFile(cacheType Type, data []byte) (string, error) {
	filePath, file, err := CreateFile(cacheType)
	if err != nil {
		return "", err
	}
	if _, err = file.Write(data); err != nil {
		file.Close()
		os.Remove(filePath)
		return "", fmt.Errorf("failed to write cache file: %w", err)
	}
	if err := file.Close(); err != nil {
		os.Remove(filePath)
		return "", fmt.Errorf("failed to close cache file: %w", err)
	}
	return filePath, nil
}

// ReadFile 读取磁盘缓存文件
func ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

// ReadFileString 读取磁盘缓存文件为字符串
func ReadFileString(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// RemoveFile 删除磁盘缓存文件
func RemoveFile(filePath string) error {
	return os.Remove(filePath)
}

// CleanupOldFiles 清理超过 maxAge 的缓存文件
func CleanupOldFiles(maxAge time.Duration) error {
	dir := Dir()
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	now := time.Now()
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			continue
		}
		if now.Sub(info.ModTime()) > maxAge {
			os.Remove(filepath.Join(dir, entry.Name()))
		}
	}
	return nil
}

// Info 获取缓存目录信息（文件数与总大小）
func Info() (fileCount int, totalSize int64, err error) {
	dir := Dir()
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, 0, nil
		}
		return 0, 0, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			continue
		}
		fileCount++
		totalSize += info.Size()
	}
	return fileCount, totalSize, nil
}

// ShouldUseCache 判断是否应该使用磁盘缓存（启用且数据量超过阈值）
func ShouldUseCache(dataSize int64) bool {
	c := GetConfig()
	if !c.Enabled {
		return false
	}
	if dataSize < c.ThresholdBytes {
		return false
	}
	_, totalSize, err := Info()
	return err == nil && totalSize+dataSize <= c.MaxSizeBytes
}
