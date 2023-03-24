// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package usecase

import (
	"context"
	"github.com/fajarabdillahfn/shoping-gql/internal/model"
	"sync"
)

// Ensure, that UseCaseMock does implement UseCase.
// If this is not the case, regenerate this file with moq.
var _ UseCase = &UseCaseMock{}

// UseCaseMock is a mock implementation of UseCase.
//
//	func TestSomethingThatUsesUseCase(t *testing.T) {
//
//		// make and configure a mocked UseCase
//		mockedUseCase := &UseCaseMock{
//			CheckoutFunc: func(ctx context.Context, productsBought map[string]int) (*model.Cart, error) {
//				panic("mock out the Checkout method")
//			},
//		}
//
//		// use mockedUseCase in code that requires UseCase
//		// and then make assertions.
//
//	}
type UseCaseMock struct {
	// CheckoutFunc mocks the Checkout method.
	CheckoutFunc func(ctx context.Context, productsBought map[string]int) (*model.Cart, error)

	// calls tracks calls to the methods.
	calls struct {
		// Checkout holds details about calls to the Checkout method.
		Checkout []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ProductsBought is the productsBought argument value.
			ProductsBought map[string]int
		}
	}
	lockCheckout sync.RWMutex
}

// Checkout calls CheckoutFunc.
func (mock *UseCaseMock) Checkout(ctx context.Context, productsBought map[string]int) (*model.Cart, error) {
	if mock.CheckoutFunc == nil {
		panic("UseCaseMock.CheckoutFunc: method is nil but UseCase.Checkout was just called")
	}
	callInfo := struct {
		Ctx            context.Context
		ProductsBought map[string]int
	}{
		Ctx:            ctx,
		ProductsBought: productsBought,
	}
	mock.lockCheckout.Lock()
	mock.calls.Checkout = append(mock.calls.Checkout, callInfo)
	mock.lockCheckout.Unlock()
	return mock.CheckoutFunc(ctx, productsBought)
}

// CheckoutCalls gets all the calls that were made to Checkout.
// Check the length with:
//
//	len(mockedUseCase.CheckoutCalls())
func (mock *UseCaseMock) CheckoutCalls() []struct {
	Ctx            context.Context
	ProductsBought map[string]int
} {
	var calls []struct {
		Ctx            context.Context
		ProductsBought map[string]int
	}
	mock.lockCheckout.RLock()
	calls = mock.calls.Checkout
	mock.lockCheckout.RUnlock()
	return calls
}
