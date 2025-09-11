package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Data []int

func (d Data) Len()int {
	return len(d)
}

func (d Data) Less(i, j int) bool {
	return d[i] < d[j]
}

func(d Data) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func main() {
	fmt.Println("use of sort interface")

	howMany := 10000000
	data := make(Data, howMany)
	for i := range howMany {
		data[i] = rand.Intn(howMany) + 1
	}
	fmt.Println(data[0:10])
	sort.Sort(data)
	fmt.Println(data[0:10])
}



