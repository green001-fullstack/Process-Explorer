package main

import (
    "fmt"
    "os"
    "strings"
	"strconv"
	"path/filepath"
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

func commandLine(pid int)string{
	path := filepath.Join("/proc/", strconv.Itoa(pid), "/cmdline")
	file, err := os.ReadFile(path)
	if err != nil{
		fmt.Println("Error reading file:", err)
		return ""
	}
	command := string(file) 
	return command
	}

func main() {
    // Read /proc/self/status
	file, err := os.ReadFile("/proc/self/status")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	content := string(file)
	// fmt.Println(content)
    
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
	command := strings.ReplaceAll(commandLine(data.PID), "\x00", " ")
	data.Command = strings.TrimSpace(command)
	fmt.Println(command)
	fmt.Printf("%+v\n", data)

	
}