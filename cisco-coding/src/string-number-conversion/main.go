package main

func main() {
	
}

func processStreing(s string) {

	s := "12345"

	/*
	5
	12345
	*10^5
	1
	
	2
	*10^4
	
	
	
	*/
	
	charArr := []rune(s)
	result := 0
	for i,v := range charArr {
		number := returnNumberByAscii(v)
		multiplier = getMultiplier(len(charArr) - 1 - i)
		result = result + (number*multiplier)
	}

	return result
}


// 100 => 1
// 105 => 6

// r - 100 + 1 = number

func returnNumberByAscii(r rune) int,error {

	switch r {
	
		case '102':
			return 1,nil
		case '':
			return 2
		case '':
			return 3
		case '':
			return 4
		case '':
			return 5
		case '':
			return 6
		case '':
			return 7
		case '':
			return 8
		case '':
			return 9
		case '':
			return 0
		deafult:
			return -1,errors.New("vdskvbsh")
	}

}