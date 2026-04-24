package main

import (
	"fmt"
	"os"
	"strconv"
	// "strings"
	"process-explorer/functions"
	"flag"
)

func main() {
	sortBy := flag.String("sort", "pid", "sort by: pid mem")
	flag.Parse()

	args := flag.Args()

	if len(args) == 1 {
		// if you want to check only self or specific number
		if args[0] == "self" {
			process, err := functions.ReadProcess(os.Getpid())
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			// fmt.Printf("%-7s %-7s %-7s %-9s %s\n",
			// "PID", "PPID", "STATE", "MEM(KB)", "COMMAND")
			// fmt.Println(strings.Repeat("-", 70))
			functions.PrintHeader()
			functions.PrintProcessLine(process)
		} else{
			pidNum, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid PID:" + args[0] + "— please provide a number")
			return
		}
		if args[0] == "0" {
			fmt.Println("Invalid number, input values above 0")
			return
		}

		process, err := functions.ReadProcess(pidNum)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// fmt.Printf("%-7s %-7s %-7s %-9s %s\n",
		// 	"PID", "PPID", "STATE", "MEM(KB)", "COMMAND")
		// fmt.Println(strings.Repeat("-", 70))
		functions.PrintHeader()
		functions.PrintProcessLine(process)
		}
	} else {
		
		processes, err := functions.GetAllProcesses()
		if err != nil {
			fmt.Println("Error collecting processes", err)
			return
		}
		fmt.Printf("Found %d processes\n", len(processes))

		// fmt.Printf("%-7s %-7s %-7s %-9s %s\n",
		// 	"PID", "PPID", "STATE", "MEM(KB)", "COMMAND")
		// fmt.Println(strings.Repeat("-", 70))
		functions.PrintHeader()


		functions.SortProcesses(processes, *sortBy)
		for _, p := range processes {
			functions.PrintProcessLine(p)
		}
	}

}
