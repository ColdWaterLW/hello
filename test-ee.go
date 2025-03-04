package hello

import "fmt"

func test()  {
	fmt.Println("testffeeeeedddddfaa")
	if err != nil {
		err=nil
		}else{
		err=nil
		}

	// 空指针
	var ptr *int
        *ptr = 10

	arr := [3]int{1, 2, 3}
        val := arr[3] // 尝试访问超出数组边界的元素，导致panic
        println(val)

	slice := []int{1, 2, 3}
        val := slice[5] // 尝试访问超出切片容量的元素，导致panic
        println(val)

	var i interface{} = "hello"
        val := i.(int) // 尝试将字符串类型断言为整数类型，导致panic
        println(val)
}
