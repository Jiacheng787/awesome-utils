package file

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"
)

type Package struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Private         bool              `json:"private"`
	Scripts         map[string]string `json:"scripts"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
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

	content, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Fatal(err)
	}

	// 注意两种初始化对象方式的区别：
	// var pkg *Package 初始化一个空指针 nil
	// &Package{} 初始化一个空对象指针，内部属性都是初始值
	pkg := &Package{}
	if err := json.Unmarshal(content, pkg); err != nil {
		log.Fatal(err)
	}
	log.Println(pkg)
	log.Printf("version: %s", pkg.Version)
	log.Printf("dependencies: %v", pkg.Dependencies)
	for item := range pkg.Dependencies {
		log.Printf("key: %s, value: %s", item, pkg.Dependencies[item])
	}
}
