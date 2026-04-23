package functions

import(
	"strconv"
	"path/filepath"
	"os"
	"fmt"
	"strings"
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