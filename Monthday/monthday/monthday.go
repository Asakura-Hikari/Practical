package monthday

func Monthday(year, month int) int {
	var day int

	switch month {
	case 1:
		day = 31
	case 2:
		if year%4 == 0 && year%100 != 0 {
			day = 29
		} else {
			day = 28
		}
	case 3:
		day = 31
	case 4:
		day = 30
	case 5:
		day = 31
	case 6:
		day = 30
	case 7:
		day = 31
	case 8:
		day = 31
	case 9:
		day = 30
	case 10:
		day = 31
	case 11:
		day = 30
	case 12:
		day = 31
	default:
		println("Invalid Day")
	}

	return day
}