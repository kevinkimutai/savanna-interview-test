package ports

import "github.com/kevinkimutai/savanna-app/internal/app/core/domain"

type QueuePort interface {
	SendSMSQueue(order domain.Order, phoneNumber uint, customerName string)
}
