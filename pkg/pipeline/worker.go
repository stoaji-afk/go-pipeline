package pipeline

import (
	"sync"

	"device-pipeline/pkg/types"
)

// worker — горутинa, читающая входной канал, вызывающая step и пишущая результат в выходной канал.
func worker(in, out chan interface{}, step types.StepFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)

	for item := range in {
		result := step(item)
		if result != nil {
			out <- result
		}
	}
}
