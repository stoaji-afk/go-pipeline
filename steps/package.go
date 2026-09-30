package steps

import (
	"log"
	"time"

	"device-pipeline/pkg/types"
)

// PackageStep — финальный этап: упаковка + генерация отчёта (100 мс).
func PackageStep(in interface{}) interface{} {
	device, ok := in.(types.Device)
	if !ok {
		log.Printf(types.MsgPackageUnknownType, in)
		return nil
	}

	time.Sleep(100 * time.Millisecond)

	report := types.Report{
		DeviceID:      device.ID,
		DeviceModel:   device.Model,
		AssembledAt:   device.AssembledAt,
		AssembledBy:   "assembly-line-1",
		FWVersion:     device.FW.Version,
		FWPublishDate: device.FW.PublishDate,
		FWUploadedAt:  device.FWUploadedAt,
		QCPassed:      device.QCPassed,
		QCAt:          device.QCAt,
		TotalTime:     time.Since(device.CreatedAt()),
	}

	pkg := types.Package{
		Device: device,
		Report: report,
	}

	log.Printf("📦 Упаковано: устройство %s, отчёт #%s", device.ID, report.DeviceID)
	return pkg
}
