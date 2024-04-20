package ordertaking

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
	CustomerInfo    any // undefined yet
	ShippingAddress any // undefined yet
	BillingAddress  any // undefined yet
	Price           any // undefined yet
	BillingAmount   any // undefined yet
)

type Order[T1 ProductCode, T2 OrderQuantity] struct{
	ID OrderID
	CustomerID CustomerID
	ShippingAddress ShippingAddress
	BillingAddress BillingAddress
	OrderLines []OrderLine[T1, T2]
	AmountToBill BillingAmount
}

type OrderLine[T1 ProductCode, T2 OrderQuantity] struct{
	ID OrderLineID
	OrderID OrderID
	ProductCode T1
	OrderQuantity T2
	Priced Price
}
