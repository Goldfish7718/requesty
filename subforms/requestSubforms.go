package subforms

import (
	"log"
	"requesty/requests"

	"github.com/charmbracelet/huh"
)

func RequestSubform() {
	var requestType string
	var route string
	var save bool

	// baseUrl := utils.GetBaseUrl()
	baseUrl := "http://localhost:3000"

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select Request type").
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

	completeUrl := baseUrl + route

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
		requests.SaveRequest(requestType, route)
	}
}
