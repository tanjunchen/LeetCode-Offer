package main

import (
	"fmt"
	"sync"
	"time"
)

// 需求如下 A、B 两个协程分别打印 1、2、3、4 和 A，B，C，D

func test1() {
	channelA := make(chan bool, 1)
	channelB := make(chan bool)
	exit := make(chan bool)
	go func() {
		arr := []int{1, 2, 3, 4}
		for i := 0; i < len(arr); i++ {
			if ok := <-channelA; ok {
				println(arr[i])
				channelB <- true
			}
		}
	}()

	go func() {
		defer func() {
			exit <- true
		}()
		arr := []string{"A", "B", "C", "D"}
		for i := 0; i < len(arr); i++ {
			if ok := <-channelB; ok {
				println(arr[i])
				channelA <- true
			}
		}
	}()

	channelA <- true
	<-exit
}

// 定义 A、B 两个 channal，开 A、B 两个协程，A 协程输出奇数、B 协程输出偶数，通过两个独立的 channal 控制顺序，交替输出。

func test2() {
	channelA := make(chan bool, 1)
	channelB := make(chan bool)
	exit := make(chan bool)
	go func() {
		for i := 1; i <= 10; i++ {
			if ok := <-channelA; ok {
				println(2*i - 1)
				channelB <- true
			}
		}
	}()

	go func() {
		defer func() {
			exit <- true
		}()
		for i := 1; i <= 10; i++ {
			if ok := <-channelB; ok {
				println(2 * i)
				channelA <- true
			}
		}
	}()

	channelA <- true
	<-exit
}

// 定义 A、B、C 三个 channal，开 A、B、C 三个协程，依次输出 A B C 交替输出 10 次
// channel 实现版本
func test3() {
	channelA := make(chan bool, 1)
	channelB := make(chan bool)
	channelC := make(chan bool)
	exit := make(chan bool)
	go func() {
		for i := 0; i < 3; i++ {
			if ok := <-channelA; ok {
				println("A")
				channelB <- true
			}
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			if ok := <-channelB; ok {
				println("B")
				channelC <- true
			}
		}
	}()

	go func() {
		defer func() {
			close(exit)
		}()
		for i := 0; i < 3; i++ {
			if ok := <-channelC; ok {
				println("C")
				channelA <- true
			}
		}
	}()

	channelA <- true
	<-exit
}

func test4() {
	var l = sync.Mutex{}
	var cond = sync.NewCond(&l)
	var turn = 0
	s := []string{"A", "B", "C"}
	for _, item := range s {
		go func(x string) {
			for i := 0; i < 10; i++ {
				l.Lock()
				for s[turn] != x {
					cond.Wait()
				}
				// my turn
				fmt.Println(i, x)
				turn = (turn + 1) % 3
				cond.Broadcast()
				cond.Wait()
				l.Unlock()
			}
		}(item)
	}
	time.Sleep(time.Hour)
}

func test5() {
	var mutex = new(sync.Mutex)
	var i = 0
	var count = 3
	var wg sync.WaitGroup
	wg.Add(3)
	var printMsg func(mutex *sync.Mutex, wg *sync.WaitGroup, content string, divNum int)
	printMsg = func(mutex *sync.Mutex, wg *sync.WaitGroup, content string, divNum int) {
		for ; i <= count; {
			mutex.Lock()
			if i%3 == divNum {
				fmt.Println(content)
				i++
			}
			mutex.Unlock()
		}
		wg.Done()
	}

	go printMsg(mutex, &wg, "A", 0)
	go printMsg(mutex, &wg, "B", 1)
	go printMsg(mutex, &wg, "C", 2)
	wg.Wait()
}

var m = new(sync.Mutex)

var i = 0
var max = 3

func Print(content string, divNum int, wg *sync.WaitGroup) {
	for ; i <= max; {
		m.Lock()
		if i%3 == divNum {
			fmt.Println(content)
			i++
		}
		m.Unlock()
		// 添加 runtime.Gosched(),主动让出 CPU
		// runtime.Gosched()
	}
	wg.Done()
}

func test6() {
	var wg sync.WaitGroup
	wg.Add(3)

	go Print("A", 0, &wg)
	go Print("B", 1, &wg)
	go Print("C", 2, &wg)

	wg.Wait()
}

var num int = 2

// 用两个chan进行通知
func Print2(content string, ch chan int, ch2 chan int, wg *sync.WaitGroup) {
	if "A" == content {
		for i := 0; i <= num; i++ {
			<-ch2
			if i < num {
				fmt.Println(content)
				ch <- i
			}
		}
	} else {
		for i := 0; i < num; i++ {
			<-ch2
			fmt.Println(content)
			ch <- i
		}
	}
	wg.Done()
}

func test7() {
	var wg sync.WaitGroup
	wg.Add(3)

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go Print2("A", ch1, ch3, &wg)
	go Print2("B", ch2, ch1, &wg)
	go Print2("C", ch3, ch2, &wg)

	// 先触发
	ch3 <- 1

	wg.Wait()
}

func gen(v string, times int) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < times; i++ {
			ch <- v
		}
	}()
	return ch
}

func fanIn(times int, inputs []<-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < times; i++ {
			for _, input := range inputs {
				v := <-input
				ch <- v
			}
		}
	}()
	return ch
}

// Go 中的扇入方式 通用的解决方式
func test8() {
	M, N := 5, 2
	times := M
	inputs := make([]<-chan string, 0, N)
	for i := 0; i < N; i++ {
		threadName := string('A' + i)
		inputs = append(inputs, gen(threadName, times))
	}
	for char := range fanIn(times, inputs) {
		fmt.Print(char)
	}
}

func test9() {
	goroutineNum := 5
	count := 2
	inputs := make([]<-chan string, 0, goroutineNum)
	for i := 0; i < goroutineNum; i++ {
		inputs = append(inputs, generateGoroutine(string('A'+i), count))
	}

	for i := range print9(count, inputs) {
		fmt.Print(i)
	}
}

func print9(count int, inputs []<-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < count; i++ {
			for _, input := range inputs {
				v := <-input
				ch <- v
			}
		}
	}()
	return ch
}

func generateGoroutine(s string, count int) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < count; i++ {
			ch <- s
		}
	}()
	return ch
}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()
	// test7()
	// test8()
	test9()
}
