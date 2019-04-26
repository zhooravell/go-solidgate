package solidgate

type InitPaymentResponse struct {
	Order   Order   `json:"order"`
	PayForm PayForm `json:"pay_form"`
}

type ChargeResponse struct {
	Transactions   map[string]Transaction `json:"transactions"`
	Order          Order                  `json:"order"`
	Transaction    Transaction            `json:"transaction"`
	PaymentAdviser PaymentAdviser         `json:"payment_adviser"`
}

type RecurringResponse struct {
	Transactions   map[string]Transaction `json:"transactions"`
	Order          Order                  `json:"order"`
	Transaction    Transaction            `json:"transaction"`
	PaymentAdviser PaymentAdviser         `json:"payment_adviser"`
}

type RefundResponse struct {
	Order        Order                  `json:"order"`
	Transaction  Transaction            `json:"transaction"`
	Transactions map[string]Transaction `json:"transactions"`
}

type StatusResponse struct {
	Transactions   map[string]Transaction `json:"transactions"`
	Chargebacks    map[string]Chargeback  `json:"chargebacks"`
	Order          Order                  `json:"order"`
	PaymentAdviser `json:"payment_adviser"`
}

type Order struct {
	OrderID            string `json:"order_id"`
	Amount             int    `json:"amount"`
	Currency           string `json:"currency"`
	Fraudulent         bool   `json:"fraudulent"`
	MarketingAmount    int    `json:"marketing_amount,omitempty"`
	MarketingCurrency  string `json:"marketing_currency,omitempty"`
	ProcessingAmount   int    `json:"processing_amount,omitempty"`
	ProcessingCurrency string `json:"processing_currency,omitempty"`
	Status             string `json:"status"`
	RefundedAmount     int    `json:"refunded_amount,omitempty"`
	TotalFeeAmount     int    `json:"total_fee_amount"`
	FeeCurrency        string `json:"fee_currency,omitempty"`
	Descriptor         string `json:"descriptor,omitempty"`
}

type PayForm struct {
	Token      string `json:"token"`
	DesignName string `json:"design_name"`
}

type Card struct {
	Bank         string    `json:"bank"`
	Bin          string    `json:"bin"`
	Brand        string    `json:"brand"`
	Country      string    `json:"country"`
	Number       string    `json:"number"`
	CardExpMonth string    `json:"card_exp_month"`
	CardExpYear  int       `json:"card_exp_year"`
	CardType     string    `json:"card_type"`
	CardToken    CardToken `json:"card_token,omitempty"`
}

type CardToken struct {
	Token string `json:"token"`
}

type Transaction struct {
	ID        string    `json:"id"`
	Operation string    `json:"operation"`
	Status    string    `json:"status"`
	Amount    int       `json:"amount"`
	Currency  string    `json:"currency"`
	Card      Card      `json:"card"`
	CardToken CardToken `json:"card_token,omitempty"`
	Fee       Fee       `json:"fee,omitempty"`
}

type PaymentAdviser struct {
	Advise string `json:"advise"`
}

type Fee struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

type Chargeback struct {
	ID             int    `json:"id"`
	DisputeDate    string `json:"dispute_date"`
	SettlementDate string `json:"settlement_date"`
	Amount         int    `json:"amount"`
	Currency       string `json:"currency"`
	ReasonCode     string `json:"reason_code"`
	Status         string `json:"status"`
}
