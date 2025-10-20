package data

import "fmt"

var Countries [10]string

func init() {
	Countries[0] = "Argentina"
	Countries[1] = "jamaica"
	Countries[2] = "Naija"
	Countries[3] = "GHana"
	Countries[7] = "Portugal"

	fmt.Println("Countries Saved ")
}
