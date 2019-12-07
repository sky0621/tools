package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

const (
	max          = 50000
	templatePath = "./json.tmpl"
	outputPath   = "./data.json"
)

func abs(path string) string {
	p, err := filepath.Abs(filepath.Clean(path))
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func main() {
	fmt.Println("Start")

	var items []*Item
	for i := 0; i < max; i++ {
		comma := ""
		if i < max-1 {
			comma = ","
		}
		items = append(items, &Item{No: i, Comma: comma})
	}
	fmt.Println(len(items))

	tpath := abs(templatePath)
	fmt.Printf("templatePath: %s\n", tpath)

	tmpl := template.Must(template.ParseFiles(tpath))
	buf := &bytes.Buffer{}
	err := tmpl.Execute(buf, &Data{Items: items})
	if err != nil {
		log.Fatal(err)
	}

	opath := abs(outputPath)
	fmt.Printf("outputPath: %s\n", opath)

	file, err := os.OpenFile(opath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(file)
	w.WriteString(buf.String())
	w.Flush()
}

type Data struct {
	Items []*Item
}

type Item struct {
	No    int
	Comma string
}
