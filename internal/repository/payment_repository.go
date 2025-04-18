package repository

import "payment-options/internal/models"

type PaymentRepository interface {
	CallBCA() models.PaymentMethod
	CallMandiri() models.PaymentMethod
	CallBNI() models.PaymentMethod
	CallBRI() models.PaymentMethod
	CallCIMB() models.PaymentMethod
	CallPermata() models.PaymentMethod
}
