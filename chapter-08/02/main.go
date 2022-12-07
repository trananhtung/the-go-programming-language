package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jedib0t/go-pretty/v6/table"
)

// https://github.com/Forec/simple-go-ftp/blob/master/server/server.go

var ipAddress = flag.String("ip", "127.0.0.1", "IP address to listen on")
var port = flag.Int("port", 8080, "Port to listen on")
var password string
var bufferLength = 4 * 1024 * 1024 // 4MB

func getRootDir() (string, bool) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		return "", false
	}

	return dir, true
}

func listDir(path string) (info string, ok bool) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return "", false
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "File name", "Mode", "Size(byte)"})

	for i, v := range dir {
		t.AppendRow([]interface{}{i, v.Name(), v.Mode().String(), v.Size()})
		t.AppendSeparator()
	}
	return t.Render(), true

}

func main() {
	fmt.Println(os.Args)
	test, _ := getRootDir()
	info, _ := listDir(test)
	fmt.Println(info)
}
