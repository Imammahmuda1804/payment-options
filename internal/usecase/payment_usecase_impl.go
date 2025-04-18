package usecase

import (
	"payment-options/internal/models"
	"payment-options/internal/repository"
	"sync"
)

type paymentUsecase struct {
	repo repository.PaymentRepository
}

func NewPaymentUsecase(r repository.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{repo: r}
}

func (u *paymentUsecase) GetPaymentOptions() (map[string]models.PaymentMethod, error) {
	var wg sync.WaitGroup
	result := make(map[string]models.PaymentMethod)
	mu := sync.Mutex{}

	wg.Add(6)

	go func() {
		defer wg.Done()
		method := u.repo.CallBCA()
		mu.Lock()
		result["bca"] = method
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		method := u.repo.CallMandiri()
		mu.Lock()
		result["mandiri"] = method
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		method := u.repo.CallBNI()
		mu.Lock()
		result["bni"] = method
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		method := u.repo.CallBRI()
		mu.Lock()
		result["bri"] = method
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		method := u.repo.CallCIMB()
		mu.Lock()
		result["cimb"] = method
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		method := u.repo.CallPermata()
		mu.Lock()
		result["permata"] = method
		mu.Unlock()
	}()

	wg.Wait()
	return result, nil
}
