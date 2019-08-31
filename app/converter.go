package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Circles Circles  `xml:"circle"`
}

type Circles []struct {
	X string `xml:"cx,attr"`
	Y string `xml:"cy,attr"`
}

func convert(svg string, tsp string) (bool, error) {

	// open svg file
	svgFile, err := os.Open(svg)
	if err != nil {
		return false, errors.Wrapf(err, "Cannot open SVG file by path %s", svg)
	}
	defer svgFile.Close()

	// create tsp file
	tspFile, err := getTSP(tsp)
	if err != nil {
		return false, err
	}
	defer tspFile.Close()

	// get svg coordinates
	coors := getCoordinates(svgFile)

	// store coordinates in tsp
	data := strings.Join(coors, "\n")
	if err := writeInFile(tspFile, data); err != nil {
		return false, err
	}

	return true, nil
}

func getTSP(path string) (*os.File, error) {

	tspFile, err := os.Create(path)
	if err != nil {
		return &os.File{}, errors.Wrapf(err, "Cannot create a .TSP file by path %s", path)
	}
	tspHeaders := []string{
		"NAME: output",
		"COMMENT: svg2tsp",
		"TYPE: TSP",
		"DIMENSION: 1500", // maybe need to make it configurable
		"EDGE_WEIGHT_TYPE: EUC_2D",
		"NODE_COORD_SECTION",
	}

	for _, line := range tspHeaders {
		if err := writeInFile(tspFile, line); err != nil {
			tspFile.Close()
			return &os.File{}, err
		}
	}

	return tspFile, nil
}

func getCoordinates(file *os.File) []string {

	decoder := xml.NewDecoder(file)
	lines := []string{}
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}

		switch startEl := token.(type) {
		case xml.StartElement:
			if startEl.Name.Local == "svg" {
				var root SVG
				decoder.DecodeElement(&root, &startEl)

				for i, crl := range root.Circles {
					line := fmt.Sprintf("%d %s %s", i+1, crl.X, crl.Y)
					lines = append(lines, line)
				}
			}
		default:
		}
	}

	return append(lines, "EOF")
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
