package steps

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"device-pipeline/pkg/types"
)

func generateUUIDv7() string {
	b := make([]byte, 16)

	// Первые 48 бит — Unix epoch milliseconds
	ms := time.Now().UnixMilli()
	binary.BigEndian.PutUint64(b[0:8], uint64(ms))

	// Остальные байты — случайные
	_, _ = rand.Read(b[6:16])

	// Version 7
	b[6] = (b[6] & 0x0F) | 0x70
	// Variant RFC 4122
	b[8] = (b[8] & 0x3F) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

func AssembleStep(in interface{}) interface{} {
	order, ok := in.(types.DeviceOrder)
	if !ok {
		log.Printf(types.MsgAssembleUnknownType, in)
		return nil
	}

	time.Sleep(400 * time.Millisecond)

	now := time.Now()
	device := types.NewDevice(
		generateUUIDv7(),
		order.Model,
		now,
		order.CreatedAt,
	)

	log.Printf(types.MsgAssembled, device.ID, device.Model)
	return device
}
