package point_demo

import "fmt"

func TestPoint() {
	var count int = 20
	var countPoint *int
	countPoint = &count
	fmt.Printf("count  address ： %x \n", &count)
	fmt.Printf("countPoint address : %x \n", countPoint)
	fmt.Printf("countPoint 指针指向地址的值 : %d\n", *countPoint)

	var nilPoint *int
	fmt.Println("nilPoint %d", nilPoint)
	fmt.Printf("nilPoint 指针指向地址的值 : %d\n", nilPoint)

	if nilPoint == nil {
		fmt.Println("is nil  ")
	}

}

// 指针数组
func TestPointArr() {
	a, b := 1, 2
	pointArrary := [...]*int{&a, &b}
	fmt.Println("指针数组", pointArrary)
}

// 数组指针
func TestArrPoint() {
	arr := [...]int{2, 3, 4}
	arrPoint := &arr

	fmt.Println("数组指针 arrPoint", arrPoint)

}
