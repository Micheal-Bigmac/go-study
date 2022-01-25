package main

import (
	"awesomeProject/interface_demo"
	"awesomeProject/json_demo"
	"errors"
	"fmt"
	"reflect"
	"time"
)

func main() {
	//makeSlice()
	//makeMap()
	//makeChannel()

	//appendElementForSlice()

	//copyForSlice()
	//newMap()

	//deleteFromMap()
	//receivePanic()
	//getLen()
	//closeChannel()
	//testForStruct
	//struct_demo.TestForStruct()
	//struct_demo.TestForStruct2()
	//struct_demo.TestForStruct3()
	//one := new(struct_demo.SelftDeft)
	//one.ID = 0
	//one.Name = "string"
	//one.Age = 2
	//
	//one.Run()
	//dog := new(oritend_object.Dog)
	//dog.Id = 1
	//dog.Name = "kitty"
	//dog.Age = 10
	//dog.Colour = "red"
	//dog.Run()
	//dog.Eat()
	//dog := new(oritend_object.Dog)
	//dog.Id = 0
	//dog.Name = "dog"
	//dog.Colour = "black"
	//
	//action(dog)
	//cat := new(oritend_object.Cat)
	//cat.Id = 1
	//cat.Name = "cat"
	//cat.Colour = "white"
	//action(cat)

	// 并发

	/*fmt.Printf("%d", runtime.NumCPU())
	runtime.GOMAXPROCS(2)*/

	/*//启动发送数据协程
	go goruntime.Send()
	//启动接受数据协程
	go goruntime.Receive()*/

	/*	// 测试协程的同步
		goruntime.Read()     // 主线程读取文件
		go goruntime.Write() // 开启协程去写入
		goruntime.SWG.Wait() // 相当于Java 的wait ，等待写的协程全部完成
		fmt.Println("All done !")*/

	/*// 测试指针
	point_demo.TestPoint()*/

	/*// 测试指针数组，数组指针
	point_demo.TestPointArr()
	point_demo.TestArrPoint()*/

	/*//测试 json 序列化
	json_demo.TestSerializeStruct()
	json_demo.TestSerializeMap()*/

	// 测试json 反序列化
	json_demo.TestDeserializeStruct()
	//json_demo.TestDeSerializeMap()
	time.Sleep(time.Second * 60)
}

func action(b interface_demo.Behavior) string {
	b.Run()
	b.Eat()
	return ""
}

// defer
func closeChannel() {
	//defer coverPanic()
	mChan := make(chan int, 1)
	defer close(mChan)
	mChan <- 1 // 往channel 写数据
	//close(mChan)
	//	 bussiness logical
}

func getLen() {
	mSlice := make([]string, 2, 5)
	mSlice[0] = "111"
	mSlice[1] = "222"
	fmt.Println("len : ", len(mSlice))
	fmt.Println("cap : ", cap(mSlice))
}
func deleteFromMap() {
	mMap := make(map[int]string)
	mMap[1] = "11"
	mMap[2] = "22"
	mMap[3] = "33"

	delete(mMap, 1)
	delete(mMap, 2)
	fmt.Println(mMap)

	receivePanic()
}

func receivePanic() {
	//defer func() {
	//	message := recover()
	//	fmt.Println("catch error : ", message)
	//}()
	defer coverPanic()
	//panic(" throws exception ")

	panic(errors.New("a error "))
}

//封装异常捕获
func coverPanic() {
	message := recover()
	switch message.(type) {
	case string:
		fmt.Println("string message :", message)
	case error:
		fmt.Println("error message : ", message)
	default:
		fmt.Println("unknow painc  : ", message)
	}
}

// 测试 slice 复制
func copyForSlice() {
	sSlice := make([]string, 3)
	sSlice[0] = "one"
	sSlice[1] = "two"
	sSlice[2] = "three"

	dSlice := make([]string, 2)
	dSlice[0] = "three"
	dSlice[1] = "four"

	copy(sSlice, dSlice)

	fmt.Println(sSlice)

}

// 测试slice 扩容
func appendElementForSlice() {
	mSlice := make([]string, 2)
	mSlice[0] = "one"
	mSlice[1] = "two"
	fmt.Println("len : ", len(mSlice))
	fmt.Println("cap : ", cap(mSlice))

	mSlice = append(mSlice, "three")
	fmt.Println(mSlice)
	fmt.Println("after len :", len(mSlice))
	fmt.Println("after cap :", cap(mSlice))
}

// 使用new 方式创建Map
func newMap() {
	mMap := new(map[int]string)
	fmt.Println(reflect.TypeOf(mMap))
}

// 创建切片
func makeSlice() {
	mSlice := make([]string, 3)
	mSlice[0] = "dog"
	mSlice[1] = "cat"
	mSlice[2] = "tiger"

	fmt.Println(mSlice)

}

// 创建Map
func makeMap() {
	mMap := make(map[int]string)
	mMap[1] = "11"
	mMap[2] = "22"
	mMap[3] = "33"
	fmt.Println(mMap)
	fmt.Println(reflect.TypeOf(mMap))
}

//创建没有缓存的 channel
func makeChannel() {
	mChannel := make(chan int)
	close(mChannel)
}
