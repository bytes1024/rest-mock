package example

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var min int64 = 1

func TestDoubleAverage(t *testing.T) {
	x := doubleAverage(10, 1000)
	fmt.Println(x)

	x = doubleAverage(9, 1000-x)
	fmt.Println(x)
}

func doubleAverage(count, amount int64) int64 {
	if count <= 0 {
		return 0
	}

	if count == 1 {
		return amount
	}

	max := amount - min*count

	avg := max / count

	avg2 := 2*avg + min
	rand.Seed(time.Now().UnixNano())
	x := rand.Int63n(avg2) + min

	return x
}
