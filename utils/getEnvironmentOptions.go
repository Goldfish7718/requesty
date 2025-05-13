package utils

import (
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
)

func GetEnvironmentsOptions() []huh.Option[string] {
	var environmentOptions []huh.Option[string]

	folderPath := "data/environments"

	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal("Error reading directory", folderPath)
	}

	for _, file := range files {
		environmentName := strings.TrimSuffix(file.Name(), ".json")
		environmentOptions = append(environmentOptions, huh.NewOption(environmentName, environmentName))
	}

	return environmentOptions
}
