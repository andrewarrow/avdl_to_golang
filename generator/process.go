package generator

import "fmt"
import "strings"

type FieldsAndName struct {
	f_strings  []string
	f_floats   []string
	f_longs    []string
	recordName string
}

func GetFieldsAndName(lines []string) FieldsAndName {
	recordOn := false
	recordName := ""
	f_strings := []string{}
	f_floats := []string{}
	f_longs := []string{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasSuffix(line, "}") {
			recordOn = false
		}
		if recordOn {
			tokens := strings.Split(line, " ")
			dataType := tokens[0]
			name := tokens[1][0 : len(tokens[1])-1]
			if dataType == "string" {
				f_strings = append(f_strings, name)
			} else if dataType == "float" {
				f_floats = append(f_floats, name)
			} else if dataType == "long" {
				f_longs = append(f_longs, name)
			}
		}
		if strings.HasPrefix(line, "record") && strings.HasSuffix(line, "{") {
			tokens := strings.Split(line, " ")
			recordName = tokens[1]
			recordOn = true
		}
	}
	f := FieldsAndName{}

	f.f_longs = f_longs
	f.f_strings = f_strings
	f.f_floats = f_floats
	f.recordName = recordName
	return f
}

func ProcessLines(lines []string) string {
	f := GetFieldsAndName(lines)

	theFields := `"strings": "%s", "floats": "%s", "longs": "%s"`
	filledIn := fmt.Sprintf(theFields, strings.Join(f.f_strings, ","),
		strings.Join(f.f_floats, ","), strings.Join(f.f_longs, ","))

	content := fmt.Sprintf("\"%s\": newFields(map[string]string{%s}),",
		f.recordName, filledIn)

	return content
}
func ProcessLinesForStructs(lines []string) string {

	f := GetFieldsAndName(lines)

	content := fmt.Sprintf("\n\ntype %s struct {\n", f.recordName)

	for _, f := range f.f_strings {
		content = content + fmt.Sprintf("  %s string\n", f)
	}
	for _, f := range f.f_floats {
		content = content + fmt.Sprintf("  %s float32\n", f)
	}
	for _, f := range f.f_longs {
		content = content + fmt.Sprintf("  %s int64\n", f)
	}
	content = content + fmt.Sprintf("}\n")

	return content
}
