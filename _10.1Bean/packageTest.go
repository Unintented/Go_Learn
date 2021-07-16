package bean

import "fmt"

var V1 float32 = 42
var v2 int8 = 10

func TestPackageVisit() float32 {
	fmt.Printf("Data from packageTest %d\n", v2)
	return V1
}
