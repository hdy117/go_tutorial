// ============================================
// Go 反射教程
// ============================================
//
// 本文件涵盖 Go 语言反射的核心功能：
// - reflect.Type 和 reflect.Value
// - 类型检查与转换
// - 值的操作（读取、修改）
// - 结构体反射（字段、标签）
// - 方法反射与调用
// - 创建新值
// - 反射的性能考量
//
// 最佳实践：
// 1. 尽量避免使用反射，它会降低性能并丧失类型安全
// 2. 反射代码难以阅读和维护，只在必要场景使用
// 3. 必须使用时，确保有充分的测试覆盖
// 4. 反射操作需要检查合法性，避免 panic
// 5. 结构体标签解析是反射的常见用途
// ============================================

package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// ============================================
// 1. 基础反射操作
// ============================================

func demonstrateBasicReflection() {
	fmt.Println("=== 基础反射 ===")
	
	x := 42
	
	// 获取 Type
	t := reflect.TypeOf(x)
	fmt.Printf("TypeOf(%v) = %v\n", x, t)
	fmt.Printf("Type name: %s\n", t.Name())
	fmt.Printf("Type kind: %v\n", t.Kind())
	
	// 获取 Value
	v := reflect.ValueOf(x)
	fmt.Printf("ValueOf(%v) = %v\n", x, v)
	fmt.Printf("Value type: %v\n", v.Type())
	fmt.Printf("Value kind: %v\n", v.Kind())
	fmt.Printf("Interface: %v\n", v.Interface())
	fmt.Printf("Int: %d\n", v.Int())
	
	// 字符串
	str := "hello"
	vStr := reflect.ValueOf(str)
	fmt.Printf("String value: %s\n", vStr.String())
	
	// 检查是否可设置
	fmt.Printf("CanSet: %v\n", v.CanSet())  // false，因为不是指针
}

// ============================================
// 2. 修改值（通过指针）
// ============================================

func demonstrateModifyValue() {
	fmt.Println("\n=== 修改值 ===")
	
	x := 42
	
	// 获取指针的 Value
	v := reflect.ValueOf(&x)
	fmt.Printf("Pointer value: %v\n", v)
	fmt.Printf("Pointer kind: %v\n", v.Kind())
	
	// 解引用获取指向的值
	elem := v.Elem()
	fmt.Printf("Elem type: %v\n", elem.Type())
	fmt.Printf("Elem kind: %v\n", elem.Kind())
	fmt.Printf("Elem can set: %v\n", elem.CanSet())
	
	// 修改值
	if elem.CanSet() {
		elem.SetInt(100)
	}
	fmt.Printf("x after set: %d\n", x)
	
	// 修改字符串
	str := "hello"
	vStr := reflect.ValueOf(&str).Elem()
	vStr.SetString("world")
	fmt.Printf("str after set: %s\n", str)
}

// ============================================
// 3. 类型检查与转换
// ============================================

func demonstrateTypeInspection() {
	fmt.Println("\n=== 类型检查 ===")
	
	// 检查类型
	checkType := func(v interface{}) {
		t := reflect.TypeOf(v)
		fmt.Printf("%T: Kind=%v, Name=%s\n", v, t.Kind(), t.Name())
	}
	
	checkType(42)
	checkType(3.14)
	checkType("hello")
	checkType(true)
	checkType([]int{1, 2, 3})
	checkType(map[string]int{"a": 1})
	checkType(struct{}{})
	
	// 类型转换
	var i interface{} = int64(42)
	
	// 检查底层类型
	v := reflect.ValueOf(i)
	fmt.Printf("Interface value kind: %v\n", v.Kind())
	
	// 转换为具体类型
	if v.Kind() == reflect.Int64 {
		n := v.Int()
		fmt.Printf("Converted to int64: %d\n", n)
	}
}

// ============================================
// 4. 结构体反射 ⭐
// ============================================

type Person struct {
	Name    string `json:"name" validate:"required"`
	Age     int    `json:"age" validate:"min=0,max=150"`
	Email   string `json:"email,omitempty"`
	private string // 未导出字段
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s", p.Name)
}

func (p *Person) HaveBirthday() {
	p.Age++
}

func demonstrateStructReflection() {
	fmt.Println("\n=== 结构体反射 ===")
	
	p := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
	
	// 获取 Type 和 Value
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)
	
	fmt.Printf("Type: %v, NumField: %d\n", t, t.NumField())
	
	// 遍历字段
	fmt.Println("\n字段信息:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		
		fmt.Printf("  Field %d:\n", i)
		fmt.Printf("    Name: %s\n", field.Name)
		fmt.Printf("    Type: %v\n", field.Type)
		fmt.Printf("    Tag: %s\n", field.Tag)
		fmt.Printf("    Value: %v\n", value.Interface())
		fmt.Printf("    Exported: %v\n", field.PkgPath == "")  // 空表示导出
	}
	
	// 通过名称获取字段
	if nameField, ok := t.FieldByName("Name"); ok {
		fmt.Printf("\nName field tag: %s\n", nameField.Tag)
		
		// 获取标签值
		jsonTag := nameField.Tag.Get("json")
		fmt.Printf("JSON tag: %s\n", jsonTag)
	}
	
	// 修改结构体（通过指针）
	vPtr := reflect.ValueOf(&p)
	elem := vPtr.Elem()
	
	if nameField := elem.FieldByName("Name"); nameField.IsValid() && nameField.CanSet() {
		nameField.SetString("Bob")
	}
	fmt.Printf("Modified person: %+v\n", p)
}

// ============================================
// 5. 结构体标签解析
// ============================================

// 标签解析器
type FieldInfo struct {
	Name      string
	Type      string
	JSONName  string
	Required  bool
	OmitEmpty bool
	Validate  string
}

func parseStructTags(t reflect.Type) []FieldInfo {
	var fields []FieldInfo
	
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		
		// 跳过未导出字段
		if field.PkgPath != "" {
			continue
		}
		
		info := FieldInfo{
			Name: field.Name,
			Type: field.Type.String(),
		}
		
		// 解析 json 标签
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			parts := strings.Split(jsonTag, ",")
			info.JSONName = parts[0]
			for _, part := range parts[1:] {
				if part == "omitempty" {
					info.OmitEmpty = true
				}
			}
			if info.JSONName == "-" {
				continue  // 跳过此字段
			}
		}
		
		// 解析 validate 标签
		info.Validate = field.Tag.Get("validate")
		info.Required = strings.Contains(info.Validate, "required")
		
		fields = append(fields, info)
	}
	
	return fields
}

func demonstrateTagParsing() {
	fmt.Println("\n=== 标签解析 ===")
	
	t := reflect.TypeOf(Person{})
	fields := parseStructTags(t)
	
	for _, f := range fields {
		fmt.Printf("Field: %s\n", f.Name)
		fmt.Printf("  JSON Name: %s\n", f.JSONName)
		fmt.Printf("  Required: %v\n", f.Required)
		fmt.Printf("  OmitEmpty: %v\n", f.OmitEmpty)
		fmt.Printf("  Validate: %s\n", f.Validate)
	}
}

// ============================================
// 6. 方法反射
// ============================================

func demonstrateMethodReflection() {
	fmt.Println("\n=== 方法反射 ===")
	
	p := Person{Name: "Alice", Age: 30}
	
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)
	
	fmt.Printf("NumMethod: %d\n", t.NumMethod())
	
	// 遍历方法（值接收者方法）
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("Method %d: %s\n", i, method.Name)
		fmt.Printf("  Type: %v\n", method.Type)
		fmt.Printf("  NumIn: %d, NumOut: %d\n", method.Type.NumIn(), method.Type.NumOut())
	}
	
	// 调用方法
	if greetMethod := v.MethodByName("Greet"); greetMethod.IsValid() {
		results := greetMethod.Call(nil)
		fmt.Printf("Greet result: %s\n", results[0].String())
	}
	
	// 指针类型可以调用所有方法
	vPtr := reflect.ValueOf(&p)
	tPtr := reflect.TypeOf(&p)
	
	fmt.Printf("\nPointer NumMethod: %d\n", tPtr.NumMethod())
	
	// 调用指针接收者方法
	if birthdayMethod := vPtr.MethodByName("HaveBirthday"); birthdayMethod.IsValid() {
		birthdayMethod.Call(nil)
		fmt.Printf("Age after birthday: %d\n", p.Age)
	}
}

// ============================================
// 7. 切片和 Map 反射
// ============================================

func demonstrateSliceMapReflection() {
	fmt.Println("\n=== 切片和 Map 反射 ===")
	
	// 切片
	nums := []int{1, 2, 3, 4, 5}
	v := reflect.ValueOf(nums)
	
	fmt.Printf("Slice kind: %v\n", v.Kind())
	fmt.Printf("Slice len: %d\n", v.Len())
	fmt.Printf("Slice cap: %d\n", v.Cap())
	
	// 访问元素
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		fmt.Printf("  [%d] = %d\n", i, elem.Int())
	}
	
	// 修改元素（通过指针）
	vPtr := reflect.ValueOf(&nums).Elem()
	vPtr.Index(0).SetInt(100)
	fmt.Printf("Modified slice: %v\n", nums)
	
	// 创建新切片
	newSlice := reflect.MakeSlice(v.Type(), 3, 5)
	fmt.Printf("New slice: %v, len=%d, cap=%d\n", newSlice, newSlice.Len(), newSlice.Cap())
	
	// Map
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	vMap := reflect.ValueOf(m)
	
	fmt.Printf("\nMap kind: %v\n", vMap.Kind())
	fmt.Printf("Map len: %d\n", vMap.Len())
	
	// 遍历 map
	for _, key := range vMap.MapKeys() {
		value := vMap.MapIndex(key)
		fmt.Printf("  %s: %d\n", key.String(), value.Int())
	}
	
	// 修改 map（通过指针）
	vMapPtr := reflect.ValueOf(&m).Elem()
	vMapPtr.SetMapIndex(reflect.ValueOf("d"), reflect.ValueOf(4))
	fmt.Printf("Modified map: %v\n", m)
	
	// 删除元素
	vMapPtr.SetMapIndex(reflect.ValueOf("a"), reflect.Value{})  // 空值表示删除
	fmt.Printf("After delete: %v\n", m)
	
	// 创建新 map
	newMap := reflect.MakeMap(vMap.Type())
	newMap.SetMapIndex(reflect.ValueOf("x"), reflect.ValueOf(10))
	fmt.Printf("New map: %v\n", newMap.Interface())
}

// ============================================
// 8. 创建新值
// ============================================

func demonstrateCreateValues() {
	fmt.Println("\n=== 创建新值 ===")
	
	// 创建基本类型
	intType := reflect.TypeOf(0)
	newInt := reflect.New(intType)  // 创建 *int
	newInt.Elem().SetInt(42)
	fmt.Printf("New int: %d\n", newInt.Elem().Int())
	
	// 创建结构体
	personType := reflect.TypeOf(Person{})
	newPerson := reflect.New(personType).Elem()
	
	newPerson.FieldByName("Name").SetString("Charlie")
	newPerson.FieldByName("Age").SetInt(25)
	newPerson.FieldByName("Email").SetString("charlie@example.com")
	
	fmt.Printf("New person: %+v\n", newPerson.Interface())
	
	// 创建切片
	sliceType := reflect.TypeOf([]int{})
	newSlice := reflect.MakeSlice(sliceType, 0, 10)
	
	newSlice = reflect.Append(newSlice, reflect.ValueOf(1))
	newSlice = reflect.Append(newSlice, reflect.ValueOf(2))
	newSlice = reflect.Append(newSlice, reflect.ValueOf(3))
	
	fmt.Printf("New slice: %v\n", newSlice.Interface())
	
	// 创建 map
	mapType := reflect.TypeOf(map[string]int{})
	newMap := reflect.MakeMap(mapType)
	
	newMap.SetMapIndex(reflect.ValueOf("one"), reflect.ValueOf(1))
	newMap.SetMapIndex(reflect.ValueOf("two"), reflect.ValueOf(2))
	
	fmt.Printf("New map: %v\n", newMap.Interface())
}

// ============================================
// 9. 实用工具：深拷贝
// ============================================

func deepCopy(dst, src interface{}) error {
	dstVal := reflect.ValueOf(dst)
	srcVal := reflect.ValueOf(src)
	
	if dstVal.Kind() != reflect.Ptr || dstVal.IsNil() {
		return fmt.Errorf("dst must be a non-nil pointer")
	}
	
	dstElem := dstVal.Elem()
	if srcVal.Type() != dstElem.Type() {
		return fmt.Errorf("src and dst must have the same type")
	}
	
	deepCopyValue(dstElem, srcVal)
	return nil
}

func deepCopyValue(dst, src reflect.Value) {
	switch src.Kind() {
	case reflect.Ptr:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.New(src.Elem().Type()))
		deepCopyValue(dst.Elem(), src.Elem())
		
	case reflect.Struct:
		for i := 0; i < src.NumField(); i++ {
			deepCopyValue(dst.Field(i), src.Field(i))
		}
		
	case reflect.Slice:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.MakeSlice(src.Type(), src.Len(), src.Cap()))
		for i := 0; i < src.Len(); i++ {
			deepCopyValue(dst.Index(i), src.Index(i))
		}
		
	case reflect.Map:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.MakeMap(src.Type()))
		for _, key := range src.MapKeys() {
			srcValue := src.MapIndex(key)
			dstValue := reflect.New(srcValue.Type()).Elem()
			deepCopyValue(dstValue, srcValue)
			dst.SetMapIndex(key, dstValue)
		}
		
	default:
		if dst.CanSet() {
			dst.Set(src)
		}
	}
}

func demonstrateDeepCopy() {
	fmt.Println("\n=== 深拷贝 ===")
	
	type Node struct {
		Value int
		Next  *Node
	}
	
	original := &Node{
		Value: 1,
		Next: &Node{
			Value: 2,
			Next: &Node{
				Value: 3,
			},
		},
	}
	
	var copied Node
	if err := deepCopy(&copied, original); err != nil {
		fmt.Printf("Copy error: %v\n", err)
		return
	}
	
	// 修改原值
	original.Next.Value = 200
	
	fmt.Printf("Original: %d -> %d -> %d\n", original.Value, original.Next.Value, original.Next.Next.Value)
	fmt.Printf("Copied: %d -> %d -> %d\n", copied.Value, copied.Next.Value, copied.Next.Next.Value)
}

// ============================================
// 10. 实用工具：结构体验证
// ============================================

func validateStruct(s interface{}) error {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("expected struct, got %v", v.Kind())
	}
	
	t := v.Type()
	
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		
		// 检查 required
		tag := field.Tag.Get("validate")
		if strings.Contains(tag, "required") {
			if isZeroValue(value) {
				return fmt.Errorf("field %s is required", field.Name)
			}
		}
		
		// 检查数值范围
		if strings.Contains(tag, "min=") && (value.Kind() == reflect.Int || value.Kind() == reflect.Float64) {
			minStr := extractTagValue(tag, "min=")
			if minStr != "" {
				min, _ := strconv.Atoi(minStr)
				if int(value.Int()) < min {
					return fmt.Errorf("field %s must be >= %d", field.Name, min)
				}
			}
		}
		
		if strings.Contains(tag, "max=") && (value.Kind() == reflect.Int || value.Kind() == reflect.Float64) {
			maxStr := extractTagValue(tag, "max=")
			if maxStr != "" {
				max, _ := strconv.Atoi(maxStr)
				if int(value.Int()) > max {
					return fmt.Errorf("field %s must be <= %d", field.Name, max)
				}
			}
		}
	}
	
	return nil
}

func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Ptr, reflect.Slice, reflect.Map:
		return v.IsNil()
	default:
		return false
	}
}

func extractTagValue(tag, key string) string {
	idx := strings.Index(tag, key)
	if idx == -1 {
		return ""
	}
	
	start := idx + len(key)
	end := start
	for end < len(tag) && tag[end] != ',' && tag[end] != ' ' {
		end++
	}
	
	return tag[start:end]
}

func demonstrateValidation() {
	fmt.Println("\n=== 结构体验证 ===")
	
	validPerson := Person{Name: "Alice", Age: 30}
	if err := validateStruct(validPerson); err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Println("Valid person: OK")
	}
	
	invalidPerson := Person{Name: "", Age: 200}
	if err := validateStruct(invalidPerson); err != nil {
		fmt.Printf("Validation error: %v\n", err)
	}
}

// ============================================
// 主函数
// ============================================

func main() {
	demonstrateBasicReflection()
	demonstrateModifyValue()
	demonstrateTypeInspection()
	demonstrateStructReflection()
	demonstrateTagParsing()
	demonstrateMethodReflection()
	demonstrateSliceMapReflection()
	demonstrateCreateValues()
	demonstrateDeepCopy()
	demonstrateValidation()
	
	// ============================================
	// 练习题
	// ============================================
	//
	// 练习 1：实现一个通用的 Map 转换函数
	//   func TransformMap(input interface{}, fn func(interface{}) interface{}) interface{}
	//   - 支持任意类型的 map
	//   - 对每个值应用转换函数
	//
	// 练习 2：实现结构体到 map 的转换
	//   func StructToMap(s interface{}) map[string]interface{}
	//   - 只处理导出字段
	//   - 使用 json tag 作为 key
	//   - 递归处理嵌套结构体
	//
	// 练习 3：实现 map 到结构体的转换
	//   func MapToStruct(m map[string]interface{}, s interface{}) error
	//   - 使用反射设置结构体字段
	//   - 处理类型转换
	//   - 支持嵌套结构体
	//
	// 练习 4：实现一个依赖注入容器（使用反射）
	//   type DIContainer struct { ... }
	//   - Register(constructor interface{}) 注册构造函数
	//   - Resolve(target interface{}) error 解析依赖
	//   - 自动注入构造函数参数
	//
	// 练习 5：实现 RPC 调用器
	//   type RPCClient struct { ... }
	//   - Call(method string, args []interface{}, reply interface{}) error
	//   - 使用反射检查方法签名
	//   - 验证参数数量和类型
	//
	// 练习 6：实现一个 ORM 风格的查询构建器
	//   type Query struct { ... }
	//   - Where(field string, op string, value interface{}) *Query
	//   - Find(dest interface{}) error
	//   - 使用反射填充结果到结构体切片
	//
	// 练习 7：实现一个 JSON Schema 生成器
	//   func GenerateSchema(t interface{}) map[string]interface{}
	//   - 从结构体标签生成 JSON Schema
	//   - 支持 required、type、format 等字段
}
