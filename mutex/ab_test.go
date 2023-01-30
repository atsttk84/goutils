package mutex

import (
	"fmt"
	"testing"
)

func TestSafeCounterTest(t *testing.T) {
	sc := SafeCounter{v: make(map[int]int)}
	for i := 0; i < 11; i++ {
		fmt.Println(sc.Draw(1))
	}
}

func TestAbTest(t *testing.T) {
	ab := AbTest{r: 30, t: 10, c: make(map[int]int), l: make(map[int][]int)}
	for i := 0; i < 11; i++ {
		fmt.Println(ab.Draw(1))
		fmt.Println(ab.Next(1))
	}
}

func TestInitList(t *testing.T) {
	fmt.Println(initList(20, 10))
}

func BenchmarkUnSafeCounter(b *testing.B) {
	sc := UnSafeCounter{v: make(map[int]int)}
	for i := 0; i < b.N; i++ {
		sc.Draw(1)
	}
}


func BenchmarkSafeCounter(b *testing.B) {
	sc := SafeCounter{v: make(map[int]int)}
	for i := 0; i < b.N; i++ {
		sc.Draw(1)
	}
}

func BenchmarkAbTest100(b *testing.B) {
	ab := AbTest{r: 5, t: 100, c: make(map[int]int), l: make(map[int][]int)}
	for i := 0; i < b.N; i++ {
		ab.Draw(1)
	}
}

func BenchmarkAbTest10000(b *testing.B) {
	ab := AbTest{r: 5, t: 10000, c: make(map[int]int), l: make(map[int][]int)}
	for i := 0; i < b.N; i++ {
		ab.Draw(1)
	}
}

func BenchmarkAbTest100000(b *testing.B) {
	ab := AbTest{r: 5, t: 100000, c: make(map[int]int), l: make(map[int][]int)}
	for i := 0; i < b.N; i++ {
		ab.Draw(1)
	}
}

func BenchmarkAbTest1000000(b *testing.B) {
	ab := AbTest{r: 5, t: 1000000, c: make(map[int]int), l: make(map[int][]int)}
	for i := 0; i < b.N; i++ {
		ab.Draw(1)
	}
}
