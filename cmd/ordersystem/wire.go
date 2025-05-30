//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/dchfarah/gocourse-challenge-03/internal/entity"
	"github.com/dchfarah/gocourse-challenge-03/internal/event"
	"github.com/dchfarah/gocourse-challenge-03/internal/infra/database"
	"github.com/dchfarah/gocourse-challenge-03/internal/infra/web/handlers"
	"github.com/dchfarah/gocourse-challenge-03/internal/usecase"
	"github.com/dchfarah/gocourse-challenge-03/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	setOrderCreatedEvent,
	setOrdersListedEvent,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

var setOrdersListedEvent = wire.NewSet(
	event.NewOrdersListed,
	wire.Bind(new(events.EventInterface), new(*event.OrdersListed)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewListOrdersUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.ListOrdersUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrdersListedEvent,
		usecase.NewListOrdersUseCase,
	)
	return &usecase.ListOrdersUseCase{}
}

func NewWebCreateOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *handlers.WebCreateOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		handlers.NewWebCreateOrderHandler,
	)
	return &handlers.WebCreateOrderHandler{}
}

func NewWebListOrdersHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *handlers.WebListOrdersHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrdersListedEvent,
		handlers.NewWebListOrdersHandler,
	)
	return &handlers.WebListOrdersHandler{}
}
