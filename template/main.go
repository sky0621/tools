package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func abs(path string) string {
	p, err := filepath.Abs(filepath.Clean(path))
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func main() {
	var count *int = flag.Int("c", 10, "データ件数")
	var templatePath *string = flag.String("t", "./json.tmpl", "テンプレートファイルパス")
	var outputPath *string = flag.String("o", "./data.json", "生成JSONパス")
	flag.Parse()

	fmt.Println("Start")
	fmt.Printf("[Args][count:%d][templatePath:%s][outputPath:%s]\n", *count, *templatePath, *outputPath)

	var items []*Item
	for i := 1; i < *count+1; i++ {
		comma := ""
		if i < *count {
			comma = ","
		}
		items = append(items, &Item{No: i, Comma: comma})
	}
	fmt.Printf("item length: %d\n", len(items))

	tpath := abs(*templatePath)
	fmt.Printf("templatePath: %s\n", tpath)

	tmpl := template.Must(template.ParseFiles(tpath))
	buf := &bytes.Buffer{}
	err := tmpl.Execute(buf, items)
	if err != nil {
		log.Fatal(err)
	}

	opath := abs(*outputPath)
	_, err = os.Stat(opath)
	if !os.IsNotExist(err) {
		if err := os.Remove(opath); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("outputPath: %s\n", opath)

	file, err := os.OpenFile(opath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(file)
	w.WriteString(buf.String())
	w.Flush()

	fmt.Println("End")
}

type Item struct {
	No    int
	Comma string
}
