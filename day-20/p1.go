package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Image struct {
	pixels  []string
	padding int
	fill    byte
}

func (image *Image) Print() {
	for line := image.padding; line < len(image.pixels)-image.padding; line++ {
		for pixel := image.padding; pixel < len(image.pixels[line])-image.padding; pixel++ {
			switch image.pixels[line][pixel] {
			case '0':
				print(".")
				break
			case '1':
				print("#")
				break
			}
		}

		print("\n")
	}
}

func addPadding(input *Image, padding int) *Image {
	inputHeight, inputWidth := len(input.pixels), len(input.pixels[0])
	outputHeight, outputWidth := inputHeight+(padding*2), inputWidth+(padding*2)

	outputPixels := make([]string, outputHeight)

	topBottomPadding := strings.Repeat(string(input.fill), outputWidth)
	leftRightPadding := topBottomPadding[0:padding]

	for i := 0; i < padding; i++ {
		outputPixels[i] = topBottomPadding
	}

	for i := 0; i < inputHeight; i++ {
		var builder strings.Builder

		builder.WriteString(leftRightPadding)
		builder.WriteString(input.pixels[i])
		builder.WriteString(leftRightPadding)

		outputPixels[i+padding] = builder.String()
	}

	for i := inputHeight + padding; i < outputHeight; i++ {
		outputPixels[i] = topBottomPadding
	}

	return &Image{
		pixels:  outputPixels,
		padding: padding,
	}
}

func enhance(input *Image, algorithm string) *Image {
	outputFill, _ := strconv.ParseInt(strings.Repeat(string(input.fill), 9), 2, 32)
	output := &Image{
		pixels: make([]string, len(input.pixels)+2),
		fill:   algorithm[outputFill],
	}

	padded := addPadding(input, 2)

	height, width := len(padded.pixels)-2, len(padded.pixels[0])-2

	for y := 0; y < height; y++ {
		outputRow := make([]byte, width)
		for x := 0; x < width; x++ {
			square := padded.pixels[y][x:x+3] +
				padded.pixels[y+1][x:x+3] +
				padded.pixels[y+2][x:x+3]

			i, _ := strconv.ParseInt(square, 2, 32)

			outputRow[x] = algorithm[i]
		}
		output.pixels[y] = string(outputRow)
	}

	return output
}

func lightPixelCount(input *Image) int {
	height, width := len(input.pixels), len(input.pixels[0])

	var count int
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if input.pixels[y][x] == '1' {
				count++
			}
		}
	}
	return count
}

func main() {
	image, algorithm := parseInput()
	for i := 0; i < 2; i++ {
		image = enhance(image, algorithm)
	}

	fmt.Printf("Lit pixels in the resulting image: %d\n", lightPixelCount(image))

	// Answer: 5573
}

func parseInput() (*Image, string) {
	name := "input.txt"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	log.Printf("Input file: %s\n", name)

	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	algorithm := toBits(scanner.Text())
	scanner.Scan()

	var pixels []string
	for scanner.Scan() {
		pixels = append(pixels, toBits(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &Image{pixels: pixels, padding: 0, fill: '0'}, algorithm
}

func toBits(s string) string {
	b := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '.':
			b[i] = '0'
			break
		case '#':
			b[i] = '1'
			break
		}
	}

	return string(b)
}
