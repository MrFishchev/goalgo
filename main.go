package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"strings"

	"./gosort"
	vs "./visualizer"
)

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
		return &vs.GifVisualizer{}
	default:
		return nil
	}
}

func generateRandomArray(count int, max int) []int {
	result := make([]int, count)

	for i := 0; i < count; i++ {
		b := make([]byte, 1)
		rand.Read(b)
		number := float64(b[0])
		result[i] = int(number / 255 * float64(max))
	}

	return result
}

func runSort(visMode string, algo string, sortingFunc gosort.Sorter) {
	visualizer := getVisualizer(visMode)
	visualizer.Setup(algo)
	dataForSorting := generateRandomArray(gosort.NumberOfElements, vs.MaxHeight)

	sortingFunc(dataForSorting, visualizer.AddFrame)
	visualizer.Complete()

	fmt.Println("Done!")
}

func main() {
	var selectedAlgo string
	var selectedOutput string

	algoMap := map[string]gosort.Sorter{
		"bubble": gosort.BubbleSort,
	}

	flag.StringVar(&selectedAlgo, "a",
		"bubble", "Select sorting algorithm [all | "+getNamesOfAlgorithms(algoMap)+"]")
	flag.IntVar(&vs.Fps, "f", 10, "Frames per second")
	flag.IntVar(&vs.Mode, "v", 1, "Visualization mode")
	flag.IntVar(&vs.MaxHeight, "m", 9, "Highest value")
	flag.IntVar(&gosort.NumberOfElements, "c", 30, "Number of values")
	flag.StringVar(&selectedOutput, "o", "stdout", "Select output [stdout]/gif")
	flag.Parse()

	fmt.Printf("Sorting by: %v\nNumber of elements: %v\nHighest value: %v\n",
		selectedAlgo, gosort.NumberOfElements, vs.MaxHeight)

	sortingFunc := algoMap[selectedAlgo]
	if sortingFunc != nil {
		runSort(selectedOutput, selectedAlgo, sortingFunc)
	}
}
