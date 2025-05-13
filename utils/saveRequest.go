package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"requesty/types"

	"github.com/charmbracelet/huh"
)

func SaveRequest(reqType string, route string) {
	var environmentName string

	options := GetEnvironmentsOptions()
	if len(options) == 0 {
		fmt.Println("No saved environments found!")
		return
	}

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select environment to save this request:").
				Options(options...).
				Value(&environmentName),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	folderPath := "data/environments"
	filePath := filepath.Join(folderPath, environmentName+".json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	var environment types.Environment

	err = json.Unmarshal(data, &environment)
	if err != nil {
		log.Fatal("Error Unmarshalling JSON", err)
	}

	newRequest := types.Request{
		ReqType: reqType,
		Route:   route,
	}

	environment.Requests = append(environment.Requests, newRequest)

	updatedData, err := json.MarshalIndent(environment, "", "  ")
	if err != nil {
		log.Fatal("Error marshalling data", err)
	}

	err = os.WriteFile(filePath, updatedData, 0644)
	if err != nil {
		log.Fatal("Error writing to file", err)
	}

	fmt.Println("Request saved successfully")
}
