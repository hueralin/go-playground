package main

import "fmt"

// 数组

func arrayTest01() {
	// 声明
	var arr [5]int
	fmt.Println(arr) // 0 0 0 0 0
	// 赋值
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	arr[3] = 3
	arr[4] = 4
	fmt.Println(arr) // 0 1 2 3 4
}

func arrayTest02() {
	// 声明并初始化
	var arr = [5]int{10, 15, 20, 25, 30}
	fmt.Println(arr)

	// 短声明
	arr2 := [5]int{10, 15, 20, 25, 30}
	fmt.Println(arr2)

	// 部分初始化, 未初始化的为 0 值
	arr3 := [5]int{10, 15}
	fmt.Println(arr3) // 10 15 0 0 0

	// 指定索引赋值, 未初始化的为 0 值
	arr4 := [5]int{1: 10, 2: 20}
	fmt.Println(arr4) // 0 10 20 0 0

	// 通过 ... 让 Go 计算数组长度
	arr5 := [...]int{10, 20, 30}
	fmt.Println(arr5) // 10 20 30
}

func arrayTest03() {
	// 注意：数组的长度是数组类型的一部分，所以 [3]int 和 [5]int 不是同一种类型
	arr1 := [3]int{1, 2, 3}
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr1's type is %T\n", arr1) // [3]int
	fmt.Printf("arr2's type is %T\n", arr2) // [5]int
}

func arrayTest04() {
	// 多维数组
	arr := [3][2]string{
		{"1", "Vue"},
		{"2", "React"},
		{"3", "Angular"},
	}
	fmt.Println(arr)
}

func arrayTest05() {
	arr := [...]int{1, 2, 3} // 自动计算数组长度
	fmt.Println(len(arr))    // 3
}

func arrayTest06() {
	arr := [3]int{1, 2, 3}
	for index, value := range arr {
		fmt.Println(index, value)
	}

	for _, value := range arr {
		fmt.Println(value)
	}
}

func arrayTest07() {
	// Go 中的数组是值类型，而不是引用类型，这一点和 c 语言不一样
	arr := [3]int{1, 2, 3}
	// 数组的赋值会产生一个新的拷贝（深拷贝），修改新数组不会影响到原有数组
	arrCopy := arr
	arrCopy[0] = 0
	arrCopy2 := &arr
	arrCopy2[0] = 4
	fmt.Println(arr)      // 4, 2, 3
	fmt.Println(arrCopy)  // 0, 2, 3
	fmt.Println(arrCopy2) // 4, 2, 3
}

// 切片
// 切片是对数组的一个连续片段的引用，所以切片是一个引用类型
// 切片本身不拥有任何数据，它们只是对现有数组的引用，每个切片值都会将数组作为其底层的数据结构
// slice 的语法和数组很像，只是没有固定长度而已

func sliceTest01() {
	// 方法一：声明整型切片，注意和数组的声明不同，数组的声明带长度，而切片不需要
	var list []int
	fmt.Printf("list's type is %T\n", list) // []int
	fmt.Println(list)
	// 方法二：声明并初始化一个空切片（GoLand 更推荐方法一）
	var listEmpty = []int{}
	fmt.Println(listEmpty)
	// 方法三：make 声明 make([]Type, size, cap)
	var listMake = make([]int, 3, 5)
	// 指针: 指向第一个 “切片元素” 对应的底层数组元素的地址（注意：第一个切片元素不一定是第一个数组元素）
	// 长度: 切片中的元素个数
	// 容量: 从切片的开始位置到底层数组的结束位置
	fmt.Println(listMake)

	arr := [5]int{1, 2, 3, 4, 5}
	var s1 = arr[1:4]                                    // 数组变量[开始位置:结束位置]，左闭右开
	fmt.Printf("s1's type is %T, value is %v\n", s1, s1) // s1's type is []int, value is [2 3 4]
	// 注意：第一个切片元素不一定是第一个数组元素
	fmt.Printf("arr[0] is %d, s1[0] is %d\n", arr[0], s1[0]) // 1 2
}

func sliceTest02() {
	s := make([]int, 3, 5)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

func sliceTest03() {
	s := make([]int, 3, 5)
	// panic: runtime error: index out of range [10] with length 3
	fmt.Println(s[10])
}

func sliceTest04() {
	var s []int
	// 切片未初始化时，其值为 nil，引用类型的变量均如此
	fmt.Println(s == nil) // true
	// 判断切片是否为空
	fmt.Println(len(s) == 0) // true
}

func sliceTest05() {
	// 切片本身并不包含任何元素，对切片的操作都会反映在底层数组中
	arr := [5]int{1, 2, 3, 4, 5}
	// 不写开始和结束位置，表明获取数组的全部元素
	s := arr[:]
	fmt.Println(arr) // 1 2 3 4 5
	fmt.Println(s)   // 1 2 3 4 5

	s[0] = 10086
	fmt.Println(arr) // 10086 2 3 4 5
	fmt.Println(s)   // 10086 2 3 4 5
}

func sliceTest06() {
	// 声明一个切片
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)      // 1 2 3 4 5
	fmt.Println(len(s)) // 5
	fmt.Println(cap(s)) // 5
	// 追加元素
	s = append(s, 6, 7)
	fmt.Println(s)      // 1 2 3 4 5 6 7
	fmt.Println(len(s)) // 7
	fmt.Println(cap(s)) // 10
	// 追加一个切片，使用 ... 解包一个切片
	s = append(s, []int{8, 9, 10, 11}...)
	fmt.Println(s)      // 1 2 3 4 5 6 7 8 9 10 11
	fmt.Println(len(s)) // 11
	fmt.Println(cap(s)) // 20

	// 往切片追加元素时，如果容量不够，
	// Go 会新创建一个底层数组，现有数组中的元素会被拷贝到新数组，并返回新数组的引用
	// 扩容一般是扩两倍
}

func sliceTest07() {
	// 多维切片
	s := [][]string{
		{"1", "Vue"},
		{"2", "React"},
		{"3", "Angular"},
	}
	fmt.Println(s)
}

// 映射(map)
// 在Go 语言中，map 是散列表(哈希表)的引用
// 它是一个拥有键值对元素的无序集合，在这个集合中，键是唯一的，可以通过键来获取、更新或移除操作
// 无论这个散列表有多大，这些操作基本上是通过常量时间完成的
// 所有可比较的类型，如 整型字符串 等，都可以作为 key

func mapTest01() {
	// make(map[keyType]valueType)
	person := make(map[string]string)
	fmt.Println(person) // map[]

	scores := make(map[string]int)
	fmt.Println(scores) // map[]

	// 通过字面量创建 map（也可以短声明）
	var framework = map[string]string{
		"1": "Vue",
		"2": "React",
		"3": "Angular",
	}
	fmt.Println(framework)

	// 新建 or 修改 kv
	person["name"] = "spider"
	person["home"] = "beijing"
	fmt.Println(person) // map[home:beijing name:spider]

	// 删除 kv
	delete(person, "home")
	fmt.Println(person) // map[name:spider]

	// 判断 kv 是否存在: value, ok := map[key]
	name, isNameExisted := person["name"]
	fmt.Println(name)          // "spider"
	fmt.Println(isNameExisted) // true

	name2, isNameExisted2 := person["hello"]
	fmt.Println(name2)          // "空字符串"
	fmt.Println(isNameExisted2) // false

	// 遍历 map
	person["home"] = "beijing"
	for key, val := range person {
		fmt.Println(key, val)
	}

	// 获取 map 的长度
	fmt.Println(len(person)) // 2

	// map 是引用类型
	newPerson := person
	fmt.Println(person)
	fmt.Println(newPerson)
	newPerson["name"] = "baga"
	newPerson["home"] = "yalu"
	fmt.Println(person)
	fmt.Println(newPerson)
}

func main() {
	//arrayTest01()
	//arrayTest02()
	//arrayTest03()
	//arrayTest04()
	//arrayTest05()
	//arrayTest06()
	//arrayTest07()

	//sliceTest01()
	//sliceTest02()
	//sliceTest03()
	//sliceTest04()
	//sliceTest05()
	//sliceTest06()
	//sliceTest07()

	//mapTest01()
	var m = map[string]string{"name": "tom", "age": "18"}

	for k := range m {
		// age
		// name
		fmt.Printf("key is %v\n", k)
	}

	for _, v := range m {
		// 18
		// tom
		fmt.Printf("value is %v\n", v)
	}

	for k, v := range m {
		// age 18
		// name tom
		fmt.Printf("key is %v, value is %v\n", k, v)
	}
}
