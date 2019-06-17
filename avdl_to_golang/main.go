package main

import "io/ioutil"
import "fmt"
import "strings"
import "os"

import "github.com/andrewarrow/avdl_to_golang/generator"

func main() {

	if len(os.Args) < 3 {
		fmt.Println("./avdl_to_golang <dir> <package_name>")
		return
	}

	dir := os.Args[1]
	pname := os.Args[2]
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	items := []string{}
	struct_items := []string{}
	for _, f := range files {
		schema, err := ioutil.ReadFile(dir + "/" + f.Name())
		if err != nil {
			fmt.Printf("Error reading file %q - %v\n", f.Name(), err)
			return
		}
		lines := strings.Split(string(schema), "\n")
		items = append(items, generator.ProcessLines(lines))
		struct_items = append(struct_items, generator.ProcessLinesForStructs(lines))
	}
	generator.WriteSchemaDotGo(pname, items, struct_items)
}
