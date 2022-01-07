package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

type InputError struct{ Errstr string }

func (i *InputError) Error() string {
	return i.Errstr
}

func rule110() []int {
	return []int{0, 1, 1, 1, 0, 1, 1, 0}
}

//flag can only consist of several distinct options
func flagOutCheck(flagOut *string) {
	//check submitted flags
	if *flagOut != "b" && *flagOut != "s" && *flagOut != "d" && *flagOut != "a" {
		err := InputError{"-out flag input malformed."}
		fmt.Println(err)
		panic(err)
	}
}

//flag can only consist of several distinct options
func flagLatCheck(flagLat *string) {
	//check submitted flags
	if *flagLat != "t" && *flagLat != "z" && *flagLat != "o" && *flagLat != "u" && *flagLat != "a" {
		err := InputError{"-lat flag input malformed."}
		fmt.Println(*flagLat)
		fmt.Println(err)
		panic(err)
	}
}

//an ic can only consist of <0>'s and <1>'s
func flagIcCheck(ic string) {
	for i := 0; i < len(ic); i++ {
		cell := ic[i]
		if cell != '0' && cell != '1' {
			err := InputError{"-ic flag input malformed."}
			fmt.Println(err)
			panic(err)
		}
	}
}

//gens must be greater than 0 and a valid integer
func flagGenCheck(gens *int) {
	if *gens <= 0 {
		err := InputError{"gen must be greater than 0 and a valid integer"}
		fmt.Println(err)
		panic(err)
	}
}

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

func main() {
	//graph output - binary/state/domain/all
	flagOut := flag.String("out", "b", "sets output parameters -(b)inary -(s)tate -(d)omain -(a)ll")

	//lattice - toroidal, zero edged, 1 edged, unbounded
	flagLat := flag.String("lat", "t", "sets lattice properties -(t)oroidal -(z)ero-edged -(o)ne-edged -(u)nbounded -(a)ll")

	//number of generations to compute
	flagGens := flag.Int("gens", 10, "sets number of generations to compute before stopping")

	//initial condition of 1's and 0's
	flagIc := flag.String("ic", "00000100000", "initial condition of ca")
	flag.Parse()

	flagOutCheck(flagOut)
	flagLatCheck(flagLat)
	flagGenCheck(flagGens)
	flagIcCheck(*flagIc)

	//generate state ca and work off of that
	lattice := InitLattice(*flagIc)
	for i := 1; i < *flagGens; i++ {
		*lattice = append(*lattice, ParseRowToroid((*lattice)[0]))
	}

	//output

	fmt.Println("Program Arguments")
	fmt.Println("-out: " + *flagOut)
	fmt.Println("-lat: " + *flagLat)
	fmt.Println("-gens" + strconv.Itoa(*flagGens))
	fmt.Println("IC: " + *flagIc)
	fmt.Println(time.Now())
	fmt.Println(lattice)

}
