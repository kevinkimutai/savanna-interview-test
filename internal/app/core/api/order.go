package application

import (
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

func (r *OrderRepo) CreateOrder(orderItems []domain.OrderItem, customerID string) (domain.Order, error) {

	//Start Tx

	//Save Order
	order, err := r.db.CreateOrder(orderItems[0].OrderID, customerID)
	if err != nil {
		return order, err
	}

	//Save each orderItem
	var orders []domain.OrderItem

	for _, item := range orderItems {
		orderitem, err := r.db.CreateOrderItem(item)
		if err != nil {
			return domain.Order{}, err
		}

		orders = append(orders, orderitem)
	}

	//Calculate Total Price
	totalPrice, err := r.db.GetTotalPrice(order.OrderID)
	if err != nil {
		return order, err
	}

	//Update Order
	updatedOrder, err := r.db.UpdateOrderTotalPrice(order.OrderID, totalPrice)
	if err != nil {
		return order, err
	}
	//Tx Close

	//rabbitmqMsg
	//TODO:GET CUSTOMER DETAILS/PHONE_NUMBER/NAME
	r.queue.SendSMSQueue(updatedOrder, 254722670831, "Kevin Kimutai")

	//Return Order
	return updatedOrder, nil

}
func (r *OrderRepo) GetOrderByID(orderID string) (domain.Order, error) {
	//TODO:HANDLE ERRORS
	order, err := r.db.GetOrderByID(orderID)

	return order, err
}
func (r *OrderRepo) DeleteOrder(orderID string) error {
	//TODO:HANDLE ERRORS
	err := r.db.DeleteOrder(orderID)

	return err
}

func (r *OrderRepo) GetAllOrders(orderParams domain.OrderParams) (domain.OrdersFetch, error) {
	params := utils.GetOrderAPIParams(orderParams)

	data, err := r.db.GetAllOrders(params)

	return data, err

}
