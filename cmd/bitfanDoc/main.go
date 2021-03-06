package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/clbanning/mxj"
	"github.com/fatih/structs"
	"github.com/k0kubun/pp"
	"github.com/vjeantet/bitfan/processors/doc"
)

var (
	codecStructName     = flag.String("codec", "", "name of BitFan Codec")
	processorStructName = flag.String("processor", "processor", "name of BitFan Processor")
	output              = flag.String("output", "docdoc.go", "output file name; default docdoc.go")
)

func buildProcessorDoc() {
	// récupérer le path
	cwd, _ := os.Getwd()
	var dataJsonDocsPath string = ""

	// charger un doc.Processor
	dataJsonDocsPath = cwd + "../../../docs/data/processors/"
	dp, _ := doc.NewProcessor(cwd)

	JsonFilePath := filepath.Join(dataJsonDocsPath, strings.ToLower(dp.Name)+".json")
	dataMap := mxj.Map(structs.Map(dp))
	jsonBytes, _ := dataMap.JsonIndent("", "  ", true)
	if err := ioutil.WriteFile(filepath.Clean(JsonFilePath), jsonBytes, 0644); err != nil {
		log.Printf("writing output: %v", err)
	}

	g := &Generator{
		buf: bytes.Buffer{},
	}

	// Génère le contenu
	// Print the header and package clause.
	g.Printf("// Code generated by \"bitfanDoc %s\"; DO NOT EDIT\n", strings.Join(os.Args[1:], " "))
	g.Printf("package %s\n\n", dp.Name)
	g.Printf("import \"github.com/vjeantet/bitfan/processors/doc\"\n\n") // Used by all methods.

	g.Printf("func (p *%s) Doc() *doc.Processor {\n\treturn ", *processorStructName)

	w := bufio.NewWriter(&g.buf)
	pp.ColoringEnabled = false
	pp.Fprint(w, dp)
	w.Flush()
	g.Printf("\n}")

	outputName := *output
	outputName = filepath.Join(cwd, strings.ToLower(*output))
	err := ioutil.WriteFile(outputName, g.buf.Bytes(), 0644)
	if err != nil {
		log.Fatalf("writing output: %v", err)
	}

	outputName = filepath.Join(cwd, strings.ToLower("README.md"))
	err = ioutil.WriteFile(outputName, dp.GenMarkdown("logstash"), 0644)
	if err != nil {
		log.Fatalf("writing output: %v", err)
	}

}

func buildCodecDoc(name string) {
	// récupérer le path
	cwd, _ := os.Getwd()
	var dataJsonDocsPath string = ""

	// charger un doc.Codec
	dataJsonDocsPath = cwd + "../../../docs/data/codecs/"
	dc, _ := doc.NewCodec(cwd)
	dc.Name = name

	JsonFilePath := filepath.Join(dataJsonDocsPath, strings.ToLower(dc.Name+".json"))
	dataMap := mxj.Map(structs.Map(dc))
	jsonBytes, _ := dataMap.JsonIndent("", "  ", true)
	if err := ioutil.WriteFile(filepath.Clean(JsonFilePath), jsonBytes, 0644); err != nil {
		log.Printf("writing output: %v", err)
	}
	g := &Generator{
		buf: bytes.Buffer{},
	}

	// Génère le contenu
	// Print the header and package clause.
	g.Printf("// Code generated by \"bitfanDoc %s\"; DO NOT EDIT\n", strings.Join(os.Args[1:], " "))
	g.Printf("package %s\n\n", dc.PkgName)
	g.Printf("import \"github.com/vjeantet/bitfan/processors/doc\"\n\n") // Used by all methods.

	pp.ColoringEnabled = false

	g.Printf("func Doc() *doc.Codec {\n\treturn ")
	w := bufio.NewWriter(&g.buf)
	pp.Fprint(w, dc)
	w.Flush()
	g.Printf("\n}")

	outputName := *output
	outputName = filepath.Join(cwd, strings.ToLower(*output))
	err := ioutil.WriteFile(outputName, g.buf.Bytes(), 0644)
	if err != nil {
		log.Fatalf("writing output: %v", err)
	}

}

func main() {
	flag.Parse()

	if *codecStructName != "" {
		buildCodecDoc(*codecStructName)
	} else {
		buildProcessorDoc()
	}
}

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	buf bytes.Buffer // Accumulated output.
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}
