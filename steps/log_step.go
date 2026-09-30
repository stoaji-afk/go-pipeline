package steps

import (
	"log"

	"device-pipeline/pkg/types"
)

func LogStep(in interface{}) interface{} {
	order, ok := in.(types.DeviceOrder)
	if !ok {
		log.Printf("log: неожиданный тип %T", in)
		return nil
	}
	log.Printf(types.MsgLogOrder,
		order.ID, order.Model, order.CreatedAt.Format("15:04:05.000"))
	return order
}
