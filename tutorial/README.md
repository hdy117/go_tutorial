# Go 语言核心特性教程

本教程包含 10 个教学文件，涵盖 Go 语言的核心特性，每个文件都包含详细的注释、示例代码和练习题。

## 文件结构

```
tutorial/
├── README.md              # 本文件
├── 01_basic_syntax.go     # 基础语法（变量、类型、控制流、数组、切片、Map）
├── 02_functions.go        # 函数特性（多返回值、闭包、defer、递归）
├── 03_struct_method.go    # 结构体与方法（值/指针接收者、嵌入）
├── 04_interface.go        # 接口（隐式实现、类型断言、空接口）
├── 05_concurrency.go      # 并发编程（Goroutine、Channel、并发模式）
├── 06_sync_context.go     # 同步原语与 Context（Mutex、WaitGroup、Context）
├── 07_error_handling.go   # 错误处理（自定义错误、错误链、panic/recover）
├── 08_generics.go         # 泛型编程（类型参数、约束、泛型容器）
├── 09_reflect.go          # 反射（类型检查、值操作、结构体反射）
├── 10_standard_lib.go     # 标准库常用包
└── exercises.md           # 练习题汇总
```

## 学习路线

### 第一阶段：基础入门（第 1-3 周）
1. **01_basic_syntax.go** - 掌握 Go 的基础语法
2. **02_functions.go** - 理解函数的高级特性
3. **03_struct_method.go** - 学习面向对象编程

### 第二阶段：进阶核心（第 4-6 周）
4. **04_interface.go** - 掌握接口和鸭子类型
5. **05_concurrency.go** - Go 的核心特色：并发编程
6. **06_sync_context.go** - 同步原语和上下文控制

### 第三阶段：工程实践（第 7-9 周）
7. **07_error_handling.go** - 错误处理最佳实践
8. **08_generics.go** - 泛型编程（Go 1.18+）
9. **09_reflect.go** - 反射的使用和注意事项
10. **10_standard_lib.go** - 标准库常用包

## 如何使用

### 运行教学文件

每个教学文件都可以直接运行，查看示例输出：

```bash
cd tutorial
go run 01_basic_syntax.go
go run 02_functions.go
# ... 以此类推
```

### 完成练习题

1. 打开 `exercises.md` 查看练习题
2. 在对应的教学文件中，找到练习题部分
3. 根据注释提示，完成代码实现
4. 运行测试验证结果

例如，在 `01_basic_syntax.go` 文件末尾：

```go
// 练习题（请在此文件基础上完成）
//
// 练习 1：创建一个 map 存储学生姓名和分数...
```

### 编写测试

建议为每个练习题编写单元测试：

```go
func TestRemoveDuplicates(t *testing.T) {
    input := []int{1, 2, 2, 3, 3, 3}
    expected := []int{1, 2, 3}
    result := removeDuplicates(input)
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}
```

运行测试：
```bash
go test -v
```

## 文件内容说明

### 01_basic_syntax.go
- 变量声明与初始化
- 常量与 iota
- 基本数据类型
- 控制流程（if、for、switch）
- 数组和切片（Slice）
- Map
- range 遍历

### 02_functions.go
- 函数定义与调用
- 多返回值 ⭐
- 命名返回值
- 变长参数
- 函数作为值和类型
- 闭包（Closure）⭐
- defer 延迟执行 ⭐
- 递归和 init 函数

### 03_struct_method.go
- 结构体定义与初始化
- 方法定义
- 值接收者 vs 指针接收者 ⭐
- 结构体嵌入（Embedding）⭐
- 结构体标签（Tag）
- 方法集
- 完整示例：银行账户

### 04_interface.go
- 接口定义与隐式实现 ⭐
- 空接口（interface{} / any）
- 类型断言（Type Assertion）⭐
- 类型开关（Type Switch）
- 接口组合
- 自定义错误类型
- 依赖注入示例

### 05_concurrency.go
- Goroutine 基础 ⭐
- Channel 基础 ⭐
- 无缓冲 vs 有缓冲 Channel
- 单向 Channel
- Select 多路复用 ⭐
- Worker Pool 模式 ⭐
- Pipeline 模式 ⭐
- Fan-out / Fan-in
- 常见陷阱与注意事项

### 06_sync_context.go
- Mutex / RWMutex ⭐
- WaitGroup ⭐
- Once（单例模式）
- Pool（对象池）
- Map（并发安全 Map）
- Atomic（原子操作）
- Context（上下文控制）⭐
- 综合示例：任务队列

### 07_error_handling.go
- error 接口
- 创建错误
- 自定义错误类型
- 错误链（Error Wrapping）⭐
- errors.Is 和 errors.As
- panic 和 recover
- 错误处理模式
- 多重错误和重试

### 08_generics.go
- 泛型函数
- 类型约束（Constraints）⭐
- 自定义约束
- 泛型类型（Stack、Queue、Set）⭐
- 泛型接口
- 类型推导
- 实用模式（Option、Result）

### 09_reflect.go
- reflect.Type 和 reflect.Value
- 修改值
- 类型检查与转换
- 结构体反射 ⭐
- 结构体标签解析
- 方法反射
- 切片和 Map 反射
- 实用工具（深拷贝、验证器）

### 10_standard_lib.go
- fmt - 格式化 I/O
- strings/bytes - 字符串操作
- strconv - 类型转换
- time - 时间处理
- os/filepath - 文件系统
- io/bufio - I/O 操作
- encoding/json - JSON 处理
- net/http - HTTP 服务
- sort - 排序
- regexp - 正则表达式

## 练习题难度

- ⭐ 初级：适合刚学完相关概念
- ⭐⭐ 中级：需要综合运用多个知识点
- ⭐⭐⭐ 高级：需要深入理解底层原理或设计模式
- ⭐⭐⭐⭐ 专家：需要综合运用多个模块的知识

## 学习建议

1. **循序渐进**：按照文件顺序学习，不要跳过一个文件
2. **动手实践**：不要只是阅读，要动手运行和修改代码
3. **完成练习**：练习题是巩固知识的关键
4. **编写测试**：养成编写单元测试的习惯
5. **深入理解**：对于核心概念（如接口、并发），要深入理解其设计原理
6. **阅读源码**：学习标准库的源码实现

## 参考资料

- [Go 官方文档](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go 语言设计与实现](https://draveness.me/golang/)

## 贡献

欢迎提交 Issue 和 PR 来改进本教程！

## 许可证

MIT License
