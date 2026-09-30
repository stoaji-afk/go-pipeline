package steps

import (
	"fmt"
	"log"
	"time"

	"device-pipeline/pkg/types"
)

func FirmwareStep(in interface{}) interface{} {
	device, ok := in.(types.Device)
	if !ok {
		log.Printf(types.MsgFirmwareUnknownType, in)
		return nil
	}

	time.Sleep(200 * time.Millisecond)

	device.FW = types.Firmware{
		Version:     fmt.Sprintf("v1.%d.%d", time.Now().Month(), device.AssembledAt.Day()),
		PublishDate: time.Now().Add(-48 * time.Hour),
	}
	device.FWUploadedAt = time.Now()

	log.Printf("💾 Прошито: устройство %s, прошивка %s", device.ID, device.FW.Version)
	return device
}
