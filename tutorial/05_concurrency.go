// ============================================
// Go 并发编程教程 - Goroutine 与 Channel
// ============================================
//
// 本文件涵盖 Go 语言并发编程的核心：
// - Goroutine ⭐轻量级线程
// - Channel ⭐协程间通信
// - 无缓冲 vs 有缓冲 Channel
// - 单向 Channel
// - select 多路复用
// - 关闭 Channel
// - for-range 遍历 Channel
// - 并发模式（Worker Pool、Pipeline、Fan-out/Fan-in）
//
// 最佳实践：
// 1. 不要通过共享内存来通信，而要通过通信来共享内存
// 2. Channel 的拥有者应该是写入方，负责关闭
// 3. 不要从接收方关闭 channel，不要关闭已经关闭的 channel
// 4. 使用有缓冲 channel 提高性能，但要注意缓冲区大小
// 5. 使用 select 处理多个 channel 操作
// 6. 总是考虑 goroutine 泄漏问题
// ============================================

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ============================================
// 1. Goroutine 基础
// ============================================
//
// Goroutine 是 Go 运行时管理的轻量级线程
// 使用 go 关键字启动
// 调度器使用 GMP 模型（Goroutine - Machine - Processor）

func sayHello() {
	fmt.Println("Hello from goroutine!")
}

func demonstrateGoroutine() {
	fmt.Println("=== Goroutine 基础 ===")
	
	// 启动 goroutine
	go sayHello()
	
	// 使用匿名函数
	go func() {
		fmt.Println("匿名函数 goroutine")
	}()
	
	// goroutine 闭包注意事项：传递参数
	for i := 0; i < 3; i++ {
		// 错误方式：所有 goroutine 共享同一个 i
		// go func() {
		//     fmt.Println(i)  // 可能都打印相同的值
		// }()
		
		// 正确方式：传递参数
		go func(n int) {
			fmt.Printf("Goroutine %d\n", n)
		}(i)
	}
	
	// 等待 goroutine 执行完成
	time.Sleep(100 * time.Millisecond)
}

// ============================================
// 2. Channel 基础
// ============================================
//
// Channel 是 goroutine 之间通信和同步的机制
// 类型：chan T
// 操作：ch <- v（发送），v <- ch（接收）

func demonstrateChannel() {
	fmt.Println("\n=== Channel 基础 ===")
	
	// 创建无缓冲 channel
	ch := make(chan int)
	
	// 启动发送 goroutine
	go func() {
		fmt.Println("发送: 42")
		ch <- 42  // 发送，会阻塞直到有接收者
	}()
	
	// 接收（阻塞直到有数据）
	value := <-ch
	fmt.Printf("接收: %d\n", value)
}

// ============================================
// 3. 无缓冲 vs 有缓冲 Channel
// ============================================

// 无缓冲 channel：同步通信
// 发送和接收必须同时准备好，否则会阻塞
func demonstrateUnbuffered() {
	fmt.Println("\n=== 无缓冲 Channel ===")
	
	ch := make(chan string)
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("发送方：准备发送")
		ch <- "消息"  // 阻塞，直到接收方准备好
		fmt.Println("发送方：发送完成")
	}()
	
	fmt.Println("主线程：等待接收...")
	msg := <-ch  // 阻塞，直到发送方准备好
	fmt.Printf("主线程：收到 %s\n", msg)
}

// 有缓冲 channel：异步通信
// 发送在缓冲区满时阻塞，接收在缓冲区空时阻塞
func demonstrateBuffered() {
	fmt.Println("\n=== 有缓冲 Channel ===")
	
	ch := make(chan int, 3)  // 缓冲区大小为 3
	
	// 发送不会阻塞（缓冲区未满）
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("发送了 3 个值，未阻塞")
	
	// 缓冲区已满，再发送会阻塞
	// ch <- 4  // 会阻塞！
	
	// 接收
	fmt.Printf("接收: %d\n", <-ch)
	fmt.Printf("接收: %d\n", <-ch)
	fmt.Printf("接收: %d\n", <-ch)
}

// ============================================
// 4. 单向 Channel
// ============================================
//
// 只发送：chan<- T
// 只接收：<-chan T

// 只发送的 channel
func producer(ch chan<- int, name string) {
	for i := 0; i < 3; i++ {
		ch <- i
		fmt.Printf("%s 生产: %d\n", name, i)
		time.Sleep(50 * time.Millisecond)
	}
	close(ch)  // 生产者负责关闭
}

// 只接收的 channel
func consumer(ch <-chan int, name string) {
	for val := range ch {  // range 在 channel 关闭时自动退出
		fmt.Printf("%s 消费: %d\n", name, val)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("%s: channel 已关闭\n", name)
}

func demonstrateDirectional() {
	fmt.Println("\n=== 单向 Channel ===")
	
	ch := make(chan int, 2)
	
	go producer(ch, "生产者")
	consumer(ch, "消费者")
}

// ============================================
// 5. Select 多路复用 ⭐
// ============================================
//
// select 同时监听多个 channel 操作
// 随机选择一个可执行的 case
// 配合 default 实现非阻塞操作

func demonstrateSelect() {
	fmt.Println("\n=== Select 多路复用 ===")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	// 启动两个 goroutine
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "来自 ch1"
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "来自 ch2"
	}()
	
	// 同时监听两个 channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("收到:", msg1)
		case msg2 := <-ch2:
			fmt.Println("收到:", msg2)
		}
	}
}

// 非阻塞 select（使用 default）
func demonstrateNonBlocking() {
	fmt.Println("\n=== 非阻塞操作 ===")
	
	ch := make(chan int)
	
	// 非阻塞发送
	select {
	case ch <- 1:
		fmt.Println("发送成功")
	default:
		fmt.Println("发送会被阻塞，执行 default")
	}
	
	// 非阻塞接收
	select {
	case val := <-ch:
		fmt.Println("接收到:", val)
	default:
		fmt.Println("接收会被阻塞，执行 default")
	}
}

// 超时控制
func demonstrateTimeout() {
	fmt.Println("\n=== 超时控制 ===")
	
	ch := make(chan string)
	
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "结果"
	}()
	
	select {
	case result := <-ch:
		fmt.Println("收到结果:", result)
	case <-time.After(1 * time.Second):
		fmt.Println("超时！等待超过 1 秒")
	}
}

// ============================================
// 6. 并发模式：Worker Pool ⭐
// ============================================
//
// 固定数量的 worker 处理任务队列

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for job := range jobs {
		fmt.Printf("Worker %d 开始处理任务 %d\n", id, job)
		
		// 模拟处理时间
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		
		result := job * job  // 计算平方
		results <- result
		
		fmt.Printf("Worker %d 完成任务 %d\n", id, job)
	}
}

func demonstrateWorkerPool() {
	fmt.Println("\n=== Worker Pool ===")
	
	const numJobs = 10
	const numWorkers = 3
	
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	
	var wg sync.WaitGroup
	
	// 启动 workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}
	
	// 发送任务
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)  // 关闭 jobs，worker 会退出
	
	// 等待所有 worker 完成
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// 收集结果
	for result := range results {
		fmt.Printf("结果: %d\n", result)
	}
}

// ============================================
// 7. 并发模式：Pipeline ⭐
// ============================================
//
// 多个处理阶段串联，数据流式处理

// 阶段 1：生成数字
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// 阶段 2：平方
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// 阶段 3：过滤偶数
func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func demonstratePipeline() {
	fmt.Println("\n=== Pipeline 模式 ===")
	
	// 构建 pipeline: gen -> sq -> filterEven
	c := gen(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	c = sq(c)
	c = filterEven(c)
	
	// 消费结果
	for result := range c {
		fmt.Printf("结果: %d\n", result)
	}
}

// ============================================
// 8. 并发模式：Fan-out / Fan-in
// ============================================
//
// Fan-out：多个 goroutine 从同一个 channel 读取
// Fan-in：多个 channel 合并到一个 channel

// fan-in：合并多个 channel
func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	
	// 为每个输入 channel 启动一个 goroutine
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}
	
	// 等待所有输入关闭后关闭输出
	go func() {
		wg.Wait()
		close(out)
	}()
	
	return out
}

func doWork(id int, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for val := range in {
			// 模拟处理
			result := val * 2
			fmt.Printf("Worker %d 处理 %d -> %d\n", id, val, result)
			out <- result
		}
		close(out)
	}()
	return out
}

func demonstrateFanOutFanIn() {
	fmt.Println("\n=== Fan-out / Fan-in ===")
	
	in := gen(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	
	// Fan-out：启动 3 个 worker
	c1 := doWork(1, in)
	c2 := doWork(2, in)
	c3 := doWork(3, in)
	
	// Fan-in：合并结果
	for result := range fanIn(c1, c2, c3) {
		fmt.Printf("最终结果: %d\n", result)
	}
}

// ============================================
// 9. 优雅退出与 Context（预告，详见 06_sync_context.go）
// ============================================

func workerWithQuit(ch chan int, quit chan bool) {
	for {
		select {
		case val := <-ch:
			fmt.Printf("处理: %d\n", val)
		case <-quit:
			fmt.Println("Worker 收到退出信号")
			return
		}
	}
}

func demonstrateGracefulShutdown() {
	fmt.Println("\n=== 优雅退出 ===")
	
	ch := make(chan int)
	quit := make(chan bool)
	
	go workerWithQuit(ch, quit)
	
	// 发送一些任务
	for i := 0; i < 3; i++ {
		ch <- i
	}
	
	time.Sleep(100 * time.Millisecond)
	quit <- true  // 发送退出信号
	time.Sleep(50 * time.Millisecond)
}

// ============================================
// 10. 常见陷阱与注意事项
// ============================================

func demonstratePitfalls() {
	fmt.Println("\n=== 常见陷阱 ===")
	
	// 1. 向 nil channel 发送会永远阻塞
	var ch chan int  // nil channel
	// ch <- 1  // 永远阻塞！
	
	// 2. 关闭 nil channel 会 panic
	// close(ch)  // panic!
	
	// 3. 向已关闭的 channel 发送会 panic
	ch2 := make(chan int)
	close(ch2)
	// ch2 <- 1  // panic!
	
	// 4. 重复关闭 channel 会 panic
	// close(ch2)  // panic!
	
	// 5. 从已关闭的 channel 接收会立即返回零值
	v, ok := <-ch2
	fmt.Printf("从关闭 channel 接收: %d, ok=%v\n", v, ok)  // 0, false
	
	// 6. range 遍历已关闭 channel 会正常退出
	ch3 := make(chan int, 3)
	ch3 <- 1
	ch3 <- 2
	close(ch3)
	
	fmt.Println("Range 遍历已关闭 channel:")
	for v := range ch3 {
		fmt.Println(v)
	}
	fmt.Println("Range 正常退出")
}

// ============================================
// 主函数
// ============================================

func main() {
	rand.Seed(time.Now().UnixNano())
	
	demonstrateGoroutine()
	demonstrateChannel()
	demonstrateUnbuffered()
	demonstrateBuffered()
	demonstrateDirectional()
	demonstrateSelect()
	demonstrateNonBlocking()
	demonstrateTimeout()
	demonstrateWorkerPool()
	demonstratePipeline()
	demonstrateFanOutFanIn()
	demonstrateGracefulShutdown()
	demonstratePitfalls()
	
	// ============================================
	// 练习题
	// ============================================
	//
	// 练习 1：实现一个并发素数筛（Sieve of Eratosthenes）
	//   - 使用 pipeline 模式
	//   - 每个阶段过滤一个素数的倍数
	//   - 生成前 100 个素数
	//
	// 练习 2：实现一个带并发限制的 HTTP 爬虫
	//   - 接收 URL 列表
	//   - 使用 worker pool 限制并发数（如最多 5 个并发）
	//   - 返回每个 URL 的内容长度
	//   - 支持超时控制
	//
	// 练习 3：实现一个并发安全的计数器
	//   type Counter struct { count int }
	//   - 使用 channel 实现（不要使用 mutex）
	//   - 支持 Inc() 和 Get() 操作
	//   - 支持 Reset()
	//
	// 练习 4：实现一个广播系统
	//   - 一个发送者，多个接收者
	//   - 每个接收者都能收到所有消息
	//   - 支持动态添加/移除接收者
	//
	// 练习 5：实现一个任务调度器
	//   - 可以提交延迟执行的任务
	//   - 支持取消未执行的任务
	//   - 使用优先队列（可用 time.After）
	//
	// 练习 6：实现一个速率限制器（Token Bucket）
	//   - 使用 channel 作为令牌桶
	//   - 控制请求的速率
	//   - 支持突发流量
	//
	// 练习 7：实现一个并行归并排序
	//   - 对切片进行排序
	//   - 使用 goroutine 并行处理子数组
	//   - 设置阈值，小数组使用普通排序
}
