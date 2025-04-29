# Go 程序性能分析

## pprof

pprof 是 Go 语言内置的性能分析工具，可以帮助分析 CPU 使用率、内存分配、goroutine 阻塞等问题。

### 1. 添加 pprof 支持到程序中

```go
import (
    "net/http"
    _ "net/http/pprof"
)

func main() {
    // 启动 pprof 服务
    go func() {
        http.ListenAndServe("localhost:6060", nil)
    }()
    
    // 你的程序代码...
}
```

### 2. 收集和分析数据

**CPU 分析**：
```
go tool pprof http://localhost:6060/debug/pprof/profile
```

**内存分析**：
```
go tool pprof http://localhost:6060/debug/pprof/heap
```

**阻塞分析**：
```
go tool pprof http://localhost:6060/debug/pprof/block
```

### 3. 常用的 pprof 命令

在 pprof 交互模式下，可以使用：
- `top`：显示消耗资源最多的函数
- `list <函数名>`：显示特定函数的详细信息
- `web`：在浏览器中查看性能图表（需要安装 graphviz）
- `pdf`：生成 PDF 报告

## 使用 trace 工具

trace 可以帮助分析程序在并发执行时的行为：

```
curl -o trace.out http://localhost:6060/debug/pprof/trace?seconds=5
go tool trace trace.out
```

## 使用基准测试

```go
func BenchmarkXxx(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // 待测试代码
    }
}
```

运行：
```
go test -bench=. -benchmem
```

## 其他实用技巧

1. **使用 runtime/pprof 库**：不需要 HTTP 服务器的场景

```go
import "runtime/pprof"

// CPU 分析
f, _ := os.Create("cpu.prof")
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()

// 内存分析
f, _ := os.Create("mem.prof")
pprof.WriteHeapProfile(f)
defer f.Close()
```

2. **使用 -memprofile 和 -cpuprofile**：
```
go test -memprofile=mem.prof -cpuprofile=cpu.prof
```

# Go 程序 CPU 性能分析示例

这个项目包含两个演示如何使用 pprof 进行 CPU 性能分析的示例程序。

## 快速开始

运行主程序并选择要运行的示例：

```bash
go run main.go
```

## 示例一：使用 HTTP 服务器

位于 `cmd/demo1/main.go` 的示例展示如何通过 HTTP 服务器收集和查看性能分析数据。

### 直接运行方法：

```bash
go run cmd/demo1/main.go
```

程序启动后会在后台运行，并启动一个 HTTP 服务器。你可以通过以下方式查看分析数据：

1. **通过浏览器**：
   访问 http://localhost:6060/debug/pprof/

2. **通过命令行**（需要另开一个终端）：
   ```bash
   # 收集 30 秒的 CPU 分析数据
   go tool pprof http://localhost:6060/debug/pprof/profile
   ```

## 示例二：直接生成分析文件

位于 `cmd/demo2/main.go` 的示例展示如何直接将性能分析数据保存到文件。

### 直接运行方法：

```bash
go run cmd/demo2/main.go
```

程序会运行约 10 秒钟，然后自动停止并生成 `cpu.prof` 文件。

### 分析数据：

```bash
go tool pprof cpu.prof
```

## 常用的 pprof 命令

进入 pprof 交互模式后，可以使用以下命令分析性能数据：

- `top`：显示占用 CPU 时间最多的函数
- `top 10`：显示前 10 个函数
- `list <函数名>`：显示函数的详细调用情况
- `web`：在浏览器中以图形方式查看调用图（需要安装 Graphviz）
- `pdf`：生成 PDF 格式的调用图
- `png`：生成 PNG 格式的调用图

## 安装 Graphviz（用于图形化展示）

使用 `web`、`pdf` 等命令需要安装 Graphviz：

- macOS: `brew install graphviz`
- Ubuntu/Debian: `apt-get install graphviz`
- Windows: 从 https://graphviz.org/download/ 下载并安装

## 解读分析结果

分析结果中主要关注以下几点：

1. **flat**：函数自身占用的 CPU 时间
2. **cum**：函数及其调用的函数占用的 CPU 总时间
3. **调用图**：显示函数之间的调用关系和每个调用占用的时间

## 性能优化建议

1. 识别热点函数（CPU 占用最高的函数）
2. 优化算法复杂度
3. 减少内存分配和垃圾回收
4. 使用并发（适当情况下）
5. 使用更高效的数据结构