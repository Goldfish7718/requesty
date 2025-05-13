package utils

import (
	// "fmt"
	"encoding/json"
	"log"
	"os"
	"requesty/types"
)

func GetEnvironment() types.Environment {
	var environment types.Environment
	envPath := "data/currentenv.json"

	file, err := os.ReadFile(envPath)
	if err != nil {
		if os.IsNotExist(err) {
			return types.Environment{} // return empty environment
		}
		log.Fatal("Error reading file", err)
	}

	err = json.Unmarshal(file, &environment)
	if err != nil {
		log.Fatal("Error unmarshalling JSON", err)
	}

	return environment
}
