# Go 语言核心技术脑图

```
Go 语言核心技术
│
├── 1. 基础语法 (Basic Syntax)
│   ├── 变量与常量
│   │   ├── var / const / :=
│   │   ├── iota 枚举
│   │   └── 类型推断
│   ├── 数据类型
│   │   ├── 基本类型: int, float64, string, bool
│   │   ├── 复合类型: array, slice, map, struct
│   │   ├── 指针 (*T)
│   │   └── 函数类型
│   ├── 流程控制
│   │   ├── if / else / switch / fallthrough
│   │   ├── for / range
│   │   ├── break / continue / goto
│   │   └── defer (延迟执行)
│   └── 函数
│       ├── 多返回值
│       ├── 命名返回值
│       ├── 变长参数 (...T)
│       ├── 闭包 (Closure)
│       └── 递归
│
├── 2. 面向对象 (OOP in Go)
│   ├── 结构体 (Struct)
│   │   ├── 定义与初始化
│   │   ├── 嵌入类型 (Embedding)
│   │   └── 标签 (Tag)
│   ├── 方法 (Method)
│   │   ├── 值接收者 vs 指针接收者
│   │   └── 方法集
│   └── 接口 (Interface)
│       ├── 隐式实现
│       ├── 空接口 (interface{} / any)
│       ├── 类型断言 (Type Assertion)
│       ├── 类型开关 (Type Switch)
│       └── 接口组合
│
├── 3. 并发编程 (Concurrency) ⭐核心特色
│   ├── Goroutine
│   │   ├── 轻量级线程
│   │   ├── go 关键字
│   │   └── 调度器 (GMP模型)
│   ├── Channel
│   │   ├── 无缓冲 Channel
│   │   ├── 有缓冲 Channel
│   │   ├── 单向 Channel (<-chan, chan<-)
│   │   ├── select 多路复用
│   │   └── 关闭与range遍历
│   ├── 同步原语 (sync包)
│   │   ├── Mutex / RWMutex (互斥锁)
│   │   ├── WaitGroup (等待组)
│   │   ├── Once (一次性执行)
│   │   ├── Pool (对象池)
│   │   └── Map (并发安全Map)
│   └── Context
│       ├── 超时控制
│       ├── 取消信号
│       └── 值传递
│
├── 4. 错误处理 (Error Handling)
│   ├── error 接口
│   ├── 自定义错误类型
│   ├── errors.New / fmt.Errorf
│   ├── 错误链 (Error Wrapping) - Go 1.13+
│   │   ├── %w 格式化动词
│   │   ├── errors.Is
│   │   └── errors.As
│   └── panic / recover
│       ├── 何时使用panic
│       └── recover的注意事项
│
├── 5. 反射 (Reflect)
│   ├── reflect.Type / reflect.Value
│   ├── 类型检查与转换
│   ├── 动态创建对象
│   ├── 修改变量值
│   └── 结构体标签解析
│
├── 6. 标准库 (Standard Library)
│   ├── 常用包
│   │   ├── fmt - 格式化I/O
│   │   ├── strings / bytes - 字符串处理
│   │   ├── strconv - 类型转换
│   │   ├── time - 时间处理
│   │   ├── os / path/filepath - 文件系统
│   │   ├── io / bufio - I/O操作
│   │   ├── encoding/json - JSON处理
│   │   ├── net/http - HTTP服务
│   │   ├── database/sql - 数据库
│   │   └── regexp - 正则表达式
│   └── 常用工具
│       ├── sort - 排序
│       ├── container (heap, list, ring)
│       └── math / math/rand
│
├── 7. 测试与调试 (Testing)
│   ├── 单元测试 (testing.T)
│   ├── 基准测试 (testing.B)
│   ├── 模糊测试 (testing.F) - Go 1.18+
│   ├── 覆盖率测试
│   ├── 表驱动测试
│   ├── Mock / Stub
│   └── pprof 性能分析
│
├── 8. 工程实践 (Engineering)
│   ├── 包管理 (Go Modules)
│   │   ├── go.mod / go.sum
│   │   ├── 版本管理
│   │   ├── 私有仓库
│   │   └── 代理配置
│   ├── 代码组织
│   │   ├── 包设计原则
│   │   ├── 可见性 (大写导出)
│   │   └── 内聚与耦合
│   ├── 设计模式
│   │   ├── 工厂模式
│   │   ├── 单例模式
│   │   ├── 观察者模式
│   │   └── 策略模式
│   └── 性能优化
│       ├── 内存分配
│       ├── GC调优
│       ├── 逃逸分析
│       └── Benchmark
│
└── 9. 高级特性 (Advanced)
    ├── 泛型 (Generics) - Go 1.18+
    │   ├── 类型参数
    │   ├── 类型约束
    │   └── 泛型函数/类型
    ├── CGO
    │   ├── 调用C代码
    │   └── 跨平台编译
    ├── 编译与构建
    │   ├── go build / install
    │   ├── 交叉编译
    │   ├── Build Tags
    │   └── ldflags
    └── 运行时
        ├── GMP调度模型
        ├── 内存管理
        └── GC机制
```

---

## 核心代码示例


---

## 核心代码示例

### 示例1: Goroutine + Channel (并发的核心)

```go
package main

import (
	"fmt"
	"time"
)

// 生产者-消费者模式
func producer(ch chan<- int, name string) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("%s 生产: %d\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch <-chan int, name string) {
	for val := range ch {
		fmt.Printf("%s 消费: %d\n", name, val)
		time.Sleep(150 * time.Millisecond)
	}
}

func main() {
	ch := make(chan int, 3) // 有缓冲channel
	
	go producer(ch, "生产者")
	consumer(ch, "消费者")
}
```

### 示例2: select 多路复用

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "来自channel 1"
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "来自channel 2"
	}()
	
	// 同时监听多个channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-time.After(300 * time.Millisecond):
			fmt.Println("超时!")
		}
	}
}
```

### 示例3: 接口与多态

```go
package main

import "fmt"

// 定义接口
type Animal interface {
	Speak() string
}

// Dog 实现
type Dog struct{}

func (d Dog) Speak() string {
	return "汪汪!"
}

// Cat 实现
type Cat struct{}

func (c Cat) Speak() string {
	return "喵喵!"
}

// 多态函数
func MakeSound(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	animals := []Animal{Dog{}, Cat{}}
	for _, animal := range animals {
		MakeSound(animal) // 多态调用
	}
}
```

### 示例4: 并发控制 (WaitGroup + Mutex)

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := 0
	
	// 启动10个goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			mu.Lock()
			counter++
			fmt.Printf("Goroutine %d: counter = %d\n", id, counter)
			mu.Unlock()
		}(i)
	}
	
	wg.Wait() // 等待所有goroutine完成
	fmt.Printf("最终计数: %d\n", counter)
}
```

### 示例5: Context 超时控制

```go
package main

import (
	"context"
	"fmt"
	"time"
)

// 模拟耗时操作
func slowOperation(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		return fmt.Errorf("操作完成")
	case <-ctx.Done():
		return ctx.Err() // context deadline exceeded
	}
}

func main() {
	// 设置2秒超时
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	if err := slowOperation(ctx); err != nil {
		fmt.Println("错误:", err)
	}
}
```

### 示例6: 泛型 (Go 1.18+)

```go
package main

import "fmt"

// 泛型函数: 求最大值
func Max[T comparable](a, b T) T {
	// 注意: comparable 只支持 == 和 !=
	// 这里仅作示例，实际数值比较需要 constraints.Ordered
	if a > b {
		return a
	}
	return b
}

// 泛型栈
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func main() {
	// 整型栈
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	fmt.Println(intStack.Pop()) // 2
	
	// 字符串栈
	strStack := Stack[string]{}
	strStack.Push("hello")
	fmt.Println(strStack.Pop()) // hello
}
```

### 示例7: 错误处理最佳实践

```go
package main

import (
	"errors"
	"fmt"
)

// 自定义错误类型
type NotFoundError struct {
	Resource string
	ID       int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID %d not found", e.Resource, e.ID)
}

// 包装错误
func findUser(id int) error {
	// 模拟查找失败
	err := errors.New("database connection failed")
	return fmt.Errorf("failed to find user %d: %w", id, err)
}

func main() {
	err := findUser(42)
	
	// 错误链检查
	if errors.Is(err, errors.New("database connection failed")) {
		// 检查是否是特定错误
	}
	
	// 类型断言
	var notFound *NotFoundError
	if errors.As(err, &notFound) {
		fmt.Printf("未找到资源: %s\n", notFound.Resource)
	}
	
	fmt.Println(err)
}
```

### 示例8: 反射应用

```go
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	
	// 获取类型信息
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)
	
	fmt.Println("类型:", t.Name())
	fmt.Println("字段数:", t.NumField())
	
	// 遍历字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("字段: %s, 值: %v, Tag: %s\n", 
			field.Name, value, field.Tag.Get("json"))
	}
	
	// 修改值 (需要指针)
	v2 := reflect.ValueOf(&p).Elem()
	v2.FieldByName("Age").SetInt(31)
	fmt.Println(p) // {Alice 31}
}
```

---

## 学习路线图

```
第一阶段: 基础入门 (1-2周)
├── 环境搭建与基本语法
├── 变量、类型、控制结构
├── 函数、数组、切片、map
└── 结构体与方法

第二阶段: 进阶核心 (2-3周)
├── 接口与多态 ⭐
├── Goroutine与Channel ⭐⭐
├── 标准库常用包
└── 错误处理与defer

第三阶段: 工程实践 (2-3周)
├── Go Modules包管理
├── 单元测试与基准测试
├── 并发模式与同步原语
└── 性能分析与优化

第四阶段: 高级特性 (持续学习)
├── 泛型编程
├── 反射与 unsafe
├── CGO与底层原理
└── 源码阅读 (runtime, net/http等)
```

---

*此脑图涵盖了Go语言的核心技术栈，建议结合官方文档和实际项目深入学习。*
