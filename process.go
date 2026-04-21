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

    fmt.Printf("%-10s %s\n", "Name:", p.Name)
	fmt.Printf("%-10s %d\n", "PID:",  p.PID)
	fmt.Printf("%-10s %d\n", "PPID", p.PPID)
	fmt.Printf("%-10s %s\n", "State", p.State)
	fmt.Printf("%-10s %d\n", "Threads", p.Threads)
	fmt.Printf("%-10s %d KB\n", "VmSize", p.VmSize)
	fmt.Printf("%-10s %d KB\n", "VmRSS", p.VmRSS)
	fmt.Printf("%-10s %d\n", "UID", p.UID)
	fmt.Printf("%-10s %s\n", "Command", command)
}

func main() {
	args := os.Args
	if len(args) == 2{
		pidNum, err := strconv.Atoi(os.Args[1])
		if err != nil{
			fmt.Println("Invalid PID:" + args[1] + "— please provide a number")
			return
		}
		if args[1] == "0"{
			fmt.Println("Invalid number, input values above 0")
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