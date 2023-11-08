package objectives

func Obj1() []string {
	var menu []string
	var result []string
	menu = append(menu, "Hamburger")
	menu = append(menu, "Salad")
	for i := 0; i < len(menu); i++ {
		//fmt.Println("Food: " + Menu[i])
		result = append(result, "Food: "+menu[i])
	}
	//result = append(result, "Food: beans")
	return result
}
