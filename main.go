package main

import (
	"awesome-utils/collection/promise"
	"fmt"
	"sync"
	"time"
)

func main() {
	//l := list.New([]int{2, 4, 6, 8})
	//l.ForEach(func(item int, index int) {
	//	fmt.Println("===ForEach", item, index)
	//})
	//res := l.Map(func(item int, index int) int {
	//	fmt.Println("===Map", item, index)
	//	return item + 1
	//}).Filter(func(item int, index int) bool {
	//	fmt.Println("===Filter", item, index)
	//	return item%2 == 0
	//})
	//fmt.Println(res.ToPrimitive())
	wg := new(sync.WaitGroup)

	p := promise.New(func(resolve func(promise.T), reject func(error)) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			resolve(2333)

			// error 是 go 的一种内置类型，本质上是一个接口
			// 通常有两种方式可以定制
			// 一种是通过 errors.New()
			// 另一种是 fmt.Errorf()
			//reject(fmt.Errorf("请求失败"))
		}()
	})

	p.Then(func(res promise.T) promise.T {
		fmt.Println("===success1", res)
		return 666
	}, func(err error) promise.T {
		fmt.Println("===error1", err)
		return nil
	}).Then(func(res promise.T) promise.T {
		fmt.Println("===success2", res)
		return nil
	}, func(err error) promise.T {
		fmt.Println("===error2", err)
		return nil
	})

	// 保证所有子线程都执行完成后再退出 main 函数
	// 如果没有 wg.Wait() 则 main 函数执行完成就直接退出
	wg.Wait()
}
