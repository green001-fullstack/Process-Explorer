package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"strconv"
)

func CommandLine(pid int)string{
	path := filepath.Join("/proc", strconv.Itoa(pid), "cmdline")
	file, err := os.ReadFile(path)
	if err != nil{
		fmt.Println("Error reading file:", err)
		return ""
	}
	cmd := strings.ReplaceAll(string(file), "\x00", " ") 
	return strings.TrimSpace(cmd)
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
			data.PID = extractNumber(value)
		case "PPid":
			data.PPID = extractNumber(value)
		case "State":
			data.State = stateExtraction(value)
		case "Threads":
			data.Threads = extractNumber(value)
		case "VmSize":
			data.VmSize = extractNumber(value)
		case "VmRSS":
			data.VmRSS = extractNumber(value)
		case "Uid":
			data.UID = uidValue(value)
		}
	}
	// Add the command part of the struct
	data.Command = CommandLine(pid)
	return &data, nil
}