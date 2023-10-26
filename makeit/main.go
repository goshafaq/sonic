package main

import (
	"fmt"
	//_ "github.com/go-shafaq/encoding-json"
	"github.com/goshafaq/sonic"
	"io/fs"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	//MakeIt("./")
	CheckIt()
}

func MakeIt(rootPath string) {
	fsys := os.DirFS(rootPath)
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if (!d.IsDir()) && AvailableFile(path) {
			println(path)
			ReplaceAll(rootPath + path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("\n===Done===\n")
}

const (
	standsJsonPkg = `encoding/json`
	customJsonPkg = `github.com/go-shafaq/encoding-json`
	go_shafaqPkg  = `github.com/go-shafaq/sonic`
	bytedancePkg  = `github.com/bytedance/sonic`
	currFile      = `makeit/main.go`
)

func ReplaceAll(path string) {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	output := strings.ReplaceAll(string(input), `bytedance`, `goshafaq`)
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		panic(err)
	}
}

var notAvailableRgx = regexp.MustCompile(`/?\..*/`)

func AvailableFile(path string) bool {
	available := !notAvailableRgx.MatchString(path)
	if !available {
		return false
	}

	return !strings.Contains(path, currFile)
}

func CheckIt() {
	//sonic.DefCase(func(tag string) func(pkgPath string, name string) string {
	//	return func(string, name string) string {
	//		return strings.ToUpper(name)
	//	}
	//})
	data, err := sonic.Marshal(&struct {
		UserId int
	}{13})
	if err != nil {
		panic(err)
	}

	println(string(data))
}
