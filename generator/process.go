package generator

import "fmt"
import "strings"

func ProcessLines(lines []string) {
	recordOn := false
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "record") && strings.HasSuffix(line, "{") {
			recordOn = true
		}
		if strings.HasSuffix(line, "}") {
			recordOn = false
		}
		if recordOn {
			fmt.Println(line)
		}
	}
}
