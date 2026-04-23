package functions

import "fmt"


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