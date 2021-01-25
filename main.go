package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"strings"
	"time"

	"./gosort"
	vs "./visualizer"
)

//region Private Methods

//Return available algorithms from a map
func getNamesOfAlgorithms(m map[string]gosort.Sorter) string {
	names := make([]string, 0, len(m))
	for name := range m {
		names = append(names, name)
	}
	return strings.Join(names, " | ")
}

func getVisualizer(mode string) vs.Visualizer {
	switch mode {
	case "stdout":
		return vs.FrameGen(vs.WriteStdout)
	case "gif":
		return new(vs.GifVisualizer)
	default:
		return nil
	}
}

//Returns an array of N elements with unificated values
//from 0 to Max
func generateRandomArray(count int, max int) []int {
	result := make([]int, count)

	for i := 0; i < count; i++ {
		b := make([]byte, 1)
		rand.Read(b) //get random value from 0 to 255
		number := float64(b[0])

		//unificating values from 0 to max
		result[i] = int(number / 255 * float64(max))

		if result[i] == 0 {
			result[i]++
		}
	}

	return result
}

//Runs algorithm and prepares a visualizer
func runSort(visMode string, algo string, sortingFunc gosort.Sorter) {
	visualizer := getVisualizer(visMode)
	visualizer.Setup(algo)
	dataForSorting := generateRandomArray(gosort.NumberOfElements, vs.MaxHeight)

	sortingFunc(dataForSorting, visualizer.AddFrame)
	visualizer.Complete()
}

//endregion

func main() {
	var selectedAlgo string
	var selectedOutput string

	algoMap := map[string]gosort.Sorter{
		"bubble":    gosort.BubbleSort,
		"selection": gosort.SelectionSort,
		"insertion": gosort.InsertionSort,
		"bogo":      gosort.BogoSort,
		"cocktail":  gosort.CocktailSort,
		"shell":     gosort.ShellSort,
	}

	flag.StringVar(&selectedAlgo, "a",
		"bubble", "Select sorting algorithm [all | "+getNamesOfAlgorithms(algoMap)+"]")
	flag.IntVar(&vs.Fps, "f", 10, "Frames per second")
	flag.IntVar(&vs.Mode, "v", 1, "Visualization mode 1-dots 2-solid")
	flag.IntVar(&vs.MaxHeight, "m", 9, "Highest value")
	flag.IntVar(&gosort.NumberOfElements, "c", 30, "Number of values")
	flag.StringVar(&selectedOutput, "o", "stdout", "Select output [stdout]/gif")
	flag.Parse()

	fmt.Printf("Sorting by: %v\nNumber of elements: %v\nHighest value: %v\n",
		selectedAlgo, gosort.NumberOfElements, vs.MaxHeight)

	sortingFunc := algoMap[selectedAlgo]
	if sortingFunc != nil {
		runSort(selectedOutput, selectedAlgo, sortingFunc)
	} else if selectedAlgo == "all" {
		for name, sortFunc := range algoMap {

			if name == "bogo" {
				continue
			}

			runSort(selectedOutput, name, sortFunc)
			fmt.Println(name)

			if selectedOutput == "stdout" {
				time.Sleep(1 * time.Second)
			}
		}
	} else {
		flag.PrintDefaults()
	}
}
