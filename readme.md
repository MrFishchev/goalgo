# goalgo
Visualization of algorithms (sorting, serching) in Go language

- [ ] Search visualizer
- [ ] Barley-break visualizer (https://github.com/Bjarten/alvito)
- [ ] More sorting algorithms
- [x] Sorting visualizer

## Sorting Algorithms Visualization

### BogoSort

[![Bogo Sort Visualization](https://mrfishchev.github.io/goalgo/gosort/demo/bogo.gif)](https://en.wikipedia.org/wiki/Bogosort) 

### BubbleSort

[![Bubble Sort Visualization](https://mrfishchev.github.io/goalgo/gosort/demo/bubble.gif)](https://en.wikipedia.org/wiki/Bubble_sort) 

### CocktailSort

[![Cocktail Sort Visualization](https://mrfishchev.github.io/goalgo/gosort/demo/cocktail.gif)](https://en.wikipedia.org/wiki/Cocktail_shaker_sort) 

### InsertionSort

[![Insertion Sort Visualization](https://mrfishchev.github.io/goalgo/gosort/demo/insertion.gif)](https://en.wikipedia.org/wiki/Insertion_sort) 

### SelectionSort

[![Selection Sort Visualization](https://mrfishchev.github.io/goalgo/gosort/demo/selection.gif)](https://en.wikipedia.org/wiki/Selection_sort) 

### ShellSort

[![Shell Sort Visualization](https://mrfishchev.github.io/goalgo/gosort/demo/shell.gif)](https://en.wikipedia.org/wiki/Shellsort)

## How To Run

$ go run main.go
    -a string
            Select sorting algorithm [all | bubble | ...] (default "bubble")
    -c int
            Number of values (default 30)
    -f int
            Frames per second (default 10)
    -m int
            Highest value (default 9)
    -o string
            Select output stdout/gif (default "stdout")
    -v int
            Visualization mode 1-dots 2-solid (default 1)

## License

[MIT](https://github.com/mrfishchev/goalgo/blob/main/LICENSE)