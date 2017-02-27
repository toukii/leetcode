package main

import (
	"bytes"
	"html/template"
	"io"
	"os"
	// "testing"
	"fmt"
	"io/ioutil"
	"bufio"

	qby "github.com/qiniu/bytes"
	"github.com/everfore/rpcsv"
)

var (
	tpl *template.Template
	data map[string]interface{}
)

func init() {
	tpl, _ = template.New("tpl").Parse("Hello, {{.Content}}.")
	data = make(map[string]interface{})
}

func Stdout(data map[string]interface{})  {
	data["Content"] = "World"
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// Stdout(data)
	// Pipe(data)
	// PipeWriter(data)
	Qiniu()
	Rpcsv()
}

func Rpcsv()  {
	data["Content"]="RPCSV"
	w:=rpcsv.NewBufWriter()
	err:=tpl.Execute(w,data)
	fmt.Println(err)
	fmt.Printf("%v",w.Bytes())
}

func Qiniu()  {
	data["Content"] = "Qiniu"
	buf:=make([]byte,100,100)
	r:=qby.NewReader(buf)
	w:=qby.NewWriter(buf)
	err:=tpl.Execute(w, data)
	fmt.Println("error:",err)
	bs:=r.Bytes()
	fmt.Println(bs)
	fmt.Println("reader bytes:",string(bs),"|")
	fmt.Println("origin bytes:",string(buf),"|")
	data["Content"]="Ya"
	w.Reset()
	tpl.Execute(w,data)
	fmt.Println("reader2 bytes:",string(r.Bytes()),"|",string(w.Bytes()),"|")
}

func PipeWriter(data map[string]interface{})  {
	w := io.PipeWriter{}
	defer w.Close()
	w.Write([]byte("pipeWriter"))
		// err := tpl.Execute(w, data)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		
	
}

func Pipe(data map[string]interface{})  {
	r, w := io.Pipe()
	c := make(chan int, 1)
	data["Content"] = "Golang"
	go write(w, data, c)
	defer r.Close()
	defer w.Close()
	// bufReader(r)
	// ioutilRead(r)
	bytesRead(r)
	// <-c
}

func bytesRead(r io.Reader)  {
	fmt.Println("bytesReader")
	buf:=bytes.NewBuffer(make([]byte,0,100))
	n,err:=buf.ReadFrom(r)
	fmt.Println(n,err)
	fmt.Println(string(buf.Bytes()))
	fmt.Println("end bytesReader")
}

func bufReader(r io.Reader)  {
	var buf = make([]byte, 64)
	n, err := r.Read(buf)
	if err != nil {
		fmt.Printf("read: %v", err)
	}
	fmt.Printf("%q", buf[0:n+6])
}

func ioutilRead(r io.Reader)  {
	fmt.Println("start ioutilRead")
	bs,err:=ioutil.ReadAll(r)
	fmt.Println(string(bs),err)
	fmt.Println("end ioutilRead")
}

func write(w io.Writer, data map[string]interface{}, c chan int) {
	fmt.Println("write start")
	// defer func() {
	// }()
	fmt.Println(data)
// c<-0
	err := tpl.Execute(w, data)
	fmt.Println("tpl.Execute")
	wr:=bufio.NewWriter(w)
	wr.WriteString("\n")
	wr.Flush()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("write over")
}
