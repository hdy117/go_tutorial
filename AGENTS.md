# Go 语言教程项目 - AI 代理指南

## 项目概述

本项目是一个**Go 语言核心特性教程**，旨在帮助学习者系统掌握 Go 语言的关键概念和编程技巧。项目采用中文作为主要文档和注释语言，包含 10 个循序渐进的教学文件，涵盖从基础语法到高级特性的完整学习路径。

**项目元数据**：
- 模块名称：`c03`
- 语言版本：Go 1.25.5
- 文档语言：中文（注释和文档主要使用中文）
- 学习方式：每个教学文件可独立运行，包含详细注释和练习题

## 项目结构

```
go_tutorial/
├── main.go                    # 项目入口（极简示例，仅输出 "end of main function"）
├── go.mod                     # Go 模块定义（模块名 c03，Go 1.25.5）
├── go.sum                     # 依赖校验和
├── .gitignore                 # Git 忽略配置（仅忽略 build 目录）
├── README.md                  # 项目主文档（Go 核心技术脑图，含代码示例和学习路线）
├── AGENTS.md                  # 本文件
│
├── tutorial/                  # 核心教程目录（10 个教学文件，共约 6200+ 行代码）
│   ├── README.md              # 教程使用指南（文件说明、学习路线、使用方法）
│   ├── exercises.md           # 练习题汇总（约 70 道练习题，按难度分级）
│   ├── user.json              # 示例数据文件（用于 JSON 处理示例）
│   │
│   ├── 01_basic_syntax.go     # 基础语法（514 行）- 变量、类型、控制流、数组、切片、Map
│   ├── 02_functions.go        # 函数特性（527 行）- 多返回值、闭包、defer、递归
│   ├── 03_struct_method.go    # 结构体与方法（898 行）- 值/指针接收者、嵌入
│   ├── 04_interface.go        # 接口（555 行）- 隐式实现、类型断言、空接口
│   ├── 05_concurrency.go      # 并发编程（567 行）- Goroutine、Channel、并发模式
│   ├── 06_sync_context.go     # 同步原语与 Context（637 行）- Mutex、WaitGroup、Context
│   ├── 07_error_handling.go   # 错误处理（560 行）- 自定义错误、错误链、panic/recover
│   ├── 08_generics.go         # 泛型编程（719 行）- 类型参数、约束、泛型容器
│   ├── 09_reflect.go          # 反射（662 行）- 类型检查、值操作、结构体反射
│   └── 10_standard_lib.go     # 标准库常用包（634 行）- fmt、strings、time、os、net/http 等
│
└── skills/golang/             # Go 开发技能库
    ├── SKILL.md               # Go 开发指南（Effective Go 速查、常见模式、最佳实践）
    ├── assets/                # 资源文件（空目录）
    ├── references/            # 参考文档
    │   ├── effective-go.md    # Effective Go 速查（约 60 行）
    │   ├── pitfalls.md        # 常见陷阱（约 80 行）
    │   └── performance.md     # 性能优化提示（约 70 行）
    └── scripts/
        └── init_module.py     # 新项目初始化脚本（196 行 Python 脚本）
```

## 技术栈

### Go 版本与依赖
- **Go 版本**：1.25.5
- **外部依赖**：
  - `github.com/google/uuid v1.6.0` - UUID 生成
  - `golang.org/x/exp v0.0.0-20260112195511-716be5621a96` - Go 扩展包

### 标准库覆盖范围
教学文件涵盖了以下标准库包：
- `fmt`, `strings`, `bytes`, `strconv` - 格式化与字符串处理
- `time` - 时间处理
- `os`, `path/filepath` - 文件系统操作
- `io`, `bufio` - I/O 操作
- `encoding/json` - JSON 处理
- `net/http` - HTTP 服务
- `sync`, `context` - 并发控制
- `reflect` - 反射
- `errors` - 错误处理

## 构建与运行

### 运行教学文件
每个教学文件都是独立的可执行程序：

```bash
# 运行特定教学文件
go run tutorial/01_basic_syntax.go
go run tutorial/02_functions.go
# ... 以此类推

# 构建可执行文件
go build -o build/output tutorial/01_basic_syntax.go
```

### 主程序
```bash
# 运行项目入口（极简示例）
go run main.go
```

### 模块管理
```bash
# 下载依赖
go mod download

# 整理模块
go mod tidy

# 查看依赖
go list -m all
```

## 代码组织规范

### 教学文件结构
每个教学文件（`tutorial/XX_*.go`）遵循统一的组织模式：

```go
// ============================================
// 文件标题和简介
// ============================================
//
// 本文件涵盖的主题列表
//
// 最佳实践说明
// ============================================

package main

import (...)

// 按主题组织的代码示例
// 每个主题包含：概念说明 + 代码示例

func main() {
    // 演示代码
    // 练习题（通常在文件末尾）
}
```

### 命名约定
- **导出标识符**：PascalCase（如 `StudentMap`）
- **私有标识符**：camelCase（如 `studentID`）
- **缩写词**：全大写（如 `HTTPRequest`, `URLString`）
- **类型别名**：使用有意义的名称（如 `StudentID string`）

### 注释规范
- 使用中文注释解释概念
- 关键概念用 `⭐` 标记
- 多行注释使用块注释格式
- 代码示例配合详细说明

## 学习路线

### 第一阶段：基础入门（第 1-3 周）
1. **01_basic_syntax.go** - 基础语法
2. **02_functions.go** - 函数特性
3. **03_struct_method.go** - 面向对象编程

### 第二阶段：进阶核心（第 4-6 周）
4. **04_interface.go** - 接口和鸭子类型
5. **05_concurrency.go** - 并发编程（Go 核心特色）
6. **06_sync_context.go** - 同步原语和上下文控制

### 第三阶段：工程实践（第 7-9 周）
7. **07_error_handling.go** - 错误处理最佳实践
8. **08_generics.go** - 泛型编程（Go 1.18+）
9. **09_reflect.go** - 反射的使用
10. **10_standard_lib.go** - 标准库常用包

## 练习题系统

### 难度分级
- ⭐ **初级**：刚学完相关概念即可完成
- ⭐⭐ **中级**：需要综合运用多个知识点
- ⭐⭐⭐ **高级**：需要深入理解底层原理或设计模式
- ⭐⭐⭐⭐ **专家**：需要综合运用多个模块的知识

### 练习题位置
- 每个教学文件末尾包含嵌入式练习题（以注释形式给出）
- `tutorial/exercises.md` 汇总所有练习题（约 70 道，按文件和难度分类）

### 练习模式
1. 阅读教学文件中的概念说明
2. 查看文件末尾的练习题注释
3. 根据注释提示完成代码实现
4. 运行验证结果

## 关键概念速查

### Go 语言特色
1. **多返回值**：`(result, error)` 模式是标准做法
2. **隐式接口**：无需显式声明 `implements`
3. **Goroutine**：轻量级线程，使用 `go` 关键字启动
4. **Channel**：goroutine 间通信机制
5. **defer**：延迟执行，常用于资源清理

### 并发编程原则
```
不要通过共享内存来通信，而要通过通信来共享内存

- Channel 的拥有者负责关闭
- 不要从接收方关闭 channel
- 使用 select 处理多个 channel 操作
- 总是考虑 goroutine 泄漏问题
```

### 错误处理模式
```go
// 错误包装
if err != nil {
    return fmt.Errorf("context: %w", err)
}

// 错误检查
if errors.Is(err, ErrNotFound) { ... }

// 类型断言
var notFound *NotFoundError
if errors.As(err, &notFound) { ... }
```

## 工具脚本

### 新项目初始化
```bash
# 使用 skills 目录下的脚本创建标准 Go 项目结构
python3 skills/golang/scripts/init_module.py <module-name>
```

该脚本会创建：
- 标准目录结构（cmd/, internal/, pkg/, api/, configs/, scripts/）
- 基础 main.go 模板
- 完整的 .gitignore 配置
- Makefile 构建脚本（含 build、test、lint、run 等目标）
- README.md 模板

## 参考资料

### 项目内参考
- `skills/golang/references/effective-go.md` - Effective Go 速查
- `skills/golang/references/pitfalls.md` - 常见陷阱
- `skills/golang/references/performance.md` - 性能优化
- `skills/golang/SKILL.md` - Go 开发技能指南（含模式、测试、常用库推荐）

### 外部资源
- [Go 官方文档](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)

## AI 代理开发指南

### 修改建议
1. **保持中文注释**：所有新添加的代码注释应使用中文
2. **统一文件格式**：教学文件使用 `package main`，包含 `main()` 函数
3. **添加练习题**：如新增教学内容，请在文件末尾添加相应练习题
4. **难度标记**：关键概念用 `⭐` 标记，练习题标注难度等级

### 添加新教学文件
如需添加新的教学文件（如 `11_advanced_patterns.go`）：
1. 放置在 `tutorial/` 目录下
2. 遵循 `XX_topic_name.go` 命名格式
3. 使用标准文件头注释模板
4. 在 `tutorial/README.md` 中更新文件列表
5. 在 `tutorial/exercises.md` 中添加相应练习题

### 代码审查清单
- [ ] 代码使用 `gofmt` 格式化
- [ ] 中文注释清晰准确
- [ ] 关键概念有 `⭐` 标记
- [ ] 包含可运行的示例代码
- [ ] 文件末尾包含练习题（如适用）
- [ ] 导入的包都被使用
- [ ] 错误处理遵循项目规范
