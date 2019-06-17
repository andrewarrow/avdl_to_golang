package generator

import "io/ioutil"
import "strings"

func WriteSchemaDotGo(pname string, items []string, struct_items []string) {
	fields := `
	import "strings"

	type Fields struct {
		stringFields []string
		floatFields  []string
		longFields   []string
	}
	type ValueFields struct {
		stringFields []string
		floatFields  []float32
		longFields   []int64
		name         string
	}


	func newFields(m map[string]string) Fields {
		f := Fields{}
		if m["strings"] != "" {
			for _, s := range strings.Split(m["strings"], ",") {
				f.stringFields = append(f.stringFields, s)
			}
		}
		if m["floats"] != "" {
			for _, s := range strings.Split(m["floats"], ",") {
				f.floatFields = append(f.floatFields, s)
			}
		}
		if m["longs"] != "" {
			for _, s := range strings.Split(m["longs"], ",") {
				f.longFields = append(f.longFields, s)
			}
		}
		return f
	}
	`
	content := "package " + pname + "\n\n"
	content = content + "// Code generated by github.com/andrewarrow/avdl_to_golang. DO NOT EDIT.\n\n"

	content = content + fields
	content = content + "var schema map[string]Fields = map[string]Fields{" +
		strings.Join(items, "\n") + "\n}\n"
	content = content + strings.Join(struct_items, "\n")
	ioutil.WriteFile("schema.go", []byte(content), 0644)
}
