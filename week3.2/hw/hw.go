// Program for printing volumes based on file input
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Solid interface {
	Volume() float64
}

type ShapeEntry struct {
	Type rune
	Dim1 float64
	Dim2 float64
}

type Sphere struct {
	Radius float64 
}

type Cube struct {
	Length float64
}

type Pyramid struct {
	Base float64
	Height float64
}

func (s Sphere) Volume() float64 {
	return (4.0/3.0) * math.Pi	* (s.Radius * s.Radius * s.Radius)
}

func (c Cube) Volume() float64 {
	return c.Length * c.Length * c.Length
}

func (p Pyramid) Volume() float64 {
	return (p.Base * p.Base * p.Height) / 3
}

func readFile(fileName string) (map[string]Solid, error) {
	volume := make(map[string]Solid)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	entries := make([]ShapeEntry, 0)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
    		return nil, err
		}
		line := scanner.Text()
		parts := strings.Fields(line)
		shapeType := rune(parts[0][0])
		shapeDim1, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, err
		}
		shapeDim2, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			return nil, err
		}
		newShape := ShapeEntry{
			Type : shapeType,
			Dim1 : shapeDim1, 
			Dim2 : shapeDim2}
		entries = append(entries, newShape)
	}
	var cubeCount, sphereCount, pyramidCount int
	for _, entry := range entries {
		switch entry.Type {
		case 'C' : 
			cubeCount++
			Key := fmt.Sprintf("Cube%d", cubeCount)
			Value := Cube{Length :entry.Dim1}
			volume[Key] = Value
		case 'S' :
			sphereCount++
			Key := fmt.Sprintf("Sphere%d", sphereCount)
			Value := Sphere{Radius : entry.Dim1}
			volume[Key] = Value
		case 'P' :
			pyramidCount++
			Key := fmt.Sprintf("Pyramid%d", pyramidCount)
			Value := Pyramid{Base : entry.Dim1, Height : entry.Dim2}
			volume[Key] = Value
		}
	}
	return volume, nil
}

func main() {
	volume, err := readFile("data.txt")
	if err != nil {
		return
	}
	keys := make([]string,0)
	for key := range volume {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _,key := range keys {
		mass := volume[key].Volume()
		fmt.Printf("%s volume = %.4f\n", key ,mass)
	}
	
}