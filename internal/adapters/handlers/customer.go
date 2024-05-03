package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
	"github.com/kevinkimutai/savanna-app/internal/ports"
)

type CustomerService struct {
	api ports.CustomerApiPort
}

func NewCustomerService(api ports.CustomerApiPort) *CustomerService {
	return &CustomerService{api: api}
}

func (s *CustomerService) CreateCustomer(c *fiber.Ctx) error {
	newCustomer := domain.Customer{}

	//Bind To struct
	if err := c.BodyParser(&newCustomer); err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Check missing values
	customer, err := domain.NewCustomerDomain(newCustomer)
	if err != nil {
		return c.Status(400).JSON(
			domain.ErrorResponse{
				StatusCode: 400,
				Message:    err.Error(),
			})
	}

	//Create Customer API
	cus, err := s.api.CreateCustomer(customer)

	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	return c.Status(201).JSON(
		domain.CustomerResponse{
			StatusCode: 201,
			Message:    "Customer created successfully",
			Data:       cus,
		})
}

func (s *CustomerService) GetAllCustomers(c *fiber.Ctx) error {
	//Get Query Params
	m := c.Queries()

	//Bind To CustomerParams
	cusParams := domain.CheckCustomerParams(m)

	//Get All Customers API
	cus, err := s.api.GetAllCustomers(cusParams)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})

	}
	return c.Status(200).JSON(
		domain.CustomersResponse{
			StatusCode:    200,
			Message:       "Successfully retrieved customers",
			Page:          cus.Page,
			NumberOfPages: cus.NumberOfPages,
			Total:         cus.Total,
			Data:          cus.Data,
		})
}

func (s *CustomerService) GetCustomerByID(c *fiber.Ctx) error {
	customerID := c.Params("customerID")

	//Check if id is present
	if customerID == "" {
		return c.Status(400).JSON(
			domain.ErrorResponse{
				StatusCode: 400,
				Message:    "Missing customer id",
			})
	}

	//Get Customer API
	customer, err := s.api.GetCustomerByID(customerID)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	return c.Status(200).JSON(
		domain.CustomerResponse{
			StatusCode: 200,
			Message:    "Successfully retrieved customer",
			Data:       customer})

}

func (s *CustomerService) DeleteCustomer(c *fiber.Ctx) error {
	customerID := c.Params("customerID")

	//Check if id is present
	if customerID == "" {
		return c.Status(400).JSON(
			domain.ErrorResponse{
				StatusCode: 400,
				Message:    "Missing customer id",
			})
	}

	//Delete Customer API
	err := s.api.DeleteCustomer(customerID)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	return c.Status(204).JSON(
		domain.CustomerResponse{
			StatusCode: 204,
			Message:    "Successfully Deleted Customer",
		})

}
