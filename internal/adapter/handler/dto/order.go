package dto

import (
	"fmt"

	"github.com/fiap/challenge-gofood/internal/domain/entity"
)

type OrderResponse struct {
	ID            uint                 `json:"id"`
	CustomerID    uint                 `json:"customerID"`
	CustomerCPF   string               `json:"customerCPF"`
	CustomerName  string               `json:"customerName"`
	AttendantID   uint                 `json:"attendantID"`
	AttendantName string               `json:"attendantName"`
	Amount        string               `json:"amount"`
	ItemsTotal    int                  `json:"itemsTotal"`
	Status        string               `json:"status"`
	Payment       OrderPaymentResponse `json:"payment"`
	Items         []OrderItemResponse
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}

func ToOrderResponse(order *entity.Order) OrderResponse {
	response := OrderResponse{
		ID:            order.ID,
		CustomerCPF:   order.Customer.CPF,
		CustomerName:  order.Customer.Name,
		AttendantID:   order.Attendant.ID,
		AttendantName: order.Attendant.Name,
		Amount:        fmt.Sprintf("%.2f", order.Amount()),
		ItemsTotal:    order.ItemsQuantity(),
		Status:        order.Status.ToString(),
		Payment: OrderPaymentResponse{
			Status: order.Payment.Status.ToString(),
			Method: order.Payment.Method.ToString(),
		},
		Items:     []OrderItemResponse{},
		CreatedAt: order.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: order.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	for _, item := range order.Items {
		response.Items = append(response.Items, OrderItemResponse{
			ProductID:   item.Product.ID,
			ProductName: item.Product.Name,
			Quantity:    item.Quantity,
			UnitPrice:   item.UnitPrice,
		})
	}
	return response
}

func ToOrderResponses(orders []*entity.Order) []OrderResponse {
	var response []OrderResponse
	for _, entity := range orders {
		response = append(response, ToOrderResponse(entity))
	}
	return response
}

type OrderPaymentResponse struct {
	Status string `json:"status"`
	Method string `json:"method"`
}

type OrderItemResponse struct {
	ProductID   uint    `json:"productID"`
	ProductName string  `json:"productName"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unitPrice"`
}

type StartOrderRequest struct {
	CustomerCPF string `json:"customerCPF"`
	AttendantID uint   `json:"attendantID"`
}

type StartOrderResponse struct {
	ID uint `json:"id"`
}

func ToStartOrderResponse(order *entity.Order) StartOrderResponse {
	return StartOrderResponse{
		ID: order.ID,
	}
}

type AddItemToOrderRequest struct {
	ProductID uint `json:"productID"`
	Quantity  int  `json:"quantity"`
}

type PaymentOrderRequest struct {
	PaymentMethod string `json:"paymentMethod"`
}
