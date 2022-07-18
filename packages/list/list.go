package list

type List struct {
	arr []int
}

func New(arr []int) *List {
	ins := &List{arr: arr}
	return ins
}

func (l *List) ForEach(fn func(int, int)) {
	for index, item := range l.arr {
		fn(item, index)
	}
}

func (l *List) Map(fn func(int, int) int) *List {
	res := make([]int, 0)
	for index, item := range l.arr {
		res = append(res, fn(item, index))
	}
	return &List{arr: res}
}

func (l *List) Filter(fn func(int, int) bool) *List {
	res := make([]int, 0)
	for index, item := range l.arr {
		if fn(item, index) {
			res = append(res, item)
		}
	}
	return &List{arr: res}
}

func (l *List) ToPrimitive() []int {
	return l.arr
}
