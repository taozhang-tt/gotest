package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Go 程序 CPU 性能分析示例")
	fmt.Println("=========================")
	fmt.Println("1. 运行 HTTP 服务器示例 (使用 net/http/pprof)")
	fmt.Println("2. 运行文件生成示例 (使用 runtime/pprof)")
	fmt.Println("请选择 (1-2):")

	var choice string
	fmt.Scanln(&choice)

	var cmd *exec.Cmd

	switch choice {
	case "1":
		fmt.Println("\n运行 HTTP 服务器示例...")
		cmd = exec.Command("go", "run", "cmd/demo1/main.go")
	case "2":
		fmt.Println("\n运行文件生成示例...")
		cmd = exec.Command("go", "run", "cmd/demo2/main.go")
	default:
		fmt.Println("无效选择，请输入 1 或 2")
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("运行示例失败: %v\n", err)
	}
}
