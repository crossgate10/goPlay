package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

var (
	no = flag.String("no", "", "")
)

func main() {
	flag.Parse()

	var tmpl = template.Must(template.ParseFiles("../exp.tmpl"))
	f, err := os.Create("exp.go")
	if err != nil {
		fmt.Println("err: ", err)
	}
	err = tmpl.Execute(f, struct {
		No string
	}{
		No: *no,
	})
	if err != nil {
		fmt.Println("err: ", err)
	}
}
