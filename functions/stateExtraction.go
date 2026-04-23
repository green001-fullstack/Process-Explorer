package functions

import(
	"strings"
)

func StateExtraction(str string) string {
	word := strings.Fields(str)
	if len(word) > 0 {
		return word[0]
	}
	return ""
}