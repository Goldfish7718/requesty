package subforms

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"requesty/requests"
	"requesty/types"
	"requesty/utils"

	"github.com/charmbracelet/huh"
)

func SavedRequestSubform() {
	projectName := utils.GetEnvironment().ProjectName
	if projectName == "" {
		fmt.Println("No environment selected!")
		return
	}

	folderPath := "data/projects"
	filePath := filepath.Join(folderPath, projectName+".json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	var project types.Project
	var requestOptions []huh.Option[int]
	var requestIndex int

	err = json.Unmarshal(data, &project)
	if err != nil {
		log.Fatal("Error Unmarshalling JSON", err)
	}

	if len(project.Requests) == 0 {
		fmt.Println("No saved request found!")
		return
	}

	for index, request := range project.Requests {
		requestOption := fmt.Sprintf("%s %s", request.ReqType, request.Route)
		requestOptions = append(requestOptions, huh.NewOption(requestOption, index))
	}

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Select request to perform:\nProject: " + projectName + "\nBase URL: " + project.BaseUrl).
				Options(requestOptions...).
				Value(&requestIndex),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	selectedRequest := project.Requests[requestIndex]
	completeUrl := project.BaseUrl + selectedRequest.Route

	switch selectedRequest.ReqType {
	case "GET":
		requests.Get(completeUrl)

	case "POST":
		requests.Post(completeUrl)

	case "PUT":
		requests.Put(completeUrl)

	case "DELETE":
		requests.Delete(completeUrl)
	}
}
