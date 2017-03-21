package golang

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T) {

	c := make(chan string, 1)

	go func() {
		for {
			select {
			case v, ok := <-c:
				if !ok {
					fmt.Println("1:close")
					return
				}
				fmt.Println("1:" + v)
			default:
			}
		}

	}()

	go func() {
		for {
			select {
			case v, ok := <-c:
				if !ok {
					fmt.Println("2:close")
					return
				}
				fmt.Println("2:" + v)
			default:
			}
		}

	}()

	go func() {
		close(c)
	}()

	select {
	case <-c:
	}
}
