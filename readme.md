# svg2tsp

CLI tool for convert SVG circle coordinates in TSP file format.
App expects to receive a two arguments: path to svg source file and path to tsp result file.

Execution example: `./svg2tsp src/example.svg dist/example.tsp`

# Requirements

SVG file should contain only `<circle>` elements with `X` and `Y` coordinates, all other tags will be omitted.
