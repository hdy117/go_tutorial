// ============================================
// Go 泛型教程
// ============================================
//
// 本文件涵盖 Go 1.18+ 泛型编程：
// - 泛型函数
// - 泛型类型
// - 类型参数
// - 类型约束（Type Constraints）⭐
// - 类型集（Type Sets）
// - 泛型接口
// - 类型推导
//
// 最佳实践：
// 1. 只在真正需要时使用泛型，不要为了用而用
// 2. 约束应该尽可能小（小接口原则）
// 3. 优先考虑标准库中的约束（constraints 包）
// 4. 类型参数命名应简洁（T, K, V, E 等）
// 5. 泛型会增加编译时间和二进制大小，谨慎使用
// ============================================

package main

import (
	"cmp"
	"fmt"
	"golang.org/x/exp/constraints"
)

// ============================================
// 1. 泛型函数
// ============================================
//
// 语法：func 函数名[T 约束](参数 T) 返回值

// 最简单的泛型函数 - 返回零值
func zeroValue[T any]() T {
	var zero T
	return zero
}

// 泛型 Print
func Print[T any](v T) {
	fmt.Printf("值: %v, 类型: %T\n", v, v)
}

// 泛型 Map 函数
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// 泛型 Filter 函数
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// 泛型 Reduce 函数
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

func demonstrateGenericFunctions() {
	fmt.Println("=== 泛型函数 ===")
	
	// 返回零值
	fmt.Printf("int 零值: %d\n", zeroValue[int]())
	fmt.Printf("string 零值: %q\n", zeroValue[string]())
	fmt.Printf("bool 零值: %v\n", zeroValue[bool]())
	
	// Print
	Print(42)
	Print("hello")
	Print(3.14)
	
	// Map
	nums := []int{1, 2, 3, 4, 5}
	doubles := Map(nums, func(n int) int { return n * 2 })
	fmt.Printf(" doubles: %v\n", doubles)
	
	strings := Map(nums, func(n int) string { return fmt.Sprintf("num%d", n) })
	fmt.Printf(" strings: %v\n", strings)
	
	// Filter
	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Printf(" evens: %v\n", evens)
	
	// Reduce
	sum := Reduce(nums, 0, func(acc, n int) int { return acc + n })
	fmt.Printf(" sum: %d\n", sum)
	
	product := Reduce(nums, 1, func(acc, n int) int { return acc * n })
	fmt.Printf(" product: %d\n", product)
}

// ============================================
// 2. 类型约束（Type Constraints）
// ============================================
//
// 约束定义了类型参数必须满足的条件

// 数值类型约束
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// 求和（约束为整数或浮点数）
func Sum[T constraints.Integer | constraints.Float](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// 使用 cmp.Ordered（Go 1.21+）
func Compare[T cmp.Ordered](a, b T) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func demonstrateConstraints() {
	fmt.Println("\n=== 类型约束 ===")
	
	// Max/Min 可以用于任何可比较类型
	fmt.Printf("Max(3, 5) = %d\n", Max(3, 5))
	fmt.Printf("Max(3.14, 2.71) = %f\n", Max(3.14, 2.71))
	fmt.Printf("Max(\"apple\", \"banana\") = %s\n", Max("apple", "banana"))
	
	// Sum 只能用于数值类型
	ints := []int{1, 2, 3, 4, 5}
	floats := []float64{1.1, 2.2, 3.3}
	
	fmt.Printf("Sum(ints) = %d\n", Sum(ints))
	fmt.Printf("Sum(floats) = %f\n", Sum(floats))
	
	// Compare
	fmt.Printf("Compare(3, 5) = %d\n", Compare(3, 5))
	fmt.Printf("Compare(\"a\", \"a\") = %d\n", Compare("a", "a"))
}

// ============================================
// 3. 自定义约束
// ============================================

// 定义整数约束
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// 定义可以相加的类型
type Addable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// 使用 ~ 表示底层类型也满足约束
type MyInt int  // MyInt 的底层类型是 int

func Add[T Addable](a, b T) T {
	return a + b
}

// 带有方法的约束
type Stringer interface {
	String() string
}

func ToStrings[T Stringer](items []T) []string {
	result := make([]string, len(items))
	for i, item := range items {
		result[i] = item.String()
	}
	return result
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

func demonstrateCustomConstraints() {
	fmt.Println("\n=== 自定义约束 ===")
	
	// Add 可以用于所有 Addable 类型
	fmt.Printf("Add(1, 2) = %d\n", Add(1, 2))
	fmt.Printf("Add(1.5, 2.5) = %f\n", Add(1.5, 2.5))
	fmt.Printf("Add(\"Hello, \", \"World\") = %s\n", Add("Hello, ", "World"))
	
	// MyInt 也满足约束（因为有 ~）
	var a MyInt = 10
	var b MyInt = 20
	fmt.Printf("Add(MyInt) = %d\n", Add(a, b))
	
	// ToStrings
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
	}
	strings := ToStrings(people)
	fmt.Printf("ToStrings: %v\n", strings)
}

// ============================================
// 4. 泛型类型
// ============================================
//
// 类型也可以是泛型的

// 泛型栈
type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

// 泛型队列
type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if len(q.items) == 0 {
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// 泛型集合（基于 map）
type Set[T comparable] struct {
	items map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]struct{})}
}

func (s *Set[T]) Add(item T) {
	s.items[item] = struct{}{}
}

func (s *Set[T]) Remove(item T) {
	delete(s.items, item)
}

func (s *Set[T]) Contains(item T) bool {
	_, ok := s.items[item]
	return ok
}

func (s *Set[T]) Size() int {
	return len(s.items)
}

func (s *Set[T]) ToSlice() []T {
	result := make([]T, 0, len(s.items))
	for item := range s.items {
		result = append(result, item)
	}
	return result
}

// 泛型链表节点
type ListNode[T any] struct {
	Value T
	Next  *ListNode[T]
}

type LinkedList[T any] struct {
	head *ListNode[T]
	size int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Append(value T) {
	newNode := &ListNode[T]{Value: value}
	
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	l.size++
}

func demonstrateGenericTypes() {
	fmt.Println("\n=== 泛型类型 ===")
	
	// Stack
	intStack := NewStack[int]()
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	
	if val, ok := intStack.Pop(); ok {
		fmt.Printf("Pop: %d\n", val)
	}
	fmt.Printf("Stack size: %d\n", intStack.Size())
	
	// 字符串栈
	strStack := NewStack[string]()
	strStack.Push("hello")
	strStack.Push("world")
	
	// Queue
	queue := NewQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	
	if val, ok := queue.Dequeue(); ok {
		fmt.Printf("Dequeue: %d\n", val)
	}
	
	// Set
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(2)  // 重复
	
	fmt.Printf("Set size: %d\n", set.Size())
	fmt.Printf("Contains 2: %v\n", set.Contains(2))
	fmt.Printf("Contains 5: %v\n", set.Contains(5))
	
	// LinkedList
	list := NewLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	fmt.Printf("LinkedList size: %d\n", list.size)
}

// ============================================
// 5. 泛型接口
// ============================================

// 可比较接口（Go 1.20+）
type Comparable[T any] interface {
	Compare(other T) int  // -1: less, 0: equal, 1: greater
}

// 泛型排序接口
type Sorter[T any] interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// 泛型二分查找（要求约束为 Ordered）
func BinarySearch[T constraints.Ordered](slice []T, target T) (int, bool) {
	left, right := 0, len(slice)-1
	
	for left <= right {
		mid := left + (right-left)/2
		if slice[mid] == target {
			return mid, true
		}
		if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1, false
}

// 泛型树节点
type TreeNode[T any] struct {
	Value    T
	Children []*TreeNode[T]
}

func NewTreeNode[T any](value T) *TreeNode[T] {
	return &TreeNode[T]{Value: value, Children: make([]*TreeNode[T], 0)}
}

func (n *TreeNode[T]) AddChild(child *TreeNode[T]) {
	n.Children = append(n.Children, child)
}

// 深度优先遍历
func (n *TreeNode[T]) DFS(visit func(T)) {
	if n == nil {
		return
	}
	visit(n.Value)
	for _, child := range n.Children {
		child.DFS(visit)
	}
}

func demonstrateGenericInterfaces() {
	fmt.Println("\n=== 泛型接口 ===")
	
	// BinarySearch
	nums := []int{1, 3, 5, 7, 9, 11, 13}
	if idx, found := BinarySearch(nums, 7); found {
		fmt.Printf("Found 7 at index %d\n", idx)
	}
	if _, found := BinarySearch(nums, 6); !found {
		fmt.Println("6 not found")
	}
	
	// Tree
	root := NewTreeNode("root")
	child1 := NewTreeNode("child1")
	child2 := NewTreeNode("child2")
	grandchild := NewTreeNode("grandchild")
	
	root.AddChild(child1)
	root.AddChild(child2)
	child1.AddChild(grandchild)
	
	fmt.Println("DFS traversal:")
	root.DFS(func(value string) {
		fmt.Printf("  %s\n", value)
	})
}

// ============================================
// 6. 类型推导
// ============================================

func demonstrateTypeInference() {
	fmt.Println("\n=== 类型推导 ===")
	
	// 显式指定类型参数
	Print[int](42)
	
	// 编译器推导类型参数
	Print(42)        // 推导为 Print[int]
	Print("hello")   // 推导为 Print[string]
	
	// 从参数推导
	x := Max(3, 5)   // 推导 T 为 int
	y := Max(1.5, 2.5)  // 推导 T 为 float64
	fmt.Printf("x=%v, y=%v\n", x, y)
	
	// 泛型类型推导
	stack := NewStack[int]()  // 必须显式指定，无法推导
	stack.Push(1)
	
	// 从字面量推导
	nums := []int{1, 2, 3}
	doubled := Map(nums, func(n int) int { return n * 2 })
	// 编译器从 nums 推导 T 为 int，从返回值推导 U 为 int
	fmt.Printf("doubled: %v\n", doubled)
}

// ============================================
// 7. 实用泛型模式
// ============================================

// 可选值类型（类似 Rust 的 Option）
type Option[T any] struct {
	value   T
	present bool
}

func Some[T any](v T) Option[T] {
	return Option[T]{value: v, present: true}
}

func None[T any]() Option[T] {
	return Option[T]{present: false}
}

func (o Option[T]) IsSome() bool {
	return o.present
}

func (o Option[T]) IsNone() bool {
	return !o.present
}

func (o Option[T]) Unwrap() T {
	if !o.present {
		panic("called Unwrap on None")
	}
	return o.value
}

func (o Option[T]) UnwrapOr(defaultValue T) T {
	if o.present {
		return o.value
	}
	return defaultValue
}

// 结果类型（类似 Rust 的 Result）
type Result[T any] struct {
	value T
	err   error
}

func Ok[T any](v T) Result[T] {
	return Result[T]{value: v, err: nil}
}

func Err[T any](e error) Result[T] {
	var zero T
	return Result[T]{value: zero, err: e}
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

func (r Result[T]) IsErr() bool {
	return r.err != nil
}

func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic(r.err)
	}
	return r.value
}

func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.err == nil {
		return r.value
	}
	return defaultValue
}

func (r Result[T]) Error() error {
	return r.err
}

// Pair 类型
type Pair[A, B any] struct {
	First  A
	Second B
}

func NewPair[A, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{First: a, Second: b}
}

func demonstrateUtilityPatterns() {
	fmt.Println("\n=== 实用泛型模式 ===")
	
	// Option
	maybeValue := Some(42)
	if maybeValue.IsSome() {
		fmt.Printf("Value: %d\n", maybeValue.Unwrap())
	}
	
	noValue := None[int]()
	fmt.Printf("Or default: %d\n", noValue.UnwrapOr(0))
	
	// Result
	success := Ok(42)
	failure := Err[int](fmt.Errorf("something went wrong"))
	
	if success.IsOk() {
		fmt.Printf("Success: %d\n", success.Unwrap())
	}
	
	if failure.IsErr() {
		fmt.Printf("Error: %v\n", failure.Error())
	}
	fmt.Printf("Or default: %d\n", failure.UnwrapOr(0))
	
	// Pair
	pair := NewPair("answer", 42)
	fmt.Printf("Pair: (%v, %v)\n", pair.First, pair.Second)
}

// ============================================
// 主函数
// ============================================

func main() {
	demonstrateGenericFunctions()
	demonstrateConstraints()
	demonstrateCustomConstraints()
	demonstrateGenericTypes()
	demonstrateGenericInterfaces()
	demonstrateTypeInference()
	demonstrateUtilityPatterns()
	
	// ============================================
	// 练习题
	// ============================================
	//
	// 练习 1：实现泛型的 Map、Filter、Reduce
	//   - Map 将 []T 转换为 []U
	//   - Filter 根据条件过滤元素
	//   - Reduce 将切片归约为单个值
	//   - 编写测试验证功能
	//
	// 练习 2：实现泛型的缓存
	//   type Cache[K comparable, V any] struct { ... }
	//   - Set(key K, value V, ttl time.Duration)
	//   - Get(key K) (V, bool)
	//   - Delete(key K)
	//   - 支持 TTL 自动过期
	//
	// 练习 3：实现泛型的 Channel 操作函数
	//   func MapChan[T, U any](input <-chan T, fn func(T) U) <-chan U
	//   func FilterChan[T any](input <-chan T, predicate func(T) bool) <-chan T
	//   func ReduceChan[T, U any](input <-chan T, initial U, fn func(U, T) U) U
	//
	// 练习 4：实现泛型的排序算法
	//   func QuickSort[T constraints.Ordered](slice []T)
	//   func MergeSort[T constraints.Ordered](slice []T)
	//   - 支持任意可排序类型
	//   - 与 sort.Slice 性能对比
	//
	// 练习 5：实现泛型的函数组合
	//   func Compose[A, B, C any](f func(B) C, g func(A) B) func(A) C
	//   func Pipe[A, B, C any](f func(A) B, g func(B) C) func(A) C
	//   func Curry[A, B, C any](f func(A, B) C) func(A) func(B) C
	//   - 验证函数组合的正确性
	//
	// 练习 6：实现泛型的状态机
	//   type StateMachine[S comparable, E any] struct { ... }
	//   - AddTransition(from S, event E, to S)
	//   - Trigger(event E) error
	//   - 支持状态转换验证
	//
	// 练习 7：实现泛型的依赖注入容器
	//   type Container struct { ... }
	//   - Register[T any](constructor func(...) T)
	//   - Resolve[T any]() (T, error)
	//   - 支持单例和瞬态生命周期
}
