package main

import "fmt"

func main() {
	var arr [10]int                                 // 声明了一个int类型的数组
	arr[0] = 42                                     // 数组下标是从0开始的
	arr[1] = 13                                     // 赋值操作
	fmt.Printf("The first element is %d\n", arr[0]) // 获取数据，返回42
	fmt.Printf("The last element is %d\n", arr[9])  //返回未赋值的最后一个元素，默认返回0

	// 数组可以使用另一种:=来声明
	a := [3]int{1, 2, 3}   // 声明了一个长度为3的int数组
	b := [10]int{1, 2, 3}  // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0
	c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度

	fmt.Printf("a: %v,\n b: %v,\n c: %v\n", a, b, c)

	// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

	// 上面的声明可以简化，直接忽略内部的类型
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Printf("doubleArray: %v, easyArray: %v\n", doubleArray, easyArray)

	// 和声明array一样，只是少了长度
	var fslice []int
	fmt.Printf("fslice: %T\n", fslice)

	slice := []byte{'a', 'b', 'c', 'd'}
	fmt.Printf("slice: %v\n", slice)
	// slice可以从一个数组或一个已经存在的slice中再次声明。slice通过array[i:j]来获取，
	// 其中i是数组的开始位置，j是结束位置，但不包含array[j]，它的长度是j-i。

	// 声明一个含有10个元素元素类型为byte的数组
	var ar = [10]byte{'a', 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l'}
	// 声明两个含有byte的slice
	var as, bs []byte

	// as 指向数组的第三个元素开始，并到第五个元素结束，
	as = ar[2:5]
	//现在as含有的元素: ar[2]、ar[3]和ar[4]
	fmt.Printf("as: %v\n", as)

	// bs是数组ar的另一个slice
	bs = ar[3:5]
	// bs 的元素时 ar[3], ar[4]
	fmt.Printf("bs: %v\n", bs)

	// 注意slice和数组在声明时的区别：声明数组时，
	// 方括号内写明了数组的长度或使用...自动计算长度，
	// 而声明slice时，方括号内没有任何字符。

	/*
		slice有一些简便的操作
		slice的默认开始位置是0，ar[:n]等价于ar[0:n]
		slice的第二个序列默认是数组的长度，ar[n:]等价于ar[n:len(ar)]
		如果从一个数组里面直接获取slice，可以这样ar[:]，因为默认第一个序列是0，第二个是数组的长度，即等价于ar[0:len(ar)]
	*/

	// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	var numbers map[string]int
	numbers = make(map[string]int)
	numbers["one"] = 1  //赋值
	numbers["ten"] = 10 //赋值
	numbers["three"] = 3
	fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据
	// 打印出来如:第三个数字是: 3

	/*
		map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
		map的长度是不固定的，也就是和slice一样，也是一种引用类型
		内置的len函数同样适用于map，返回map拥有的key的数量
		map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11
		map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制
	*/

	// 初始化一个字典
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("C# have no rating associated with C# in the map")
	}
	delete(rating, "C") // 删除key为C的元素

	// map也是一种引用类型，如果两个map同时指向一个底层，那么一个改变，另一个也相应的改变：
	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	fmt.Printf("m: %v\n", m)
	m1 := m
	m1["Hello"] = "Salut"
	fmt.Printf("after, m: %v\n", m)

	// make 用于内建类型（map、slice 和 channel）的内存分配。new 用于各种类型的内存分配。

	/*
		内建函数new本质上说跟其它语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。有一点非常重要：
		new返回指针。
	*/

	/*
			内建函数make(T, args)与new(T)有着不同的功能，make只能创建slice、map和channel，并且返回一个有初始值(非零)的T类型，而不是*T。
			本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。例如，一个slice，是一个包含指向数据（内部array）
			的指针、长度和容量的三项描述符；在这些项目被初始化之前，slice为nil。对于slice、map和channel来说，make初始化了内部的数据结构，填充适当的值。

		make返回初始化后的（非零）值。
	*/
}
