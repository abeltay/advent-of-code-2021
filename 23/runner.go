package aoc

import (
	"container/heap"
	"fmt"
)

// Runner runs the algorithm to get the answer
func Runner(data space) int {
	spaceMap := make(map[string]int)
	h := &spaceHeap{data}
	heap.Init(h)
	for h.Len() > 0 {
		// for i := 0; i < 10; i++ {
		s1 := heap.Pop(h)
		s2 := s1.(space)
		// printSpace(s2)
		cost := solve(s2, h, spaceMap)
		if cost != 0 {
			return cost
		}
	}
	return 0
}

func printSpace(s space) {
	fmt.Println("-----------")
	for i := range s.hallway {
		if s.hallway[i] == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(s.hallway[i])
		}
	}
	fmt.Println()
	for i := len(s.rooms[0]) - 1; i >= 0; i-- {
		fmt.Print(" ")
		for j := range s.rooms {
			fmt.Print(" ")
			if s.rooms[j][i] == -1 {
				fmt.Print(".")
			} else {
				fmt.Print(s.rooms[j][i])
			}
		}
		fmt.Println()
	}
	fmt.Println(s.cost)
}

type space struct {
	hallway []int
	rooms   [][]int
	cost    int
}

func solve(s space, h *spaceHeap, spaceMap map[string]int) int {
	var (
		cost    = []int{1, 10, 100, 1000}
		doorPos = []int{2, 4, 6, 8}
	)
	var notSolved bool
	for i := range s.rooms {
		for j := range s.rooms[i] {
			if s.rooms[i][j] != i {
				notSolved = true
				break
			}
		}
		if notSolved {
			break
		}
	}
	if !notSolved {
		return s.cost
	}

	// Hallway to home
	for i := range s.hallway {
		if s.hallway[i] > -1 {
			amphipod := s.hallway[i]
			// Check if we can move to home
			var cannotGoHome bool
			for j := range s.rooms[amphipod] {
				if s.rooms[amphipod][j] == -1 {
					break
				}
				if s.rooms[amphipod][j] != amphipod {
					cannotGoHome = true
					break
				}
			}
			if cannotGoHome {
				continue
			}
			// Move it home
			min := doorPos[amphipod]
			max := i
			if max < min {
				min, max = max, min
			}
			this := cloneSpace(s)
			// Count steps including intersection
			moves := 1
			var blocked bool
			for j := min + 1; j < max; j++ {
				if this.hallway[j] != -1 {
					blocked = true
					break
				}
				moves++
			}
			if blocked {
				continue
			}
			for j := range s.rooms[amphipod] {
				if s.rooms[amphipod][j] == -1 {
					this.rooms[amphipod][j], this.hallway[i] = this.hallway[i], this.rooms[amphipod][j]
					moves += len(this.rooms[amphipod]) - j
					break
				}
			}
			this.cost += moves * cost[amphipod]
			// printSpace(this)
			if shouldAdd(this, spaceMap) {
				heap.Push(h, this)
			}
			return 0
		}
	}

	// Home to hallway
	for i := range s.rooms {
		for j := range s.rooms[i] {
			if s.rooms[i][j] == -1 {
				break
			}
			if s.rooms[i][j] != i {
				spaces := moveRoomToHallway(s, i)
				for k := range spaces {
					// printSpace(spaces[k])
					if shouldAdd(spaces[k], spaceMap) {
						heap.Push(h, spaces[k])
					}
				}
				break
			}
		}
	}
	// printSpace(s)
	return 0
}

func moveRoomToHallway(s space, room int) []space {
	var (
		cost    = []int{1, 10, 100, 1000}
		doorPos = []int{2, 4, 6, 8}
	)
	loc := len(s.rooms[room]) - 1
	for ; loc >= 0; loc-- {
		if s.rooms[room][loc] != -1 {
			break
		}
	}
	if loc < 0 {
		return nil
	}
	var newSpaces []space
	moves := len(s.rooms[room]) - loc
	// Move it to open space
	for i := doorPos[room] + 1; i < len(s.hallway); i++ {
		moves++
		if i == 2 || i == 4 || i == 6 || i == 8 {
			continue
		}
		if s.hallway[i] != -1 {
			break
		}
		this := cloneSpace(s)
		this.rooms[room][loc], this.hallway[i] = this.hallway[i], this.rooms[room][loc]
		this.cost += moves * cost[s.rooms[room][loc]]
		newSpaces = append(newSpaces, this)
	}
	moves = len(s.rooms[room]) - loc
	for i := doorPos[room] - 1; i >= 0; i-- {
		moves++
		if i == 2 || i == 4 || i == 6 || i == 8 {
			continue
		}
		if s.hallway[i] != -1 {
			break
		}
		this := cloneSpace(s)
		this.rooms[room][loc], this.hallway[i] = this.hallway[i], this.rooms[room][loc]
		this.cost += moves * cost[s.rooms[room][loc]]
		newSpaces = append(newSpaces, this)
	}
	return newSpaces
}

func cloneSpace(s space) space {
	var s2 space
	s2.hallway = make([]int, len(s.hallway))
	for i := range s.hallway {
		s2.hallway[i] = s.hallway[i]
	}

	s2.rooms = make([][]int, len(s.rooms))
	for i := range s.rooms {
		room := make([]int, len(s.rooms[i]))
		for j := range s.rooms[i] {
			room[j] = s.rooms[i][j]
		}
		s2.rooms[i] = room
	}
	s2.cost = s.cost
	return s2
}

func shouldAdd(s space, spaceMap map[string]int) bool {
	str := fmt.Sprintf("%v%v", s.hallway, s.rooms)
	v, ok := spaceMap[str]
	if !ok {
		spaceMap[str] = s.cost
		return true
	}
	if s.cost < v {
		spaceMap[str] = s.cost
		return true
	}
	return false
}

type spaceHeap []space

func (h spaceHeap) Len() int           { return len(h) }
func (h spaceHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h spaceHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *spaceHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(space))
}

func (h *spaceHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
