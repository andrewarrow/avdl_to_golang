package generator

import "io/ioutil"
import "strings"

func WriteSchemaDotGo(pname string, items []string) {
	content := "package " + pname + "\n\n"
	content = content + "var schema map[string]fields = map[string]fields{" +
		strings.Join(items, "\n") + "\n}"
	ioutil.WriteFile("schema.go", []byte(content), 0644)
}
