// ============================================
// Go 标准库常用包教程
// ============================================
//
// 本文件涵盖 Go 标准库中常用的包：
// - fmt - 格式化 I/O
// - strings/bytes - 字符串和字节操作
// - strconv - 类型转换
// - time - 时间处理
// - os/path/filepath - 文件系统
// - io/bufio - I/O 操作
// - encoding/json - JSON 处理
// - net/http - HTTP 服务
// - sync - 同步原语（已在 06_sync_context.go 覆盖）
// - sort - 排序
// - regexp - 正则表达式
//
// 最佳实践：
// 1. 优先使用标准库，避免不必要的第三方依赖
// 2. 了解每个包的适用场景和限制
// 3. 注意资源释放（Close）和错误处理
// 4. 使用 context 控制超时和取消
// ============================================

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ============================================
// 1. fmt 包 - 格式化 I/O
// ============================================

func demonstrateFmt() {
	fmt.Println("=== fmt 包 ===")
	
	// Println - 自动添加空格和换行
	fmt.Println("Hello", "World")  // Hello World
	
	// Printf - 格式化输出
	name := "Go"
	version := 1.22
	fmt.Printf("语言: %s, 版本: %.2f\n", name, version)
	
	// 常用格式化动词
	fmt.Printf("字符串: %s / %q\n", "hello", "hello")  // %q 带引号
	fmt.Printf("整数: %d / %x / %b\n", 42, 42, 42)     // 十进制/十六进制/二进制
	fmt.Printf("浮点数: %f / %.2f / %e\n", 3.14159, 3.14159, 3.14159)
	fmt.Printf("布尔: %t\n", true)
	fmt.Printf("任意类型: %v / %+v / %#v\n", 
		struct{A int}{42},
		struct{A int}{42},
		struct{A int}{42})
	
	// Sprintf - 返回字符串
	s := fmt.Sprintf("格式化后的字符串: %d", 42)
	fmt.Println(s)
	
	// Fprintf - 写入 io.Writer
	fmt.Fprintf(os.Stdout, "写入标准输出: %s\n", "test")
	
	// Scanf - 格式化输入
	// var input string
	// fmt.Print("输入: ")
	// fmt.Scanf("%s", &input)
	
	// Errorf - 创建格式化错误
	err := fmt.Errorf("发生错误，代码: %d", 500)
	fmt.Printf("错误: %v\n", err)
}

// ============================================
// 2. strings 包 - 字符串操作
// ============================================

func demonstrateStrings() {
	fmt.Println("\n=== strings 包 ===")
	
	s := "  Hello, World!  "
	
	// 基本操作
	fmt.Printf("Contains 'World': %v\n", strings.Contains(s, "World"))
	fmt.Printf("HasPrefix 'Hello': %v\n", strings.HasPrefix(strings.TrimSpace(s), "Hello"))
	fmt.Printf("HasSuffix '!': %v\n", strings.HasSuffix(strings.TrimSpace(s), "!"))
	
	// 查找
	fmt.Printf("Index 'World': %d\n", strings.Index(s, "World"))
	fmt.Printf("Count 'l': %d\n", strings.Count(s, "l"))
	
	// 大小写
	fmt.Printf("ToUpper: %s\n", strings.ToUpper(s))
	fmt.Printf("ToLower: %s\n", strings.ToLower(s))
	
	// 修剪
	fmt.Printf("TrimSpace: %q\n", strings.TrimSpace(s))
	fmt.Printf("Trim: %q\n", strings.Trim(s, " !"))
	
	// 替换
	fmt.Printf("Replace: %s\n", strings.Replace(s, "World", "Go", 1))
	fmt.Printf("ReplaceAll: %s\n", strings.ReplaceAll(s, "l", "L"))
	
	// 分割和连接
	parts := strings.Split("a,b,c,d", ",")
	fmt.Printf("Split: %v\n", parts)
	
	joined := strings.Join(parts, "-")
	fmt.Printf("Join: %s\n", joined)
	
	// Fields - 按空白分割
	fields := strings.Fields("  hello   world  ")
	fmt.Printf("Fields: %v\n", fields)
	
	// Builder - 高效字符串构建（推荐用于大量拼接）
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteByte(' ')
	builder.WriteString("World")
	fmt.Printf("Builder: %s\n", builder.String())
	
	// Reader - 字符串作为 io.Reader
	reader := strings.NewReader("Hello, Reader!")
	buf := make([]byte, 5)
	reader.Read(buf)
	fmt.Printf("Reader: %s\n", string(buf))
}

// ============================================
// 3. strconv 包 - 类型转换
// ============================================

func demonstrateStrconv() {
	fmt.Println("\n=== strconv 包 ===")
	
	// 字符串转整数
	i, err := strconv.Atoi("42")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Atoi: %d\n", i)
	
	// 带基数的转换
	i64, _ := strconv.ParseInt("1010", 2, 64)  // 二进制
	fmt.Printf("ParseInt (binary): %d\n", i64)
	
	i64, _ = strconv.ParseInt("FF", 16, 64)    // 十六进制
	fmt.Printf("ParseInt (hex): %d\n", i64)
	
	// 字符串转浮点数
	f, _ := strconv.ParseFloat("3.14159", 64)
	fmt.Printf("ParseFloat: %f\n", f)
	
	// 字符串转布尔值
	b, _ := strconv.ParseBool("true")
	fmt.Printf("ParseBool: %v\n", b)
	
	// 整数转字符串
	s := strconv.Itoa(42)
	fmt.Printf("Itoa: %s\n", s)
	
	// 带格式转换
	s = strconv.FormatInt(42, 16)  // 十六进制
	fmt.Printf("FormatInt (hex): %s\n", s)
	
	s = strconv.FormatFloat(3.14159, 'f', 2, 64)
	fmt.Printf("FormatFloat: %s\n", s)
	
	// 引用和取消引用
	quoted := strconv.Quote("Hello\nWorld!")
	fmt.Printf("Quote: %s\n", quoted)
	
	unquoted, _ := strconv.Unquote(`"Hello\nWorld!"`)
	fmt.Printf("Unquote: %q\n", unquoted)
}

// ============================================
// 4. time 包 - 时间处理
// ============================================

func demonstrateTime() {
	fmt.Println("\n=== time 包 ===")
	
	// 当前时间
	now := time.Now()
	fmt.Printf("Now: %v\n", now)
	fmt.Printf("Formatted: %s\n", now.Format("2006-01-02 15:04:05"))
	
	// 时间解析
	t, _ := time.Parse("2006-01-02", "2024-03-15")
	fmt.Printf("Parsed: %v\n", t)
	
	// 时间戳
	timestamp := now.Unix()
	fmt.Printf("Unix timestamp: %d\n", timestamp)
	
	// 时间计算
	tomorrow := now.Add(24 * time.Hour)
	fmt.Printf("Tomorrow: %v\n", tomorrow)
	
	diff := tomorrow.Sub(now)
	fmt.Printf("Difference: %v\n", diff)
	fmt.Printf("Hours: %f\n", diff.Hours())
	
	// 时间比较
	fmt.Printf("Before: %v\n", now.Before(tomorrow))
	fmt.Printf("After: %v\n", now.After(tomorrow))
	fmt.Printf("Equal: %v\n", now.Equal(tomorrow))
	
	// 定时器
	timer := time.NewTimer(100 * time.Millisecond)
	<-timer.C
	fmt.Println("Timer expired")
	
	// Ticker
	ticker := time.NewTicker(50 * time.Millisecond)
	go func() {
		count := 0
		for range ticker.C {
			count++
			if count >= 3 {
				ticker.Stop()
				return
			}
			fmt.Println("Tick")
		}
	}()
	time.Sleep(200 * time.Millisecond)
	
	// 睡眠
	start := time.Now()
	time.Sleep(50 * time.Millisecond)
	fmt.Printf("Slept for: %v\n", time.Since(start))
	
	// 时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	shanghaiTime := now.In(loc)
	fmt.Printf("Shanghai time: %v\n", shanghaiTime)
}

// ============================================
// 5. os 和 filepath 包 - 文件系统
// ============================================

func demonstrateOS() {
	fmt.Println("\n=== os/filepath 包 ===")
	
	// 获取当前工作目录
	wd, _ := os.Getwd()
	fmt.Printf("Working directory: %s\n", wd)
	
	// 创建临时文件
	tmpfile, _ := os.CreateTemp("", "example-*.txt")
	defer os.Remove(tmpfile.Name())
	
	tmpfile.WriteString("Hello, World!")
	tmpfile.Close()
	fmt.Printf("Temp file: %s\n", tmpfile.Name())
	
	// 文件操作
	filename := "test.txt"
	
	// 写入文件
	os.WriteFile(filename, []byte("Hello, File!"), 0644)
	
	// 读取文件
	content, _ := os.ReadFile(filename)
	fmt.Printf("File content: %s\n", string(content))
	
	// 删除文件
	os.Remove(filename)
	
	// 目录操作
	os.Mkdir("testdir", 0755)
	os.MkdirAll("testdir/subdir1/subdir2", 0755)
	
	// 读取目录
	entries, _ := os.ReadDir(".")
	fmt.Printf("Entries in current dir: %d\n", len(entries))
	
	// 清理
	os.RemoveAll("testdir")
	
	// 路径操作
	path := "/usr/local/bin/go"
	fmt.Printf("Dir: %s\n", filepath.Dir(path))
	fmt.Printf("Base: %s\n", filepath.Base(path))
	fmt.Printf("Ext: %s\n", filepath.Ext(path))
	fmt.Printf("Clean: %s\n", filepath.Clean("/usr//local/../bin/go"))
	
	// 拼接路径
	newPath := filepath.Join("usr", "local", "bin", "go")
	fmt.Printf("Join: %s\n", newPath)
	
	// 路径分隔符转换
	fmt.Printf("ToSlash: %s\n", filepath.ToSlash(`C:\Users\name`))
}

// ============================================
// 6. io 和 bufio 包 - I/O 操作
// ============================================

func demonstrateIO() {
	fmt.Println("\n=== io/bufio 包 ===")
	
	// 使用 bytes.Buffer 作为 io.Writer
	var buf bytes.Buffer
	buf.WriteString("Hello")
	buf.WriteByte(' ')
	buf.Write([]byte("World"))
	fmt.Printf("Buffer: %s\n", buf.String())
	
	// 使用 strings.Reader 作为 io.Reader
	reader := strings.NewReader("Hello, Reader!")
	data := make([]byte, 5)
	n, _ := reader.Read(data)
	fmt.Printf("Read: %s (%d bytes)\n", string(data[:n]), n)
	
	// io.Copy
	reader2 := strings.NewReader("Copy this text")
	var dest bytes.Buffer
	io.Copy(&dest, reader2)
	fmt.Printf("Copied: %s\n", dest.String())
	
	// bufio.Scanner - 逐行读取
	input := "line1\nline2\nline3"
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		fmt.Printf("Scanned: %s\n", scanner.Text())
	}
	
	// bufio.Reader - 带缓冲的读取
	br := bufio.NewReader(strings.NewReader("Hello\nWorld"))
	line, _ := br.ReadString('\n')
	fmt.Printf("ReadString: %s\n", line)
	
	// bufio.Writer - 带缓冲的写入
	var out bytes.Buffer
	bw := bufio.NewWriter(&out)
	bw.WriteString("Buffered write")
	bw.Flush()  // 必须 Flush
	fmt.Printf("Writer output: %s\n", out.String())
}

// ============================================
// 7. encoding/json 包 - JSON 处理
// ============================================

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email,omitempty"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	IsActive  bool      `json:"is_active"`
}

func demonstrateJSON() {
	fmt.Println("\n=== encoding/json 包 ===")
	
	// 结构体转 JSON（编码）
	user := User{
		ID:        1,
		Name:      "Alice",
		Email:     "alice@example.com",
		Age:       30,
		CreatedAt: time.Now(),
		IsActive:  true,
	}
	
	// 紧凑格式
	jsonData, _ := json.Marshal(user)
	fmt.Printf("JSON: %s\n", string(jsonData))
	
	// 缩进格式
	jsonPretty, _ := json.MarshalIndent(user, "", "  ")
	fmt.Printf("Pretty JSON:\n%s\n", string(jsonPretty))
	
	// JSON 转结构体（解码）
	jsonInput := `{"id":2,"name":"Bob","age":25,"created_at":"2024-01-15T10:30:00Z","is_active":false}`
	
	var decoded User
	if err := json.Unmarshal([]byte(jsonInput), &decoded); err != nil {
		fmt.Printf("Unmarshal error: %v\n", err)
		return
	}
	fmt.Printf("Decoded: %+v\n", decoded)
	
	// 处理未知结构（使用 map）
	var generic map[string]interface{}
	json.Unmarshal([]byte(jsonInput), &generic)
	fmt.Printf("Generic: %v\n", generic)
	
	// 流式解码
	decoder := json.NewDecoder(strings.NewReader(jsonInput))
	var streamUser User
	decoder.Decode(&streamUser)
	fmt.Printf("Stream decoded: %+v\n", streamUser)
	
	// 流式编码
	var out bytes.Buffer
	encoder := json.NewEncoder(&out)
	encoder.SetIndent("", "  ")
	encoder.Encode(user)
	fmt.Printf("Stream encoded:\n%s", out.String())
}

// ============================================
// 8. net/http 包 - HTTP 服务
// ============================================

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	response := map[string]interface{}{
		"method": r.Method,
		"path":   r.URL.Path,
		"query":  r.URL.Query(),
	}
	
	json.NewEncoder(w).Encode(response)
}

func demonstrateHTTP() {
	fmt.Println("\n=== net/http 包 ===")
	
	// 注册处理器
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/api", jsonHandler)
	
	// 启动服务器（在后台）
	go func() {
		// 实际运行会阻塞，这里只在注释中展示
		// log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	
	// HTTP 客户端示例
	fmt.Println("HTTP Client examples:")
	
	// GET 请求
	resp, err := http.Get("https://api.github.com/users/github")
	if err != nil {
		fmt.Printf("GET error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	
	// 读取响应（部分）
	body := make([]byte, 200)
	resp.Body.Read(body)
	fmt.Printf("Body (first 200 bytes): %s...\n", string(body))
}

// ============================================
// 9. sort 包 - 排序
// ============================================

func demonstrateSort() {
	fmt.Println("\n=== sort 包 ===")
	
	// 整数切片排序
	ints := []int{3, 1, 4, 1, 5, 9, 2, 6}
	sort.Ints(ints)
	fmt.Printf("Sorted ints: %v\n", ints)
	
	// 字符串切片排序
	strs := []string{"banana", "apple", "cherry"}
	sort.Strings(strs)
	fmt.Printf("Sorted strings: %v\n", strs)
	
	// 检查是否已排序
	fmt.Printf("Ints are sorted: %v\n", sort.IntsAreSorted(ints))
	
	// 二分查找
	idx := sort.SearchInts(ints, 5)
	fmt.Printf("Found 5 at index: %d\n", idx)
	
	// 自定义排序
	type Person struct {
		Name string
		Age  int
	}
	
	people := []Person{
		{"Bob", 25},
		{"Alice", 30},
		{"Charlie", 20},
	}
	
	// 使用 sort.Slice
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Printf("Sorted by age: %+v\n", people)
	
	// 实现 sort.Interface
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Printf("Sorted by name: %+v\n", people)
}

// ============================================
// 10. regexp 包 - 正则表达式
// ============================================

func demonstrateRegexp() {
	fmt.Println("\n=== regexp 包 ===")
	
	// 编译正则表达式
	re := regexp.MustCompile(`\b\w+@\w+\.\w+\b`)
	
	// 匹配检查
	text := "Contact us at support@example.com or sales@company.org"
	fmt.Printf("Contains email: %v\n", re.MatchString(text))
	
	// 查找匹配
	match := re.FindString(text)
	fmt.Printf("First match: %s\n", match)
	
	// 查找所有匹配
	matches := re.FindAllString(text, -1)
	fmt.Printf("All matches: %v\n", matches)
	
	// 提取子匹配
	re2 := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
	submatches := re2.FindStringSubmatch("user@example.com")
	fmt.Printf("Submatches: %v\n", submatches)
	
	// 替换
	replaced := re.ReplaceAllString(text, "[EMAIL]")
	fmt.Printf("Replaced: %s\n", replaced)
	
	// 使用函数替换
	masked := re.ReplaceAllStringFunc(text, func(s string) string {
		parts := strings.Split(s, "@")
		return "***@" + parts[1]
	})
	fmt.Printf("Masked: %s\n", masked)
	
	// 分割
	re3 := regexp.MustCompile(`\s+`)
	parts := re3.Split("hello   world  go", -1)
	fmt.Printf("Split: %v\n", parts)
	
	// 查找索引
	idx := re.FindStringIndex(text)
	fmt.Printf("Match index: %v, matched: %s\n", idx, text[idx[0]:idx[1]])
}

// ============================================
// 主函数
// ============================================

func main() {
	demonstrateFmt()
	demonstrateStrings()
	demonstrateStrconv()
	demonstrateTime()
	demonstrateOS()
	demonstrateIO()
	demonstrateJSON()
	demonstrateHTTP()
	demonstrateSort()
	demonstrateRegexp()
	
	// ============================================
	// 练习题
	// ============================================
	//
	// 练习 1：实现一个日志分析工具
	//   - 读取日志文件
	//   - 使用 regexp 解析日志格式
	//   - 统计各种级别的日志数量（INFO, WARN, ERROR）
	//   - 按时间范围过滤日志
	//
	// 练习 2：实现一个简单的 Web 爬虫
	//   - 接收起始 URL
	//   - 使用 http.Get 获取页面
	//   - 使用 regexp 提取所有链接
	//   - 递归爬取（限制深度）
	//   - 保存页面内容到文件
	//
	// 练习 3：实现一个配置文件解析器
	//   - 支持 JSON 格式
	//   - 支持环境变量替换（${VAR}）
	//   - 支持默认值（${VAR:-default}）
	//   - 将配置加载到结构体
	//
	// 练习 4：实现一个 CSV 处理工具
	//   - 读取 CSV 文件
	//   - 解析为结构体切片
	//   - 支持类型转换（使用 strconv）
	//   - 写入 CSV 文件
	//   - 支持过滤和排序
	//
	// 练习 5：实现一个文件同步工具
	//   - 比较两个目录的内容
	//   - 使用 filepath.Walk 遍历
	//   - 根据修改时间决定同步方向
	//   - 支持排除某些文件模式
	//
	// 练习 6：实现一个 HTTP 中间件链
	//   - LoggingMiddleware - 记录请求日志
	//   - AuthMiddleware - 简单的 Token 验证
	//   - RateLimitMiddleware - 限流
	//   - RecoveryMiddleware - panic 恢复
	//   - 使用函数式编程组合中间件
	//
	// 练习 7：实现一个模板引擎（简化版）
	//   - 支持变量替换 {{.Name}}
	//   - 支持条件语句 {{if .Condition}}...{{end}}
	//   - 支持循环 {{range .Items}}...{{end}}
	//   - 使用 regexp 和 strings 实现
}
