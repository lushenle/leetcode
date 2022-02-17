// https://zh.wikipedia.org/wiki/%E8%94%A1%E5%8B%92%E5%85%AC%E5%BC%8F

package zeller

func zeller(year, month, day int) int {
	if month == 1 {
		month = 13
		year = year - 1
	}

	if month == 2 {
		month = 14
		year = year - 1
	}

	c := year / 100
	year = year % 100
	return (year + (year / 4) + (c / 4) - 2*c + (26 * (month + 1) / 10) + day + 700 - 1) % 7
}

//func main() {
//	weekday := make(map[int]string)
//
//	weekday = map[int]string{
//		0: "星期日",
//		1: "星期一",
//		2: "星期二",
//		3: "星期三",
//		4: "星期四",
//		5: "星期五",
//		6: "星期六",
//	}
//
//	var (
//		year  int
//		month int
//		day   int
//	)
//
//	fmt.Println("enter year: ")
//	fmt.Scan(&year)
//	fmt.Println("enter month: ")
//	fmt.Scan(&month)
//	fmt.Println("enter day: ")
//	fmt.Scan(&day)
//
//	w := Zeller(year, month, day)
//	fmt.Println(weekday[int(w)])
//}
