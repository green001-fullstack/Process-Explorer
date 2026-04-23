package functions

import(
	"strings"
	"strconv"
	"fmt"
)

func UidValue(value string) int {
	result := strings.Fields(value)
	if len(result) > 0 {
		val, err := strconv.Atoi(result[0])
		if err != nil {
			fmt.Println("Error parsing Uid value with strconvAtoi", err)
			return 0
		}
		return val
	}
	return 0
}