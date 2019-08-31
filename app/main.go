package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		raiseError(errors.New("Some app arguments are missing"))
	}

	svg, err := getFilePathSVG()
	raiseError(err)

	tsp, err := getFilePathTSP()
	raiseError(err)

	_, err = convert(svg, tsp)
	raiseError(err)
}

func getFilePathSVG() (string, error) {

	if os.Args[1] == "" {
		return "", errors.New("Invalid SVG path")
	}
	return os.Args[1], nil
}

func getFilePathTSP() (string, error) {

	if os.Args[2] == "" {
		return "", errors.New("Invalid TSP path")
	}
	return os.Args[2], nil
}

func raiseError(err error) {
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err.Error())
		os.Exit(1)
	}
}
