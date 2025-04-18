package repository

import (
	"payment-options/internal/models"
	"time"
)

type paymentRepo struct{}

func NewPaymentRepo() PaymentRepository {
	return &paymentRepo{}
}

func (r *paymentRepo) CallBCA() models.PaymentMethod {
	time.Sleep(1 * time.Second)
	return models.PaymentMethod{
		Account: "1234567890",
		Status:  "Active",
		Balance: "100000",
		Icon:    "https://sampleurl.com/bca.jpg",
	}
}

func (r *paymentRepo) CallMandiri() models.PaymentMethod {
	time.Sleep(1 * time.Second)
	return models.PaymentMethod{
		Account: "2345678901",
		Status:  "Active",
		Balance: "200000",
		Icon:    "https://sampleurl.com/mandiri.jpg",
	}
}

func (r *paymentRepo) CallBNI() models.PaymentMethod {
	time.Sleep(1 * time.Second)
	return models.PaymentMethod{
		Account: "3456789012",
		Status:  "Active",
		Balance: "300000",
		Icon:    "https://sampleurl.com/bni.jpg",
	}
}

func (r *paymentRepo) CallBRI() models.PaymentMethod {
	time.Sleep(1 * time.Second)
	return models.PaymentMethod{
		Account: "4567890123",
		Status:  "Active",
		Balance: "400000",
		Icon:    "https://sampleurl.com/bri.jpg",
	}
}

func (r *paymentRepo) CallCIMB() models.PaymentMethod {
	time.Sleep(1 * time.Second)
	return models.PaymentMethod{
		Account: "5678901234",
		Status:  "Active",
		Balance: "500000",
		Icon:    "https://sampleurl.com/cimb.jpg",
	}
}

func (r *paymentRepo) CallPermata() models.PaymentMethod {
	time.Sleep(1 * time.Second)
	return models.PaymentMethod{
		Account: "6789012345",
		Status:  "Active",
		Balance: "600000",
		Icon:    "https://sampleurl.com/permata.jpg",
	}
}
