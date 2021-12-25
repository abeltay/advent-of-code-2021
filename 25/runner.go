package aoc

const (
	east  = 1
	south = 2
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	arr := make([][]int, len(data))
	for i := range data {
		arr2 := make([]int, len(data[i]))
		for j := range data[i] {
			switch data[i][j] {
			case '>':
				arr2[j] = east
			case 'v':
				arr2[j] = south
			}
		}
		arr[i] = arr2
	}
	// for i := range arr {
	// 	fmt.Println(arr[i])
	// }
	move := true
	var ans int
	for move {
		var movedEast, movedSouth bool
		arr, movedEast = moveEast(arr)
		// fmt.Println("----------")
		// for i := range arr {
		// 	fmt.Println(arr[i])
		// }
		arr, movedSouth = moveSouth(arr)
		// fmt.Println("----------")
		// for i := range arr {
		// 	fmt.Println(arr[i])
		// }
		move = movedEast || movedSouth
		ans++
	}
	return ans
}

func moveEast(data [][]int) ([][]int, bool) {
	var moved bool
	arr := make([][]int, len(data))
	for i := range data {
		arr2 := make([]int, len(data[i]))
		for j := range data[i] {
			if data[i][j] == east {
				next := j + 1
				if next >= len(data[i]) {
					next = 0
				}
				if data[i][next] == 0 {
					arr2[next] = data[i][j]
					moved = true
				} else {
					arr2[j] = data[i][j]
				}
				continue
			}
			if data[i][j] > 0 {
				arr2[j] = data[i][j]
			}
		}
		arr[i] = arr2
	}
	return arr, moved
}

func moveSouth(data [][]int) ([][]int, bool) {
	var moved bool
	arr := make([][]int, len(data))
	for i := range data {
		arr2 := make([]int, len(data[i]))
		arr[i] = arr2
	}
	for j := range data[0] {
		for i := range data {
			if data[i][j] == south {
				next := i + 1
				if next >= len(data) {
					next = 0
				}
				if data[next][j] == 0 {
					arr[next][j] = data[i][j]
					moved = true
				} else {
					arr[i][j] = data[i][j]
				}
				continue
			}
			if data[i][j] > 0 {
				arr[i][j] = data[i][j]
			}
		}
	}
	return arr, moved
}
