package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Process struct {
	Name    string
	PID     int
	PPID    int
	State   string
	Threads int
	VmSize  int
	VmRSS   int
	UID     int
	Command string
}

func GetAllPIDs() ([]int, error) {
	entries, err := os.ReadDir("/proc")
	if err != nil {
		return nil, err
	}

	var pids []int
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		entryName, err := strconv.Atoi(entry.Name())
		if err != nil {
			continue
		}

		pids = append(pids, entryName)
	}

	return pids, nil
}

func GetAllProcesses() ([]*Process, error) {
	// 1. Call GetAllPIDs()
	pids, err := GetAllPIDs()
	if err != nil {
		return nil, err
	}

	var allProcessStruct []*Process
	// 2. Loop through each PID
	for _, process := range pids {
		// 3. Call ReadProcess() for each PID
		eachProcessStruct, err := ReadProcess(process)
		if err != nil {
			continue
		}
		allProcessStruct = append(allProcessStruct, eachProcessStruct)
	}

	return allProcessStruct, nil
}

func PrintProcessLine(p *Process) {
	command := p.Command
	if command == "" {
		command = "[kernel]"
	}

	if len(command) > 15{
		command = command[:12] + "..."
	}

	fmt.Printf("%-7d %-7d %-7s %-9d %s\n",
		p.PID,
		p.PPID,
		p.State,
		p.VmRSS,
		command,
	)
}

func stateExtraction(str string) string {
	word := strings.Fields(str)
	if len(word) > 0 {
		return word[0]
	}
	return ""
}

func extractNumber(value string) int {
	valueField := strings.Fields(value)
	if len(valueField) > 0 {
		number, _ := strconv.Atoi(valueField[0])
		return number
	}
	return 0
}

func uidValue(value string) int {
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

func PrintProcess(p *Process) {
	// Handle empty command (e.g., kernel threads)
	command := p.Command
	if command == "" {
		command = "[kernel]"
	}

	fmt.Printf("%-10s %s\n", "Name:", p.Name)
	fmt.Printf("%-10s %d\n", "PID:", p.PID)
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
	if len(args) == 2 {
		// if you want to check only self or specific number
		if args[1] == "self" {
			process, err := ReadProcess(os.Getpid())
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Printf("%-7s %-7s %-7s %-9s %s\n",
			"PID", "PPID", "STATE", "MEM(KB)", "COMMAND")
			fmt.Println(strings.Repeat("-", 70))
			PrintProcessLine(process)
		} else{
			pidNum, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Invalid PID:" + args[1] + "— please provide a number")
			return
		}
		if args[1] == "0" {
			fmt.Println("Invalid number, input values above 0")
			return
		}

		process, err := ReadProcess(pidNum)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("%-7s %-7s %-7s %-9s %s\n",
			"PID", "PPID", "STATE", "MEM(KB)", "COMMAND")
		fmt.Println(strings.Repeat("-", 70))
		PrintProcessLine(process)
		}
	} else {
		processes, err := GetAllProcesses()
		if err != nil {
			fmt.Println("Error collecting processes", err)
			return
		}
		fmt.Printf("Found %d processes\n", len(processes))

		fmt.Printf("%-7s %-7s %-7s %-9s %s\n",
			"PID", "PPID", "STATE", "MEM(KB)", "COMMAND")
		fmt.Println(strings.Repeat("-", 70))

		for _, p := range processes {
			PrintProcessLine(p)
		}
	}

}
