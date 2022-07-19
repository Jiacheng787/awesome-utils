# awesome-utils

⭐️ JS wheels that implements in golang

Inspired by [wheel-awesome](https://github.com/su37josephxia/wheel-awesome)

## 1. Promise

由于 Golang 不支持面向对象，所以这里手动模拟 `new` 操作：

```go
func New(executor Executor) *Promise {
    p := &Promise{
        status:      PENDING,
        onFulFilled: make([]func(), 0),
        onRejected:  make([]func(), 0),
    }
    executor(p._resolve, p._reject)
    return p
}
```

使用 Golang 协程模拟 JS 事件循环，正常情况下主线程代码执行完毕，进程直接退出，不会等待子线程执行完毕。这边使用 `WaitGroup` ，保证所有子线程都执行完成后再退出 `main` 函数：

```go
wg := new(sync.WaitGroup)

p := promise.New(func(resolve func(promise.T), reject func(error)) {
    wg.Add(1)
    go func() {
        defer wg.Done()
        time.Sleep(1 * time.Second)
        resolve(2333)
    }()
})

// 保证所有子线程都执行完成后再退出 main 函数
// 如果没有 wg.Wait() 则 main 函数执行完成就直接退出
wg.Wait()
```

## 2. Array Methods

