package chans

import (
	"container/list"
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
		wg.Wait()
		close(out)
	}()

	return out
}

// Splits the output of one channel into multiple channels, guarentees
func Split[T any](c <-chan T, outs ...chan<- T) {
	go func() {
		defer func() {
			for _, e := range outs {
				close(e)
			}
		}()

		cur := 0
		for i := range c {
			outs[cur] <- i
			cur++
			if cur > len(outs)-1 {
				cur = 0
			}
		}
	}()
}

// Broadcasts a message to a list of channe. If aggresive is set to true all output channels are guarenteed to receive every item sent, otherwise it will do its best
func Broadcast[T any](c <-chan T, aggresive bool, outs ...chan<- T) {
	go func() {
		defer func() {
			for _, e := range outs {
				close(e)
			}
		}()

		for x := range c {
			for _, o := range outs {
				if aggresive {
					o <- x
				} else {
					select {
					case o <- x:
					default:
					}
				}
			}
		}
	}()
}

// Runs a function in a goroutine and returns the result type and error as channels, saves the boiler plate
func Fork[T any](handler func() (T, error)) (out chan T, e chan error) {
	out = make(chan T)
	e = make(chan error)

	go func() {
		defer close(e)
		defer close(out)
		x, err := handler()
		e <- err
		out <- x
	}()

	return
}

func GiantChan[T any](in <-chan T, out chan<- T, limit uint64) (count *uint64) {
	l := list.New()
	give := sync.NewCond(&sync.Mutex{})
	take := sync.NewCond(&sync.Mutex{})
	open := true
	count = new(uint64)
	*count = 0

	give.L.Lock()

	go func() {
		defer func() { open = false }()

		for i := range in {
			*count++
			if *count > limit {
				take.Wait()
			}
			l.PushBack(i)
			give.Broadcast()
		}
	}()

	go func() {
		give.Wait()
		e := l.Front()

		for {
			if !open {
				break
			}
			if *count == 0 {
				give.Wait()
			}

			out <- e.Value.(T)

			e = e.Next()
			if e == nil {
				continue
			}

			take.Broadcast()
			*count--
		}
	}()

	return
}
