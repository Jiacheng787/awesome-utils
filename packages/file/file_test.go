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

	pkg := &Package{}
	err = json.Unmarshal(content, pkg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pkg)
	log.Printf("version: %s", pkg.Version)
	log.Printf("dependencies: %v", pkg.Dependencies)
	for item := range pkg.Dependencies {
		log.Printf("key: %s, value: %s", item, pkg.Dependencies[item])
	}
}
