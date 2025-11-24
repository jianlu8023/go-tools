package jq

import (
	"github.com/elgs/gojq"
)

// JQ JSON查询结构体，封装gojq库
type JQ struct {
	*gojq.JQ
}

// NewQuery 从已解析的JSON数据创建JQ实例
func NewQuery(jsonObject interface{}) *JQ {
	return &JQ{JQ: gojq.NewQuery(jsonObject)}
}

// NewStringQuery 从JSON字符串创建JQ实例
func NewStringQuery(jsonString string) (*JQ, error) {
	jq, err := gojq.NewStringQuery(jsonString)
	if err != nil {
		return nil, err
	}
	return &JQ{JQ: jq}, nil
}

// NewFileQuery 从JSON文件创建JQ实例
func NewFileQuery(jsonFile string) (*JQ, error) {
	jq, err := gojq.NewFileQuery(jsonFile)
	if err != nil {
		return nil, err
	}
	return &JQ{JQ: jq}, nil
}

// Query 根据表达式查询JSON数据
func (jq *JQ) Query(exp string) (interface{}, error) {
	return jq.JQ.Query(exp)
}

// QueryToString 查询并转换为字符串
func (jq *JQ) QueryToString(exp string) (string, error) {
	return jq.JQ.QueryToString(exp)
}

// QueryToInt64 查询并转换为int64
func (jq *JQ) QueryToInt64(exp string) (int64, error) {
	return jq.JQ.QueryToInt64(exp)
}

// QueryToFloat64 查询并转换为float64
func (jq *JQ) QueryToFloat64(exp string) (float64, error) {
	return jq.JQ.QueryToFloat64(exp)
}

// QueryToBool 查询并转换为bool
func (jq *JQ) QueryToBool(exp string) (bool, error) {
	return jq.JQ.QueryToBool(exp)
}

// QueryToArray 查询并转换为数组
func (jq *JQ) QueryToArray(exp string) ([]interface{}, error) {
	return jq.JQ.QueryToArray(exp)
}

// QueryToMap 查询并转换为map
func (jq *JQ) QueryToMap(exp string) (map[string]interface{}, error) {
	return jq.JQ.QueryToMap(exp)
}

// 默认客户端函数

func Query(data interface{}, exp string) (interface{}, error) {
	jq := NewQuery(data)
	return jq.Query(exp)
}

func QueryString(data interface{}, exp string) (string, error) {
	jq := NewQuery(data)
	return jq.QueryToString(exp)
}

func QueryInt64(data interface{}, exp string) (int64, error) {
	jq := NewQuery(data)
	return jq.QueryToInt64(exp)
}

func QueryFloat64(data interface{}, exp string) (float64, error) {
	jq := NewQuery(data)
	return jq.QueryToFloat64(exp)
}

func QueryBool(data interface{}, exp string) (bool, error) {
	jq := NewQuery(data)
	return jq.QueryToBool(exp)
}

func QueryArray(data interface{}, exp string) ([]interface{}, error) {
	jq := NewQuery(data)
	return jq.QueryToArray(exp)
}

func QueryMap(data interface{}, exp string) (map[string]interface{}, error) {
	jq := NewQuery(data)
	return jq.QueryToMap(exp)
}
