package concurrent

import "fmt"

var c chan int

func printA() {
	c <- 1
	fmt.Println("A")
}

func printB() {
	_ = <-c
	fmt.Println("B")
}

func main() {
	go printA()
	go printB()
}



func read(urlChan <- chan string) {
	for {
		data := <- urlChan
		// do http()
	}
}

func write(urlChan string <- chan) {
	for {
		// 锁文件
		a := 0
		urlChan <- a
	}
}

func merge(a,b <- chan int) chan int {
	result := make(chan int)
	go func() {
		v1, ok1 := <- a
		v2, ok2 := <- b
		for v1 || v2 {
			if !ok2 || (ok1 && v1 < v2) {
				result <- v1
			} else {
				result <- v2
			}
		}
		close(result)
	}()
	return result
}

func p(c chan <- int) {
	for i:=0; i < 10; i++ {
		c <- i
	}
}

func p(c <- chan int) {
	for i:=0; i < 10; i++ {
		fmt.Println( <- c)
	}
}
