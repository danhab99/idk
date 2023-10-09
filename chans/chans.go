package chans

import (
	"sync"
)

// Accumulates the content of a channel into an array
func Accumulate[T any](c <-chan T) chan []T {
	out := make(chan []T)
	go func() {
		defer close(out)
		var a []T
		for e := range c {
			a = append(a, e)
		}
		out <- a
	}()
	return out
}

type SortItem[T any] struct {
	Value T
	Index int
}

// Buffers items from a channel into memory and ensures that the items are sent in order. If limit > 0 and too many items were buffered than the next item will leak into the overflow channel.
func Sort[T any](c <-chan SortItem[T], limit int) (out chan T, overflow chan SortItem[T]) {
	out = make(chan T)
	overflow = make(chan SortItem[T])

	go func() {
		defer close(out)
		defer close(overflow)

		buff := make(map[int]T)
		count := 0

		for x := range c {
			if limit > 0 && len(buff) > limit {
				overflow <- x
				continue
			}

			buff[x.Index] = x.Value

			for buf, ok := buff[count]; ok; {
				out <- buf
				count++
			}
		}
	}()
	return
}

// Merges multiple channels into one channel
func Merge[T any](chans ...<-chan T) chan T {
	var wg sync.WaitGroup
	wg.Add(len(chans))

	out := make(chan T)

	for _, c := range chans {
		go func(c <-chan T) {
			defer wg.Done()

			for x := range c {
				out <- x
			}
		}(c)
	}

	go func() {
		wg.Done()
		close(out)
	}()

	return out
}

// Splits the output of one channel into 2 different channels
func Split[T any](c <-chan T) (left, right chan T) {
	left = make(chan T)
	right = make(chan T)

	go func() {
		defer close(left)
		defer close(right)

		dir := true
		for x := range c {
			if dir {
				left <- x
			} else {
				right <- x
			}
			dir = !dir
		}
	}()

	return
}

// Given a channel (buffered, but not required) that you expect to send items in bursts, you get a channel that constantly outputs content until the input channel closes or until it goes quiet.
func Burst[T any](c <-chan T) (out chan T) {
	out = make(chan T)

	go func() {
		defer close(out)

		for {
			select {
			case x, ok := <-c:
				if ok {
					out <- x
				} else {
					return
				}
			default:
				return
			}
		}
	}()

	return
}
