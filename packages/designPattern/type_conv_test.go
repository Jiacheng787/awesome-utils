package designPattern

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestTypeConv(t *testing.T) {
	x := "煎鱼"
	var v interface{} = x
	fmt.Println(v)
}

func TestSortMethod(t *testing.T) {
	// 初始化空 slice
	var ms []int
	// 初始化随机数种子
	// 如果这一步不执行，则每次生成随机数序列都是一样的
	// 例如 [81 87 47 59 81 18 25 40 56 0]
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		// rand.Intn(n) 用于生成 [0,n) 范围的 int 类型伪随机数
		ms = append(ms, rand.Intn(100))
	}
	log.Println(ms)
	sort.Ints(ms)
	log.Println(ms)
}
