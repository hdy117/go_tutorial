package main

import (
	"fmt"
	"testing"
)

// 最佳实践 1: 使用表驱动测试 (Table-Driven Tests)
// 这是 Go 社区最推荐的测试方式，可以轻松添加更多测试用例
func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		a        float32
		b        float32
		want     float32
		wantErr  bool
		checkErr func(error) bool // 可选：检查特定的错误信息
	}{
		{
			name:    "正常除法",
			a:       10.0,
			b:       2.0,
			want:    5.0,
			wantErr: false,
		},
		{
			name:    "除零错误",
			a:       10.0,
			b:       0.0,
			want:    0.0,
			wantErr: true,
			checkErr: func(err error) bool {
				return err != nil && err.Error() == "second value should not be 0"
			},
		},
		{
			name:    "小数除法",
			a:       7.5,
			b:       2.5,
			want:    3.0,
			wantErr: false,
		},
		{
			name:    "负数除法",
			a:       -10.0,
			b:       2.0,
			want:    -5.0,
			wantErr: false,
		},
		{
			name:    "被除数为0",
			a:       0.0,
			b:       5.0,
			want:    0.0,
			wantErr: false,
		},
	}

	// 最佳实践 2: 使用子测试 (Subtests)
	// 可以单独运行某个测试用例，输出更清晰
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)

			// 检查错误情况
			if (err != nil) != tt.wantErr {
				t.Errorf("Divide(%v, %v) 错误 = %v, wantErr %v", tt.a, tt.b, err, tt.wantErr)
				return
			}

			// 如果有自定义错误检查函数，使用它
			if tt.wantErr && tt.checkErr != nil {
				if !tt.checkErr(err) {
					t.Errorf("Divide(%v, %v) 错误信息不符合预期: %v", tt.a, tt.b, err)
				}
			}

			// 检查返回值（只在没有错误时检查）
			if !tt.wantErr && got != tt.want {
				t.Errorf("Divide(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// 最佳实践 3: 为每个函数编写测试
func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"正数相加", 2, 3, 5},
		{"负数相加", -2, -3, -5},
		{"正负相加", 5, -3, 2},
		{"零值相加", 0, 5, 5},
		{"两个零", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// 最佳实践 4: 使用辅助函数减少重复代码
func assertNoError(t *testing.T, err error) {
	t.Helper() // 标记为辅助函数，错误信息会指向调用者
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

// 最佳实践 5: 基准测试 (Benchmark Tests)
func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Divide(10.0, 2.0)
	}
}

// 最佳实践 6: 示例函数 (Example Functions)
// 这些会出现在文档中，并且可以验证输出
func ExampleDivide() {
	result, err := Divide(10.0, 2.0)
	if err != nil {
		return
	}
	fmt.Println(result)
	// Output: 5
}
