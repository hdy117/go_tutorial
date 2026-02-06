// ============================================
// Go 基础语法教程
// ============================================
// 
// 本文件涵盖 Go 语言的基础语法特性：
// - 变量声明与初始化
// - 常量与 iota
// - 数据类型
// - 控制流程
// - 数组、切片、Map
//
// 最佳实践：
// 1. 优先使用短变量声明 :=
// 2. 使用有意义的变量名，Go 倾向于简短但清晰的命名
// 3. 避免使用全局变量
// 4. 错误处理优先返回 error，而非使用异常机制
// ============================================

package main

import "fmt"

func main() {
	// ============================================
	// 1. 变量声明
	// ============================================
	
	// 方式1：var 声明（显式类型）
	var name string = "Go"
	var age int = 15
	
	// 方式2：var 声明（类型推断）
	var language = "Golang"  // 编译器自动推断为 string
	
	// 方式3：短变量声明（最常用）
	// 只能在函数内部使用，自动推断类型
	year := 2009
	isAwesome := true
	
	fmt.Printf("语言: %s, 年龄: %d, 发布年份: %d\n", name, age, year)
	fmt.Printf("%s 很棒? %v\n", language, isAwesome)
	
	// 多变量声明
	var a, b, c int = 1, 2, 3
	x, y, z := "hello", 42, 3.14
	fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)
	fmt.Printf("x=%s, y=%d, z=%f\n", x, y, z)
	
	// ============================================
	// 2. 常量与 iota
	// ============================================
	// 
	// 常量使用 const 声明，编译期确定，不可修改
	// iota 是常量计数器，从 0 开始，每行递增 1
	
	const Pi = 3.14159
	const (
		Monday = iota    // 0
		Tuesday          // 1
		Wednesday        // 2
		Thursday         // 3
		Friday           // 4
		Saturday         // 5
		Sunday           // 6
	)
	
	// iota 技巧：位运算定义权限
	const (
		Read = 1 << iota   // 1 (0001)
		Write              // 2 (0010)
		Execute            // 4 (0100)
	)
	
	fmt.Printf("Monday=%d, Sunday=%d\n", Monday, Sunday)
	fmt.Printf("Read=%d, Write=%d, Execute=%d\n", Read, Write, Execute)
	
	// ============================================
	// 3. 基本数据类型
	// ============================================
	// 
	// 整数: int, int8, int16, int32, int64
	//       uint, uint8(byte), uint16, uint32, uint64
	// 浮点: float32, float64（默认）
	// 布尔: bool
	// 字符串: string（不可变）
	// 复数: complex64, complex128
	
	var intVal int = 100
	var floatVal float64 = 3.14159
	var boolVal bool = true
	var strVal string = "Hello, 世界"
	var complexVal complex128 = 1 + 2i
	
	fmt.Printf("int: %d, float: %f, bool: %v\n", intVal, floatVal, boolVal)
	fmt.Printf("string: %s, complex: %v\n", strVal, complexVal)
	
	// 零值（未初始化变量的默认值）
	var zeroInt int       // 0
	var zeroString string // "" (空字符串)
	var zeroBool bool     // false
	var zeroPtr *int      // nil
	fmt.Printf("零值: int=%d, string=%q, bool=%v, ptr=%v\n", 
		zeroInt, zeroString, zeroBool, zeroPtr)
	
	// ============================================
	// 4. 控制流程
	// ============================================
	
	// if - 不需要括号，支持初始化语句
	score := 85
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else {
		fmt.Println("继续加油")
	}
	
	// if 初始化语句（作用域仅限于 if 块）
	if v := score * 2; v > 150 {
		fmt.Printf("双倍分数 %d 超过 150\n", v)
	}
	
	// for 循环（Go 只有 for，没有 while）
	fmt.Println("\n--- for 循环 ---")
	for i := 0; i < 3; i++ {
		fmt.Printf("i=%d ", i)
	}
	fmt.Println()
	
	// 相当于 while
	count := 0
	for count < 3 {
		fmt.Printf("count=%d ", count)
		count++
	}
	fmt.Println()
	
	// 无限循环
	// for {
	//     // 无限循环，需要用 break 退出
	// }
	
	// switch - 自动 break，可用 fallthrough
	fmt.Println("\n--- switch ---")
	day := Wednesday
	switch day {
	case Monday:
		fmt.Println("周一")
	case Tuesday:
		fmt.Println("周二")
	case Wednesday, Thursday:
		fmt.Println("周中") // 多个 case
	default:
		fmt.Println("其他")
	}
	
	// switch 初始化 + 无条件（替代 if-else if）
	switch num := 10; {
	case num < 0:
		fmt.Println("负数")
	case num == 0:
		fmt.Println("零")
	case num > 0 && num < 10:
		fmt.Println("个位数")
	default:
		fmt.Println("大于等于10")
	}
	
	// ============================================
	// 5. 数组（固定长度）
	// ============================================
	// 
	// 数组是值类型，赋值会复制整个数组
	// 实际开发中更常用切片（slice）
	
	var arr1 [3]int = [3]int{1, 2, 3}
	arr2 := [5]int{1, 2, 3}           // 未初始化的为 0
	arr3 := [...]int{1, 2, 3, 4, 5}   // 自动推断长度
	
	fmt.Printf("arr1: %v, 长度: %d\n", arr1, len(arr1))
	fmt.Printf("arr2: %v\n", arr2)
	fmt.Printf("arr3: %v, 长度: %d\n", arr3, len(arr3))
	
	// 多维数组
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("matrix: %v\n", matrix)
	
	// ============================================
	// 6. 切片（Slice）- 动态数组 ⭐核心概念
	// ============================================
	// 
	// 切片是对数组的引用，包含三个部分：指针、长度、容量
	// len() - 长度（当前元素个数）
	// cap() - 容量（底层数组大小）
	
	// 方式1：从数组创建
	arr := [5]int{1, 2, 3, 4, 5}
	s1 := arr[1:4]  // [2, 3, 4]，左闭右开
	fmt.Printf("s1: %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
	
	// 方式2：make 创建
	s2 := make([]int, 3, 5)  // 长度3，容量5
	fmt.Printf("s2: %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
	
	// 方式3：字面量
	s3 := []int{10, 20, 30}
	
	// 追加元素（append 可能触发重新分配）
	s3 = append(s3, 40)
	s3 = append(s3, 50, 60)  // 追加多个
	fmt.Printf("s3 after append: %v\n", s3)
	
	// 复制切片
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Printf("dst after copy: %v\n", dst)
	
	// 切片共享底层数组（注意修改影响）
	original := []int{1, 2, 3, 4, 5}
	ref := original[1:3]  // [2, 3]
	ref[0] = 100          // 修改会影响 original
	fmt.Printf("original after modify ref: %v\n", original)
	
	// ============================================
	// 7. Map（哈希表）
	// ============================================
	// 
	// 无序的键值对集合
	// 键必须是可比较类型（不能是 slice, map, function）
	
	// 创建
	m1 := make(map[string]int)
	m1["alice"] = 25
	m1["bob"] = 30
	
	// 字面量创建
	m2 := map[string]int{
		"go":     2009,
		"python": 1991,
		"java":   1995,
	}
	
	// 取值（第二个返回值表示是否存在）
	age, exists := m1["alice"]
	if exists {
		fmt.Printf("alice's age: %d\n", age)
	}
	
	// 删除
	delete(m1, "bob")
	
	// 遍历（无序）
	fmt.Println("\n--- map 遍历 ---")
	for lang, year := range m2 {
		fmt.Printf("%s: %d\n", lang, year)
	}
	
	// ============================================
	// 8. range 遍历
	// ============================================
	
	// 遍历切片
	nums := []int{10, 20, 30}
	fmt.Println("\n--- range slice ---")
	for index, value := range nums {
		fmt.Printf("index=%d, value=%d\n", index, value)
	}
	
	// 只需要索引
	for i := range nums {
		fmt.Printf("index=%d\n", i)
	}
	
	// 只需要值（用 _ 忽略索引）
	for _, v := range nums {
		fmt.Printf("value=%d\n", v)
	}
	
	// 遍历字符串（按 rune）
	for i, r := range "Hello, 世界" {
		fmt.Printf("index=%d, rune=%c\n", i, r)
	}
	
	// ============================================
	// 练习题（请在此文件基础上完成）
	// ============================================
	//
	// 练习 1：创建一个 map 存储学生姓名和分数，实现以下功能：
	//   - 添加 3 个学生
	//   - 查询某个学生的分数
	//   - 计算平均分
	//   - 删除分数低于 60 分的学生
	//
	// 练习 2：编写函数实现切片的去重
	//   func removeDuplicates(nums []int) []int
	//
	// 练习 3：使用 iota 定义文件权限常量（类似 Linux）
	//   - Owner 可读可写可执行
	//   - Group 可读可执行
	//   - Other 只读
	//
	// 练习 4：编写函数找出切片中的最大值和最小值
	//   func findMinMax(nums []int) (min, max int)
	//
	// 练习 5：实现一个简单的猜数字游戏
	//   - 随机生成 1-100 的数字
	//   - 用户输入猜测，程序提示"太大"或"太小"
	//   - 使用循环直到猜对
}
