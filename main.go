package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

func Rule110() []int {
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
		prevRow := (*lattice)[i-1]

		switch *flagLat {
		case "t":
			*lattice = append(*lattice, ParseRowToroid(prevRow))
		case "z":
			*lattice = append(*lattice, ParseRowZeroEdge(prevRow))
		case "o":
			*lattice = append(*lattice, ParseRowOneEdge(prevRow))
		case "u":
		case "a":
		default:
			*lattice = append(*lattice, "-")
		}
	}

	//output
	switch *flagOut {
	case "b":
		binlat := BinaryFilter(*lattice)
		for i := 0; i < len(*binlat); i++ {
			fmt.Println((*binlat)[i])
		}
	case "s":
		for i := 0; i < len(*lattice); i++ {
			fmt.Println((*lattice)[i])
		}
	case "d":
		//I have to figure out some algorithms to convert state
		//data into domains and particles
	case "a":
	}

	fmt.Println("Program Arguments")
	fmt.Println("-out: " + *flagOut)
	fmt.Println("-lat: " + *flagLat)
	fmt.Println("-gens: " + strconv.Itoa(*flagGens))
	fmt.Println("IC: " + *flagIc)
	fmt.Println(time.Now())

}
