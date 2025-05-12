package utils

import (
	// "fmt"
	"encoding/json"
	"log"
	"os"
	"requesty/types"
)

func GetEnvironment() types.Project {
	var environment types.Project
	envPath := "data/currentenv.json"

	file, err := os.ReadFile(envPath)
	if err != nil {
		if os.IsNotExist(err) {
			return types.Project{} // return empty project
		}
		log.Fatal("Error reading file", err)
	}

	err = json.Unmarshal(file, &environment)
	if err != nil {
		log.Fatal("Error unmarshalling JSON", err)
	}

	return environment
}
