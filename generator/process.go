package generator

import "fmt"
import "strings"

func ProcessLines(lines []string) string {
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

	theFields := `"strings": "%s", "floats": "%s", "longs": "%s"`
	filledIn := fmt.Sprintf(theFields, strings.Join(f_strings, ","),
		strings.Join(f_floats, ","), strings.Join(f_longs, ","))

	content := fmt.Sprintf("\"%s\": newFields(map[string]string{%s}),",
		recordName, filledIn)

	return content
}
