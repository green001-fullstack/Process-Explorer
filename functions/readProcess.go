package functions

import (
	"os"
	"path/filepath"
	"strings"
	"strconv"
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

func ReadProcess(pid int) (*Process, error){
    // Read /proc/pid/status
	path := filepath.Join("/proc", strconv.Itoa(pid), "status")
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	content := string(file)
    
    // Separate the content by newline
	splitLine := strings.Split(content, "\n")
	var data Process
	// Loop through each string in splitLine
	for _, line := range splitLine{
		pairs := strings.SplitN(line, ":", 2)
		if len(pairs) != 2 {
    		continue
		}
		key := pairs[0]
		value := strings.TrimSpace(pairs[1])
		// Populate the struct
		switch key {
		case "Name":
			data.Name = value	
		case "Pid":
			data.PID = ExtractNumber(value)
		case "PPid":
			data.PPID = ExtractNumber(value)
		case "State":
			data.State = StateExtraction(value)
		case "Threads":
			data.Threads = ExtractNumber(value)
		case "VmSize":
			data.VmSize = ExtractNumber(value)
		case "VmRSS":
			data.VmRSS = ExtractNumber(value)
		case "Uid":
			data.UID = UidValue(value)
		}
	}
	// Add the command part of the struct
	data.Command = CommandLine(pid)
	return &data, nil
}