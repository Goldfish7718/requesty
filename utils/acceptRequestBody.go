package utils

import (
	"bytes"
	"fmt"
	"os"

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
