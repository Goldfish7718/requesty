package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/chzyer/readline"
)

func AcceptRequestBody() (*bytes.Buffer, error) {

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

	var js json.RawMessage

	err = json.Unmarshal([]byte(jsonStr), &js)
	if err != nil {
		return nil, errors.New("Invialid JSON format")
	}

	reqBody := bytes.NewBuffer([]byte(jsonStr))
	return reqBody, nil
}
