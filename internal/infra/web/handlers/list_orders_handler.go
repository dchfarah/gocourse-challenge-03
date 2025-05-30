package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dchfarah/gocourse-challenge-03/internal/entity"
	"github.com/dchfarah/gocourse-challenge-03/internal/usecase"
	"github.com/dchfarah/gocourse-challenge-03/pkg/events"
)

type WebListOrdersHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrdersListedEvent events.EventInterface
}

func NewWebListOrdersHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrdersListedEvent events.EventInterface,
) *WebListOrdersHandler {
	return &WebListOrdersHandler{
		EventDispatcher:   EventDispatcher,
		OrderRepository:   OrderRepository,
		OrdersListedEvent: OrdersListedEvent,
	}
}

func (h *WebListOrdersHandler) List(w http.ResponseWriter, r *http.Request) {
	listOrders := usecase.NewListOrdersUseCase(h.OrderRepository, h.OrdersListedEvent, h.EventDispatcher)
	output, err := listOrders.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
