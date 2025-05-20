package updates

type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

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

type RefundedPayment struct {
	Currency                string `json:"currency"`
	TotalAmount             int    `json:"total_amount"`
	InvoicePayload          string `json:"invoice_payload"`
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeId string `json:"provider_payment_charge_id,omitempty"`
}

type OrderInfo struct {
	Name            string           `json:"name,omitempty"`
	PhoneNumber     string           `json:"phone_number,omitempty"`
	Email           string           `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line_1"`
	StreetLine2 string `json:"street_line_2"`
	PostCode    string `json:"post_code"`
}
