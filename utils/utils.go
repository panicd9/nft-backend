package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

func RunCommand(cmd string) (string, error) {
	cmdParts := strings.Fields(cmd)

	command := exec.Command(cmdParts[0], cmdParts[1:]...)

	output, err := command.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command: %v", err)
	}

	outputString := string(output)

	return outputString, nil
}

func GoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
