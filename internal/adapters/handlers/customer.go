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

// CreateCustomer godoc
// @Summary      Create a new customer
// @Description  Creates a new customer with the provided details.
// @Tags         customers
// @Accept       json
// @Produce      json
// @Param        newCustomer body domain.Customer true "New customer details"
// @Success      201  {object} domain.CustomerResponse "Customer created successfully"
// @Failure      400  {object} domain.ErrorResponse    "Bad request, invalid input"
// @Failure      500  {object} domain.ErrorResponse    "Internal server error"
// @Router       /customer [post]
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

// GetAllCustomers godoc
// @Summary      Get all customers
// @Description  Retrieves all customers based on the provided query parameters.
// @Tags         customers
// @Accept       json
// @Produce      json
// @Param        page query int false "Page number"
// @Param        limit query int false "Number of customers per page"
// @Param        sortBy query string false "Sort by field"
// @Param        sortOrder query string false "Sort order ('asc' or 'desc')"
// @Success      200  {object} domain.CustomersResponse "Successfully retrieved customers"
// @Failure      500  {object} domain.ErrorResponse    "Internal server error"
// @Router       /customer [get]
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

// GetCustomerByID godoc
// @Summary      Get customer by ID
// @Description  Retrieves a customer based on the provided customer ID.
// @Tags         customers
// @Accept       json
// @Produce      json
// @Param        customerID path string true "Customer ID"
// @Success      200  {object} domain.CustomerResponse "Successfully retrieved customer"
// @Failure      400  {object} domain.ErrorResponse   "Bad request, missing customer ID"
// @Failure      404  {object} domain.ErrorResponse   "Customer not found"
// @Failure      500  {object} domain.ErrorResponse   "Internal server error"
// @Router       /customer/{customerID} [get]
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

// DeleteCustomer godoc
// @Summary      Delete customer by ID
// @Description  Deletes a customer based on the provided customer ID.
// @Tags         customers
// @Accept       json
// @Produce      json
// @Param        customerID path string true "Customer ID"
// @Success      204  {object} domain.CustomerResponse "Successfully deleted customer"
// @Failure      400  {object} domain.ErrorResponse   "Bad request, missing customer ID"
// @Failure      404  {object} domain.ErrorResponse   "Customer not found"
// @Failure      500  {object} domain.ErrorResponse   "Internal server error"
// @Router       /customer/{customerID} [delete]
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
