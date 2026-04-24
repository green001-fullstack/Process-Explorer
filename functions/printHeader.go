package functions

import (
	"fmt"
	"strings"
)

func PrintHeader(){
	fmt.Printf("%-7s %-7s %-7s %-9s %s\n",
    "PID", "PPID", "STATE", "MEM(KB)", "COMMAND")
	fmt.Println(strings.Repeat("-", 70))
}