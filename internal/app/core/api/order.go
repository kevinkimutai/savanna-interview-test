package application

import (
	"github.com/kevinkimutai/savanna-app/internal/adapters/queries"
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
	"github.com/kevinkimutai/savanna-app/internal/ports"
	"github.com/kevinkimutai/savanna-app/internal/utils"
)

type OrderRepo struct {
	db    ports.OrderRepoPort
	queue ports.QueuePort
}

func NewOrderRepo(db ports.OrderRepoPort, queue ports.QueuePort) *OrderRepo {
	return &OrderRepo{db: db, queue: queue}
}

func (r *OrderRepo) CreateOrder(orderItems []domain.OrderItem, phonenumber int, customer queries.Customer) (domain.Order, error) {

	//CreateOrder
	order, err := r.db.CreateOrder(orderItems, customer.CustomerID)
	if err != nil {
		return order, err
	}

	//rabbitmqMsg
	r.queue.SendSMSQueue(order, uint(phonenumber), customer.Name)

	//Return Order
	return order, nil

}

func (r *OrderRepo) GetOrderByID(orderID string) (domain.Order, error) {

	order, err := r.db.GetOrderByID(orderID)

	return order, err
}

func (r *OrderRepo) DeleteOrder(orderID string) error {

	err := r.db.DeleteOrder(orderID)

	return err
}

func (r *OrderRepo) GetAllOrders(orderParams domain.OrderParams) (domain.OrdersFetch, error) {
	params := utils.GetOrderAPIParams(orderParams)

	data, err := r.db.GetAllOrders(params)

	return data, err

}
