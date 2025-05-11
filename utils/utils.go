package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"requesty/types"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/chzyer/readline"
)

func AcceptRequestBody() *bytes.Buffer {

	var jsonStr string

	rl, err := readline.New("> ")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer rl.Close()

	fmt.Println("Type your JSON (press Enter to submit):")

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		if line == "" {
			break
		}
		jsonStr += line
	}

	fmt.Println("You entered:")
	fmt.Println(jsonStr)

	reqBody := bytes.NewBuffer([]byte(jsonStr))
	return reqBody
}

func ParseResponseBody(resp *http.Response) string {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	return string(body)
}

func GetEnvironment() types.Project {
	var env types.Project

	envPath := "data/currentenv.json"
	data, err := os.ReadFile(envPath)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	err = json.Unmarshal(data, &env)
	if err != nil {
		log.Fatal("Error unmarshalling JSON", err)
	}

	return env
}

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
