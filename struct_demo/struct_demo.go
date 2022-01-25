package struct_demo

import "fmt"

// 顶一个一个结构体
type SelftDeft struct {
	Name string // 属性
	Age  int
	ID   int
}

func TestForStruct() {
	var one SelftDeft
	one.Name = "111"
	one.Age = 2
	one.ID = 0

	fmt.Println(one)
}

func TestForStruct2() {
	one := SelftDeft{ID: 0, Name: "strign", Age: 1}
	fmt.Println(one)
}

// new struct
func TestForStruct3() {
	one := new(SelftDeft)
	one.ID = 0
	one.Name = "22"
	one.Age = 4
	fmt.Println(one)
}

func (a *SelftDeft) Run() {
	fmt.Println("ID : ", a.ID, "is running ")
}

func (a *SelftDeft) run() {
	fmt.Println("ID : ", a.ID, "is running ")
}
