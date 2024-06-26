package domain

import (
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/stretchr/testify/mock"
)

type AttendantRepositoryMock struct {
	mock.Mock
}

func (r *AttendantRepositoryMock) CreateAttendant(nome string) (*entities.Attendant, error) {
	args := r.Called(nome)
	return args.Get(0).(*entities.Attendant), args.Error(1)
}
func (r *AttendantRepositoryMock) GetAttendantById(id uint) (*entities.Attendant, error) {
	args := r.Called(id)
	return args.Get(0).(*entities.Attendant), args.Error(1)
}
func (r *AttendantRepositoryMock) GetAttendantByName(name string) (*entities.Attendant, error) {
	args := r.Called(name)
	return args.Get(0).(*entities.Attendant), args.Error(1)
}
func (r *AttendantRepositoryMock) GetAttendants() ([]*entities.Attendant, error) {
	args := r.Called()
	return args.Get(0).([]*entities.Attendant), args.Error(1)
}
func (r *AttendantRepositoryMock) UpdateAttendant(attendant *entities.Attendant) (*entities.Attendant, error) {
	args := r.Called(attendant)
	return args.Get(0).(*entities.Attendant), args.Error(1)
}
func (r *AttendantRepositoryMock) DeleteAttendant(id uint) error {
	args := r.Called(id)
	return args.Error(0)
}

type CustomerRepositoryMock struct {
	mock.Mock
}

func (r *CustomerRepositoryMock) CreateCustomer(customer *entities.Customer) (*entities.Customer, error) {
	args := r.Called(customer)
	return args.Get(0).(*entities.Customer), args.Error(1)
}

func (r *CustomerRepositoryMock) GetCustomerByCPF(cpf string) (*entities.Customer, error) {
	args := r.Called(cpf)
	return args.Get(0).(*entities.Customer), args.Error(1)
}

func (r *CustomerRepositoryMock) GetCustomerById(id uint) (*entities.Customer, error) {
	args := r.Called(id)
	return args.Get(0).(*entities.Customer), args.Error(1)
}

func (r *CustomerRepositoryMock) GetCustomers() ([]*entities.Customer, error) {
	args := r.Called()
	return args.Get(0).([]*entities.Customer), args.Error(1)
}

func (r *CustomerRepositoryMock) UpdateCustomer(customer *entities.Customer) (*entities.Customer, error) {
	args := r.Called(customer)
	return args.Get(0).(*entities.Customer), args.Error(1)
}

func (r *CustomerRepositoryMock) DeleteCustomer(id uint) error {
	args := r.Called(id)
	return args.Error(0)
}

type DeliveryRepositoryMock struct {
	mock.Mock
}

func (r *DeliveryRepositoryMock) GetDeliveryById(id uint) (*entities.Delivery, error) {
	args := r.Called(id)
	return args.Get(0).(*entities.Delivery), args.Error(1)
}

func (r *DeliveryRepositoryMock) CreateDelivery(delivery *entities.Delivery) (*entities.Delivery, error) {
	args := r.Called(delivery)
	return args.Get(0).(*entities.Delivery), args.Error(1)
}

func (r *DeliveryRepositoryMock) UpdateDelivery(delivery *entities.Delivery) (*entities.Delivery, error) {
	args := r.Called(delivery)
	return args.Get(0).(*entities.Delivery), args.Error(1)
}

type KitchenClientMock struct {
	mock.Mock
}

func (c *KitchenClientMock) Preparation(order *entities.Order) error {
	args := c.Called(order)
	return args.Error(0)
}

func (c *KitchenClientMock) ReadyDelivery(orderID uint) error {
	args := c.Called(orderID)
	return args.Error(0)
}

type OrderRepositoryMock struct {
	mock.Mock
}

func (r *OrderRepositoryMock) CreateOrder(entity *entities.Order) (*entities.Order, error) {
	args := r.Called(entity)
	return args.Get(0).(*entities.Order), args.Error(1)
}

func (r *OrderRepositoryMock) GetOrderById(id uint) (*entities.Order, error) {
	args := r.Called(id)
	return args.Get(0).(*entities.Order), args.Error(1)
}

func (r *OrderRepositoryMock) GetOrders() ([]*entities.Order, error) {
	args := r.Called()
	return args.Get(0).([]*entities.Order), args.Error(1)
}

func (r *OrderRepositoryMock) UpdateOrder(order *entities.Order) (*entities.Order, error) {
	args := r.Called(order)
	return args.Get(0).(*entities.Order), args.Error(1)
}

func (r *OrderRepositoryMock) RemoveItemFromOrder(idItem uint) error {
	args := r.Called(idItem)
	return args.Error(0)
}

func (r *OrderRepositoryMock) GetOrderItemById(id uint) (*entities.OrderItem, error) {
	args := r.Called(id)
	return args.Get(0).(*entities.OrderItem), args.Error(1)
}

type PaymentRepositoryMock struct {
	mock.Mock
}

func (r *PaymentRepositoryMock) GetPaymentById(id uint) (*entities.Payment, error) {
	args := r.Called(id)
	return args.Get(0).(*entities.Payment), args.Error(1)
}

func (r *PaymentRepositoryMock) UpdatePayment(payment *entities.Payment) (*entities.Payment, error) {
	args := r.Called(payment)
	return args.Get(0).(*entities.Payment), args.Error(1)
}

type PaymentClientMock struct {
	mock.Mock
}

func (c *PaymentClientMock) Pay(order *entities.Order, paymentMethod string) error {
	args := c.Called(order, paymentMethod)
	return args.Error(0)
}

func (c *PaymentClientMock) Reverse(order *entities.Order) error {
	args := c.Called(order)
	return args.Error(0)
}

type ProductRepositoryMock struct {
	mock.Mock
}

func (r *ProductRepositoryMock) CreateProduct(name string, price float64, categoryID uint) (*entities.Product, error) {
	args := r.Called(name, price, categoryID)
	return args.Get(0).(*entities.Product), args.Error(1)
}

func (r *ProductRepositoryMock) GetProductById(id uint) (*entities.Product, error) {
	args := r.Called(id)
	return args.Get(0).(*entities.Product), args.Error(1)
}

func (r *ProductRepositoryMock) GetProductByName(name string) (*entities.Product, error) {
	args := r.Called(name)
	return args.Get(0).(*entities.Product), args.Error(1)
}

func (r *ProductRepositoryMock) GetProducts() ([]*entities.Product, error) {
	args := r.Called()
	return args.Get(0).([]*entities.Product), args.Error(1)
}

func (r *ProductRepositoryMock) UpdateProduct(product *entities.Product) (*entities.Product, error) {
	args := r.Called(product)
	return args.Get(0).(*entities.Product), args.Error(1)
}

func (r *ProductRepositoryMock) DeleteProduct(id uint) error {
	args := r.Called(id)
	return args.Error(0)
}