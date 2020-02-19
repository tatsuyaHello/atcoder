package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	// 1 <= N <= 2000
	// 1000 <= Y <= 2*10^7
	// Nは整数
	// Yは1000の倍数

	// ありえる場合は「1万円札 5千円札 千円札」の順番でかき
	// 1万円札=a, 5千円札=b, 千円札=c の枚数
	// ありえない場合は、「-1 -1 -1」で書く

	sc.Split(bufio.ScanWords)
	x := make([]int, 2)
	for i := 0; i < 2; i++ {
		x[i] = nextInt()
	}

	N := x[0]
	Y := x[1]

	count := 0
	// N = a + b + c
	// Y = 1000*a + 5000*b + 10000*c
	// Y = 1000*(N-b-c) + 5000*b + 10000*c
	for b := 0; b <= N; b++ {
		for c := 0; c <= N-b; c++ {
			if 1000*(N-b-c)+5000*b+10000*c == Y {
				fmt.Println(c, b, (N - b - c))
				count++
				break
			}
		}
		if count != 0 {
			break
		}
	}
	if count == 0 {
		fmt.Println(-1, -1, -1)
	}
}
