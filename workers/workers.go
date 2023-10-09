package workers

import "sync"

// Creates a set of goroutines to process in a function
func Worker[InutType any, OutputType any](inputChan chan InutType, workerCount int, processTask func(in InutType) OutputType) chan OutputType {
	outputChan := make(chan OutputType)
	var wg sync.WaitGroup
	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()
			for in := range inputChan {
				outputChan <- processTask(in)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outputChan)
	}()

	return outputChan
}
