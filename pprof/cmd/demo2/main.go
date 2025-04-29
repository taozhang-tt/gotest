package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

// 一个计算密集型函数，模拟 CPU 密集操作
func calculatePrimesFile(max int) []int {
	var primes []int
	for i := 2; i < max; i++ {
		isPrime := true
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}

// 模拟一些矩阵操作，会消耗大量 CPU
func matrixMultiplyFile(size int) {
	a := make([][]int, size)
	b := make([][]int, size)
	c := make([][]int, size)

	// 初始化矩阵
	for i := 0; i < size; i++ {
		a[i] = make([]int, size)
		b[i] = make([]int, size)
		c[i] = make([]int, size)
		for j := 0; j < size; j++ {
			a[i][j] = rand.Intn(100)
			b[i][j] = rand.Intn(100)
		}
	}

	// 矩阵乘法
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
}

func simulateWorkFile() {
	// 执行一些计算密集型操作
	calculatePrimesFile(5000)
	matrixMultiplyFile(100)

	// 随机延迟
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
}

func main() {
	// 创建 CPU 分析文件
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Printf("创建 CPU 分析文件失败: %v\n", err)
		return
	}
	defer cpuFile.Close()

	// 开始 CPU 分析
	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		fmt.Printf("开始 CPU 分析失败: %v\n", err)
		return
	}
	defer pprof.StopCPUProfile()

	fmt.Println("开始收集 CPU 分析数据...")
	fmt.Println("程序将运行 10 秒钟...")

	// 运行一段时间的计算密集型工作
	startTime := time.Now()
	for time.Since(startTime) < 10*time.Second {
		simulateWorkFile()
	}

	fmt.Println("CPU 分析完成，数据已保存到 cpu.prof 文件")
	fmt.Println("使用以下命令查看分析结果:")
	fmt.Println("go tool pprof cpu.prof")
	fmt.Println("在 pprof 交互界面中，可以使用 top、list、web 等命令分析结果")
}
