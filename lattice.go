package main

import (
	"strconv"
)

func InitLattice(ic string) *[]string {
	lattice := []string{}

	//apply initial condition
	lattice = append(lattice, ic)

	return &lattice
}

//toroidal
func ParseRowToroid(row string) string {
	out := ""
	for i := 0; i < len(row); i++ {
		leftdex, middex, ritdex := -1, -1, -1

		//toroid wrapping
		if (i - 1) < 0 {
			leftdex = len(row) - 1
		} else {
			leftdex = i - 1
		}

		middex = i

		if (i + 1) >= len(row) {
			ritdex = 0
		} else {
			ritdex = i + 1
		}

		//state of each neighbor
		lef, lerr := strconv.Atoi(string(row[leftdex]))
		mid, merr := strconv.Atoi(string(row[middex]))
		rit, rerr := strconv.Atoi(string(row[ritdex]))
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
