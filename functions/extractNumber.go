package functions

import(
	"strconv"
	"strings"
)

func ExtractNumber(value string) int {
	valueField := strings.Fields(value)
	if len(valueField) > 0 {
		number, _ := strconv.Atoi(valueField[0])
		return number
	}
	return 0
}