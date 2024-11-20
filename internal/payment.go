package internal

import "github.com/leedrum/ikarus_travel/model"

var PaymentMethods = [...]int{
	model.PaymentMethodCash,
	model.PaymentMethodCreditCard,
	model.PaymentMethodBankTransfer,
	model.PaymentMethodOther,
}
