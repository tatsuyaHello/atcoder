package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type point struct {
	X int
	Y int
}

func main() {
	// bufio.ScanWordsは空白区切
	sc.Split(bufio.ScanWords)

	n := nextInt()
	points := make([]point, n)

	for i := 0; i < len(points); i++ {
		points[i].X = nextInt()
		points[i].Y = nextInt()
	}

	dmax := float64(0)

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if dtemp := dis(points[i], points[j]); dmax < dtemp {
				dmax = dtemp
			}
		}
	}
	fmt.Println(dmax)
}

func dis(p1, p2 point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	a, _ := strconv.Atoi(next())
	return a
}
