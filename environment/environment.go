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
	var projectName string
	var baseUrl string

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter Project name:").
				Value(&projectName),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Enter base URL:").
				Value(&baseUrl),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	project := types.Project{
		ProjectName: projectName,
		BaseUrl:     baseUrl,
		Requests:    []types.Request{},
	}

	folderPath := "data/projects"

	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		log.Fatal("Error creating directory", err)
	}

	filePath := filepath.Join(folderPath, projectName+".json")
	fileData, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		log.Fatal("❌ Error marshalling JSON:", err)
	}

	err = os.WriteFile(filePath, fileData, 0644)
	if err != nil {
		log.Fatal("Error writing to file", err)
	}

	fmt.Printf("Succefully created project %s\n", projectName)
}

func View() {
	folderPath := "data/projects"

	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal("Error reading directory", folderPath)
	}

	if len(files) == 0 {
		fmt.Println("No saved projects")
		return
	}

	fmt.Println("Saved projects:")
	for index, file := range files {
		projectName := strings.TrimSuffix(file.Name(), ".json")
		fmt.Printf("%d. %s\n", index+1, projectName)
	}
}

func Edit() {
	projectOptions := utils.GetProjectsOptions()
	var projectToEdit string

	if len(projectOptions) == 0 {
		fmt.Println("No saved projects found")
		return
	}

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select project to Edit:").
				Options(projectOptions...).
				Value(&projectToEdit),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	folderPath := "data/projects"
	filePath := filepath.Join(folderPath, projectToEdit+".json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	var project types.Project

	err = json.Unmarshal(data, &project)
	if err != nil {
		log.Fatal("Error unmarshalling JSON", err)
	}

	fmt.Printf("Project Name: %s\nBase URL: %s\nSaved requests: (%d)\n", project.ProjectName, project.BaseUrl, len(project.Requests))
	for _, req := range project.Requests {
		fmt.Println("- ", req.ReqType, " ", req.Route)
	}

	var editAction string

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select an action").
				Options(
					huh.NewOption("Change project name", "change_name"),
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
					Title("Enter new project name:").
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

		project.ProjectName = newName

	case "change_base":
		var newBaseUrl string
		if err := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Enter new project name:").
					Value(&newBaseUrl),
			),
		).Run(); err != nil {
			log.Fatal(err)
		}

		project.BaseUrl = newBaseUrl

	case "delete_req":
		var options []huh.Option[int]
		var requestIndexToDelete int

		if len(project.Requests) == 0 {
			fmt.Println("No saved requests found")
			return
		}

		for index, req := range project.Requests {
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

		project.Requests = append(project.Requests[:requestIndexToDelete], project.Requests[requestIndexToDelete+1:]...)
	}

	fileData, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		log.Fatal("❌ Error marshalling JSON:", err)
	}

	err = os.WriteFile(filePath, fileData, 0644)
	if err != nil {
		log.Fatal("Error writing to file", err)
	}

	fmt.Println("Successfully edited project")
}

func Delete() {
	var projectToDelete string
	var projectDeleteConfirm bool
	projectOptions := utils.GetProjectsOptions()

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select project to delete:").
				Options(projectOptions...).
				Value(&projectToDelete),

			huh.NewConfirm().
				Title("Are you sure you want to delete this project?").
				Affirmative("Yes").
				Negative("No").
				Value(&projectDeleteConfirm),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	if !projectDeleteConfirm {
		return
	}

	folderPath := "data/projects"
	filepath := filepath.Join(folderPath, projectToDelete+".json")

	err := os.Remove(filepath)
	if err != nil {
		log.Fatal("Error deleting file", err)
	}

	fmt.Printf("\nProject %s deleted succesfully", projectToDelete)
}

func Select() {
	var envName string
	options := utils.GetProjectsOptions()

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

	folderPath := "data/projects"
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
		log.Fatal("Error deleting file", err)
	}

	fmt.Println("Unselected current environment")
}
