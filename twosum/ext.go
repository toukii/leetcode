package main

import (
	// "bytes"
	"html/template"
	"io"
	"os"
	// "testing"
	"fmt"
	"io/ioutil"
)

var (
	tpl *template.Template
)

func init() {
	tpl, _ = template.New("tpl").Parse("Hello, {{.Content}}.")
}

func main() {
	data := make(map[string]interface{})
	data["Content"] = "World"
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println(err)
	}
	r, w := io.Pipe()
	defer r.Close()
	// pw := io.PipeWriter{}
	go func() {
		err := tpl.Execute(w, data)
		if err != nil {
			fmt.Println(err)
		}
		w.Write([]byte("......."))
	}()
	defer w.Close()
	// io.Copy(os.Stdout, w)
	bs, err := ioutil.ReadAll(r)
	fmt.Println(string(bs), err)

}
