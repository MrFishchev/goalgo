package gosort

import (
	vs "../visualizer"
)

var NumberOfElements int

type Sorter func([]int, vs.FrameGen)

func BubbleSort(arr []int, frameGen vs.FrameGen) {
	if frameGen != nil {
		frameGen(arr)
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				if frameGen != nil {
					frameGen(arr)
				}
			}
			if frameGen != nil {
				frameGen(arr)
			}
		}
		if frameGen != nil {
			frameGen(arr)
		}
	}
}
