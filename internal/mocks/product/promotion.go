// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/product/promotion.go
//
// Generated by this command:
//
//	mockgen -source=./internal/product/promotion.go -destination=./internal/mocks/product/promotion.go
//

// Package mock_product is a generated GoMock package.
package mock_product

import (
	product "mytheresa-promotions/internal/product"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPromotion is a mock of Promotion interface.
type MockPromotion struct {
	ctrl     *gomock.Controller
	recorder *MockPromotionMockRecorder
	isgomock struct{}
}

// MockPromotionMockRecorder is the mock recorder for MockPromotion.
type MockPromotionMockRecorder struct {
	mock *MockPromotion
}

// NewMockPromotion creates a new mock instance.
func NewMockPromotion(ctrl *gomock.Controller) *MockPromotion {
	mock := &MockPromotion{ctrl: ctrl}
	mock.recorder = &MockPromotionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPromotion) EXPECT() *MockPromotionMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockPromotion) Apply(p *product.Product) *product.Product {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", p)
	ret0, _ := ret[0].(*product.Product)
	return ret0
}

// Apply indicates an expected call of Apply.
func (mr *MockPromotionMockRecorder) Apply(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockPromotion)(nil).Apply), p)
}

// CanApply mocks base method.
func (m *MockPromotion) CanApply(p *product.Product) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanApply", p)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanApply indicates an expected call of CanApply.
func (mr *MockPromotionMockRecorder) CanApply(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanApply", reflect.TypeOf((*MockPromotion)(nil).CanApply), p)
}

// Percentage mocks base method.
func (m *MockPromotion) Percentage() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Percentage")
	ret0, _ := ret[0].(int)
	return ret0
}

// Percentage indicates an expected call of Percentage.
func (mr *MockPromotionMockRecorder) Percentage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Percentage", reflect.TypeOf((*MockPromotion)(nil).Percentage))
}
