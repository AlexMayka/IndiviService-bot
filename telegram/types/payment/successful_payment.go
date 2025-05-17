package payment

type SuccessfulPayment struct {
	Currency                   string    `json:"currency"`
	TotalAmount                int       `json:"total_amount"`
	InvoicePayload             string    `json:"invoice_payload"`
	SubscriptionExpirationDate int       `json:"subscription_expiration_date,omitempty"`
	IsRecurring                bool      `json:"is_recurring,omitempty"`
	IsFirstRecurring           bool      `json:"is_first_recurring,omitempty"`
	ShippingOptionId           int       `json:"shipping_option_id,omitempty"`
	OrderInfo                  OrderInfo `json:"order_info,omitempty"`
}
