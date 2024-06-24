package test

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed test/ss.png
var logo []byte

func TestByte(t *testing.T) {

	err := os.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed test/files/a.txt
//go:embed test/files/b.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))
}

//go:embed ../files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dir, _ := os.ReadDir("files")

	for _, entry := range dir {
		if !entry.IsDir() {
			file, err := path.ReadFile("files/" + entry.Name())
			if err != nil {
				panic(err)
			}
			fmt.Println(string(file))
		}
	}
}
