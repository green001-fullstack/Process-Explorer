package functions

import(
	"os"
	"strconv"
)

func GetAllPIDs() ([]int, error) {
	entries, err := os.ReadDir("/proc")
	if err != nil {
		return nil, err
	}

	var pids []int
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		entryName, err := strconv.Atoi(entry.Name())
		if err != nil {
			continue
		}

		pids = append(pids, entryName)
	}

	return pids, nil
}