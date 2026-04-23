package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"process-explorer/functions"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		// if you want to check only self or specific number
		if args[1] == "self" {
			process, err := functions.ReadProcess(os.Getpid())
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Printf("%-7s %-7s %-7s %-9s %s\n",
			"PID", "PPID", "STATE", "MEM(KB)", "COMMAND")
			fmt.Println(strings.Repeat("-", 70))
			functions.PrintProcessLine(process)
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

		process, err := functions.ReadProcess(pidNum)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("%-7s %-7s %-7s %-9s %s\n",
			"PID", "PPID", "STATE", "MEM(KB)", "COMMAND")
		fmt.Println(strings.Repeat("-", 70))
		functions.PrintProcessLine(process)
		}
	} else {
		processes, err := functions.GetAllProcesses()
		if err != nil {
			fmt.Println("Error collecting processes", err)
			return
		}
		fmt.Printf("Found %d processes\n", len(processes))

		fmt.Printf("%-7s %-7s %-7s %-9s %s\n",
			"PID", "PPID", "STATE", "MEM(KB)", "COMMAND")
		fmt.Println(strings.Repeat("-", 70))

		for _, p := range processes {
			functions.PrintProcessLine(p)
		}
	}

}
