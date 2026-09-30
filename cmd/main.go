package main

import (
	"log"
	"time"

	"device-pipeline/pkg/pipeline"
	"device-pipeline/pkg/types"
	"device-pipeline/steps"
)

func main() {
	log.SetFlags(0)

	// Создаём пайплайн
	p := pipeline.NewPipeline()

	// Регистрируем шаги
	p.AddStep(steps.LogStep)
	p.AddStep(steps.AssembleStep)
	p.AddStep(steps.FirmwareStep)
	p.AddStep(steps.TestingStep)
	p.AddStep(steps.PackageStep)

	// Генератор заказов
	go generateOrders(p)

	// Запускаем пайплайн
	p.Close()
	p.Run()

	log.Println("🏁 Пайплайн завершён")
}

// generateOrders — имитация входящего потока заказов.
func generateOrders(p *pipeline.Pipeline) {
	models := []string{"Sensor-X1", "Gateway-Pro", "Controller-M2"}

	for i := 1; i <= 3; i++ {
		order := types.DeviceOrder{
			ID:        i,
			CreatedAt: time.Now(),
			Model:     models[(i-1)%len(models)],
		}
		log.Printf("📥 Генерация заказа #%d: модель %s", order.ID, order.Model)
		p.PushOrder(order)
		time.Sleep(100 * time.Millisecond)
	}
}
