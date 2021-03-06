package payments

type Payment struct {
	Id                   string             `json:"-"`
	Amount               string             `json:"amount"`
	BeneficiaryParty     BeneficiaryParty   `json:"beneficiary_party"`
	ChargesInformation   ChargesInformation `json:"charges_information"`
	Currency             string             `json:"currency"`
	DebtorParty          DebtorParty        `json:"debtor_party"`
	EndToEndReference    string             `json:"end_to_end_reference"`
	FX                   FX                 `json:"fx"`
	NumericReference     string             `json:"numeric_reference"`
	PaymentId            string             `json:"payment_id"`
	PaymentPurporse      string             `json:"payment_purpose"`
	PaymentScheme        string             `json:"payment_scheme"`
	PaymentType          string             `json:"payment_type"`
	ProcessingDate       string             `json:"processing_date"`
	Reference            string             `json:"reference"`
	SchemePaymentSubType string             `json:"scheme_payment_sub_type"`
	SchemePaymentType    string             `json:"scheme_payment_type"`
	SponsorParty         SponsorParty       `json:"sponsor_party"`
}

type BeneficiaryParty struct {
	AccountName       string `json:"account_name"`
	AccountNumber     string `json:"account_number"`
	AccountNumberCode string `json:"account_number_code"`
	AccountType       int    `json:"account_type"`
	Address           string `json:"address"`
	BankId            string `json:"bank_id"`
	BankIdCode        string `json:"bank_id_code"`
	Name              string `json:"name"`
}

type ChargesInformation struct {
	BearerCode              string          `json:"bearer_code"`
	SenderCharges           []SenderCharges `json:"sender_charges"`
	ReceiverChargesAmount   string          `json:"receiver_charges_amount"`
	ReceiverChargesCurrency string          `json:"receiver_charges_currency"`
}

type SenderCharges struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type DebtorParty struct {
	AccountName       string `json:"account_name"`
	AccountNumber     string `json:"account_number"`
	AccountNumberCode string `json:"account_number_code"`
	Address           string `json:"address"`
	BankId            string `json:"bank_id"`
	BankIdCode        string `json:"bank_id_code"`
	Name              string `json:"name"`
}

type FX struct {
	ContractReference string `json:"contract_reference"`
	ExchangeRate      string `json:"exchange_rate"`
	OriginalAmount    string `json:"original_amount"`
	OriginalCurrency  string `json:"original_currency"`
}

type SponsorParty struct {
	AccountNumber string `json:"account_number"`
	BankId        string `json:"bank_id"`
	BankIdCode    string `json:"bank_id_code"`
}
