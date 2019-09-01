package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type SourceSVG struct {
	XMLName xml.Name `xml:"svg"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	Circles Circles  `xml:"circle"`
}

type Circles []struct {
	X string `xml:"cx,attr"`
	Y string `xml:"cy,attr"`
}

func getTSPFromSVG(svg string, tsp string) error {

	// open svg file
	svgFile, err := os.Open(svg)
	if err != nil {
		return errors.Wrapf(err, "Cannot open .SVG file by path %s", svg)
	}
	defer svgFile.Close()

	// create tsp file
	tspFile, err := os.Create(tsp)
	if err != nil {
		return errors.Wrapf(err, "Cannot create a .TSP file by path %s", tsp)
	}
	defer tspFile.Close()

	// get svg coordinates
	source := getSourceSVG(svgFile)
	headers := getTSPHeaders(len(source.Circles))

	lines := append(
		make([]string, 0, len(source.Circles)+len(headers)),
		headers...,
	)
	for i, circle := range source.Circles {
		lines = append(lines, fmt.Sprintf("%d %s %s", i+1, circle.X, circle.Y))
	}

	// store lines in tsp
	data := strings.Join(lines, "\n") + "\nEOF"
	if err := writeInFile(tspFile, data); err != nil {
		return err
	}

	return nil
}

func getSourceSVG(file *os.File) *SourceSVG {

	decoder := xml.NewDecoder(file)

	var svg SourceSVG
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}

		switch startEl := token.(type) {
		case xml.StartElement:
			if startEl.Name.Local == "svg" {
				decoder.DecodeElement(&svg, &startEl)
			}
		default:
		}
	}

	return &svg
}

func getTSPHeaders(numPoints int) []string {

	return []string{
		"NAME: output",
		"COMMENT: svg2tsp",
		"TYPE: TSP",
		fmt.Sprintf("DIMENSION: %d", numPoints),
		"EDGE_WEIGHT_TYPE: EUC_2D",
		"NODE_COORD_SECTION",
	}
}
