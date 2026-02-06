// ============================================
// Go 错误处理教程
// ============================================
//
// 本文件涵盖 Go 语言错误处理的核心：
// - error 接口
// - 创建错误
// - 自定义错误类型
// - 错误链（Go 1.13+）⭐
// - errors.Is 和 errors.As
// - panic / recover
// - 错误处理最佳实践
//
// 最佳实践：
// 1. 错误处理是值，不是异常，显式检查
// 2. 错误信息应该是小写，不包含结尾标点
// 3. 使用 fmt.Errorf("...: %w", err) 包装错误
// 4. 错误类型应该支持 errors.Is 和 errors.As
// 5. 只在真正不可恢复的情况下使用 panic
// 6. recover 只在 defer 中使用
// 7. 不要忽略错误（不要用 _ 接收，除非确实不需要）
// ============================================

package main

import (
	"errors"
	"fmt"
	"os"
)

// ============================================
// 1. error 接口
// ============================================
//
// error 是内置接口：
// type error interface {
//     Error() string
// }

func demonstrateBasicError() {
	fmt.Println("=== 基础错误处理 ===")
	
	// 打开文件（可能返回错误）
	file, err := os.Open("non_existent_file.txt")
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
	} else {
		defer file.Close()
	}
	
	// 错误处理模式：先检查错误
	result, err := divide(10, 0)
	if err != nil {
		fmt.Printf("计算错误: %v\n", err)
	} else {
		fmt.Printf("结果: %d\n", result)
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

// ============================================
// 2. 创建错误
// ============================================

func demonstrateCreatingErrors() {
	fmt.Println("\n=== 创建错误 ===")
	
	// 方式1：errors.New（静态错误）
	err1 := errors.New("发生错误")
	fmt.Printf("errors.New: %v\n", err1)
	
	// 方式2：fmt.Errorf（格式化错误）
	code := 404
	err2 := fmt.Errorf("HTTP %d: 页面未找到", code)
	fmt.Printf("fmt.Errorf: %v\n", err2)
	
	// 方式3：fmt.Errorf + %w 包装错误（Go 1.13+）⭐
	innerErr := errors.New("数据库连接失败")
	outerErr := fmt.Errorf("查询用户失败: %w", innerErr)
	fmt.Printf("包装错误: %v\n", outerErr)
}

// ============================================
// 3. 自定义错误类型
// ============================================
//
// 通过实现 error 接口创建自定义错误

// ValidationError 验证错误
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("验证错误 [%s]: %s", e.Field, e.Message)
}

// NotFoundError 资源未找到
type NotFoundError struct {
	Resource string
	ID       int
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s (ID=%d) 未找到", e.Resource, e.ID)
}

// TimeoutError 超时错误
type TimeoutError struct {
	Operation string
	Timeout   int // 毫秒
}

func (e TimeoutError) Error() string {
	return fmt.Sprintf("操作 %s 超时 (%d ms)", e.Operation, e.Timeout)
}

// 实现临时错误接口（可选）
func (e TimeoutError) Temporary() bool {
	return true
}

type temporary interface {
	Temporary() bool
}

func demonstrateCustomErrors() {
	fmt.Println("\n=== 自定义错误类型 ===")
	
	// 创建自定义错误
	valErr := ValidationError{Field: "email", Message: "格式不正确"}
	notFoundErr := NotFoundError{Resource: "User", ID: 123}
	timeoutErr := TimeoutError{Operation: "数据库查询", Timeout: 5000}
	
	fmt.Println(valErr)
	fmt.Println(notFoundErr)
	fmt.Println(timeoutErr)
	
	// 检查是否是临时错误
	if t, ok := interface{}(timeoutErr).(temporary); ok && t.Temporary() {
		fmt.Println("这是一个临时错误，可以重试")
	}
}

// ============================================
// 4. 错误链（Go 1.13+）⭐
// ============================================
//
// 使用 %w 包装错误，形成错误链
// errors.Unwrap: 获取被包装的错误
// errors.Is: 检查错误链中是否包含特定错误
// errors.As: 将错误转换为特定类型

var (
	ErrNotFound  = errors.New("资源未找到")
	ErrInvalid   = errors.New("无效的输入")
	ErrDatabase  = errors.New("数据库错误")
	ErrNetwork   = errors.New("网络错误")
)

func queryUser(id int) error {
	// 模拟底层错误
	return fmt.Errorf("%w: 连接超时", ErrDatabase)
}

func getUser(id int) error {
	err := queryUser(id)
	if err != nil {
		return fmt.Errorf("获取用户 %d 失败: %w", id, err)
	}
	return nil
}

func demonstrateErrorChain() {
	fmt.Println("\n=== 错误链 ===")
	
	err := getUser(42)
	fmt.Printf("错误: %v\n", err)
	
	// 展开错误链
	fmt.Println("\n展开错误链:")
	for err != nil {
		fmt.Printf("  -> %v\n", err)
		err = errors.Unwrap(err)
	}
}

// errors.Is 示例
func demonstrateErrorsIs() {
	fmt.Println("\n=== errors.Is ===")
	
	err := getUser(42)
	
	// 检查错误链中是否包含特定错误
	if errors.Is(err, ErrDatabase) {
		fmt.Println("是数据库错误")
	}
	
	if errors.Is(err, ErrNotFound) {
		fmt.Println("是未找到错误")
	} else {
		fmt.Println("不是未找到错误")
	}
}

// errors.As 示例
func demonstrateErrorsAs() {
	fmt.Println("\n=== errors.As ===")
	
	// 模拟一个包含 ValidationError 的错误链
	err := fmt.Errorf("处理请求失败: %w", ValidationError{
		Field:   "age",
		Message: "必须大于 0",
	})
	
	// 尝试转换为 ValidationError
	var valErr ValidationError
	if errors.As(err, &valErr) {
		fmt.Printf("验证错误: 字段=%s, 消息=%s\n", valErr.Field, valErr.Message)
	}
	
	// 尝试转换为 NotFoundError（会失败）
	var notFoundErr NotFoundError
	if errors.As(err, &notFoundErr) {
		fmt.Println("找到 NotFoundError")
	} else {
		fmt.Println("不是 NotFoundError")
	}
}

// ============================================
// 5. panic 和 recover
// ============================================
//
// panic：停止当前函数执行，向上冒泡
// recover：捕获 panic，恢复正常执行
// 只在 defer 中有效

func demonstratePanic() {
	fmt.Println("\n=== Panic ===")
	
	// 不要这样做！
	// panic("出错了！")
	
	// 触发 panic 的常见情况
	// var p *int
	// fmt.Println(*p)  // nil 指针解引用，panic
	
	// slice := []int{1, 2, 3}
	// fmt.Println(slice[10])  // 索引越界，panic
	
	// var m map[string]int
	// m["key"] = 1  // 未初始化的 map，panic
}

// 安全地执行可能 panic 的函数
func safeExecute(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// 捕获 panic
			err = fmt.Errorf("panic 捕获: %v", r)
		}
	}()
	
	fn()
	return nil
}

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("发生严重错误！")
	}
	fmt.Println("正常执行")
}

func demonstrateRecover() {
	fmt.Println("\n=== Recover ===")
	
	// 捕获 panic
	err := safeExecute(func() {
		mightPanic(true)
	})
	if err != nil {
		fmt.Printf("捕获到错误: %v\n", err)
	}
	
	// 正常执行
	err = safeExecute(func() {
		mightPanic(false)
	})
	if err != nil {
		fmt.Printf("捕获到错误: %v\n", err)
	}
}

// ============================================
// 6. 错误处理模式
// ============================================

// 6.1 哨兵错误（Sentinel Errors）
func findUser(id int) (string, error) {
	if id <= 0 {
		return "", fmt.Errorf("%w: id=%d", ErrInvalid, id)
	}
	if id == 999 {
		return "", fmt.Errorf("%w: id=%d", ErrNotFound, id)
	}
	return fmt.Sprintf("User%d", id), nil
}

// 6.2 错误包装策略
func serviceLayer() error {
	err := databaseLayer()
	if err != nil {
		// 添加上下文，保留原始错误
		return fmt.Errorf("服务层处理失败: %w", err)
	}
	return nil
}

func databaseLayer() error {
	// 模拟数据库错误
	return fmt.Errorf("%w: 连接超时", ErrDatabase)
}

// 6.3 错误处理辅助函数
func handleError(err error) {
	if err == nil {
		return
	}
	
	// 根据错误类型处理
	switch {
	case errors.Is(err, ErrNotFound):
		fmt.Println("[404] 资源不存在")
	case errors.Is(err, ErrInvalid):
		fmt.Println("[400] 请求参数错误")
	case errors.Is(err, ErrDatabase):
		fmt.Println("[500] 数据库错误")
	default:
		fmt.Printf("[500] 未知错误: %v\n", err)
	}
}

func demonstrateErrorPatterns() {
	fmt.Println("\n=== 错误处理模式 ===")
	
	// 测试不同场景
	testCases := []int{-1, 0, 1, 999}
	
	for _, id := range testCases {
		_, err := findUser(id)
		fmt.Printf("查找用户 %d: ", id)
		handleError(err)
	}
	
	// 服务层错误
	fmt.Println("\n服务层错误:")
	if err := serviceLayer(); err != nil {
		fmt.Printf("完整错误链: %v\n", err)
		
		// 检查特定错误
		if errors.Is(err, ErrDatabase) {
			fmt.Println("需要检查数据库连接")
		}
	}
}

// ============================================
// 7. 实用的错误处理工具
// ============================================

// 多重错误（简单实现）
type MultiError struct {
	errors []error
}

func (m *MultiError) Error() string {
	if len(m.errors) == 0 {
		return "no errors"
	}
	if len(m.errors) == 1 {
		return m.errors[0].Error()
	}
	
	msg := fmt.Sprintf("%d errors occurred:", len(m.errors))
	for i, err := range m.errors {
		msg += fmt.Sprintf("\n  [%d] %s", i+1, err.Error())
	}
	return msg
}

func (m *MultiError) Add(err error) {
	if err != nil {
		m.errors = append(m.errors, err)
	}
}

func (m *MultiError) HasErrors() bool {
	return len(m.errors) > 0
}

// 重试函数
func withRetry(maxRetries int, fn func() error) error {
	var lastErr error
	
	for i := 0; i < maxRetries; i++ {
		if err := fn(); err != nil {
			lastErr = err
			// 检查是否是临时错误
			if t, ok := err.(temporary); ok && t.Temporary() {
				continue
			}
			return err
		}
		return nil
	}
	
	return fmt.Errorf("重试 %d 次后失败: %w", maxRetries, lastErr)
}

func demonstrateUtilities() {
	fmt.Println("\n=== 错误处理工具 ===")
	
	// 多重错误
	multi := &MultiError{}
	multi.Add(errors.New("错误1"))
	multi.Add(errors.New("错误2"))
	multi.Add(nil)  // 会被忽略
	
	if multi.HasErrors() {
		fmt.Println(multi.Error())
	}
	
	// 重试示例
	attempts := 0
	err := withRetry(3, func() error {
		attempts++
		if attempts < 3 {
			return TimeoutError{Operation: "测试", Timeout: 100}
		}
		return nil
	})
	
	if err != nil {
		fmt.Printf("重试结果: %v\n", err)
	} else {
		fmt.Printf("成功，尝试了 %d 次\n", attempts)
	}
}

// ============================================
// 8. HTTP 错误处理示例
// ============================================

type HTTPError struct {
	StatusCode int
	Message    string
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
}

// 模拟 HTTP 处理
func handleRequest(path string) error {
	switch path {
	case "/users":
		return nil  // 成功
	case "/admin":
		return HTTPError{StatusCode: 403, Message: "Forbidden"}
	case "/unknown":
		return HTTPError{StatusCode: 404, Message: "Not Found"}
	default:
		return HTTPError{StatusCode: 500, Message: "Internal Server Error"}
	}
}

func demonstrateHTTPError() {
	fmt.Println("\n=== HTTP 错误处理 ===")
	
	paths := []string{"/users", "/admin", "/unknown", "/error"}
	
	for _, path := range paths {
		if err := handleRequest(path); err != nil {
			var httpErr HTTPError
			if errors.As(err, &httpErr) {
				fmt.Printf("%s -> %d: %s\n", path, httpErr.StatusCode, httpErr.Message)
			}
		} else {
			fmt.Printf("%s -> 200 OK\n", path)
		}
	}
}

// ============================================
// 主函数
// ============================================

func main() {
	demonstrateBasicError()
	demonstrateCreatingErrors()
	demonstrateCustomErrors()
	demonstrateErrorChain()
	demonstrateErrorsIs()
	demonstrateErrorsAs()
	demonstratePanic()
	demonstrateRecover()
	demonstrateErrorPatterns()
	demonstrateUtilities()
	demonstrateHTTPError()
	
	// ============================================
	// 练习题
	// ============================================
	//
	// 练习 1：实现一个堆栈跟踪的错误类型
	//   type StackError struct { error; stack []byte }
	//   - 创建错误时捕获堆栈
	//   - 实现 Error() 方法，输出错误信息和堆栈
	//
	// 练习 2：实现一个错误码系统
	//   - 定义错误码常量（如 ErrCodeNotFound = 404）
	//   - 实现 CodedError 结构体，包含 Code 和 Message
	//   - 实现 FromCode(code int) 根据 HTTP 状态码创建错误
	//   - 实现 HTTPStatus() 返回对应的 HTTP 状态码
	//
	// 练习 3：实现一个批处理错误收集器
	//   type BatchProcessor struct { ... }
	//   - 处理多个项目，收集所有错误
	//   - 如果所有错误都是同一种类型，返回该类型错误
	//   - 如果有多种错误，返回 MultiError
	//
	// 练习 4：实现一个带上下文的错误
	//   type ContextError struct { error; Context map[string]interface{} }
	//   - 支持添加键值对上下文
	//   - Error() 输出时包含上下文信息
	//   - 实现 Unwrap() 支持错误链
	//
	// 练习 5：实现一个断言工具
	//   func AssertNotNil(v interface{}, msg string)
	//   func AssertTrue(condition bool, msg string)
	//   func AssertNoError(err error)
	//   - 断言失败时 panic
	//   - 在测试中使用
	//
	// 练习 6：实现一个错误重试装饰器
	//   func Retryable(fn func() error, opts RetryOptions) func() error
	//   - 支持自定义重试次数、退避策略
	//   - 支持只对特定错误重试
	//   - 支持超时
}
