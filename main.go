package main

import (
	"fmt"
	"log"
	"os"
	"requesty/subforms"
	"requesty/utils"

	"github.com/charmbracelet/huh"
)

func main() {

	folderPath := "data/environments"

	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		log.Fatal("Error creating directory", err)
	}

	var choice string
	var selectActionString string
	var options = []huh.Option[string]{
		huh.NewOption("Make a new request", "new_request"),
		huh.NewOption("Perform saved request", "perform_saved"),
		huh.NewOption("Manage Environments", "manage_env"),
		huh.NewOption("Exit", "exit"),
	}

	fmt.Println("Welcome to Requesty!")

	for {
		environmentInfo := utils.GetEnvironment()
		if environmentInfo.EnvironmentName != "" {
			selectActionString = fmt.Sprintf("(%s) Select an action:", environmentInfo.EnvironmentName)
		} else {
			selectActionString = "Select an action:"
		}

		if err := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title(selectActionString).
					Options(options...).
					Value(&choice),
			),
		).Run(); err != nil {
			log.Fatal(err)
		}

		switch choice {
		case "new_request":
			subforms.RequestSubform()

		case "perform_saved":
			subforms.SavedRequestSubform()

		case "manage_env":
			subforms.EnvironmentSubform()

		case "exit":
			fmt.Println("Bye!")
			return
		}
	}
}
