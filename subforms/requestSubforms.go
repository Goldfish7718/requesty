package subforms

import (
	"fmt"
	"log"
	"requesty/requests"
	"requesty/utils"

	"github.com/charmbracelet/huh"
)

func RequestSubform() {
	var requestType string
	var route string
	var save bool

	environment := utils.GetEnvironment()

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title(fmt.Sprintf("Select request type\nEnvironment name: %s\nBase URL: %s", environment.EnvironmentName, environment.BaseUrl)).
				Options(
					huh.NewOption("GET", "GET"),
					huh.NewOption("POST", "POST"),
					huh.NewOption("PUT", "PUT"),
					huh.NewOption("DELETE", "DELETE"),
				).
				Value(&requestType),

			huh.NewInput().
				Title("Enter route (with trailing slash; appended to base URL)").
				Value(&route),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	completeUrl := environment.BaseUrl + route

	switch requestType {
	case "GET":
		requests.Get(completeUrl)

	case "POST":
		requests.Post(completeUrl)

	case "PUT":
		requests.Put(completeUrl)

	case "DELETE":
		requests.Delete(completeUrl)
	}

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Do you want to save this request?").
				Affirmative("Yes").
				Negative("No").
				Value(&save),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	if save {
		utils.SaveRequest(requestType, route)
	}
}
