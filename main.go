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

	p := promise.New(func(resolve func(interface{}), reject func(err error)) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			//resolve(2333)
			reject(fmt.Errorf("请求失败"))
		}()
	})

	p.Then(func(res interface{}) interface{} {
		fmt.Println("===success1", res)
		return 666
	}, func(err error) interface{} {
		fmt.Println("===error1", err)
		return 0
	}).Then(func(res interface{}) interface{} {
		fmt.Println("===success2", res)
		return 0
	}, func(err error) interface{} {
		fmt.Println("===error2", err)
		return 0
	})

	// 保证所有子线程都执行完成后再退出 main 函数
	// 如果没有 wg.Wait() 则 main 函数执行完成就直接退出
	wg.Wait()
}
