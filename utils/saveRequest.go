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
	var projectName string

	options := GetProjectsOptions()
	if len(options) == 0 {
		fmt.Println("No saved environments found!")
		return
	}

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select project to save this request:").
				Options(options...).
				Value(&projectName),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	folderPath := "data/projects"
	filePath := filepath.Join(folderPath, projectName+".json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	var project types.Project

	err = json.Unmarshal(data, &project)
	if err != nil {
		log.Fatal("Error Unmarshalling JSON", err)
	}

	newRequest := types.Request{
		ReqType: reqType,
		Route:   route,
	}

	project.Requests = append(project.Requests, newRequest)

	updatedData, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		log.Fatal("Error marshalling data", err)
	}

	err = os.WriteFile(filePath, updatedData, 0644)
	if err != nil {
		log.Fatal("Error writing to file", err)
	}

	fmt.Println("Request saved successfully")
}
