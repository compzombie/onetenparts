package main

import (
	"fmt"
	"strconv"
)

func Rule110() []int {
	return []int{0, 1, 1, 1, 0, 1, 1, 0}
}

func InitLattice(ic string) *[]string {
	lattice := []string{}

	//apply initial condition
	lattice = append(lattice, ic)

	return &lattice
}

func GetCellState(index int, row string) int {
	out, err := strconv.Atoi(string(row[index]))
	if err != nil {
		panic(err)
	}

	return out
}

func ComputeCellState(lefdex int, middex int, ritdex int) int {
	return 4*Rule110()[lefdex] + 2*Rule110()[middex] + 1*Rule110()[ritdex]
}

func GetToroidNeighbors(index int, row string) (int, int, int) {
	l, m, r := 0, 0, 0

	if (index - 1) < 0 {
		l = len(row) - 1
	} else {
		l = index - 1
	}

	m = index

	if (index + 1) >= len(row) {
		r = 0
	} else {
		r = index + 1
	}

	return l, m, r
}

func GetZeroNeighbors(index int, row string) (int, int, int) {
	l, m, r := 0, 0, 0

	//edges are 0's
	if (index - 1) < 0 {
		l = 0
	} else {
		l = index - 1
	}

	m = index

	if (index + 1) >= len(row) {
		r = 0
	} else {
		r = index + 1
	}

	return l, m, r
}

func GetOneNeighbors(index int, row string) (int, int, int) {
	l, m, r := 0, 0, 0

	//edges are 0's
	if (index - 1) < 0 {
		l = 1
	} else {
		l = index - 1
	}

	m = index

	if (index + 1) >= len(row) {
		r = 1
	} else {
		r = index + 1
	}

	return l, m, r
}

//toroidal
func ParseRowToroid(row string) string {
	out := ""
	for i := 0; i < len(row); i++ {
		leftdex, middex, ritdex := GetToroidNeighbors(i, row)

		//state of each neighbor
		lef := GetCellState(leftdex, row)
		mid := GetCellState(middex, row)
		rit := GetCellState(ritdex, row)

		out += strconv.Itoa(ComputeCellState(lef, mid, rit))
	}
	return out
}

func BuildToroidHist(lattice *[]string) map[int][]string {
	history := make(map[int][]string)
	for iter := 1; iter < len(*lattice); iter++ { //i=1 because skip IC
		for cell := 0; cell < len((*lattice)[iter]); cell++ {
			//left, middle, and right indices of previous iter
			l, m, r := cell-1, cell, cell+1

			if l < 0 {
				l = len((*lattice)[iter]) - 1
			}

			if r > len((*lattice)[iter])+1 {
				r = 0
			}

			history[cell] = append(history[cell], fmt.Sprintf("%v%v%v", l, m, r))
		}
	}

	return history
}

//zero-edged
func ParseRowZeroEdge(row string) string {
	out := ""
	for i := 0; i < len(row); i++ {

		//-1 because doesn't exist in lookup table
		leftdex, middex, ritdex := -1, -1, -1

		//edges are 0's
		if (i - 1) < 0 {
			leftdex = -1
		} else {
			leftdex = i - 1
		}

		middex = i

		if (i + 1) >= len(row) {
			ritdex = -1
		} else {
			ritdex = i + 1
		}

		//cells can only have a value that is present in the rules lookup
		//e.g. if there are no 0's, then no cell can ever take the value 0
		//prob should throw an error for this when setting up general rules
		safe_zero_dex := 0

		//state of each neighbor
		lef, mid, rit := safe_zero_dex, -1, safe_zero_dex
		lerr, merr, rerr := error(nil), error(nil), error(nil)

		if leftdex != -1 {
			lef, lerr = strconv.Atoi(string(row[leftdex]))
		}

		mid, merr = strconv.Atoi(string(row[middex]))

		if ritdex != -1 {
			rit, rerr = strconv.Atoi(string(row[ritdex]))
		}

		if lerr != nil {
			panic(lerr)
		}
		if merr != nil {
			panic(merr)
		}
		if rerr != nil {
			panic(rerr)
		}

		state := 4*Rule110()[lef] + 2*Rule110()[mid] + 1*Rule110()[rit]

		out += strconv.Itoa(state)
	}

	return out
}

func BuildZeroHist(lattice *[]string) map[int][]string {
	history := make(map[int][]string)
	for iter := 1; iter < len(*lattice); iter++ { //i=1 because skip IC
		for cell := 0; cell < len((*lattice)[iter]); cell++ {
			//left, middle, and right indices of previous iter
			l, m, r := cell-1, cell, cell+1

			if l < 0 {
				l = 0
			}

			if r > len((*lattice)[iter])+1 {
				r = 0
			}

			history[cell] = append(history[cell], fmt.Sprintf("%v%v%v", l, m, r))
		}
	}

	return history
}

//one-edged
func ParseRowOneEdge(row string) string {
	out := ""
	for i := 0; i < len(row); i++ {

		//-1 because doesn't exist in lookup table
		leftdex, middex, ritdex := -1, -1, -1

		//edges are 0's
		if (i - 1) < 0 {
			leftdex = -1
		} else {
			leftdex = i - 1
		}

		middex = i

		if (i + 1) >= len(row) {
			ritdex = -1
		} else {
			ritdex = i + 1
		}

		//cells can only have a value that is present in the rules lookup
		//e.g. if there are no 0's, then no cell can ever take the value 0
		//prob should throw an error for this when setting up general rules
		safe_one_dex := 1

		//state of each neighbor
		lef, mid, rit := safe_one_dex, -1, safe_one_dex
		lerr, merr, rerr := error(nil), error(nil), error(nil)

		if leftdex != -1 {
			lef, lerr = strconv.Atoi(string(row[leftdex]))
		}

		mid, merr = strconv.Atoi(string(row[middex]))

		if ritdex != -1 {
			rit, rerr = strconv.Atoi(string(row[ritdex]))
		}

		if lerr != nil {
			panic(lerr)
		}
		if merr != nil {
			panic(merr)
		}
		if rerr != nil {
			panic(rerr)
		}

		state := 4*Rule110()[lef] + 2*Rule110()[mid] + 1*Rule110()[rit]

		out += strconv.Itoa(state)
	}

	return out
}

func BuildOneHist(lattice *[]string) map[int][]string {
	history := make(map[int][]string)
	for iter := 1; iter < len(*lattice); iter++ { //i=1 because skip IC
		for cell := 0; cell < len((*lattice)[iter]); cell++ {
			//left, middle, and right indices of previous iter
			l, m, r := cell-1, cell, cell+1

			if l < 0 {
				l = 1
			}

			if r > len((*lattice)[iter])+1 {
				r = 1
			}

			history[cell] = append(history[cell], fmt.Sprintf("%v%v%v", l, m, r))
		}
	}

	return history
}
