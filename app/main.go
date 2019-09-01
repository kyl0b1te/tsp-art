package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func main() {

	tspFl := flag.Bool("tsp", false, "generate TSP from SVG")
	artFl := flag.Bool("art", false, "generate TSP Art from SVG with CYC")
	flag.Parse()

	if *tspFl == true && *artFl == true {
		raiseError(errors.New("Specify only one flag, -tsp or -art"))
	}

	if *tspFl == true {

		if len(flag.Args()) != 1 || flag.Arg(0) == "" {
			raiseError(errors.New("Invalid input parameters"))
		}

		svg := flag.Arg(0)
		tsp := getNewFileName(svg, "", ".tsp")

		err := getTSPFromSVG(svg, tsp)
		raiseError(err)
	} else if *artFl == true {

		if len(flag.Args()) != 2 || flag.Arg(0) == "" || flag.Arg(1) == "" {
			raiseError(errors.New("Invalid input parameters"))
		}

		svg := flag.Arg(0)
		cyc := flag.Arg(1)
		res := getNewFileName(svg, "-art", ".svg")

		err := getSVGFromCYC(cyc, svg, res)
		raiseError(err)
	} else {

		help()
	}

	os.Exit(0)
}

func help() {

	lines := []string{
		"\nHow to use it:",
		"\t./tsp_art -tsp [PATH TO SVG] - generates TSP file from SVG",
		"\t./tsp_art -art [PATH TO SVG] [PATH TO CYC] - generates SVG tsp art file from SVG and CYC",
	}
	fmt.Println(strings.Join(lines, "\n"))
}

func getNewFileName(srcFile string, prefix string, ext string) string {

	srcExt := filepath.Ext(srcFile)
	trimmed := strings.TrimRight(srcFile, srcExt)
	return trimmed + prefix + ext
}

func writeInFile(file *os.File, data string) error {

	writed, err := file.Write([]byte(data + "\n"))
	if err != nil {
		return errors.Wrapf(err, "Cannot write data in file: %s", data)
	}
	if writed != len(data) {
		return errors.Wrapf(err, "Failed on writing data: %s", data)
	}
	return nil
}

func raiseError(err error) {
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err.Error())
		os.Exit(1)
	}
}
