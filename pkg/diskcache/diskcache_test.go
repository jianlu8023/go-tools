package diskcache

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestDiskCacheRoundTrip(t *testing.T) {
	dir := t.TempDir()
	SetConfig(Config{Enabled: true, ThresholdBytes: 1, MaxSizeBytes: 1024, Path: dir, DirName: "test-cache"})
	defer SetConfig(Config{}) // 还原全局配置

	path, err := WriteFile(TypeBody, []byte("hello-cache"))
	if err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	defer RemoveFile(path)

	if filepath.Dir(path) != filepath.Join(dir, "test-cache") {
		t.Fatalf("unexpected cache dir: %s", path)
	}
	data, err := ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if string(data) != "hello-cache" {
		t.Fatalf("unexpected content: %s", data)
	}
	s, err := ReadFileString(path)
	if err != nil || s != "hello-cache" {
		t.Fatalf("ReadFileString: %q, %v", s, err)
	}

	count, size, err := Info()
	if err != nil {
		t.Fatalf("Info: %v", err)
	}
	if count != 1 || size != int64(len("hello-cache")) {
		t.Fatalf("Info = %d files, %d bytes", count, size)
	}
}

func TestDirDefaults(t *testing.T) {
	SetConfig(Config{Enabled: true})
	defer SetConfig(Config{})
	if Dir() != filepath.Join(os.TempDir(), DefaultDirName) {
		t.Fatalf("unexpected default dir: %s", Dir())
	}
}

func TestCleanupOldFiles(t *testing.T) {
	dir := t.TempDir()
	SetConfig(Config{Enabled: true, Path: dir, DirName: "c"})
	defer SetConfig(Config{})

	fresh, err := WriteFile(TypeFile, []byte("x"))
	if err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	defer RemoveFile(fresh)

	// 构造一个"旧"文件
	oldPath := filepath.Join(Dir(), "old.tmp")
	if err := os.WriteFile(oldPath, []byte("y"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.Chtimes(oldPath, time.Now().Add(-24*time.Hour), time.Now().Add(-24*time.Hour)); err != nil {
		t.Fatal(err)
	}

	if err := CleanupOldFiles(time.Hour); err != nil {
		t.Fatalf("CleanupOldFiles: %v", err)
	}
	if _, err := ReadFile(oldPath); !os.IsNotExist(err) {
		t.Fatalf("old file should be removed, err=%v", err)
	}
	if _, err := ReadFile(fresh); err != nil {
		t.Fatalf("fresh file should remain: %v", err)
	}
}

func TestShouldUseCache(t *testing.T) {
	dir := t.TempDir()
	SetConfig(Config{Enabled: false, ThresholdBytes: 100, MaxSizeBytes: 1024, Path: dir})
	defer SetConfig(Config{})
	if ShouldUseCache(1000) {
		t.Fatal("disabled config should not use cache")
	}
	SetConfig(Config{Enabled: true, ThresholdBytes: 100, MaxSizeBytes: 1024, Path: dir})
	if ShouldUseCache(50) {
		t.Fatal("below threshold should not use cache")
	}
	if !ShouldUseCache(200) {
		t.Fatal("above threshold within max size should use cache")
	}
	if ShouldUseCache(2048) {
		t.Fatal("exceeding max size should not use cache")
	}
}
