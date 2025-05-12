package subforms

import (
	"log"
	"requesty/environment"

	"github.com/charmbracelet/huh"
)

func EnvironmentSubform() {
	var action string

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select Request type").
				Options(
					huh.NewOption("Create new Environment", "new"),
					huh.NewOption("View Environments", "view"),
					huh.NewOption("Edit Environments", "edit"),
					huh.NewOption("Delete Environment", "delete"),
					huh.NewOption("Select Environment", "select"),
					huh.NewOption("Unselect current Environment", "unselect"),
				).
				Value(&action),
		),
	).Run(); err != nil {
		log.Fatal(err)
	}

	switch action {
	case "new":
		environment.New()

	case "view":
		environment.View()

	case "edit":
		environment.Edit()

	case "delete":
		environment.Delete()

	case "select":
		environment.Select()

	case "unselect":
		environment.Unselect()
	}
}
