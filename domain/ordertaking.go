package domain

// ------------------
// Simple types (Value Objects)
// ------------------

// Product codes
// data ProductCode = WidgetCode OR GizmoCode
// data WidgetCode = string starting with "W" then 4 digits
// data GizmoCode = string starting with "G" then 3 digits

type (
	Widget string // constraing: string starting with "W" then 4 digits
	Gizmo  string // constraing: string starting with "G" then 3 digits

	ProductCode interface {
		Widget | Gizmo
	}
)

// Order Quantity
// data OrderQuantity = UnitQuantity OR KilogramQuantity
// data UnitQuantity = int
// data KilogramQuantity = decimal

type (
	Unit  int
	Kilos float64

	OrderQuantity interface {
		Unit | Kilos
	}
)

// ------------------
// Maintained types (Entities)
// ------------------

type (
	OrderID     int
	OrderLineID int
	CustomerID  int
)

type (
	CustomerInfo any // undefined yet

	BillingAddress any // undefined yet
	Price          any // undefined yet
	BillingAmount  any // undefined yet
)

type (
	ShippingAddress interface {
		UnvalidatedAddress | ValidatedAddress
	} // undefined yet

	UnvalidatedAddress string
	ValidatedAddress   string
)

//	type OrderLine[T1 ProductCode, T2 OrderQuantity] struct {
//		ID            OrderLineID
//		OrderID       OrderID
//		ProductCode   T1
//		OrderQuantity T2
//		Priced        Price
//	}
type OrderLine struct {
	ID            OrderLineID
	OrderID       OrderID
	ProductCode   Widget
	OrderQuantity Unit
	Priced        Price
}

// Orderの状態遷移を型で表現
// Unvalidated Order -> Validated Order -> Priced Order
type Order interface {
	UnvalidatedOrder | ValidatedOrder | PricedOrder
}

type (
	UnvalidatedOrderLine OrderLine
	ValidatedOrderLine   OrderLine // undefined yet
	PricedOrderLine      any       // undefined yet
	UnvalidatedOrder     struct {
		OrderID         OrderID
		CustomerInfo    CustomerInfo
		ShippingAddress UnvalidatedAddress
		BillingAddress  UnvalidatedAddress
		OrderLines      []UnvalidatedOrderLine
	}

	ValidatedOrder struct {
		OrderID         OrderID
		CustomerInfo    CustomerInfo
		ShippingAddress ValidatedAddress
		BillingAddress  ValidatedAddress
		OrderLines      []ValidatedOrderLine
	}

	PricedOrder struct {
		OrderID         OrderID
		CustomerInfo    CustomerInfo
		ShippingAddress ValidatedAddress
		BillingAddress  ValidatedAddress
		OrderLines      []PricedOrderLine
		AmountToBill    BillingAmount
	}
)
