package steps

import (
	"log"
	"time"

	"device-pipeline/pkg/types"
)

// TestingStep — QC-тестирование устройства (100 мс).
func TestingStep(in interface{}) interface{} {
	device, ok := in.(types.Device)
	if !ok {
		log.Printf(types.MsgTestingUnknownType, in)
		return nil
	}

	time.Sleep(100 * time.Millisecond)

	// QC всегда проходит (в реальности — рандом/условия)
	device.QCPassed = true
	device.QCAt = time.Now()

	log.Printf("🧪 QC пройден: устройство %s", device.ID)
	return device
}
