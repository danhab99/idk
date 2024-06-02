package chans

import (
	"testing"
)

// func TestGiantChan(t *testing.T) {
// 	in := make(chan uint64)
// 	out := make(chan uint64)
// 	defer close(in)
// 	defer close(out)
// 	count := GiantChan[uint64](in, out, math.MaxUint64-1)

// 	for i := uint64(0); i < math.MaxUint32+1; i++ {
// 		t.Log("Adding num", i)
// 		in <- i
// 	}

// 	if *count < math.MaxUint32 {
// 		t.Fail()
// 	}
// }

func TestAccumulate(t *testing.T) {
	c := make(chan int, 10)
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)

	arr := <-Accumulate(c)

	if len(arr) != 10 {
		t.Error("didn't collect enough items")
	}
}

func TestMerge(t *testing.T) {
	a := make(chan int, 10)
	b := make(chan int, 10)
	c := make(chan int, 10)

	for i := 0; i < 10; i++ {
		a <- i
		b <- i
		c <- i
	}

	close(a)
	close(b)
	close(c)

	o := Merge(a, b, c)
	arr := <-Accumulate(o)

	if len(arr) != 30 {
		t.Error("didn't get 30 items")
	}
}

func TestSplit(t *testing.T) {
	a := make(chan int, 1)
	b := make(chan int, 1)
	c := make(chan int, 1)

	in := make(chan int)
	Split(in, a, b, c)

	in <- 1
	in <- 2
	in <- 3
	close(in)

	x := <-a
	if x != 1 {
		t.Error("didn't get 1")
	}
	x = <-b
	if x != 2 {
		t.Error("didn't get 2")
	}
	x = <-c
	if x != 3 {
		t.Error("didn't get 3")
	}
}
