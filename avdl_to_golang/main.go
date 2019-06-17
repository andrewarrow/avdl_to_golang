package main

import "io/ioutil"
import "fmt"
import "strings"

import "github.com/andrewarrow/avdl_to_golang/generator"

func main() {

	dir := "samples"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range files {
		schema, err := ioutil.ReadFile(dir + "/" + f.Name())
		if err != nil {
			fmt.Printf("Error reading file %q - %v\n", f.Name(), err)
			return
		}
		lines := strings.Split(string(schema), "\n")
		generator.ProcessLines(lines)
	}
}
