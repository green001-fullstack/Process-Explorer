package main

import (
    "fmt"
    "os"
    "strings"
	"strconv"
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

func extractNumber(value string) int{
	valueField := strings.Fields(value) 
	if len(valueField) > 0 {
		number, _ := strconv.Atoi(valueField[0])
		return number
	}
	return 0
}

func uidValue(value string) string{
	result := strings.Fields(value)
	if len(result) > 0 {
		val := result[0]
		return val
	}
	return ""
}

func commandLine(pid int)string{
	file, err := os.ReadFile("/proc/" + strconv.Itoa(pid) + "/cmdline" )
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
    
    // Parse it
	splitLine := strings.Split(content, "\n")
	var data Process
	for _, line := range splitLine{
		pairs := strings.SplitN(line, ":", 2)
		if len(pairs) != 2 {
    		continue
		}
		key := pairs[0]
		value := strings.TrimSpace(pairs[1])
		
		switch key {
		case "Name":
			data.Name = value	
		case "Pid":
			valueInt, _ := strconv.Atoi(value)
			data.PID = valueInt

		case "PPid":
			valueInt, _ := strconv.Atoi(value)
			data.PPID = valueInt
		case "State":
			data.State = value
		case "Threads":
			valueInt, _ := strconv.Atoi(value)
			data.Threads = valueInt
		case "VmSize":
			data.VmSize = extractNumber(value)
		case "VmRSS":
			data.VmRSS = extractNumber(value)
		case "Uid":
			Uid, _ := strconv.Atoi(uidValue(value))
			data.UID = Uid
		}
	}
	command := strings.ReplaceAll(commandLine(data.PID), "\x00", " ")
	data.Command = strings.TrimSpace(command)
	fmt.Println(command)
	fmt.Printf("%+v\n", data)
    // Hint: each line is "Key:   Value"
    // Hint: strings.Split() splits a string
    // Hint: strings.TrimSpace() removes whitespace

    // Print the result

	
}