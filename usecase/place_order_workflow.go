package usecase

import (
	"fmt"

	"go-ddd-made-functional.com/domain"
)

type (
	Event              struct{}
	PlaceOrderWorkflow func(domain.UnvalidatedOrder) Event
)

// the private parts of the workflow below
type (
	// Validate Order
	checkProductCodeExists func(domain.Widget) bool
	checkAddressExists     func(domain.UnvalidatedAddress) domain.ValidatedAddress
	validateOrder          func(checkProductCodeExists, checkAddressExists, domain.UnvalidatedOrder) domain.ValidatedOrder

	// Price Order
	getProcuctPrice func(domain.Widget) domain.Price
	priceOrder      func(getProcuctPrice, domain.ValidatedOrder) domain.PricedOrder
)

func (p *PlaceOrderWorkflow) validateOrder() validateOrder {
	return func(
		checkProductCodeExists checkProductCodeExists,
		checkAddressExists checkAddressExists,
		unvalidatedOrder domain.UnvalidatedOrder,
	) domain.ValidatedOrder {
		validatedOrderID := unvalidatedOrder.OrderID           // if you want to validate, add check
		validatedCustomerInfo := unvalidatedOrder.CustomerInfo // if you want to validate, add check
		var validatedOrderLines []domain.ValidatedOrderLine
		for i, unvalidatedOrderLine := range unvalidatedOrder.OrderLines {
			if checkProductCodeExists(unvalidatedOrderLine.ProductCode) {
				validatedOrderLines = append(validatedOrderLines, domain.ValidatedOrderLine{
					ID:            domain.OrderLineID(i),
					OrderID:       unvalidatedOrder.OrderID,
					ProductCode:   unvalidatedOrderLine.ProductCode,
					OrderQuantity: unvalidatedOrderLine.OrderQuantity,
					Priced:        nil,
				})
			}
		}

		return domain.ValidatedOrder{
			OrderID:         validatedOrderID,
			CustomerInfo:    validatedCustomerInfo,
			ShippingAddress: checkAddressExists(unvalidatedOrder.ShippingAddress),
			BillingAddress:  checkAddressExists(unvalidatedOrder.BillingAddress),
			OrderLines:      validatedOrderLines,
		}
	}
}

func (p *PlaceOrderWorkflow) validateOrderV2(
	checkProductCodeExists checkProductCodeExists,
	checkAddressExists checkAddressExists,
	unvalidatedOrder domain.UnvalidatedOrder,
) domain.ValidatedOrder {
	validatedOrderID := unvalidatedOrder.OrderID           // if you want to validate, add check
	validatedCustomerInfo := unvalidatedOrder.CustomerInfo // if you want to validate, add check
	var validatedOrderLines []domain.ValidatedOrderLine
	for i, unvalidatedOrderLine := range unvalidatedOrder.OrderLines {
		if checkProductCodeExists(unvalidatedOrderLine.ProductCode) {
			validatedOrderLines = append(validatedOrderLines, domain.ValidatedOrderLine{
				ID:            domain.OrderLineID(i),
				OrderID:       unvalidatedOrder.OrderID,
				ProductCode:   unvalidatedOrderLine.ProductCode,
				OrderQuantity: unvalidatedOrderLine.OrderQuantity,
				Priced:        nil,
			})
		}
	}

	return domain.ValidatedOrder{
		OrderID:         validatedOrderID,
		CustomerInfo:    validatedCustomerInfo,
		ShippingAddress: checkAddressExists(unvalidatedOrder.ShippingAddress),
		BillingAddress:  checkAddressExists(unvalidatedOrder.BillingAddress),
		OrderLines:      validatedOrderLines,
	}
}

func (p *PlaceOrderWorkflow) priceOrder() priceOrder {
	return func(
		getProcuctPrice getProcuctPrice,
		validatedOrder domain.ValidatedOrder,
	) domain.PricedOrder {
		var pricedOrderLines []domain.PricedOrderLine

		return domain.PricedOrder{
			OrderID:         validatedOrder.OrderID,
			CustomerInfo:    validatedOrder.CustomerInfo,
			ShippingAddress: validatedOrder.ShippingAddress,
			BillingAddress:  validatedOrder.BillingAddress,
			OrderLines:      pricedOrderLines,
		}
	}
}

func (p *PlaceOrderWorkflow) PlaceOrder(
	checkProductExists checkProductCodeExists, // dependencies
	checkAddressExists checkAddressExists, // dependencies
	getProcuctPrice getProcuctPrice, // dependencies
) PlaceOrderWorkflow {
	return func(unvalidatedOrder domain.UnvalidatedOrder) Event {
		// validatedOrder := p.validateOrder()(checkProductExists, checkAddressExists, unvalidatedOrder)
		// pricedOrder := p.priceOrder()(getProcuctPrice, validatedOrder)
		// pricedOrder := p.priceOrder()(getProcuctPrice, p.validateOrder()(checkProductExists, checkAddressExists, unvalidatedOrder)) // 関数自体を型にしたい場合はこっち
		pricedOrder := p.priceOrder()(getProcuctPrice, p.validateOrderV2(checkProductExists, checkAddressExists, unvalidatedOrder)) // これでもいけるが
		fmt.Println(pricedOrder)
		return Event{}
	}
}
