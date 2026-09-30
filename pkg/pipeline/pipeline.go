package pipeline

import (
	"log"
	"sync"

	"device-pipeline/pkg/types"
)

// Pipeline — конвейер шагов обработки заказов.
type Pipeline struct {
	steps    []types.StepFunc
	orderCh  chan interface{}
	resultCh chan interface{}
	wg       sync.WaitGroup
}

// NewPipeline создаёт новый пустой пайплайн.
func NewPipeline() *Pipeline {
	return &Pipeline{
		orderCh: make(chan interface{}, 1),
	}
}

// PushOrder добавляет заказ в пайплайн.
func (p *Pipeline) PushOrder(order types.DeviceOrder) {
	p.orderCh <- order
}

// AddStep добавляет шаг в конец пайплайна.
func (p *Pipeline) AddStep(step types.StepFunc) {
	p.steps = append(p.steps, step)
}

// Close закрывает входящий канал.
func (p *Pipeline) Close() {
	close(p.orderCh)
}

// Run запускает все шаги пайплайна в горутине.
func (p *Pipeline) Run() {
	n := len(p.steps)

	// N+1 каналов: ch[0] = orderCh, ch[N] = resultCh
	channels := make([]chan interface{}, n+1)
	channels[0] = p.orderCh
	for i := 1; i <= n; i++ {
		channels[i] = make(chan interface{}, 1)
	}
	p.resultCh = channels[n]

	// Запускаем воркер на каждый шаг
	for i, step := range p.steps {
		p.wg.Add(1)
		go worker(channels[i], channels[i+1], step, &p.wg)
	}

	// Коллектор готовых пакетов
	go func() {
		for item := range p.resultCh {
			if pkg, ok := item.(types.Package); ok {
				log.Printf(types.MsgPipelineFinal,
					pkg.Report.DeviceID,
					pkg.Report.DeviceModel,
					pkg.Report.FWVersion,
					pkg.Report.QCPassed,
					pkg.Report.TotalTime,
				)
			}
		}
	}()

	p.wg.Wait()
}
