// ============================================
// Go 接口教程
// ============================================
//
// 本文件涵盖 Go 语言接口的核心特性：
// - 接口定义与实现（隐式实现）⭐
// - 空接口（interface{} / any）
// - 类型断言（Type Assertion）⭐
// - 类型开关（Type Switch）
// - 接口组合（嵌套接口）
// - 接口值与底层结构
// - 常用标准库接口
//
// 最佳实践：
// 1. 接口应该小，通常只有 1-3 个方法（小接口原则）
// 2. 不要提前定义接口，而是先写实现，需要时再抽象
// 3. 使用接口解耦代码，便于测试
// 4. 检查接口是否被正确实现：var _ Interface = (*Type)(nil)
// 5. 避免使用空接口，除非确实需要处理任意类型
// ============================================

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// ============================================
// 1. 接口定义与隐式实现 ⭐
// ============================================
//
// Go 接口是隐式实现的：
// 只要类型实现了接口的所有方法，就自动实现了该接口
// 不需要显式声明（如 Java 的 implements）

// 定义接口
type Writer interface {
	Write(p []byte) (n int, err error)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

// 组合接口
type ReadWriter interface {
	Reader
	Writer
}

// 简单接口示例
type Stringer interface {
	String() string
}

type Speaker interface {
	Speak() string
}

// ============================================
// 2. 类型实现接口
// ============================================

// Dog 类型实现 Speaker 接口
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("%s: 汪汪!", d.Name)
}

// Cat 类型实现 Speaker 接口
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("%s: 喵喵!", c.Name)
}

// 另一个类型
type Robot struct {
	Model string
}

func (r Robot) Speak() string {
	return fmt.Sprintf("%s: 你好，我是机器人", r.Model)
}

// ============================================
// 3. 接口的多态使用
// ============================================

// 接收接口参数，实现多态
func MakeSound(s Speaker) {
	fmt.Println(s.Speak())
}

// 接口切片
func MakeAllSounds(speakers []Speaker) {
	for _, s := range speakers {
		fmt.Println(s.Speak())
	}
}

func demonstratePolymorphism() {
	// 同一接口，不同类型
	dog := Dog{Name: "旺财"}
	cat := Cat{Name: "咪咪"}
	robot := Robot{Model: "R2-D2"}
	
	fmt.Println("=== 多态调用 ===")
	MakeSound(dog)
	MakeSound(cat)
	MakeSound(robot)
	
	// 接口切片
	fmt.Println("\n=== 接口切片 ===")
	animals := []Speaker{dog, cat, robot}
	MakeAllSounds(animals)
}

// ============================================
// 4. 空接口 interface{} / any
// ============================================
//
// 空接口没有任何方法，所有类型都实现了空接口
// Go 1.18+ 引入了 any 作为 interface{} 的别名

func describe(i interface{}) {
	fmt.Printf("值: %v, 类型: %T\n", i, i)
}

func demonstrateEmptyInterface() {
	fmt.Println("\n=== 空接口 ===")
	
	// 空接口可以接受任意类型
	describe(42)
	describe("hello")
	describe(3.14)
	describe([]int{1, 2, 3})
	describe(map[string]int{"a": 1})
	
	// 使用 any（Go 1.18+）
	var x any = "使用 any"
	fmt.Printf("x: %v\n", x)
	
	// 空接口的常见用途：
	// 1. 处理未知类型的数据（JSON 解码）
	// 2. 实现通用的数据结构
	// 3. fmt 包的 Print 系列函数
}

// ============================================
// 5. 类型断言（Type Assertion）⭐
// ============================================
//
// 将接口值转换回具体类型
// x.(T) 断言 x 的类型是 T
// 失败时会 panic，安全写法：v, ok := x.(T)

func demonstrateTypeAssertion() {
	fmt.Println("\n=== 类型断言 ===")
	
	var i interface{} = "hello"
	
	// 安全类型断言
	s, ok := i.(string)
	if ok {
		fmt.Printf("字符串值: %s, 长度: %d\n", s, len(s))
	}
	
	// 断言失败不会 panic
	n, ok := i.(int)
	if !ok {
		fmt.Println("i 不是 int 类型")
	} else {
		fmt.Println("整数值:", n)
	}
	
	// 空接口切片处理
	var data []interface{} = []interface{}{
		"string",
		42,
		3.14,
		true,
		Dog{Name: "Buddy"},
	}
	
	for _, item := range data {
		switch v := item.(type) {
		case string:
			fmt.Printf("字符串: %s\n", v)
		case int:
			fmt.Printf("整数: %d\n", v)
		case float64:
			fmt.Printf("浮点数: %f\n", v)
		case bool:
			fmt.Printf("布尔: %v\n", v)
		case Speaker:
			fmt.Printf("Speaker: %s\n", v.Speak())
		default:
			fmt.Printf("未知类型: %T\n", v)
		}
	}
}

// ============================================
// 6. 类型开关（Type Switch）
// ============================================

func doSomething(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("处理字符串: %s (长度 %d)\n", v, len(v))
	case int:
		fmt.Printf("处理整数: %d (两倍 %d)\n", v, v*2)
	case []int:
		fmt.Printf("处理整数切片，长度: %d\n", len(v))
	case map[string]interface{}:
		fmt.Printf("处理 map，键值对数量: %d\n", len(v))
	case nil:
		fmt.Println("值是 nil")
	default:
		fmt.Printf("未处理的类型: %T\n", v)
	}
}

func demonstrateTypeSwitch() {
	fmt.Println("\n=== 类型开关 ===")
	
	doSomething("Hello")
	doSomething(42)
	doSomething([]int{1, 2, 3})
	doSomething(map[string]interface{}{"a": 1})
	doSomething(3.14)
}

// ============================================
// 7. 接口值与底层结构
// ============================================
//
// 接口值由两部分组成：(类型, 值)
// - 类型：具体类型的信息
// - 值：具体值的副本或指针
//
// 注意：nil 接口值和值为 nil 的接口值是不同的！

func demonstrateInterfaceInternals() {
	fmt.Println("\n=== 接口值内部 ===")
	
	var p *Dog = nil
	var s Speaker
	
	// s 是 nil 接口值
	fmt.Printf("s == nil: %v\n", s == nil)
	
	// 赋值后，s 不是 nil，即使值是 nil
	s = p
	fmt.Printf("s == nil: %v (注意！不为 nil)\n", s == nil)
	fmt.Printf("s 的类型: %T, 值: %v\n", s, s)
	
	// 调用方法会 panic，因为值是 nil
	// fmt.Println(s.Speak())  // panic!
	
	// 正确检查方式
	if p != nil {
		s = p
		fmt.Println(s.Speak())
	}
}

// ============================================
// 8. 接口的最佳实践
// ============================================

// 小接口原则：接口应该小而专注
type Closer interface {
	Close() error
}

type Flusher interface {
	Flush() error
}

// 接口组合
type WriteFlusher interface {
	Writer
	Flusher
}

// 编译时检查接口实现（推荐）
// 如果不实现，编译会报错
var _ Speaker = (*Dog)(nil)
var _ Speaker = (*Cat)(nil)

// ============================================
// 9. 实用示例：自定义错误类型
// ============================================

// 实现 error 接口
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("验证错误 [%s]: %s", e.Field, e.Message)
}

// 带错误码的错误
type CodedError struct {
	Code    int
	Message string
}

func (e CodedError) Error() string {
	return fmt.Sprintf("错误码 %d: %s", e.Code, e.Message)
}

func demonstrateCustomError() {
	fmt.Println("\n=== 自定义错误 ===")
	
	err1 := ValidationError{Field: "email", Message: "格式不正确"}
	err2 := CodedError{Code: 404, Message: "页面未找到"}
	
	fmt.Println(err1)
	fmt.Println(err2)
	
	// 检查错误类型
	var valErr ValidationError
	if ok := interface{}(err1).(ValidationError); ok.Field == "email" {
		fmt.Println("是邮箱验证错误")
	}
}

// ============================================
// 10. 实用示例：自定义 Stringer
// ============================================

// 实现 fmt.Stringer 接口
type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("Point(%d, %d)", p.X, p.Y)
}

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle{width=%d, height=%d, area=%d}",
		r.Width, r.Height, r.Width*r.Height)
}

func demonstrateStringer() {
	fmt.Println("\n=== 自定义 Stringer ===")
	
	p := Point{X: 10, Y: 20}
	r := Rectangle{Width: 30, Height: 40}
	
	// 使用 %v 或 %s 会自动调用 String() 方法
	fmt.Printf("点: %v\n", p)
	fmt.Printf("矩形: %s\n", r)
}

// ============================================
// 11. 标准库常用接口
// ============================================

func demonstrateStandardInterfaces() {
	fmt.Println("\n=== 标准库接口 ===")
	
	// io.Writer 示例
	var w io.Writer = os.Stdout
	w.Write([]byte("Hello, io.Writer!\n"))
	
	// bytes.Buffer 实现了 io.Writer
	var buf bytes.Buffer
	buf.Write([]byte("写入 buffer"))
	fmt.Println(buf.String())
	
	// 使用 io.Copy
	input := bytes.NewReader([]byte("复制这段文字\n"))
	io.Copy(os.Stdout, input)
	
	// fmt.Stringer 示例
	var s fmt.Stringer = Point{X: 1, Y: 2}
	fmt.Println(s.String())
}

// ============================================
// 12. 依赖注入示例
// ============================================

// 数据存储接口
type UserRepository interface {
	GetUser(id int) (string, error)
	SaveUser(id int, name string) error
}

// 模拟实现
type MockUserRepository struct {
	users map[int]string
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[int]string),
	}
}

func (m *MockUserRepository) GetUser(id int) (string, error) {
	name, ok := m.users[id]
	if !ok {
		return "", fmt.Errorf("用户不存在")
	}
	return name, nil
}

func (m *MockUserRepository) SaveUser(id int, name string) error {
	m.users[id] = name
	return nil
}

// 服务层，依赖接口而非具体实现
type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserName(id int) (string, error) {
	return s.repo.GetUser(id)
}

func demonstrateDependencyInjection() {
	fmt.Println("\n=== 依赖注入 ===")
	
	// 使用模拟实现
	mockRepo := NewMockUserRepository()
	mockRepo.SaveUser(1, "张三")
	
	service := NewUserService(mockRepo)
	
	name, err := service.GetUserName(1)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("用户名:", name)
	}
}

// ============================================
// 主函数
// ============================================

func main() {
	demonstratePolymorphism()
	demonstrateEmptyInterface()
	demonstrateTypeAssertion()
	demonstrateTypeSwitch()
	demonstrateInterfaceInternals()
	demonstrateCustomError()
	demonstrateStringer()
	demonstrateStandardInterfaces()
	demonstrateDependencyInjection()
	
	// ============================================
	// 练习题
	// ============================================
	//
	// 练习 1：定义 Shape 接口，包含 Area() 和 Perimeter() 方法
	//   - 实现 Circle 和 Rectangle 类型
	//   - 编写函数 PrintShapeInfo(s Shape) 打印形状信息
	//   - 创建 Shape 切片，遍历并打印每个形状的信息
	//
	// 练习 2：实现一个通用的 Max 函数，使用接口比较大小
	//   - 定义 Comparable 接口，包含 Compare(other interface{}) int
	//   - 实现 Int 和 String 类型满足该接口
	//   - 实现 Max(a, b Comparable) Comparable 返回较大者
	//
	// 练习 3：实现一个简单的 HTTP Handler 接口模拟
	//   - 定义 Handler 接口，包含 ServeHTTP(request string) string
	//   - 实现 HomeHandler、AboutHandler、NotFoundHandler
	//   - 使用 map[string]Handler 实现路由
	//   - 编写函数处理请求：func Handle(path string, handlers map[string]Handler)
	//
	// 练习 4：实现一个事件系统
	//   - 定义 Event 接口，包含 Type() string 和 Data() interface{}
	//   - 实现 UserLoginEvent、OrderCreatedEvent
	//   - 定义 EventHandler 接口，包含 Handle(e Event)
	//   - 实现 EventBus，支持订阅和发布事件
	//
	// 练习 5：使用空接口实现一个泛型栈（Go 1.18 之前的做法）
	//   type Stack struct { items []interface{} }
	//   - 实现 Push(item interface{})
	//   - 实现 Pop() (interface{}, bool)
	//   - 实现 Peek() (interface{}, bool)
	//   - 实现 IsEmpty() bool
	//   - 注意：使用时需要进行类型断言
	//
	// 练习 6：实现一个可排序的接口体系
	//   - 定义 Sorter 接口，包含 Sort([]interface{}) []interface{}
	//   - 实现 BubbleSorter、QuickSorter
	//   - 实现一个通用函数，接收 Sorter 和待排序数据，返回排序结果
}
