package main

import (
	"flag"
	"fmt"
)

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
	flagTxt := flag.Bool("txt", false, "output txt file")

	flagPng := flag.Bool("png", true, "output png file")

	//graph output - binary/state/domain/all
	flagOut := flag.String("out", "b", "sets output parameters -(b)inary -(s)tate -(d)omain -(a)ll")

	//lattice - toroidal, zero edged, 1 edged, unbounded
	flagLat := flag.String("lat", "t", "sets lattice properties -(t)oroidal -(z)ero-edged -(o)ne-edged -(u)nbounded -(a)ll")

	//number of generations to compute
	flagGens := flag.Int("gens", 100, "sets number of generations to compute before stopping")

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

	//history := make(map[int][]string)

	//text output
	if *flagTxt {
		PrintTimeStamp(flagOut, flagLat, flagGens, flagIc)

		switch *flagOut {
		case "b":
			PrintBinaryFilter(lattice)

		case "s":
			PrintStateFilter(lattice)

		case "d":
			/*
				//ALL OF THIS BELOW NEEDS TO GO SOMEWHERE ELSE LIKE PARSEDOMAIN()
				//get a collection of neighborhoods for each cell by state and time
				//by making an array of maps, one spot for each cell by index that matches in lattice
				//for each spot in map array: <lattice> index->[]string <neighbor history>

				switch *flagLat {
				case "t":
					//history = BuildToroidHist(lattice)
				case "z":
					//history = BuildZeroHist(lattice)
				case "o":
					//history = BuildOneHist(lattice)
				case "u":
					//BuildUnboundHist(lattice)
				}
			*/

		case "a":
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
