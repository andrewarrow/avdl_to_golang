package generator

import "io/ioutil"
import "strings"

func WriteSchemaDotGo(items []string) {
	content := "var schema map[string]fields = map[string]fields{" +
		strings.Join(items, "\n") + "\n}"
	ioutil.WriteFile("schema.go", []byte(content), 0644)
}
