package service

import "github.com/trewanek/go-with-ddd-example/domain/entity"

type TransportService struct{}

func NewTransportService() *TransportService {
	return &TransportService{}
}

func (service *TransportService) Transport(
	from *entity.PhysicalDistributionBase,
	to *entity.PhysicalDistributionBase,
	baggage *entity.Baggage,
) {
	shippedBaggage := from.Ship(baggage)
	to.Receive(shippedBaggage)

	// 配送記録を行う
}
