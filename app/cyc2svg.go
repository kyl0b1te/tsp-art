package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func getSVGFromCYC(cyc string, originSVG string, svg string) error {

	// open cyc file
	cycFile, err := os.Open(cyc)
	if err != nil {
		return errors.Wrapf(err, "Cannot open .CYC file by path %s", cyc)
	}
	defer cycFile.Close()

	// open origin svg file
	originSvgFile, err := os.Open(originSVG)
	if err != nil {
		return errors.Wrapf(err, "Cannot open .SVG file by path %s", originSVG)
	}
	defer originSvgFile.Close()

	// create svg file
	svgFile, err := os.Create(svg)
	if err != nil {
		return errors.Wrapf(err, "Cannot create a SVG file by path %s", svg)
	}
	defer svgFile.Close()

	source := getSourceSVG(originSvgFile)
	coords, err := getPathCoordinates(cycFile, source)
	if err != nil {
		return err
	}

	// store data in svg file
	data := getSVGData(source, coords)
	return writeInFile(svgFile, data)
}

func getPathCoordinates(cycFile *os.File, source *SourceSVG) ([]string, error) {

	idx, err := getPointIndexes(cycFile)
	if err != nil {
		return []string{}, err
	}

	max := len(source.Circles)
	coords := make([]string, 0, max)

	for i, id := range idx {
		if id >= max {
			return []string{}, errors.New(fmt.Sprintf("Cannot get coordinates with #%d", id))
		}
		circle := source.Circles[id]

		prefix := "L"
		if i == 0 {
			prefix = "M"
		}

		coords = append(coords, fmt.Sprintf("%s %s %s", prefix, circle.X, circle.Y))
	}

	return coords, nil
}

func getPointIndexes(file *os.File) ([]int, error) {

	idx := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		id, err := strconv.Atoi(line)
		if err != nil {
			return []int{}, errors.Wrapf(err, "Cannot retrieve an index from %s", line)
		}

		idx = append(idx, id)
	}

	return idx, nil
}

func getSVGData(source *SourceSVG, coords []string) string {

	tpl := `<?xml version="1.0" encoding="utf-8"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg version="1.1" preserveAspectRatio="xMinYMin meet" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="%spx" height="%spx" viewBox="0 0 %s %s">
 <g fill="black" stroke="none">
 <path d="%s Z" stroke-width="0.25" stroke="black" fill="none" />
 </g>
</svg>
	`
	return fmt.Sprintf(
		tpl,
		source.Width,
		source.Height,
		source.Width,
		source.Height,
		strings.Join(coords, "\n"),
	)
}
