package functions

import "fmt"

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