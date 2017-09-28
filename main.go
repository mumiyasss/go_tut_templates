package main

import "fmt"
import "math/rand"


func main() {
	my_map();

	fmt.Println(add(1, 2, 3, 4, 5))
	
}

func my_map() {
	// полная инициализация
	var nm1 map[string]string = map[string]string{};
	nm1["test"] = "ok"
	fmt.Println(nm1)

	// сокращенная инициализация
	nm2 := map[string]string{};
	nm2["test1"] = "ok1"
	fmt.Println(nm2)

	// через make
	var nm3 = make(map[string]string)
	nm3["firstName"] = "Vasily"
	fmt.Println(nm3)

}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func rand_lesson() {
	var i int;
	var j int;
	var a [3]int;
	// 0 - Дискретная
	// 1 - Математический анализ
	// 2 - Алгбра и геометрия
	i = 0;
	for i<=6109{
		j = rand.Intn(3)
		i++;
	}
	i=0;
	for i<=10 {
		j = rand.Intn(3)
		if(a[j] == 0) {
			fmt.Println(j)
		}
		a[j]++;
		i++;
	}
}