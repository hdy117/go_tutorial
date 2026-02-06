# Go 语言核心特性练习题

本文件汇总了所有教学文件中的练习题，按照难度分级。

## 难度说明

- ⭐ 初级：适合刚学完相关概念
- ⭐⭐ 中级：需要综合运用多个知识点
- ⭐⭐⭐ 高级：需要深入理解底层原理或设计模式

---

## 01_basic_syntax.go 练习题

### 练习 1：学生成绩管理 ⭐
创建一个 map 存储学生姓名和分数，实现以下功能：
- 添加 3 个学生
- 查询某个学生的分数
- 计算平均分
- 删除分数低于 60 分的学生

### 练习 2：切片去重 ⭐
```go
func removeDuplicates(nums []int) []int
```
实现切片去重，保持原有顺序。

### 练习 3：文件权限常量 ⭐
使用 iota 定义文件权限常量（类似 Linux）：
- Owner 可读可写可执行
- Group 可读可执行
- Other 只读

### 练习 4：查找极值 ⭐
```go
func findMinMax(nums []int) (min, max int)
```
找出切片中的最大值和最小值。

### 练习 5：猜数字游戏 ⭐⭐
实现一个简单的猜数字游戏：
- 随机生成 1-100 的数字
- 用户输入猜测，程序提示"太大"或"太小"
- 使用循环直到猜对
- 记录猜测次数

---

## 02_functions.go 练习题

### 练习 1：变长参数极值 ⭐
```go
func minMax(nums ...int) (min, max int, err error)
```
- 接收任意数量的整数，返回最大值和最小值
- 错误处理：如果没有传入参数，返回错误

### 练习 2：累加器闭包 ⭐⭐
```go
func makeAccumulator(initial int) (add, sub func(int) int)
```
- add(5) 表示加5
- sub(3) 表示减3
- 两个函数共享同一个状态

### 练习 3：切片过滤 ⭐
```go
func filter(nums []int, predicate func(int) bool) []int
```
接收一个整数切片和一个过滤函数，返回满足条件的元素。

### 练习 4：函数计时器 ⭐
使用 defer 实现一个函数计时器，能够计算并打印函数执行时间。
提示：使用 time.Since

### 练习 5：记忆化函数 ⭐⭐
```go
func memoize(f func(int) int) func(int) int
```
缓存任意函数的结果，避免重复计算。

### 练习 6：函数管道 ⭐⭐
```go
func pipeline(data int, funcs ...func(int) int) int
```
示例：`pipeline(5, double, addOne, square) = ((5*2)+1)^2 = 121`

---

## 03_struct_method.go 练习题

### 练习 1：矩形结构体 ⭐
定义一个 Rectangle 结构体，包含 Width 和 Height：
- 实现 Area() 计算面积
- 实现 Perimeter() 计算周长
- 实现 Scale(factor float64) 按因子缩放（修改原值）
- 实现 IsSquare() 判断是否为正方形

### 练习 2：图书管理系统 ⭐⭐
实现一个 Book 结构体：
- 字段：Title, Author, ISBN, Price, PublishedYear
- 实现 ApplyDiscount(discountPercent float64) 打折
- 实现 GetAge() 返回书的"年龄"
- 实现 String() string 方法（格式化输出）

### 练习 3：学校人员系统 ⭐⭐
使用嵌入实现以下结构：
- 基础 Person 结构体（Name, Age）
- Student 嵌入 Person，添加 StudentID, Major, Grades([]float64)
- Teacher 嵌入 Person，添加 TeacherID, Department, Salary
- 为 Student 实现 GetAverageGrade() 方法

### 练习 4：TTL 缓存 ⭐⭐⭐
```go
type Cache struct {
    data map[string]interface{}
    ttl  map[string]time.Time
}
```
- 实现 Set(key string, value interface{}, duration time.Duration)
- 实现 Get(key string) (interface{}, bool)
- 实现 Delete(key string)
- Get 时检查是否过期

### 练习 5：链表实现 ⭐⭐⭐
```go
type Node struct {
    Value int
    Next  *Node
}
```
- 实现 Append(value int) 在尾部添加
- 实现 Insert(index, value int) 在指定位置插入
- 实现 Delete(index int) 删除指定位置
- 实现 Reverse() 反转链表
- 实现 String() 打印链表内容

---

## 04_interface.go 练习题

### 练习 1：形状接口 ⭐
定义 Shape 接口，包含 Area() 和 Perimeter() 方法：
- 实现 Circle 和 Rectangle 类型
- 编写函数 PrintShapeInfo(s Shape) 打印形状信息
- 创建 Shape 切片，遍历并打印每个形状的信息

### 练习 2：可比较接口 ⭐⭐
实现通用的 Max 函数，使用接口比较大小：
- 定义 Comparable 接口，包含 Compare(other interface{}) int
- 实现 Int 和 String 类型满足该接口
- 实现 Max(a, b Comparable) Comparable 返回较大者

### 练习 3：HTTP 路由系统 ⭐⭐⭐
实现一个简单的 HTTP Handler 接口模拟：
- 定义 Handler 接口，包含 ServeHTTP(request string) string
- 实现 HomeHandler、AboutHandler、NotFoundHandler
- 使用 map[string]Handler 实现路由
- 编写函数处理请求：func Handle(path string, handlers map[string]Handler)

### 练习 4：事件系统 ⭐⭐⭐
实现一个事件系统：
- 定义 Event 接口，包含 Type() string 和 Data() interface{}
- 实现 UserLoginEvent、OrderCreatedEvent
- 定义 EventHandler 接口，包含 Handle(e Event)
- 实现 EventBus，支持订阅和发布事件

### 练习 5：泛型栈（Go 1.18 之前做法）⭐⭐
使用空接口实现一个泛型栈：
```go
type Stack struct { items []interface{} }
```
- 实现 Push(item interface{})
- 实现 Pop() (interface{}, bool)
- 实现 Peek() (interface{}, bool)
- 实现 IsEmpty() bool
- 注意：使用时需要进行类型断言

### 练习 6：排序器接口 ⭐⭐
实现可排序的接口体系：
- 定义 Sorter 接口，包含 Sort([]interface{}) []interface{}
- 实现 BubbleSorter、QuickSorter
- 实现一个通用函数，接收 Sorter 和待排序数据，返回排序结果

---

## 05_concurrency.go 练习题

### 练习 1：并发素数筛 ⭐⭐⭐
实现一个并发素数筛（Sieve of Eratosthenes）：
- 使用 pipeline 模式
- 每个阶段过滤一个素数的倍数
- 生成前 100 个素数

### 练习 2：并发爬虫 ⭐⭐⭐
实现一个带并发限制的 HTTP 爬虫：
- 接收 URL 列表
- 使用 worker pool 限制并发数（如最多 5 个并发）
- 返回每个 URL 的内容长度
- 支持超时控制

### 练习 3：Channel 计数器 ⭐⭐
实现一个并发安全的计数器：
```go
type Counter struct { count int }
```
- 使用 channel 实现（不要使用 mutex）
- 支持 Inc() 和 Get() 操作
- 支持 Reset()

### 练习 4：广播系统 ⭐⭐⭐
实现一个广播系统：
- 一个发送者，多个接收者
- 每个接收者都能收到所有消息
- 支持动态添加/移除接收者

### 练习 5：任务调度器 ⭐⭐⭐
实现一个任务调度器：
- 可以提交延迟执行的任务
- 支持取消未执行的任务
- 使用优先队列（可用 time.After）

### 练习 6：速率限制器 ⭐⭐⭐
实现一个速率限制器（Token Bucket）：
- 使用 channel 作为令牌桶
- 控制请求的速率
- 支持突发流量

### 练习 7：并行归并排序 ⭐⭐⭐
实现一个并行归并排序：
- 对切片进行排序
- 使用 goroutine 并行处理子数组
- 设置阈值，小数组使用普通排序

---

## 06_sync_context.go 练习题

### 练习 1：环形缓冲区 ⭐⭐⭐
```go
type RingBuffer struct { ... }
```
- 使用 Mutex 保护
- 实现 Write(data []byte) (n int, err error)
- 实现 Read(p []byte) (n int, err error)
- 当缓冲区满时，Write 阻塞；空时，Read 阻塞

### 练习 2：信号量 ⭐⭐⭐
```go
type Semaphore struct { ... }
```
- 使用 Channel 实现
- Acquire() 获取许可，如果没有则阻塞
- Release() 释放许可
- TryAcquire(timeout time.Duration) bool 带超时的获取

### 练习 3：读写缓存 ⭐⭐⭐
```go
type RWCache struct { ... }
```
- 使用 RWMutex
- 支持 Set、Get、Delete
- 支持 TTL（过期时间），使用 goroutine 定期清理

### 练习 4：加权负载均衡器 ⭐⭐⭐
```go
type LoadBalancer struct { ... }
```
- 后端服务器有权重
- 使用 atomic 实现无锁的轮询
- 支持动态添加/移除后端

### 练习 5：断路器 ⭐⭐⭐⭐
```go
type CircuitBreaker struct { ... }
```
- 状态：Closed、Open、Half-Open
- 失败次数超过阈值进入 Open
- Open 状态经过超时后进入 Half-Open
- Half-Open 成功则 Closed，失败则 Open
- 使用 sync/atomic 或 Mutex 保证并发安全

### 练习 6：分布式锁 ⭐⭐⭐⭐
```go
type DistributedLock struct { ... }
```
- Lock() 获取锁，阻塞直到成功
- TryLock(timeout time.Duration) bool 带超时
- Unlock() 释放锁
- 使用 Context 支持取消

### 练习 7：限流器 ⭐⭐⭐
```go
type RateLimiter struct { ... }
```
- 使用令牌桶算法
- Allow() bool 判断是否允许通过
- Wait(ctx context.Context) error 等待直到允许通过

---

## 07_error_handling.go 练习题

### 练习 1：堆栈错误 ⭐⭐
```go
type StackError struct { error; stack []byte }
```
- 创建错误时捕获堆栈
- 实现 Error() 方法，输出错误信息和堆栈

### 练习 2：错误码系统 ⭐⭐
- 定义错误码常量（如 ErrCodeNotFound = 404）
- 实现 CodedError 结构体，包含 Code 和 Message
- 实现 FromCode(code int) 根据 HTTP 状态码创建错误
- 实现 HTTPStatus() 返回对应的 HTTP 状态码

### 练习 3：批处理错误 ⭐⭐⭐
```go
type BatchProcessor struct { ... }
```
- 处理多个项目，收集所有错误
- 如果所有错误都是同一种类型，返回该类型错误
- 如果有多种错误，返回 MultiError

### 练习 4：上下文错误 ⭐⭐
```go
type ContextError struct { error; Context map[string]interface{} }
```
- 支持添加键值对上下文
- Error() 输出时包含上下文信息
- 实现 Unwrap() 支持错误链

### 练习 5：断言工具 ⭐
```go
func AssertNotNil(v interface{}, msg string)
func AssertTrue(condition bool, msg string)
func AssertNoError(err error)
```
- 断言失败时 panic
- 在测试中使用

### 练习 6：错误重试装饰器 ⭐⭐⭐
```go
func Retryable(fn func() error, opts RetryOptions) func() error
```
- 支持自定义重试次数、退避策略
- 支持只对特定错误重试
- 支持超时

---

## 08_generics.go 练习题

### 练习 1：泛型集合操作 ⭐
- Map 将 []T 转换为 []U
- Filter 根据条件过滤元素
- Reduce 将切片归约为单个值
- 编写测试验证功能

### 练习 2：泛型缓存 ⭐⭐
```go
type Cache[K comparable, V any] struct { ... }
```
- Set(key K, value V, ttl time.Duration)
- Get(key K) (V, bool)
- Delete(key K)
- 支持 TTL 自动过期

### 练习 3：泛型 Channel 操作 ⭐⭐
```go
func MapChan[T, U any](input <-chan T, fn func(T) U) <-chan U
func FilterChan[T any](input <-chan T, predicate func(T) bool) <-chan T
func ReduceChan[T, U any](input <-chan T, initial U, fn func(U, T) U) U
```

### 练习 4：泛型排序算法 ⭐⭐
```go
func QuickSort[T constraints.Ordered](slice []T)
func MergeSort[T constraints.Ordered](slice []T)
```
- 支持任意可排序类型
- 与 sort.Slice 性能对比

### 练习 5：函数组合 ⭐⭐⭐
```go
func Compose[A, B, C any](f func(B) C, g func(A) B) func(A) C
func Pipe[A, B, C any](f func(A) B, g func(B) C) func(A) C
func Curry[A, B, C any](f func(A, B) C) func(A) func(B) C
```
- 验证函数组合的正确性

### 练习 6：泛型状态机 ⭐⭐⭐
```go
type StateMachine[S comparable, E any] struct { ... }
```
- AddTransition(from S, event E, to S)
- Trigger(event E) error
- 支持状态转换验证

### 练习 7：泛型依赖注入 ⭐⭐⭐⭐
```go
type Container struct { ... }
```
- Register[T any](constructor func(...) T)
- Resolve[T any]() (T, error)
- 支持单例和瞬态生命周期

---

## 09_reflect.go 练习题

### 练习 1：Map 转换 ⭐⭐
```go
func TransformMap(input interface{}, fn func(interface{}) interface{}) interface{}
```
- 支持任意类型的 map
- 对每个值应用转换函数

### 练习 2：结构体转 Map ⭐⭐⭐
```go
func StructToMap(s interface{}) map[string]interface{}
```
- 只处理导出字段
- 使用 json tag 作为 key
- 递归处理嵌套结构体

### 练习 3：Map 转结构体 ⭐⭐⭐
```go
func MapToStruct(m map[string]interface{}, s interface{}) error
```
- 使用反射设置结构体字段
- 处理类型转换
- 支持嵌套结构体

### 练习 4：依赖注入容器 ⭐⭐⭐⭐
```go
type DIContainer struct { ... }
```
- Register(constructor interface{}) 注册构造函数
- Resolve(target interface{}) error 解析依赖
- 自动注入构造函数参数

### 练习 5：RPC 调用器 ⭐⭐⭐⭐
```go
type RPCClient struct { ... }
```
- Call(method string, args []interface{}, reply interface{}) error
- 使用反射检查方法签名
- 验证参数数量和类型

### 练习 6：ORM 查询构建器 ⭐⭐⭐⭐
```go
type Query struct { ... }
```
- Where(field string, op string, value interface{}) *Query
- Find(dest interface{}) error
- 使用反射填充结果到结构体切片

### 练习 7：JSON Schema 生成器 ⭐⭐⭐
```go
func GenerateSchema(t interface{}) map[string]interface{}
```
- 从结构体标签生成 JSON Schema
- 支持 required、type、format 等字段

---

## 10_standard_lib.go 练习题

### 练习 1：日志分析工具 ⭐⭐
- 读取日志文件
- 使用 regexp 解析日志格式
- 统计各种级别的日志数量（INFO, WARN, ERROR）
- 按时间范围过滤日志

### 练习 2：Web 爬虫 ⭐⭐⭐
- 接收起始 URL
- 使用 http.Get 获取页面
- 使用 regexp 提取所有链接
- 递归爬取（限制深度）
- 保存页面内容到文件

### 练习 3：配置文件解析器 ⭐⭐⭐
- 支持 JSON 格式
- 支持环境变量替换（${VAR}）
- 支持默认值（${VAR:-default}）
- 将配置加载到结构体

### 练习 4：CSV 处理工具 ⭐⭐⭐
- 读取 CSV 文件
- 解析为结构体切片
- 支持类型转换（使用 strconv）
- 写入 CSV 文件
- 支持过滤和排序

### 练习 5：文件同步工具 ⭐⭐⭐
- 比较两个目录的内容
- 使用 filepath.Walk 遍历
- 根据修改时间决定同步方向
- 支持排除某些文件模式

### 练习 6：HTTP 中间件链 ⭐⭐⭐
- LoggingMiddleware - 记录请求日志
- AuthMiddleware - 简单的 Token 验证
- RateLimitMiddleware - 限流
- RecoveryMiddleware - panic 恢复
- 使用函数式编程组合中间件

### 练习 7：模板引擎（简化版）⭐⭐⭐⭐
- 支持变量替换 {{.Name}}
- 支持条件语句 {{if .Condition}}...{{end}}
- 支持循环 {{range .Items}}...{{end}}
- 使用 regexp 和 strings 实现

---

## 学习建议

1. **循序渐进**：按照文件顺序完成练习
2. **先独立完成**：遇到困难再看参考答案
3. **编写测试**：为每个练习编写单元测试
4. **代码审查**：完成后回顾代码，思考是否有改进空间
5. **扩展练习**：在基础上添加更多功能

## 参考答案

参考答案可以在 `tutorial/solutions/` 目录中找到（如需要，请告知我创建）。
