// Code generated by MockGen. DO NOT EDIT.
// Source: use_case.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	payment_service "github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/modules/response/payment_service"
	entity "github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"
	enum "github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/enum"
	gomock "github.com/golang/mock/gomock"
)

// MockOrderUseCase is a mock of OrderUseCase interface.
type MockOrderUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockOrderUseCaseMockRecorder
}

// MockOrderUseCaseMockRecorder is the mock recorder for MockOrderUseCase.
type MockOrderUseCaseMockRecorder struct {
	mock *MockOrderUseCase
}

// NewMockOrderUseCase creates a new mock instance.
func NewMockOrderUseCase(ctrl *gomock.Controller) *MockOrderUseCase {
	mock := &MockOrderUseCase{ctrl: ctrl}
	mock.recorder = &MockOrderUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderUseCase) EXPECT() *MockOrderUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrderUseCase) Create(order entity.Order, products []*entity.Product) (*entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", order, products)
	ret0, _ := ret[0].(*entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOrderUseCaseMockRecorder) Create(order, products interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrderUseCase)(nil).Create), order, products)
}

// GetAll mocks base method.
func (m *MockOrderUseCase) GetAll() (*[]entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*[]entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockOrderUseCaseMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockOrderUseCase)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockOrderUseCase) GetById(id int) (*entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockOrderUseCaseMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockOrderUseCase)(nil).GetById), id)
}

// UpdateStatusById mocks base method.
func (m *MockOrderUseCase) UpdateStatusById(id int, status enum.StatusOrder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusById", id, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatusById indicates an expected call of UpdateStatusById.
func (mr *MockOrderUseCaseMockRecorder) UpdateStatusById(id, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusById", reflect.TypeOf((*MockOrderUseCase)(nil).UpdateStatusById), id, status)
}

// MockPaymentUseCase is a mock of PaymentUseCase interface.
type MockPaymentUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentUseCaseMockRecorder
}

// MockPaymentUseCaseMockRecorder is the mock recorder for MockPaymentUseCase.
type MockPaymentUseCaseMockRecorder struct {
	mock *MockPaymentUseCase
}

// NewMockPaymentUseCase creates a new mock instance.
func NewMockPaymentUseCase(ctrl *gomock.Controller) *MockPaymentUseCase {
	mock := &MockPaymentUseCase{ctrl: ctrl}
	mock.recorder = &MockPaymentUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentUseCase) EXPECT() *MockPaymentUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPaymentUseCase) Create(payment *entity.Payment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", payment)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPaymentUseCaseMockRecorder) Create(payment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPaymentUseCase)(nil).Create), payment)
}

// CreateQRCode mocks base method.
func (m *MockPaymentUseCase) CreateQRCode(order *entity.Order) (*payment_service.CreateQRCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateQRCode", order)
	ret0, _ := ret[0].(*payment_service.CreateQRCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateQRCode indicates an expected call of CreateQRCode.
func (mr *MockPaymentUseCaseMockRecorder) CreateQRCode(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQRCode", reflect.TypeOf((*MockPaymentUseCase)(nil).CreateQRCode), order)
}

// GetLastPaymentStatus mocks base method.
func (m *MockPaymentUseCase) GetLastPaymentStatus(orderId int) (enum.PaymentStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastPaymentStatus", orderId)
	ret0, _ := ret[0].(enum.PaymentStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastPaymentStatus indicates an expected call of GetLastPaymentStatus.
func (mr *MockPaymentUseCaseMockRecorder) GetLastPaymentStatus(orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastPaymentStatus", reflect.TypeOf((*MockPaymentUseCase)(nil).GetLastPaymentStatus), orderId)
}

// PaymentNotification mocks base method.
func (m *MockPaymentUseCase) PaymentNotification(order *entity.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaymentNotification", order)
	ret0, _ := ret[0].(error)
	return ret0
}

// PaymentNotification indicates an expected call of PaymentNotification.
func (mr *MockPaymentUseCaseMockRecorder) PaymentNotification(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaymentNotification", reflect.TypeOf((*MockPaymentUseCase)(nil).PaymentNotification), order)
}

// MockCategoryUseCase is a mock of CategoryUseCase interface.
type MockCategoryUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryUseCaseMockRecorder
}

// MockCategoryUseCaseMockRecorder is the mock recorder for MockCategoryUseCase.
type MockCategoryUseCaseMockRecorder struct {
	mock *MockCategoryUseCase
}

// NewMockCategoryUseCase creates a new mock instance.
func NewMockCategoryUseCase(ctrl *gomock.Controller) *MockCategoryUseCase {
	mock := &MockCategoryUseCase{ctrl: ctrl}
	mock.recorder = &MockCategoryUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryUseCase) EXPECT() *MockCategoryUseCaseMockRecorder {
	return m.recorder
}

// GetById mocks base method.
func (m *MockCategoryUseCase) GetById(id int) (*entity.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*entity.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockCategoryUseCaseMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockCategoryUseCase)(nil).GetById), id)
}

// MockClientUseCase is a mock of ClientUseCase interface.
type MockClientUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockClientUseCaseMockRecorder
}

// MockClientUseCaseMockRecorder is the mock recorder for MockClientUseCase.
type MockClientUseCaseMockRecorder struct {
	mock *MockClientUseCase
}

// NewMockClientUseCase creates a new mock instance.
func NewMockClientUseCase(ctrl *gomock.Controller) *MockClientUseCase {
	mock := &MockClientUseCase{ctrl: ctrl}
	mock.recorder = &MockClientUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientUseCase) EXPECT() *MockClientUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockClientUseCase) Create(client entity.Client) (*entity.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", client)
	ret0, _ := ret[0].(*entity.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockClientUseCaseMockRecorder) Create(client interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClientUseCase)(nil).Create), client)
}

// GetAlreadyExists mocks base method.
func (m *MockClientUseCase) GetAlreadyExists(cpf int, email string) (*entity.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlreadyExists", cpf, email)
	ret0, _ := ret[0].(*entity.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlreadyExists indicates an expected call of GetAlreadyExists.
func (mr *MockClientUseCaseMockRecorder) GetAlreadyExists(cpf, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlreadyExists", reflect.TypeOf((*MockClientUseCase)(nil).GetAlreadyExists), cpf, email)
}

// GetClientByCpf mocks base method.
func (m *MockClientUseCase) GetClientByCpf(cpf int) (*entity.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientByCpf", cpf)
	ret0, _ := ret[0].(*entity.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientByCpf indicates an expected call of GetClientByCpf.
func (mr *MockClientUseCaseMockRecorder) GetClientByCpf(cpf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientByCpf", reflect.TypeOf((*MockClientUseCase)(nil).GetClientByCpf), cpf)
}

// GetClientById mocks base method.
func (m *MockClientUseCase) GetClientById(id *int) (*entity.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientById", id)
	ret0, _ := ret[0].(*entity.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientById indicates an expected call of GetClientById.
func (mr *MockClientUseCaseMockRecorder) GetClientById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientById", reflect.TypeOf((*MockClientUseCase)(nil).GetClientById), id)
}

// MockProductUseCase is a mock of ProductUseCase interface.
type MockProductUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockProductUseCaseMockRecorder
}

// MockProductUseCaseMockRecorder is the mock recorder for MockProductUseCase.
type MockProductUseCaseMockRecorder struct {
	mock *MockProductUseCase
}

// NewMockProductUseCase creates a new mock instance.
func NewMockProductUseCase(ctrl *gomock.Controller) *MockProductUseCase {
	mock := &MockProductUseCase{ctrl: ctrl}
	mock.recorder = &MockProductUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductUseCase) EXPECT() *MockProductUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductUseCase) Create(product entity.Product) (*entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", product)
	ret0, _ := ret[0].(*entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductUseCaseMockRecorder) Create(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductUseCase)(nil).Create), product)
}

// Delete mocks base method.
func (m *MockProductUseCase) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProductUseCaseMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductUseCase)(nil).Delete), id)
}

// GetById mocks base method.
func (m *MockProductUseCase) GetById(int int) (*entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", int)
	ret0, _ := ret[0].(*entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockProductUseCaseMockRecorder) GetById(int interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockProductUseCase)(nil).GetById), int)
}

// GetProductByCategory mocks base method.
func (m *MockProductUseCase) GetProductByCategory(categoryId int) ([]entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductByCategory", categoryId)
	ret0, _ := ret[0].([]entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductByCategory indicates an expected call of GetProductByCategory.
func (mr *MockProductUseCaseMockRecorder) GetProductByCategory(categoryId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductByCategory", reflect.TypeOf((*MockProductUseCase)(nil).GetProductByCategory), categoryId)
}

// Update mocks base method.
func (m *MockProductUseCase) Update(product entity.Product, id int) (*entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", product, id)
	ret0, _ := ret[0].(*entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockProductUseCaseMockRecorder) Update(product, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductUseCase)(nil).Update), product, id)
}
