// ============================================
// Go 结构体与方法教程
// ============================================
//
// 本文件涵盖 Go 语言面向对象编程的核心：
// - 结构体定义与初始化
// - 方法定义（值接收者 vs 指针接收者）⭐
// - 结构体嵌入（Embedding）⭐ Go 的"继承"
// - 结构体标签（Tag）
// - 匿名字段
// - 方法集
//
// 最佳实践：
// 1. 需要修改接收者状态时用指针接收者
// 2. 结构体较大时用指针接收者（避免复制开销）
// 3. 保持一致性：同一类型的方法要么全用值，要么全用指针
// 4. 嵌入用于代码复用，但不是真正的继承
// 5. 使用 JSON tag 控制序列化行为
// ============================================

package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// ============================================
// 1. 结构体定义
// ============================================
//
// 结构体是字段的集合，是值类型

// 基本结构体
type Person struct {
	Name string
	Age  int
}

// 包含多种类型的结构体
type Employee struct {
	ID        int
	Name      string
	Position  string
	Salary    float64
	HireDate  time.Time
	IsActive  bool
}

// 匿名字段（字段名即类型名）
type Anonymous struct {
	string  // 字段名是 "string"
	int     // 字段名是 "int"
}

// 嵌套结构体
type Address struct {
	City    string
	Street  string
	ZipCode string
}

type Contact struct {
	Name    string
	Email   string
	Address Address  // 嵌套结构体
}

// 带有标签的结构体（常用于 JSON/XML 序列化）
type User struct {
	ID        int       `json:"id" db:"user_id"`           // 多个标签
	Username  string    `json:"username,omitempty"`         // omitempty: 空值时省略
	Password  string    `json:"-"`                          // -: 忽略此字段
	Email     string    `json:"email" validate:"email"`
	CreatedAt time.Time `json:"created_at"`
	IsAdmin   bool      `json:"is_admin"`
}

// ============================================
// 2. 结构体初始化
// ============================================

func demonstrateStructInit() {
	// 方式1：按字段顺序初始化（不推荐，字段顺序改变会出错）
	p1 := Person{"Alice", 30}
	
	// 方式2：按字段名初始化（推荐）
	p2 := Person{
		Name: "Bob",
		Age:  25,
	}
	
	// 方式3：零值初始化
	var p3 Person  // Name="", Age=0
	
	// 方式4：new 关键字（返回指针）
	p4 := new(Person)  // *Person，字段为零值
	p4.Name = "Charlie"
	
	// 方式5：& 取地址（最常用）
	p5 := &Person{Name: "David", Age: 35}
	
	fmt.Printf("p1: %+v\n", p1)
	fmt.Printf("p2: %+v\n", p2)
	fmt.Printf("p3: %+v\n", p3)
	fmt.Printf("p4: %+v\n", *p4)
	fmt.Printf("p5: %+v\n", p5)
	
	// 嵌套结构体初始化
	contact := Contact{
		Name:  "Eve",
		Email: "eve@example.com",
		Address: Address{
			City:    "Beijing",
			Street:  "Main St",
			ZipCode: "100000",
		},
	}
	fmt.Printf("contact: %+v\n", contact)
}

// ============================================
// 3. 方法定义
// ============================================
//
// 方法是有接收者的函数
// 接收者类型前加 * 表示指针接收者

// 值接收者方法 - 操作的是副本
func (p Person) GetName() string {
	return p.Name
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// 指针接收者方法 - 可以修改原值
func (p *Person) HaveBirthday() {
	p.Age++
}

func (p *Person) ChangeName(newName string) {
	p.Name = newName
}

// ============================================
// 4. 值接收者 vs 指针接收者 ⭐
// ============================================
//
// 值接收者：
//   - 方法操作的是结构体的副本
//   - 不能修改原结构体
//   - 适用于小结构体和只读操作
//
// 指针接收者：
//   - 方法操作的是原结构体
//   - 可以修改原结构体
//   - 适用于大结构体（避免复制开销）
//   - 如果需要修改状态，必须用指针

func demonstrateReceiver() {
	p := Person{Name: "Alice", Age: 30}
	
	// 值接收者 - 操作副本
	fmt.Printf("Name: %s\n", p.GetName())
	fmt.Printf("IsAdult: %v\n", p.IsAdult())
	
	// 指针接收者 - 修改原值
	fmt.Printf("当前年龄: %d\n", p.Age)
	p.HaveBirthday()  // 自动解引用，等价于 (&p).HaveBirthday()
	fmt.Printf("过生日后: %d\n", p.Age)
	
	p.ChangeName("Alicia")
	fmt.Printf("改名后: %s\n", p.Name)
}

// ============================================
// 5. 方法集
// ============================================
//
// T 类型的方法集：所有接收者为 T 的方法
// *T 类型的方法集：所有接收者为 T 和 *T 的方法

func demonstrateMethodSet() {
	p := Person{Name: "Bob", Age: 25}
	
	// 值类型可以调用值接收者方法
	fmt.Println(p.GetName())
	
	// 值类型也可以调用指针接收者方法（Go 自动取地址）
	p.HaveBirthday()  // 等价于 (&p).HaveBirthday()
	
	// 指针类型可以调用所有方法
	ptr := &p
	fmt.Println(ptr.GetName())      // 值接收者
	ptr.HaveBirthday()              // 指针接收者
}

// ============================================
// 6. 结构体嵌入（Embedding）⭐
// ============================================
//
// Go 没有继承，使用嵌入实现代码复用
// 嵌入的字段可以直接访问（提升字段）

// 嵌入基本类型
type Engine struct {
	Power int
	Type  string
}

func (e Engine) Start() {
	fmt.Printf("%s 引擎启动，功率: %d\n", e.Type, e.Power)
}

func (e Engine) Stop() {
	fmt.Printf("%s 引擎停止\n", e.Type)
}

// Car 嵌入了 Engine
type Car struct {
	Engine    // 匿名字段，嵌入
	Brand string
	Model string
}

// 可以重写嵌入类型的方法
func (c Car) Start() {
	fmt.Printf("🚗 %s %s 准备启动...\n", c.Brand, c.Model)
	c.Engine.Start()  // 调用嵌入类型的方法
}

// 多层嵌入
type ElectricEngine struct {
	Engine
	BatteryCapacity int
}

type ElectricCar struct {
	ElectricEngine
	Brand string
}

func demonstrateEmbedding() {
	// 创建 Car
	car := Car{
		Engine: Engine{Power: 200, Type: "V8"},
		Brand:  "Toyota",
		Model:  "Camry",
	}
	
	// 直接访问嵌入字段的方法和字段（提升）
	fmt.Println("引擎功率:", car.Power)      // 等价于 car.Engine.Power
	fmt.Println("引擎类型:", car.Type)       // 等价于 car.Engine.Type
	
	car.Start()  // 调用 Car.Start()
	car.Stop()   // 调用 Engine.Stop()（被提升）
	
	// 也可以完整路径访问
	car.Engine.Start()
}

// ============================================
// 7. 结构体标签（Tag）应用
// ============================================

func demonstrateTag() {
	user := User{
		ID:        1,
		Username:  "john_doe",
		Password:  "secret123",  // 不会被序列化
		Email:     "john@example.com",
		CreatedAt: time.Now(),
		IsAdmin:   false,
	}
	
	// 序列化为 JSON
	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("JSON 编码错误:", err)
		return
	}
	fmt.Println("JSON 输出:")
	fmt.Println(string(jsonData))
	
	// 从 JSON 解码
	jsonInput := `{
		"id": 2,
		"username": "jane",
		"email": "jane@example.com",
		"created_at": "2024-01-15T10:30:00Z",
		"is_admin": true
	}`
	
	var decoded User
	if err := json.Unmarshal([]byte(jsonInput), &decoded); err != nil {
		fmt.Println("JSON 解码错误:", err)
		return
	}
	fmt.Printf("解码后: %+v\n", decoded)
}

// ============================================
// 8. 结构体比较与赋值
// ============================================

func demonstrateComparison() {
	// 结构体是可比较的（如果所有字段都可比较）
	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{Name: "Alice", Age: 30}
	p3 := Person{Name: "Bob", Age: 25}
	
	fmt.Printf("p1 == p2: %v\n", p1 == p2)  // true
	fmt.Printf("p1 == p3: %v\n", p1 == p3)  // false
	
	// 包含切片或 map 的结构体不可比较
	type Team struct {
		Name    string
		Members []string  // 切片不可比较
	}
	
	t1 := Team{Name: "A", Members: []string{"Alice", "Bob"}}
	t2 := Team{Name: "A", Members: []string{"Alice", "Bob"}}
	
	// fmt.Println(t1 == t2)  // 编译错误！
	
	// 使用 reflect.DeepEqual 比较
	fmt.Printf("DeepEqual: %v\n", fmt.Sprintf("%v", t1) == fmt.Sprintf("%v", t2))
}

// ============================================
// 9. 完整示例：银行账户
// ============================================

type BankAccount struct {
	AccountNumber string
	Owner         string
	Balance       float64
	isClosed      bool  // 小写：包内私有
}

// 构造函数（惯用法）
func NewBankAccount(accountNumber, owner string, initialBalance float64) *BankAccount {
	if initialBalance < 0 {
		initialBalance = 0
	}
	return &BankAccount{
		AccountNumber: accountNumber,
		Owner:         owner,
		Balance:       initialBalance,
		isClosed:      false,
	}
}

func (ba *BankAccount) Deposit(amount float64) error {
	if ba.isClosed {
		return fmt.Errorf("账户已关闭")
	}
	if amount <= 0 {
		return fmt.Errorf("存款金额必须大于0")
	}
	ba.Balance += amount
	return nil
}

func (ba *BankAccount) Withdraw(amount float64) error {
	if ba.isClosed {
		return fmt.Errorf("账户已关闭")
	}
	if amount <= 0 {
		return fmt.Errorf("取款金额必须大于0")
	}
	if amount > ba.Balance {
		return fmt.Errorf("余额不足")
	}
	ba.Balance -= amount
	return nil
}

func (ba BankAccount) GetBalance() float64 {
	return ba.Balance
}

func (ba *BankAccount) Close() {
	ba.isClosed = true
}

func demonstrateBankAccount() {
	fmt.Println("\n=== 银行账户示例 ===")
	
	account := NewBankAccount("10086", "张三", 1000)
	
	fmt.Printf("初始余额: %.2f\n", account.GetBalance())
	
	if err := account.Deposit(500); err != nil {
		fmt.Println("存款失败:", err)
	} else {
		fmt.Printf("存款 500 后余额: %.2f\n", account.GetBalance())
	}
	
	if err := account.Withdraw(200); err != nil {
		fmt.Println("取款失败:", err)
	} else {
		fmt.Printf("取款 200 后余额: %.2f\n", account.GetBalance())
	}
	
	// 尝试透支
	if err := account.Withdraw(2000); err != nil {
		fmt.Println("取款失败:", err)
	}
}

// ============================================
// 主函数
// ============================================

func main() {
	fmt.Println("=== 结构体初始化 ===")
	demonstrateStructInit()
	
	fmt.Println("\n=== 方法接收者 ===")
	demonstrateReceiver()
	
	fmt.Println("\n=== 方法集 ===")
	demonstrateMethodSet()
	
	fmt.Println("\n=== 结构体嵌入 ===")
	demonstrateEmbedding()
	
	fmt.Println("\n=== 结构体标签 ===")
	demonstrateTag()
	
	fmt.Println("\n=== 结构体比较 ===")
	demonstrateComparison()
	
	demonstrateBankAccount()
	
	// ============================================
	// 练习题
	// ============================================
	//
	// 练习 1：定义一个 Rectangle 结构体，包含 Width 和 Height
	//   - 实现 Area() 计算面积
	//   - 实现 Perimeter() 计算周长
	//   - 实现 Scale(factor float64) 按因子缩放（修改原值）
	//   - 实现 IsSquare() 判断是否为正方形
	//
	// 练习 2：实现一个 Book 结构体
	//   - 字段：Title, Author, ISBN, Price, PublishedYear
	//   - 实现 ApplyDiscount(discountPercent float64) 打折
	//   - 实现 GetAge() 返回书的"年龄"
	//   - 实现 String() string 方法（格式化输出）
	//
	// 练习 3：使用嵌入实现以下结构
	//   - 基础 Person 结构体（Name, Age）
	//   - Student 嵌入 Person，添加 StudentID, Major, Grades([]float64)
	//   - Teacher 嵌入 Person，添加 TeacherID, Department, Salary
	//   - 为 Student 实现 GetAverageGrade() 方法
	//
	// 练习 4：实现一个缓存结构体
	//   type Cache struct {
	//       data map[string]interface{}
	//       ttl  map[string]time.Time  // 过期时间
	//   }
	//   - 实现 Set(key string, value interface{}, duration time.Duration)
	//   - 实现 Get(key string) (interface{}, bool)
	//   - 实现 Delete(key string)
	//   - Get 时检查是否过期
	//
	// 练习 5：实现一个链表结构体
	//   type Node struct {
	//       Value int
	//       Next  *Node
	//   }
	//   - 实现 Append(value int) 在尾部添加
	//   - 实现 Insert(index, value int) 在指定位置插入
	//   - 实现 Delete(index int) 删除指定位置
	//   - 实现 Reverse() 反转链表
	//   - 实现 String() 打印链表内容
}
