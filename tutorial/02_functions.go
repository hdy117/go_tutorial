// ============================================
// Go 函数特性教程
// ============================================
//
// 本文件涵盖 Go 语言函数的核心特性：
// - 函数定义与调用
// - 多返回值 ⭐ Go 特色
// - 命名返回值
// - 变长参数
// - 函数作为值和类型
// - 闭包 (Closure)
// - defer 延迟执行 ⭐重要
// - 递归
// - init 函数
//
// 最佳实践：
// 1. 函数应当短小，只做一件事
// 2. 多返回值时，error 通常作为最后一个返回值
// 3. 使用命名返回值提高可读性，但要避免滥用
// 4. defer 常用于资源清理（关闭文件、解锁等）
// 5. 避免在热路径（hot path）中使用 defer（Go 1.14 后性能已改善）
// ============================================

package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// ============================================
// 1. 基本函数定义
// ============================================
// 格式: func 函数名(参数列表) (返回值列表) { 函数体 }

// 无参数无返回值
func sayHello() {
	fmt.Println("Hello, Go!")
}

// 有参数
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func greeting(name *string) {
	fmt.Println("hello ", *name)
}

// 多参数（同类型可省略类型声明）
func add(a, b int) int {
	return a + b
}

// ============================================
// 2. 多返回值 ⭐ Go 的重要特性
// ============================================
//
// 常见模式：(result, error)
// 可以返回任意数量的值

// 返回商和余数
func divide(dividend, divisor int) (quotient, remainder int, err error) {
	if divisor == 0 {
		return 0, 0, errors.New("除数不能为0")
	}
	quotient = dividend / divisor
	remainder = dividend % divisor
	return quotient, remainder, nil // 命名返回值可以直接使用
}

// 错误处理模式
func findUser(id int) (string, error) {
	if id <= 0 {
		return "", fmt.Errorf("无效的用户ID: %d", id)
	}
	// 模拟查找
	return fmt.Sprintf("User%d", id), nil
}

// ============================================
// 3. 命名返回值
// ============================================
//
// 优点：
//   - 代码自文档化
//   - 可以直接 return（裸返回）
// 缺点：
//   - 可能导致可读性下降（特别是长函数）
// 建议：只在简单函数中使用

// 计算矩形信息
func rectangle(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // 裸返回，自动返回命名变量
}

// ============================================
// 4. 变长参数
// ============================================
//
// ...T 表示接受任意数量的 T 类型参数
// 在函数内部作为切片处理

// 求和
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// 变长参数 + 普通参数
func printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// 展开切片传入
func useVariadic() {
	nums := []int{1, 2, 3, 4, 5}
	result := sum(nums...) // 展开操作符 ...
	fmt.Printf("sum=%d\n", result)
}

// ============================================
// 5. 函数作为值和类型
// ============================================
//
// 函数是一等公民，可以：
// - 赋值给变量
// - 作为参数传递
// - 作为返回值
// - 存储在数据结构中

// 函数类型
type Calculator func(int, int) int

// 接收函数作为参数
func operate(a, b int, op Calculator) int {
	return op(a, b)
}

// 返回函数
func makeMultiplier(factor int) Calculator {
	return func(x, y int) int {
		return (x + y) * factor
	}
}

func demonstrateFuncValue() {
	// 匿名函数赋值给变量
	multiply := func(a, b int) int {
		return a * b
	}

	// 作为参数传递
	result := operate(3, 4, multiply)
	fmt.Printf("3 * 4 = %d\n", result)

	// 直接传递匿名函数
	result = operate(10, 5, func(a, b int) int {
		return a - b
	})
	fmt.Printf("10 - 5 = %d\n", result)

	// 使用返回的函数
	double := makeMultiplier(2)
	fmt.Printf("double: (3+4)*2 = %d\n", double(3, 4))
}

// ============================================
// 6. 闭包 (Closure)
// ============================================
//
// 闭包是引用了外部变量的函数
// 闭包持有外部变量的引用，而不是值的拷贝

// 计数器工厂
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 带初始值的计数器
func makeCounterFrom(start int) func() int {
	return func() int {
		start++
		return start
	}
}

// 记忆化（Memoization）
func makeFibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func demonstrateClosure() {
	// 创建两个独立的计数器
	counter1 := makeCounter()
	counter2 := makeCounter()

	fmt.Println("=================enclosure==================")
	fmt.Println("Counter1:", counter1()) // 1
	fmt.Println("Counter1:", counter1()) // 2
	fmt.Println("Counter2:", counter2()) // 1
	fmt.Println("Counter1:", counter1()) // 3

	// Fibonacci 生成器
	fmt.Println("\nFibonacci:")
	fib := makeFibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fib())
	}
	fmt.Println()
}

// ============================================
// 7. defer 延迟执行 ⭐非常重要
// ============================================
//
// defer 语句推迟函数的执行直到上层函数返回
// 多个 defer 以 LIFO（后进先出）顺序执行
//
// 常见用途：
// - 资源清理（关闭文件、数据库连接）
// - 解锁互斥锁
// - 记录函数执行时间

func demonstrateDefer() {
	fmt.Println("函数开始")

	// defer 在函数返回前执行
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	fmt.Println("函数结束")
	// 输出顺序: defer 3, defer 2, defer 1
}

// defer 实际应用：文件处理
func readFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close() // 确保文件关闭，即使在错误路径上

	// 处理文件...
	buf := make([]byte, 100)
	n, err := file.Read(buf)
	if err != nil {
		return err
	}

	fmt.Printf("读取了 %d 字节\n", n)
	return nil
}

// defer 与返回值（重要！）
func deferAndReturn() (result int) {
	defer func() {
		result++ // 修改命名返回值
	}()
	return 10 // 实际返回 11
}

// 计算函数执行时间
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s 耗时: %s\n", name, elapsed)
}

func slowFunction() {
	defer timeTrack(time.Now(), "slowFunction")

	// 模拟耗时操作
	time.Sleep(100 * time.Millisecond)
	fmt.Println("slowFunction 执行完成")
}

// defer 中的参数求值
func demonstrateDeferArgs() {
	i := 0
	defer fmt.Printf("defer i=%d\n", i) // 参数立即求值，输出 0
	i++
	fmt.Printf("函数内 i=%d\n", i) // 输出 1
}

// ============================================
// 8. 递归
// ============================================

// 阶乘
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 尾递归优化版本（Go 不优化，但可改写为迭代）
func factorialIter(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// 斐波那契（带缓存）
var fibCache = make(map[int]int)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	if val, ok := fibCache[n]; ok {
		return val
	}
	fibCache[n] = fibonacci(n-1) + fibonacci(n-2)
	return fibCache[n]
}

// ============================================
// 9. init 函数
// ============================================
//
// 每个文件可以包含多个 init 函数
// 在包被导入时自动执行，按声明顺序执行
// 用于初始化包级变量、注册驱动等

var packageVar string

func init() {
	packageVar = "initialized in init 1"
	fmt.Println("init 1 执行")
}

func init() {
	fmt.Println("init 2 执行")
}

func Separator() {
	fmt.Println("=================================")
}

// ============================================
// 主函数
// ============================================

func main() {
	fmt.Println("=== 基本函数 ===")
	sayHello()
	var userName string = "Go开发者"
	greet(userName)
	greeting(&userName)
	fmt.Printf("3 + 5 = %d\n", add(3, 5))

	fmt.Println("\n=== 多返回值 ===")
	q, r, err := divide(17, 5)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("17 / 5 = %d 余 %d\n", q, r)
	}

	_, _, err = divide(10, 0)
	if err != nil {
		fmt.Println("除以0错误:", err)
	}

	fmt.Println("\n=== 函数作为值 ===")
	demonstrateFuncValue()

	fmt.Println("\n=== 闭包 ===")
	demonstrateClosure()

	fmt.Println("\n=== defer ===")
	demonstrateDefer()

	fmt.Println("\ndefer 和返回值:", deferAndReturn())

	fmt.Println("\n=== 递归 ===")
	fmt.Printf("5! = %d\n", factorial(5))
	fmt.Printf("fib(10) = %d\n", fibonacci(10))

	// ============================================
	// 练习题
	// ============================================
	//
	// 练习 1：实现一个函数，接收任意数量的整数，返回它们的最大值和最小值
	//   func minMax(nums ...int) (min, max int, err error)
	//   错误处理：如果没有传入参数，返回错误
	Separator()
	findMinMax := func(nums ...int) (min, max int, err error) {
		if len(nums) == 0 {
			return 0, 0, fmt.Errorf("should at least pass in one element")
		}
		min = nums[0]
		max = nums[0]
		for _, v := range nums[1:] {
			if v > max {
				max = v
			}
			if v < min {
				min = v
			}
		}
		return min, max, nil
	}
	min, max, err := findMinMax(1, 2, 3, 4, 5, 6, 8)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("min:", min, ", max:", max)
	}
	//
	// 练习 2：使用闭包实现一个累加器，支持加法和减法操作
	//   func makeAccumulator(initial int) (add, sub func(int) int)
	//   add(5) 表示加5，sub(3) 表示减3
	Separator()
	type Accumulator = func(int) int
	accumulatorMaker := func(initVal int) (addFunc, subFunc Accumulator) {
		addVal := initVal
		addFunc = func(accumulation int) int {
			addVal += accumulation
			return addVal
		}
		subVal := initVal
		subFunc = func(accumulation int) int {
			subVal -= accumulation
			return subVal
		}
		return addFunc, subFunc
	}
	addFunc, subFunc := accumulatorMaker(5)
	fmt.Println("addFunc(3)", addFunc(3))
	fmt.Println("subFunc(3)", subFunc(3))

	// 练习 3：实现一个函数，接收一个整数切片和一个过滤函数，返回满足条件的元素
	//   func filter(nums []int, predicate func(int) bool) []int
	filterFunc := func(val int) bool {
		if val == 0 {
			return true
		}
		return false
	}
	filter := func(arr []int, filterFunc func(int) bool) []int {
		outArr := make([]int, 0, len(arr))
		for _, v := range arr {
			if !filterFunc(v) {
				outArr = append(outArr, v)
			}
		}
		return outArr
	}
	arr := []int{0, 2, 3, 4, 0, 1, 3, 2, 0}
	outArr := filter(arr, filterFunc)
	fmt.Println("outArr:", outArr)

	// 练习 4：使用 defer 实现一个函数计时器，能够计算并打印函数执行时间
	//   提示：使用 time.Since
	Separator()
	slowFunc := func() {
		startTime := time.Now()
		defer func() {
			fmt.Println("func elapsed time:", time.Since(startTime))
		}()

		time.Sleep(2 * time.Second)
	}
	slowFunc()

	// 练习 5：实现一个记忆化函数，缓存任意函数的结果（进阶）
	//   func memoize(f func(int) int) func(int) int
	type OneFunc func(int) int
	memFunc := func(oneFunc OneFunc) OneFunc {
		cacheMap := map[int]int{}

		return func(key int) int {
			if v, ok := cacheMap[key]; ok {
				fmt.Println("cache hit for val:", key)
				return v
			}

			out := oneFunc(key)
			cacheMap[key] = out
			fmt.Println("no cache hit for val:", key)

			return out
		}
	}
	getcacheMapFunc := memFunc(func(val int) int { return val * val })

	Separator()
	fmt.Println("getcacheMapFunc(1):", getcacheMapFunc(1))
	fmt.Println("getcacheMapFunc(1):", getcacheMapFunc(1))

	// 练习 6：实现一个管道（pipeline）函数链
	//   func pipeline(data int, funcs ...func(int) int) int
	//   示例：pipeline(5, double, addOne, square) = ((5*2)+1)^2 = 121
	pipeline := func(val int, funcs ...OneFunc) int {
		for _, oneFunc := range funcs {
			val = oneFunc(val)
		}
		return val
	}

	Separator()
	res := pipeline(5,
		func(val int) int { return val * 2 },
		func(val int) int { return val + 1 },
		func(val int) int { return val * val })
	fmt.Println("pipeline:",res)
}
