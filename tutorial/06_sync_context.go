// ============================================
// Go 同步原语与 Context 教程
// ============================================
//
// 本文件涵盖 Go 语言并发同步的核心工具：
// - sync.Mutex / RWMutex（互斥锁）
// - sync.WaitGroup（等待组）
// - sync.Once（一次性执行）
// - sync.Pool（对象池）
// - sync.Map（并发安全 Map）
// - atomic（原子操作）
// - Context（上下文控制）⭐
//
// 最佳实践：
// 1. 优先使用 channel 进行通信，必要时使用 sync 包
// 2. 锁的粒度要小，持有锁的时间要短
// 3. 避免死锁：按固定顺序获取多个锁
// 4. 不要复制包含锁的结构体
// 5. RWMutex 在读多写少时性能更好
// 6. Context 应该作为函数的第一个参数，命名为 ctx
// 7. 不要存储 Context 在结构体中，应该显式传递
// 8. Context 的取消操作应该由创建者负责
// ============================================

package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

// ============================================
// 1. Mutex（互斥锁）
// ============================================
//
// 用于保护临界区，同一时间只有一个 goroutine 可以访问

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func demonstrateMutex() {
	fmt.Println("=== Mutex ===")
	
	var counter Counter
	var wg sync.WaitGroup
	
	// 启动 1000 个 goroutine 同时增加计数器
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}
	
	wg.Wait()
	fmt.Printf("最终计数: %d\n", counter.Get())  // 应该是 1000
}

// ============================================
// 2. RWMutex（读写锁）
// ============================================
//
// 读操作可以并发，写操作独占
// 适用于读多写少的场景

type Cache struct {
	mu    sync.RWMutex
	data  map[string]string
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	return val, ok
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func demonstrateRWMutex() {
	fmt.Println("\n=== RWMutex ===")
	
	cache := NewCache()
	var wg sync.WaitGroup
	
	// 写入
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			cache.Set(fmt.Sprintf("key%d", n), fmt.Sprintf("value%d", n))
		}(i)
	}
	
	// 读取（可以并发）
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n%10)
			if val, ok := cache.Get(key); ok {
				_ = val
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Println("Cache 操作完成")
}

// ============================================
// 3. WaitGroup（等待组）
// ============================================
//
// 等待一组 goroutine 完成

func demonstrateWaitGroup() {
	fmt.Println("\n=== WaitGroup ===")
	
	var wg sync.WaitGroup
	
	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://github.com",
	}
	
	for _, url := range urls {
		wg.Add(1)  // 增加计数器
		
		go func(u string) {
			defer wg.Done()  // 完成时减少计数器
			
			// 模拟 HTTP 请求
			resp, err := http.Get(u)
			if err != nil {
				fmt.Printf("Error fetching %s: %v\n", u, err)
				return
			}
			defer resp.Body.Close()
			
			fmt.Printf("Fetched %s: %s\n", u, resp.Status)
		}(url)
	}
	
	wg.Wait()  // 等待所有 goroutine 完成
	fmt.Println("所有请求完成")
}

// WaitGroup 常见错误：复制
func wrongWaitGroup() {
	var wg sync.WaitGroup
	
	// 错误：传递 WaitGroup 的副本
	// go func(wg sync.WaitGroup) {  // 不要这样做！
	//     defer wg.Done()
	// }(wg)
	
	// 正确：传递指针
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
	}(&wg)
}

// ============================================
// 4. Once（一次性执行）
// ============================================
//
// 保证函数只执行一次，常用于单例模式

type Singleton struct {
	data string
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("创建单例实例")
		instance = &Singleton{data: "singleton data"}
	})
	return instance
}

func demonstrateOnce() {
	fmt.Println("\n=== Once ===")
	
	var wg sync.WaitGroup
	
	// 并发获取实例
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			instance := GetInstance()
			fmt.Printf("Goroutine %d: %p\n", n, instance)
		}(i)
	}
	
	wg.Wait()
}

// ============================================
// 5. Pool（对象池）
// ============================================
//
// 用于复用临时对象，减少 GC 压力
// 适用于频繁分配和回收的对象

func demonstratePool() {
	fmt.Println("\n=== Pool ===")
	
	var bufferPool = sync.Pool{
		New: func() interface{} {
			fmt.Println("创建新 buffer")
			return make([]byte, 1024)
		},
	}
	
	// 获取对象
	buf := bufferPool.Get().([]byte)
	fmt.Printf("获取 buffer，长度: %d\n", len(buf))
	
	// 使用 buffer...
	copy(buf, "hello world")
	
	// 放回池中复用
	bufferPool.Put(buf)
	
	// 再次获取（可能是同一个对象）
	buf2 := bufferPool.Get().([]byte)
	fmt.Printf("再次获取 buffer，内容: %s\n", string(buf2))
	
	bufferPool.Put(buf2)
}

// ============================================
// 6. Map（并发安全 Map）
// ============================================
//
// 内置的 map 不是并发安全的
// sync.Map 适用于以下场景：
// 1. 只写入一次但读取多次（如缓存）
// 2. 多个 goroutine 读写不同的 key
// 3. 读取、写入、删除次数差不多

func demonstrateSyncMap() {
	fmt.Println("\n=== SyncMap ===")
	
	var m sync.Map
	var wg sync.WaitGroup
	
	// 写入
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			m.Store(fmt.Sprintf("key%d", n), n)
		}(i)
	}
	
	// 读取
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if val, ok := m.Load(fmt.Sprintf("key%d", n)); ok {
				fmt.Printf("读取 key%d: %v\n", n, val)
			}
		}(i)
	}
	
	wg.Wait()
	
	// 遍历
	fmt.Println("遍历 SyncMap:")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("  %s: %v\n", key, value)
		return true  // 继续遍历
	})
}

// ============================================
// 7. Atomic（原子操作）
// ============================================
//
// 比 Mutex 更轻量的同步原语
// 适用于简单的计数、标志位等

func demonstrateAtomic() {
	fmt.Println("\n=== Atomic ===")
	
	var counter int64 = 0
	var flag int32 = 0
	var wg sync.WaitGroup
	
	// 原子增加
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	
	wg.Wait()
	fmt.Printf("原子计数结果: %d\n", atomic.LoadInt64(&counter))
	
	// CAS 操作（Compare And Swap）
	if atomic.CompareAndSwapInt32(&flag, 0, 1) {
		fmt.Println("CAS 成功，flag 从 0 变为 1")
	}
	
	if !atomic.CompareAndSwapInt32(&flag, 0, 2) {
		fmt.Println("CAS 失败，flag 已经不是 0")
	}
}

// ============================================
// 8. Context（上下文）⭐
// ============================================
//
// 用于传递取消信号、超时、截止时间、键值对
// 是 Go 并发编程中控制生命周期的标准方式

// 8.1 取消信号
func demonstrateContextCancel() {
	fmt.Println("\n=== Context Cancel ===")
	
	ctx, cancel := context.WithCancel(context.Background())
	
	// 启动工作 goroutine
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Worker: 收到取消信号，退出")
				return
			default:
				fmt.Println("Worker: 工作中...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}(ctx)
	
	time.Sleep(1 * time.Second)
	fmt.Println("主线程：发送取消信号")
	cancel()  // 发送取消信号
	
	time.Sleep(200 * time.Millisecond)
}

// 8.2 超时控制
func demonstrateContextTimeout() {
	fmt.Println("\n=== Context Timeout ===")
	
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("操作完成")
	case <-ctx.Done():
		fmt.Println("操作超时:", ctx.Err())  // context deadline exceeded
	}
}

// 8.3 截止时间
func demonstrateContextDeadline() {
	fmt.Println("\n=== Context Deadline ===")
	
	deadline := time.Now().Add(500 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	
	if d, ok := ctx.Deadline(); ok {
		fmt.Printf("截止时间: %v\n", d)
	}
	
	<-ctx.Done()
	fmt.Println("到达截止时间:", ctx.Err())
}

// 8.4 传递值（不用于传递业务参数，只用于元数据）
func demonstrateContextValue() {
	fmt.Println("\n=== Context Value ===")
	
	type contextKey string
	const requestIDKey contextKey = "requestID"
	const userKey contextKey = "user"
	
	ctx := context.Background()
	ctx = context.WithValue(ctx, requestIDKey, "req-12345")
	ctx = context.WithValue(ctx, userKey, "alice")
	
	// 读取值
	if reqID, ok := ctx.Value(requestIDKey).(string); ok {
		fmt.Printf("Request ID: %s\n", reqID)
	}
	
	if user, ok := ctx.Value(userKey).(string); ok {
		fmt.Printf("User: %s\n", user)
	}
}

// 8.5 实际应用：HTTP 请求控制
func fetchData(ctx context.Context, url string) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}
	
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	fmt.Printf("Fetched %s: %s\n", url, resp.Status)
	return nil
}

func demonstrateContextHTTP() {
	fmt.Println("\n=== Context HTTP ===")
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	// 模拟请求
	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://github.com",
	}
	
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			if err := fetchData(ctx, u); err != nil {
				fmt.Printf("Fetch %s error: %v\n", u, err)
			}
		}(url)
	}
	
	wg.Wait()
}

// ============================================
// 9. 综合示例：并发安全的任务队列
// ============================================

type TaskQueue struct {
	mu     sync.Mutex
	tasks  []func()
	closed bool
	done   chan struct{}
}

func NewTaskQueue() *TaskQueue {
	return &TaskQueue{
		tasks: make([]func(), 0),
		done:  make(chan struct{}),
	}
}

func (q *TaskQueue) Submit(task func()) bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	
	if q.closed {
		return false
	}
	
	q.tasks = append(q.tasks, task)
	return true
}

func (q *TaskQueue) Run(workerCount int) {
	var wg sync.WaitGroup
	taskCh := make(chan func())
	
	// 启动 workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for task := range taskCh {
				fmt.Printf("Worker %d 执行任务\n", id)
				task()
			}
		}(i)
	}
	
	// 分发任务
	q.mu.Lock()
	tasks := q.tasks
	q.tasks = nil
	q.mu.Unlock()
	
	for _, task := range tasks {
		taskCh <- task
	}
	close(taskCh)
	
	wg.Wait()
	close(q.done)
}

func (q *TaskQueue) Wait() {
	<-q.done
}

func demonstrateTaskQueue() {
	fmt.Println("\n=== Task Queue ===")
	
	queue := NewTaskQueue()
	
	// 提交任务
	for i := 0; i < 5; i++ {
		n := i
		queue.Submit(func() {
			fmt.Printf("执行任务 %d\n", n)
			time.Sleep(100 * time.Millisecond)
		})
	}
	
	// 运行任务
	queue.Run(3)
	queue.Wait()
	fmt.Println("所有任务完成")
}

// ============================================
// 主函数
// ============================================

func main() {
	demonstrateMutex()
	demonstrateRWMutex()
	demonstrateWaitGroup()
	demonstrateOnce()
	demonstratePool()
	demonstrateSyncMap()
	demonstrateAtomic()
	demonstrateContextCancel()
	demonstrateContextTimeout()
	demonstrateContextDeadline()
	demonstrateContextValue()
	demonstrateContextHTTP()
	demonstrateTaskQueue()
	
	// ============================================
	// 练习题
	// ============================================
	//
	// 练习 1：实现一个并发安全的环形缓冲区
	//   type RingBuffer struct { ... }
	//   - 使用 Mutex 保护
	//   - 实现 Write(data []byte) (n int, err error)
	//   - 实现 Read(p []byte) (n int, err error)
	//   - 当缓冲区满时，Write 阻塞；空时，Read 阻塞
	//
	// 练习 2：实现一个 Semaphore（信号量）
	//   type Semaphore struct { ... }
	//   - 使用 Channel 实现
	//   - Acquire() 获取许可，如果没有则阻塞
	//   - Release() 释放许可
	//   - TryAcquire(timeout time.Duration) bool 带超时的获取
	//
	// 练习 3：实现一个读写分离的缓存
	//   type RWCache struct { ... }
	//   - 使用 RWMutex
	//   - 支持 Set、Get、Delete
	//   - 支持 TTL（过期时间），使用 goroutine 定期清理
	//
	// 练习 4：实现一个带权重的负载均衡器
	//   type LoadBalancer struct { ... }
	//   - 后端服务器有权重
	//   - 使用 atomic 实现无锁的轮询
	//   - 支持动态添加/移除后端
	//
	// 练习 5：实现一个断路器（Circuit Breaker）
	//   type CircuitBreaker struct { ... }
	//   - 状态：Closed、Open、Half-Open
	//   - 失败次数超过阈值进入 Open
	//   - Open 状态经过超时后进入 Half-Open
	//   - Half-Open 成功则 Closed，失败则 Open
	//   - 使用 sync/atomic 或 Mutex 保证并发安全
	//
	// 练习 6：实现一个分布式锁（使用文件或 Redis）
	//   type DistributedLock struct { ... }
	//   - Lock() 获取锁，阻塞直到成功
	//   - TryLock(timeout time.Duration) bool 带超时
	//   - Unlock() 释放锁
	//   - 使用 Context 支持取消
	//
	// 练习 7：实现一个限流器（Rate Limiter）
	//   type RateLimiter struct { ... }
	//   - 使用令牌桶算法
	//   - Allow() bool 判断是否允许通过
	//   - Wait(ctx context.Context) error 等待直到允许通过
}
