package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof" // 引入 pprof 包
	"time"
)

// 一个计算密集型函数，模拟 CPU 密集操作
func calculatePrimes(max int) []int {
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
func matrixMultiply(size int) {
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

func simulateWork() {
	// 执行一些计算密集型操作
	calculatePrimes(5000)
	matrixMultiply(100)

	// 随机延迟
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
}

func main() {
	// 启动 pprof HTTP 服务器
	go func() {
		fmt.Println("启动 pprof 服务器在 http://localhost:6060/debug/pprof/")
		http.ListenAndServe("localhost:6060", nil)
	}()

	fmt.Println("开始模拟 CPU 密集型工作...")
	fmt.Println("请使用以下命令进行 CPU 分析:")
	fmt.Println("go tool pprof http://localhost:6060/debug/pprof/profile")
	fmt.Println("或者通过 web 界面查看: http://localhost:6060/debug/pprof/")

	// 持续执行工作，让 pprof 有足够的数据采样
	for {
		simulateWork()
	}
}
