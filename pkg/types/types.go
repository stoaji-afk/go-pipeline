package types

import "time"

// --- Сообщения шагов пайплайна ---

const (
	MsgLogUnknownType        = "log: неожиданный тип %T"
	MsgLogOrder              = "📋 Заказ #%d: модель %s, создан %s"
	MsgAssembleUnknownType   = "assemble: неожиданный тип %T"
	MsgAssembled             = "🔧 Собрано: устройство %s, модель %s"
	MsgFirmwareUnknownType   = "firmware: неожиданный тип %T"
	MsgFirmwareUploaded      = "💾 Прошито: устройство %s, прошивка %s"
	MsgTestingUnknownType    = "testing: неожиданный тип %T"
	MsgQCPassed              = "🧪 QC пройден: устройство %s"
	MsgPackageUnknownType    = "package: неожиданный тип %T"
	MsgPacked                = "📦 Упаковано: устройство %s, отчёт #%s"
	MsgPipelineFinal         = "✅ Готов: устройство %s, модель %s, прошивка %s, QC: %v, общее время: %v"
)

// StepFunc — функция шага пайплайна: принимает вход, возвращает выход.
type StepFunc func(interface{}) interface{}

// DeviceOrder — входящий заказ на устройство.
type DeviceOrder struct {
	ID        int
	CreatedAt time.Time
	Model     string
}

// Firmware — информация о прошивке.
type Firmware struct {
	Version     string
	PublishDate time.Time
}

// Device — собранное устройство с этапами обработки.
type Device struct {
	ID             string    // UUID v7
	Model          string
	AssembledAt    time.Time
	FW             Firmware
	FWUploadedAt   time.Time
	QCPassed       bool
	QCAt           time.Time
	orderCreatedAt time.Time // unexported: для расчёта TotalTime в Report
}

// CreatedAt возвращает время создания заказа (для расчёта TotalTime).
func (d Device) CreatedAt() time.Time {
	return d.orderCreatedAt
}

// NewDevice создаёт устройство с временем создания заказа.
func NewDevice(id, model string, assembledAt, orderCreatedAt time.Time) Device {
	return Device{
		ID:             id,
		Model:          model,
		AssembledAt:    assembledAt,
		orderCreatedAt: orderCreatedAt,
	}
}

// Report — итоговый отчёт по устройству.
type Report struct {
	DeviceID      string
	DeviceModel   string
	AssembledAt   time.Time
	AssembledBy   string
	FWVersion     string
	FWPublishDate time.Time
	FWUploadedAt  time.Time
	QCPassed      bool
	QCAt          time.Time
	TotalTime     time.Duration
}

// Package — результат пайплайна: устройство + отчёт.
type Package struct {
	Device Device
	Report Report
}
