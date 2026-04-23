package functions


func GetAllProcesses() ([]*Process, error) {
	// 1. Call GetAllPIDs()
	pids, err := GetAllPIDs()
	if err != nil {
		return nil, err
	}

	var allProcessStruct []*Process
	// 2. Loop through each PID
	for _, process := range pids {
		// 3. Call ReadProcess() for each PID
		eachProcessStruct, err := ReadProcess(process)
		if err != nil {
			continue
		}
		allProcessStruct = append(allProcessStruct, eachProcessStruct)
	}

	return allProcessStruct, nil
}