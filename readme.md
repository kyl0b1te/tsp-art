# tsp-art

Simple CLI tool assistance for [Concorde](https://www.tsp.gatech.edu/index.html) application.
Provides possibility to generate TSP files from SVG and TSP Art SVG's from CYC.

## Requirements

Source SVG file should contain only `<circle>` elements with `X` and `Y` coordinates, all other tags will be omitted.

## Hot to use it

- `tsp_art -tsp [PATH TO SVG]` - generates TSP file from SVG
- `tsp_art -art [PATH TO SVG] [PATH TO CYC]` - generates SVG tsp art file from SVG and CYC
