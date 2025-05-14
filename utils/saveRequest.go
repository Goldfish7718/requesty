package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"requesty/types"
)

func SaveRequest(reqType string, route string) {
	currentEnvironment := GetEnvironment()
	if currentEnvironment.EnvironmentName == "" {
		fmt.Println("No saved environments found!")
		return
	}

	folderPath := "data/environments"
	currentEnvPath := "data/currentenv.json"

	filePath := filepath.Join(folderPath, currentEnvironment.EnvironmentName+".json")

	newRequest := types.Request{
		ReqType: reqType,
		Route:   route,
	}

	currentEnvironment.Requests = append(currentEnvironment.Requests, newRequest)

	updatedData, err := json.MarshalIndent(currentEnvironment, "", "  ")
	if err != nil {
		log.Fatal("Error marshalling data", err)
	}

	// WRITE TO ACTUAL ENVIRONMENT FILE
	err = os.WriteFile(filePath, updatedData, 0644)
	if err != nil {
		log.Fatal("Error writing to file", err)
	}

	// WRITE TO SELECTED ENVIRONENT FILE
	err = os.WriteFile(currentEnvPath, updatedData, 0644)
	if err != nil {
		log.Fatal("Error writing to file", err)
	}

	fmt.Println("Request saved successfully")
}
