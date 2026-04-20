package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// YOU define what fields this needs
// based on what you saw in /proc/PID/status
type Process struct {
	Name string
	PID int
	PPID int
	State string
	Threads int
	VmSize int
	VmRSS int
	UID int
	Command string
}

func stateExtraction(str string) string{
	word := strings.Fields(str)
	if len(word) > 0{
		return word[0]
	}
	return ""
}

func extractNumber(value string) int{
	valueField := strings.Fields(value) 
	if len(valueField) > 0 {
		number, _ := strconv.Atoi(valueField[0])
		return number
	}
	return 0
}

func uidValue(value string) int{
	result := strings.Fields(value)
	if len(result) > 0 {
		val, err := strconv.Atoi(result[0])
		if err != nil{
			fmt.Println("Error parsing Uid value with strconvAtoi", err)
			return 0
		}
		return val
	}
	return 0
}

func PrintProcess(p *Process){
	// Handle empty command (e.g., kernel threads)
    command := p.Command
    if command == "" {
        command = "[kernel]"
    }

    fmt.Printf("Name:		%s\n", p.Name)
	fmt.Printf("PID:		%d\n", p.PID)
	fmt.Printf("PPID:		%d\n", p.PPID)
	fmt.Printf("State:		%s\n", p.State)
	fmt.Printf("Threads:	%d\n", p.Threads)
	fmt.Printf("VmSize:		%d\n", p.VmSize)
	fmt.Printf("VmRSS:		%d\n", p.VmRSS)
	fmt.Printf("UID:		%d\n", p.UID)
	fmt.Printf("Command:	%s\n", command)
}

func main() {
	args := os.Args
	if len(args) == 2{
		pidNum, err := strconv.Atoi(os.Args[1])
		if err != nil{
			fmt.Println("Invalid PID: abc — please provide a number", err)
			return
		}

		process, err := ReadProcess(pidNum)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		PrintProcess(process)
	} else {
		process, err := ReadProcess(os.Getpid())
		if err != nil {
    		fmt.Println("Error:", err)
    		return
		}
		PrintProcess(process)
	}
}