package environment

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"requesty/types"
	"requesty/utils"
	"strings"

	"github.com/charmbracelet/huh"
)

func New() {
	var environmentName string
	var baseUrl string

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter Environment name:").
				Value(&environmentName),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Enter base URL:").
				Value(&baseUrl),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	environment := types.Environment{
		EnvironmentName: environmentName,
		BaseUrl:         baseUrl,
		Requests:        []types.Request{},
	}

	folderPath := "data/environments"

	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		log.Fatal("Error creating directory", err)
	}

	filePath := filepath.Join(folderPath, environmentName+".json")
	fileData, err := json.MarshalIndent(environment, "", "  ")
	if err != nil {
		log.Fatal("❌ Error marshalling JSON:", err)
	}

	err = os.WriteFile(filePath, fileData, 0644)
	if err != nil {
		log.Fatal("Error writing to file", err)
	}

	fmt.Printf("Succefully created environment %s\n", environmentName)
}

func View() {
	folderPath := "data/environments"

	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal("Error reading directory", folderPath)
	}

	if len(files) == 0 {
		fmt.Println("No saved environments")
		return
	}

	fmt.Println("Saved environments:")
	for index, file := range files {
		environmentName := strings.TrimSuffix(file.Name(), ".json")
		fmt.Printf("%d. %s\n", index+1, environmentName)
	}
}

func Edit() {
	environmentOptions := utils.GetEnvironmentsOptions()
	currentEnvironment := utils.GetEnvironment()

	var environmentToEdit string

	if len(environmentOptions) == 0 {
		fmt.Println("No saved environments found")
		return
	}

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select environment to Edit:").
				Options(environmentOptions...).
				Value(&environmentToEdit),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	folderPath := "data/environments"
	currentEnvPath := "data/currentenv.json"

	filePath := filepath.Join(folderPath, environmentToEdit+".json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	var environment types.Environment

	err = json.Unmarshal(data, &environment)
	if err != nil {
		log.Fatal("Error unmarshalling JSON", err)
	}

	oldEnvironmentName := environment.EnvironmentName

	fmt.Printf("Environment Name: %s\nBase URL: %s\nSaved requests: (%d)\n", environment.EnvironmentName, environment.BaseUrl, len(environment.Requests))
	for _, req := range environment.Requests {
		fmt.Println("- ", req.ReqType, " ", req.Route)
	}

	var editAction string

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select an action").
				Options(
					huh.NewOption("Change environment name", "change_name"),
					huh.NewOption("Change Base URL", "change_base"),
					huh.NewOption("Delete a request", "delete_req"),
				).
				Value(&editAction),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	switch editAction {
	case "change_name":
		var newName string
		if err := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Enter new environment name:").
					Value(&newName),
			),
		).Run(); err != nil {
			log.Fatal(err)
		}

		newFilePath := filepath.Join(folderPath, newName+".json")
		oldFilePath := filePath
		filePath = newFilePath

		err = os.Rename(oldFilePath, newFilePath)
		if err != nil {
			log.Fatal("Error renaming file", err)
		}

		environment.EnvironmentName = newName

	case "change_base":
		var newBaseUrl string
		if err := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Enter new environment name:").
					Value(&newBaseUrl),
			),
		).Run(); err != nil {
			log.Fatal(err)
		}

		environment.BaseUrl = newBaseUrl

	case "delete_req":
		var options []huh.Option[int]
		var requestIndexToDelete int

		if len(environment.Requests) == 0 {
			fmt.Println("No saved requests found")
			return
		}

		for index, req := range environment.Requests {
			options = append(options, huh.NewOption(fmt.Sprintf("- %s %s", req.ReqType, req.Route), index))
		}

		if err := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[int]().
					Title("Select request to delete").
					Options(options...).
					Value(&requestIndexToDelete),
			),
		).Run(); err != nil {
			log.Fatal(err)
		}

		environment.Requests = append(environment.Requests[:requestIndexToDelete], environment.Requests[requestIndexToDelete+1:]...)
	}

	fileData, err := json.MarshalIndent(environment, "", "  ")
	if err != nil {
		log.Fatal("❌ Error marshalling JSON:", err)
	}

	err = os.WriteFile(filePath, fileData, 0644)
	if err != nil {
		log.Fatal("Error writing to file", err)
	}

	if oldEnvironmentName == currentEnvironment.EnvironmentName {
		err = os.WriteFile(currentEnvPath, fileData, 0644)
		if err != nil {
			log.Fatal("Error writing to file", err)
		}
	}

	fmt.Println("Successfully edited environment")
}

func Delete() {
	var environmentToDelete string
	var environmentDeleteConfirm bool

	environmentOptions := utils.GetEnvironmentsOptions()
	currentEnvName := utils.GetEnvironment().EnvironmentName

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select environment to delete:").
				Options(environmentOptions...).
				Value(&environmentToDelete),

			huh.NewConfirm().
				Title("Are you sure you want to delete this environment?").
				Affirmative("Yes").
				Negative("No").
				Value(&environmentDeleteConfirm),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	if currentEnvName == environmentToDelete {
		fmt.Println("Cannot delete currently selected environment (Unselect first)")
		return
	}

	if !environmentDeleteConfirm {
		return
	}

	folderPath := "data/environments"
	filepath := filepath.Join(folderPath, environmentToDelete+".json")

	err := os.Remove(filepath)
	if err != nil {
		log.Fatal("Error deleting file", err)
	}

	fmt.Printf("\nEnvironment %s deleted succesfully", environmentToDelete)
}

func Select() {
	var envName string
	options := utils.GetEnvironmentsOptions()

	if len(options) == 0 {
		fmt.Println("No saved environments!")
		return
	}

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select current environment:").
				Options(options...).
				Value(&envName),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	folderPath := "data/environments"
	envPath := filepath.Join("data/", "currentenv.json")
	filePath := filepath.Join(folderPath, envName+".json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	err = os.WriteFile(envPath, data, 0644)
	if err != nil {
		log.Fatal("Error writing file", err)
	}
}

func Unselect() {
	envPath := "data/currentenv.json"
	err := os.Remove(envPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No environment currently selected")
		}
		log.Fatal("Error deleting current environment file ", err)
	}

	fmt.Println("Unselected current environment")
}
