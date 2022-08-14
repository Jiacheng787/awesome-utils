package file

import (
	"encoding/json"
	"io/fs"
	"log"
	"os"
	"path"
	"testing"
)

var externals = [...]string{
	"react",
	"react-dom",
	"antd",
	"clsx",
}

type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Private bool   `json:"private"`
	//Scripts         map[string]string `json:"scripts"`
	Dependencies map[string]string `json:"dependencies"`
	//DevDependencies  map[string]string `json:"devDependencies"`
	PeerDependencies map[string]string `json:"peerDependencies"`
}

func TestGetCwd(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		// log.Fatal() 相当于先调用 Print() 再调用 os.Exit(1)
		// 因此无需再加 return 语句
		log.Fatal(err)
	}

	absPath := path.Join(cwd, "package.json")
	log.Println(absPath)

	// 注意 ioutil.ReadFile() 方法已经废弃了
	// 直接用 os.ReadFile() 替代
	content, err := os.ReadFile(absPath)
	if err != nil {
		log.Fatal(err)
	}

	// 注意两种初始化对象方式的区别：
	// var pkg *Package 初始化一个空指针 nil
	// &Package{} 初始化一个空对象指针，内部属性都是初始值

	// 以下几种写法都可以：
	// var pkg interface{}
	// err := json.Unmarshal(content, &pkg)
	// var pkg Package
	// err := json.Unmarshal(content, &pkg)
	// var pkg *Package
	// err := json.Unmarshal(content, &pkg)
	// pkg := &Package{}
	// err := json.Unmarshal(content, &pkg)

	var pkg Package
	if err := json.Unmarshal(content, &pkg); err != nil {
		log.Fatal(err)
	}
	//log.Println(pkg)

	// 原始类型、map、slice 值传递
	// 实际上 map 和 slice 存在内部指针，不需要手动获取指针
	// struct、array 引用传递
	// 由于 array 用的很少，因此指针主要就是 struct

	// 创建 slice：
	// var s []int
	// 初始化 slice：
	// s = make([]int, 0)
	// 初始化并赋值：
	// s = []int{1, 2}

	// 创建 array：
	// var a [10]int
	// 初始化 array：
	// a = [2]int{1, 2}
	// 让编译器自动推断数组长度：
	// a = [...]int{1, 2}

	// map 有可能没有初始化
	// 例如 var m map[string]string
	// 未初始化的 map 零值是 nil，与一个空 map 基本等价，打印出来也是 map[]
	// 未初始化的 map 取值不会报错，但是添加值会报错（注意未初始化的 slice 通过 append 可正常添加）
	// 需要通过 make 方法分配确定的内存地址进行初始化
	// 例如 m := make(map[string]string)
	if pkg.PeerDependencies == nil {
		pkg.PeerDependencies = make(map[string]string)
	}

	dependencies := pkg.Dependencies
	peerDependencies := pkg.PeerDependencies

	for _, dep := range externals {
		temp := dependencies[dep]
		peerDependencies[dep] = temp
		delete(dependencies, dep)
	}
	bytes, err := json.Marshal(pkg)
	if err != nil {
		log.Fatal(err)
	}

	output := path.Join(cwd, "_package")
	err = os.WriteFile(output, bytes, fs.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
