package generator

import "fmt"
import "strings"

func ProcessLines(lines []string) {
	recordOn := false
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasSuffix(line, "}") {
			recordOn = false
		}
		if recordOn {
			tokens := strings.Split(line, " ")
			dataType := tokens[0]
			name := tokens[1][0 : len(tokens[1])-1]
			fmt.Println(name, dataType)
		}
		if strings.HasPrefix(line, "record") && strings.HasSuffix(line, "{") {
			recordOn = true
		}
	}
}
