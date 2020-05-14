package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1. Decalare struct Hero
type Hero struct {
	Name string
	Age  int
}

// 2. Decalare slice HeroSlice
type HeroSlice []Hero

// 3. Implement Interface by create three methods
func (hs HeroSlice) Len() int {
	return len(hs)
}
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}
func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	heroes := HeroSlice{}
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("Hero[%d]", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	fmt.Println("original:")
	for _, v := range heroes {
		fmt.Println(v.Name, v.Age)
	}
	sort.Sort(heroes)
	fmt.Println("sorted:")
	for _, v := range heroes {
		fmt.Println(v)
	}
}
