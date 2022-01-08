package main

func BinaryFilter(lattice []string) *[]string {
	out := []string{}
	for _, s := range lattice {

		newS := ""
		for i := 0; i < len(s); i++ {
			switch s[i] {
			case '0':
				newS += "0"
			case '1':
				newS += "1"
			case '2':
				newS += "1"
			case '3':
				newS += "1"
			case '4':
				newS += "0"
			case '5':
				newS += "1"
			case '6':
				newS += "1"
			case '7':
				newS += "0"
			default:
				newS += "-"
			}
		}
		out = append(out, newS)
	}

	return &out
}
