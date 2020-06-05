package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/windmilleng/tilt/pkg/tiltextension"
)

const tiltfileName = "Tiltfile"

type extensionError struct {
	extensionName string
	err           error
}

func main() {
	args := os.Args[1:]
	extensionsDir := args[0]

	extensionNames, err := getNamesOfAllExtensions(extensionsDir)
	if err != nil {
		fmt.Printf("Error getting names of all extensions: %v\n", err)
		os.Exit(1)
	}

	errors := []extensionError{}
	for _, e := range extensionNames {
		// validate extension name
		err = tiltextension.ValidateName(e)
		if err != nil {
			errors = append(errors, extensionError{e, err})
		}

		// extension should have Tiltfile
		expectedTiltfilePath := filepath.Join(extensionsDir, e, tiltfileName)
		_, err := os.Stat(expectedTiltfilePath)
		if os.IsNotExist(err) {
			errors = append(errors, extensionError{e, fmt.Errorf("No Tiltfile found at %s", expectedTiltfilePath)})
		}
	}

	for _, e := range errors {
		fmt.Printf("Error validating extension named %s: %v\n", e.extensionName, e.err)
	}

	if len(errors) > 0 {
		os.Exit(1)
	}

	fmt.Println("All good!")
}

func getNamesOfAllExtensions(extensionsDir string) ([]string, error) {
	files, err := ioutil.ReadDir(extensionsDir)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		fileNames = append(fileNames, file.Name())
	}

	return fileNames, nil
}