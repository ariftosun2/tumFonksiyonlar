package main



func StringParse(strmath string) map[int][]int {
	package_count := make(map[int][]int)

	for a := 0; a <= len(strmath); a++ {

		if strmath[a] == '(' {
			for b := 0; b <= len(strmath); b++ {

			}
		}
	}

	for i := 0; i <= len(strmath); i++ {
		thisNote := map[int]int{
			i: strmath[i],
		}

		package_count[i] = append(package_count[i], thisNote[i])
		start_count += consecutive_count

	}

}

func CalculateString(stringmath string) int {
	result := 0

}
