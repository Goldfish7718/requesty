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
	environmentName := utils.GetEnvironment().EnvironmentName
	if environmentName == "" {
		fmt.Println("No environment selected!")
		return
	}

	folderPath := "data/environments"
	filePath := filepath.Join(folderPath, environmentName+".json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	var environment types.Environment
	var requestOptions []huh.Option[int]
	var requestIndex int

	err = json.Unmarshal(data, &environment)
	if err != nil {
		log.Fatal("Error Unmarshalling JSON", err)
	}

	if len(environment.Requests) == 0 {
		fmt.Println("No saved request found!")
		return
	}

	for index, request := range environment.Requests {
		requestOption := fmt.Sprintf("%s %s", request.ReqType, request.Route)
		requestOptions = append(requestOptions, huh.NewOption(requestOption, index))
	}

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Select request to perform:\nEnvironment: " + environmentName + "\nBase URL: " + environment.BaseUrl).
				Options(requestOptions...).
				Value(&requestIndex),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	selectedRequest := environment.Requests[requestIndex]
	completeUrl := environment.BaseUrl + selectedRequest.Route

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
