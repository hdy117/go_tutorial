package main

import (
	"errors"
)

// 函数定义
// 参数类型和返回值类型
func Add(a int, b int) int {
	return a + b
}

// 函数返回多个值
func Divide(a float32, b float32) (float32, error) {
	if b == 0 {
		return 0.0, errors.New("second value should not be 0")
	}
	return a / b, nil
}
