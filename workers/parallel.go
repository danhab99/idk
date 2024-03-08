package workers

import "sync"

func Parallel[IN, OUT any](in <-chan IN, out chan<- OUT, count int, process func(in IN) OUT) {
	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			for e := range in {
				out <- process(e)
			}
		}()
	}

	wg.Wait()
}
