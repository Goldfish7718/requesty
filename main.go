package main

import (
	"fmt"
	"log"
	"os"
	"requesty/subforms"

	"github.com/charmbracelet/huh"
)

type Options struct {
	Label huh.Option[string]
	Value string
}

func main() {

	folderPath := "data/projects"

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
		// environmentInfo := environment.GetEnvironmentInfo()
		// if environmentInfo != "" {
		// 	selectActionString = fmt.Sprintf("(%s) Select an action:", environmentInfo)
		// } else {
		// 	selectActionString = "Select an action:"
		// }

		selectActionString = "Select an action:"

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

		// case "perform_saved":
		// 	projects.PerformSavedRequest()

		case "manage_env":
			subforms.EnvironmentSubform()

		case "exit":
			fmt.Println("Bye!")
			return
		}
	}
}
