package promise

const (
	PENDING int = iota
	FULFILLED
	REJECTED
)

type T = interface{} // 空接口代表 any 类型

type Executor = func(func(T), func(error))

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
	executor(p.Resolve, p.Reject)
	return p
}

func (p *Promise) Resolve(value T) {
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

func (p *Promise) Reject(reason error) {
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

func (p *Promise) Then(onFulFilled func(T) T, onRejected func(error) T) *Promise {
	return New(func(resolve func(T), reject func(error)) {
		if p.status == FULFILLED {
			resolve(onFulFilled(p.value))
		} else if p.status == REJECTED {
			resolve(onRejected(p.reason))
		} else {
			p.onFulFilled = append(p.onFulFilled, func() {
				resolve(onFulFilled(p.value))
			})
			p.onRejected = append(p.onRejected, func() {
				resolve(onRejected(p.reason))
			})
		}
	})
}
