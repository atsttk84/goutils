package mutex

import (
	"math/rand"
	"sync"
	"time"
)

type UnSafeCounter struct {
	v  map[int]int
}

func (c *UnSafeCounter) Draw(key int) (v int) {
	c.v[key]++
	v = c.v[key]
	return
}

/*
 https://go.dev/tour/concurrency/9
*/

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[int]int
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Draw(key int) (v int) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	c.v[key]++
	v = c.v[key]
	return
}

type AbTest struct {
	mu sync.Mutex
	c  map[int]int
	l  map[int][]int
	r  int
	t  int
}

func (ab *AbTest) Draw(key int) (v int) {
	ab.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer ab.mu.Unlock()
	if _, ok := ab.l[key]; !ok {
		ab.l[key] = initList(ab.r, ab.t)
	}
	v = ab.l[key][ab.c[key]]
	ab.c[key]++
	if len(ab.l[key]) <= ab.c[key] {
		ab.c[key] = 0
		ab.l[key] = initList(ab.r, ab.t)
	} else {
	}
	return
}

func (ab *AbTest) Next(key int) (v, c int, l []int) {
	ab.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer ab.mu.Unlock()
	l = ab.l[key]
	c = ab.c[key]
	v = ab.l[key][ab.c[key]]
	return
}

func initList(true_rate, total int) (list []int) {
  true_cnt := total * true_rate / 100
	false_list := make([]int, total-true_cnt)
	true_list := make([]int, true_cnt)
	for i := range true_list {
		true_list[i] = 1
	}
	list = append(false_list, true_list...)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
	return
}
