package aoc

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) {
	var scanners []scanner
	p := -1
	for i := range data {
		if strings.HasPrefix(data[i], "---") {
			m := scanner{}
			scanners = append(scanners, m)
			p++
			continue
		}
		if data[i] == "" {
			continue
		}
		s := strings.Split(data[i], ",")
		num, err := strconv.Atoi(s[0])
		if err != nil {
			log.Fatal(err)
		}
		num2, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}
		num3, err := strconv.Atoi(s[2])
		if err != nil {
			log.Fatal(err)
		}
		scanners[p].beacons = append(scanners[p].beacons, [3]int{num, num2, num3})
	}
	// for i := range scanners {
	// 	fmt.Println(scanners[i])
	// }
	fmt.Println(solve(scanners))
}

// Copied from: https://gist.github.com/znkr/208741c3da17bea9b319620837f59cd9

type point [3]int

type scanner struct {
	beacons []point
}

type projections [6][]int

func computeProjections(s scanner) projections {
	var p projections
	for i := range p {
		p[i] = make([]int, len(s.beacons))
	}
	for i, b := range s.beacons {
		p[0][i] = b[0]
		p[1][i] = b[1]
		p[2][i] = b[2]
		p[3][i] = -b[0]
		p[4][i] = -b[1]
		p[5][i] = -b[2]
	}
	for i := range p {
		sort.Ints(p[i])
	}
	return p
}

// findTranslations finds translations for which s and t have more than 12 matches
func findTranslations(s, t []int, res []int) []int {
	sn := len(s)
	tn := len(t)
	for si0 := 0; si0 < sn-11; si0++ {
		for ti0 := 0; ti0 < tn-11; ti0++ {
			matches := 0
			sv0, tv0 := s[si0], t[ti0]
			tr := sv0 - tv0
			si := si0
			ti := ti0
			for si < sn && ti < tn {
				sv, tv := s[si], t[ti]+tr
				if sv < tv {
					si++
				} else if tv < sv {
					ti++
				} else {
					matches++
					si++
					ti++
				}
			}

			if matches >= 12 {
				res = append(res, tr)
			}
		}
	}
	return res
}

type transform [12]int

var idTransform = transform{
	1, 0, 0, 0,
	0, 1, 0, 0,
	0, 0, 1, 0,
}

func (t transform) transform(p point) point {
	return point{
		t[0]*p[0] + t[1]*p[1] + t[2]*p[2] + t[3],
		t[4]*p[0] + t[5]*p[1] + t[6]*p[2] + t[7],
		t[8]*p[0] + t[9]*p[1] + t[10]*p[2] + t[11],
	}
}

func (t transform) inverse() transform {
	return transform{
		t[0], t[4], t[8], -t[0]*t[3] - t[4]*t[7] - t[8]*t[11],
		t[1], t[5], t[9], -t[1]*t[3] - t[5]*t[7] - t[9]*t[11],
		t[2], t[6], t[10], -t[2]*t[3] - t[6]*t[7] - t[10]*t[11],
	}
}

func (t transform) then(s transform) transform {
	return transform{
		s[0]*t[0] + s[1]*t[4] + s[2]*t[8],
		s[0]*t[1] + s[1]*t[5] + s[2]*t[9],
		s[0]*t[2] + s[1]*t[6] + s[2]*t[10],
		s[0]*t[3] + s[1]*t[7] + s[2]*t[11] + s[3],

		s[4]*t[0] + s[5]*t[4] + s[6]*t[8],
		s[4]*t[1] + s[5]*t[5] + s[6]*t[9],
		s[4]*t[2] + s[5]*t[6] + s[6]*t[10],
		s[4]*t[3] + s[5]*t[7] + s[6]*t[11] + s[7],

		s[8]*t[0] + s[9]*t[4] + s[10]*t[8],
		s[8]*t[1] + s[9]*t[5] + s[10]*t[9],
		s[8]*t[2] + s[9]*t[6] + s[10]*t[10],
		s[8]*t[3] + s[9]*t[7] + s[10]*t[11] + s[11],
	}
}

var prototransforms = [][4]int{
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 1, 0},
	{-1, 0, 0, 0},
	{0, -1, 0, 0},
	{0, 0, -1, 0},
}

func match(s, t projections, res []transform) []transform {
	var ptsbuf [3][1][4]int
	var trsbuf [1]int
	var pts [3][][4]int
	pts[0] = ptsbuf[0][:0]
	pts[1] = ptsbuf[1][:0]
	pts[2] = ptsbuf[2][:0]
	for i, st := range s[:3] {
		for j, tt := range t {
			trs := findTranslations(st, tt, trsbuf[:0])
			for _, tr := range trs {
				pt := prototransforms[j]
				pt[3] = tr
				pts[i] = append(pts[i], pt)
			}
		}
	}

	for _, x := range pts[0] {
		for _, y := range pts[1] {
			for _, z := range pts[2] {
				res = append(res, transform{
					x[0], x[1], x[2], x[3],
					y[0], y[1], y[2], y[3],
					z[0], z[1], z[2], z[3],
				})
			}
		}
	}

	return res
}

type key struct {
	from, to int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve(scanners []scanner) (int, int) {
	var wg sync.WaitGroup

	proj := make([]projections, len(scanners))
	wg.Add(len(scanners))
	for i, s := range scanners {
		i, s := i, s
		go func() {
			proj[i] = computeProjections(s)
			wg.Done()
		}()
	}
	wg.Wait()

	m := make(map[key]transform)
	var mmtx sync.Mutex
	for i0 := 0; i0 < len(scanners); i0++ {
		for i1 := 0; i1 < i0; i1++ {
			wg.Add(1)
			i0, i1 := i0, i1
			go func() {
				var buf [8]transform
				t := match(proj[i0], proj[i1], buf[:0])
				// find correct transform
				for _, t := range t {
					count := 0
					for _, b0 := range scanners[i0].beacons {
						for _, b1 := range scanners[i1].beacons {
							if b0 == t.transform(b1) {
								count++
							}
						}
					}
					if count >= 12 {
						mmtx.Lock()
						m[key{from: i1, to: i0}] = t
						m[key{from: i0, to: i1}] = t.inverse()
						mmtx.Unlock()
						break
					}
				}
				wg.Done()
			}()
		}
	}
	wg.Wait()

	// BFS to find transforms into scanner0 coordinate space
	q := append(make([]int, 0, len(scanners)), 0)
	transforms := make([]*transform, len(scanners))
	transforms[0] = &idTransform
	for len(q) > 0 {
		k := q[0]
		q = q[1:]
		for i := 0; i < len(scanners); i++ {
			if transforms[i] != nil {
				continue
			}
			if tik, ok := m[key{from: i, to: k}]; ok {
				tk0 := transforms[k]
				ti0 := tik.then(*tk0)
				transforms[i] = &ti0
				q = append(q, i)
			}
		}
	}

	for _, t := range transforms {
		if t == nil {
			panic("not solveable")
		}
	}

	numBeacons := 0
	maxDist := 0

	wg.Add(2)
	go func() {
		// Solve part 1
		normalizedBeacons := make(map[point]bool, 500)
		for i, s := range scanners {
			t := transforms[i]
			for _, b := range s.beacons {
				normalizedBeacons[t.transform(b)] = true
			}
		}
		numBeacons = len(normalizedBeacons)
		wg.Done()
	}()
	go func() {
		// Solve part 2
		scannerPos := make([]point, len(scanners))
		for i, t := range transforms {
			scannerPos[i] = t.transform(point{})
		}
		for _, p := range scannerPos {
			for _, q := range scannerPos {
				d := abs(p[0]-q[0]) + abs(p[1]-q[1]) + abs(p[2]-q[2])
				if d > maxDist {
					maxDist = d
				}
			}
		}
		wg.Done()
	}()
	wg.Wait()

	return numBeacons, maxDist
}
