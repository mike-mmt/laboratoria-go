package funkcje

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func TestFib() {
	prev, err := time.ParseDuration("1ns")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 40; i += 1 {
		start := time.Now()
		Fibonacci(i)
		duration := time.Since(start)
		fmt.Printf("Fib(%d) : %15v | %.2f times slower than previous\n", i, duration, float64(duration.Nanoseconds())/float64(prev.Nanoseconds()))
		prev = duration

	}
}

func MeanFib() {
	var sum time.Duration = 0
	count := 10000
	for i := 0; i < count; i++ {
		fmt.Print("\n\033[1A\033[K")
		numberOfDashes := int(math.Abs(float64(i / (count / 100))))
		fmt.Printf("[%s%s]", strings.Repeat("-", numberOfDashes), strings.Repeat(" ", 99-numberOfDashes))
		start := time.Now()
		Fibonacci(30)
		duration := time.Since(start)
		sum += duration
	}
	fmt.Printf("\nMean time for Fibonacci(30) is %v\n", sum/time.Duration(count))
}
