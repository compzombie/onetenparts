package main

import "strconv"

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

		state := 4*lef + 2*mid + 1*rit

		out += strconv.Itoa(state)
	}
	return out
}
