package main

import (
	"flag"
	"fmt"
)
//set flags for program
func declareFlags() {
	//TODO make this a path
	flagTxt := flag.Bool("txt", false, "output path for a txt file")

	//TODO make this a path
	flagPng := flag.Bool("png", false, "output path for a png file")
 
	//graph output - binary/state/domain/all
	flagOut := flag.String("out", "b", "sets output parameters -(b)inary -(s)tate")

	//lattice - toroidal, zero edged, 1 edged, unbounded
	flagLat := flag.String("lat", "t", "sets lattice properties -(t)oroidal -(z)ero-edged -(o)ne-edged")

	//number of generations to compute
	flagGens := flag.Int("gens", 10, "sets number of generations to compute before stopping")

	//initial condition of 1's and 0's
	flagIc := flag.String("ic", "00000100000", "initial condition of ca")
	flag.Parse()

	flagOutCheck(flagOut)
	flagLatCheck(flagLat)
	flagGenCheck(flagGens)
	flagIcCheck(*flagIc)
}

//flag can only consist of several distinct options
func flagOutCheck(flagOut *string) {
	//check submitted flags
	if *flagOut != "b" && *flagOut != "s"{
		err := InputError{"-out flag input malformed."}
		fmt.Println(err)
		panic(err)
	}
}

//flag can only consist of several distinct options
func flagLatCheck(flagLat *string) {
	//check submitted flags
	if *flagLat != "t" && *flagLat != "z" && *flagLat != "o" {
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
	declareFlags()



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
		}
	}

	//history := make(map[int][]string)

	//text output
	if *flagTxt {
		PrintTimeStamp(flagOut, flagLat, flagGens, flagIc)

		switch *flagOut {
		case "b":
			PrintBinaryFilter(lattice)

		case "s":
			PrintStateFilter(lattice)

		}
	}

	if *flagPng {
		switch *flagOut {
		case "b":
			PngBinaryFilter(lattice)
		case "s":
			PngStateFilter(lattice)
		}
	}

}
