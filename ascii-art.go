package main

import (
	"fmt"
	"os"
	"strings"
)

func GenerateAsciiArt(text string, renderingType string) string {

	result := ""

	// split the input by "\n"
	textLines := strings.Split(text, "\n")

	artList := " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"

	// read the format file and split each ascii-art to slice
	artByte, _ := os.ReadFile(renderingType)
	artLines := strings.Split(string(artByte), "\n")

	// treat the input by lines
	splitResult := make([]string, len(textLines))
	for tli, tls := range textLines {

		// look for necessary art (lines) and put them into slice by art
		artSlice := make([][]string, len(tls))
		for i := range artSlice {
			artSlice[i] = make([]string, 8)
		}
		asi := 0
		for _, tb := range tls {
			for ai, ab := range artList {
				if tb == ab {
					for i := 0; i < 8; i++ {
						artSlice[asi][i] = artLines[1+i+ai*9]
					}
					asi++
				}
			}
		}

		// place the result in suitable order to print
		for ln := 0; ln < 8; ln++ {
			for i := 0; i < len(artSlice); i++ {
				splitResult[tli] += strings.TrimRight(artSlice[i][ln], "\r")
			}
			splitResult[tli] += "\n"
		}
	}

	// conclude the result in a string
	for sri, srs := range splitResult {
		// new line in the input became "\n\n\n\n\n\n\n\n" so dealt with
		if srs == "\n\n\n\n\n\n\n\n" {
			result = result + "\n"
			continue
		}
		// do not add "\n" on last line
		if sri == len(splitResult)-1 {
			result = result + strings.TrimRight(splitResult[sri], "\n")
		} else {
			result = result + splitResult[sri]
		}
	}

	fmt.Println(result)
	return result
}
