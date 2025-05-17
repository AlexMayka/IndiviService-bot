package payment

type RefundedPayment struct {
	Currency                string `json:"currency"`
	TotalAmount             int    `json:"total_amount"`
	InvoicePayload          string `json:"invoice_payload"`
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeId string `json:"provider_payment_charge_id,omitempty"`
}
