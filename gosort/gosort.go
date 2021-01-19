//Provides visualized algorithms
//for sorting and searching
package gosort

import (
	vs "../visualizer"
)

//region Members

var NumberOfElements int

type Sorter func([]int, vs.FrameGen)

//endregion

//region Public Methods

func BubbleSort(arr []int, frameGen vs.FrameGen) {
	generateFrame(arr, frameGen)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				generateFrame(arr, frameGen)
			}
			generateFrame(arr, frameGen)
		}
		generateFrame(arr, frameGen)
	}
}

//endregion

//region Private Methods

func generateFrame(arr []int, frameGen vs.FrameGen) {
	if frameGen != nil {
		frameGen(arr)
	}
}

//endregion
