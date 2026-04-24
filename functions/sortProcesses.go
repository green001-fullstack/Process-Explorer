package functions

import "sort"

func SortProcesses(processes []*Process, by string) {
    switch by {
    case "mem":
        // sort by VmRSS, highest first
		sort.Slice(processes, func(i int, j int) bool{
			return processes[i].VmRSS > processes[j].VmRSS
		})
    case "name":
        // sort by Name, alphabetically
		sort.Slice(processes, func(i, j int)bool{
			return processes[i].Name > processes[j].Name
		})
    default:
        // sort by PID, lowest first
		sort.Slice(processes, func(i, j int) bool {
			return processes[i].PID < processes[j].PID
		})
    }
}