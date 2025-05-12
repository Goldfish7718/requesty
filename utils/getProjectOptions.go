package utils

import (
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
)

func GetProjectsOptions() []huh.Option[string] {
	var projectOptions []huh.Option[string]

	folderPath := "data/projects"

	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal("Error reading directory", folderPath)
	}

	for _, file := range files {
		projectName := strings.TrimSuffix(file.Name(), ".json")
		projectOptions = append(projectOptions, huh.NewOption(projectName, projectName))
	}

	return projectOptions
}
