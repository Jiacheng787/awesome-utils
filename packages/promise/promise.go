package promise

import (
	"fmt"
)

const (
	PENDING int = iota
	FULFILLED
	REJECTED
)

type T interface{} // 空接口代表 any 类型

type Executor func(func(T), func(error))

// 结构体属性名开头大写相当于 public，可以在模块外部访问
// 属性名开头小写相当于 private，无法在模块外部访问

type Promise struct {
	status      int
	value       T
	reason      error
	onFulFilled []func()
	onRejected  []func()
}

func New(executor Executor) *Promise {
	p := &Promise{
		status:      PENDING,
		onFulFilled: make([]func(), 0),
		onRejected:  make([]func(), 0),
	}
	executor(p._resolve, p._reject)
	return p
}

func (p *Promise) _resolve(value T) {
	if p.status != PENDING {
		return
	}
	p.status = FULFILLED
	p.value = value
	for _, callback := range p.onFulFilled {
		callback()
	}
	p.onFulFilled = p.onFulFilled[:0]
}

func (p *Promise) _reject(reason error) {
	if p.status != PENDING {
		return
	}
	p.status = REJECTED
	p.reason = reason
	for _, callback := range p.onRejected {
		callback()
	}
	p.onRejected = p.onRejected[:0]
}

//func instanceOf(a, b interface{}) bool {
//	return reflect.TypeOf(a) == reflect.TypeOf(b)
//}

func resolvePromise(p *Promise, result T, resolve func(T), reject func(error)) {
	if p == result {
		// 构造 error 对象，也可以用 errors.New()
		reject(fmt.Errorf("chaining cycle detected for promise #<Promise>"))
		return
	}
	// Golang 类型断言两种用法
	// 第一种 value, ok := result.(*Promise)
	// 注意这边需要接受参数 ok
	// 否则断言失败会直接抛出 panic
	//if value, ok := result.(*Promise); ok {
	//	value.Then(func(res T) T {
	//		resolve(res)
	//		return nil
	//	}, func(err error) T {
	//		reject(err)
	//		return nil
	//	})
	//} else {
	//	resolve(result)
	//}

	// 第二种配合 switch 使用
	switch value := result.(type) {
	case *Promise:
		value.Then(func(res T) T {
			resolve(res)
			return nil
		}, func(err error) T {
			reject(err)
			return nil
		})
	default:
		resolve(value)
	}
}

func (p *Promise) Then(onFulFilled func(T) T, onRejected func(error) T) *Promise {
	return New(func(resolve func(T), reject func(error)) {
		if p.status == FULFILLED {
			resolvePromise(p, onFulFilled(p.value), resolve, reject)
		} else if p.status == REJECTED {
			resolvePromise(p, onRejected(p.reason), resolve, reject)
		} else {
			p.onFulFilled = append(p.onFulFilled, func() {
				resolvePromise(p, onFulFilled(p.value), resolve, reject)
			})
			p.onRejected = append(p.onRejected, func() {
				resolvePromise(p, onRejected(p.reason), resolve, reject)
			})
		}
	})
}

func Resolve(value T) *Promise {
	return New(func(resolve func(T), reject func(error)) {
		resolve(value)
	})
}
