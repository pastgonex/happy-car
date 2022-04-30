package main

import (
	"context"
	"fmt"
	"time"
)

type paramKey struct{}

func main() {
	// 使用paramKey作为key，将参数传递给context， 好处是不用每次取值的时候拼一遍"param"
	c := context.WithValue(context.Background(),
		paramKey{}, "abc") // 带有abc参数的context
	c, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	// 同一个任务，不同的context，可以获取到不同的参数
	mainTask(c)

	// 手动cancel
	//var cmd string
	//for {
	//	fmt.Scan(&cmd)
	//	if cmd == "c" {
	//		cancel()
	//		break
	//	}
	//}
}

func mainTask(c context.Context) {
	fmt.Printf("main task started with param %q\n", c.Value(paramKey{}))

	//c1, cancel := context.WithTimeout(c, 2*time.Second)
	//defer cancel()

	//smallTask(context.Background(), "task1") // 子任务， 当前任务的后台任务 == 子任务
	//go smallTask(context.Background(), "task1", 4*time.Second) // 这样总共的context给的时间为 5s， 但是task1用了4s， 而task完成任务的时间为2s， 因此在5s的时候就超时了， 不能继续执行
	//go smallTask(c1, "task1", 9*time.Second) // 完不成是因为 上面defer cancel()了， 解决方法是开一个go func
	go func() {
		c1, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		smallTask(c1, "task1", 9*time.Second)
	}()
	smallTask(c, "task2", 2*time.Second)
}

func smallTask(c context.Context, name string, d time.Duration) {
	fmt.Printf("%s started with param %q\n", name, c.Value(paramKey{}))
	select {
	case <-time.After(d):
		fmt.Printf("%s done\n", name)
	case <-c.Done(): // 可以不理会， 然后一直做下去. 收到值了说明 超时了（时间到了）
		fmt.Printf("%s cancelled\n", name)
	}
}
