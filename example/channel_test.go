package example

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {

	ch := make(chan interface{}, 10)

	ch2 := make(chan interface{})

	go func() {
		tk := time.NewTicker(time.Second * 3)
		for range tk.C {
			ch <- strconv.Itoa(rand.Intn(100)) + "  hello"
		}
	}()

	go func() {
		tk := time.NewTicker(time.Second * 10)
		for range tk.C {
			ch2 <- "refresh"
		}
	}()

	for {
		select {
		case str := <-ch:
			fmt.Println("接收数据:", str)
		case str := <-ch2:
			fmt.Println("refresh: ", str)
		default:
		}
	}

}
