// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	entity "github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	enum "github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/enum"
	gomock "github.com/golang/mock/gomock"
)

// MockCategoryRepository is a mock of CategoryRepository interface.
type MockCategoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryRepositoryMockRecorder
}

// MockCategoryRepositoryMockRecorder is the mock recorder for MockCategoryRepository.
type MockCategoryRepositoryMockRecorder struct {
	mock *MockCategoryRepository
}

// NewMockCategoryRepository creates a new mock instance.
func NewMockCategoryRepository(ctrl *gomock.Controller) *MockCategoryRepository {
	mock := &MockCategoryRepository{ctrl: ctrl}
	mock.recorder = &MockCategoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryRepository) EXPECT() *MockCategoryRepositoryMockRecorder {
	return m.recorder
}

// GetById mocks base method.
func (m *MockCategoryRepository) GetById(id int) (*entity.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*entity.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockCategoryRepositoryMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockCategoryRepository)(nil).GetById), id)
}

// MockClientRepository is a mock of ClientRepository interface.
type MockClientRepository struct {
	ctrl     *gomock.Controller
	recorder *MockClientRepositoryMockRecorder
}

// MockClientRepositoryMockRecorder is the mock recorder for MockClientRepository.
type MockClientRepositoryMockRecorder struct {
	mock *MockClientRepository
}

// NewMockClientRepository creates a new mock instance.
func NewMockClientRepository(ctrl *gomock.Controller) *MockClientRepository {
	mock := &MockClientRepository{ctrl: ctrl}
	mock.recorder = &MockClientRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientRepository) EXPECT() *MockClientRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockClientRepository) Create(client entity.Client) (*entity.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", client)
	ret0, _ := ret[0].(*entity.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockClientRepositoryMockRecorder) Create(client interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClientRepository)(nil).Create), client)
}

// DeleteClientByCpf mocks base method.
func (m *MockClientRepository) DeleteClientByCpf(cpf int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteClientByCpf", cpf)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteClientByCpf indicates an expected call of DeleteClientByCpf.
func (mr *MockClientRepositoryMockRecorder) DeleteClientByCpf(cpf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClientByCpf", reflect.TypeOf((*MockClientRepository)(nil).DeleteClientByCpf), cpf)
}

// GetByCpf mocks base method.
func (m *MockClientRepository) GetByCpf(cpf int) (*entity.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCpf", cpf)
	ret0, _ := ret[0].(*entity.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCpf indicates an expected call of GetByCpf.
func (mr *MockClientRepositoryMockRecorder) GetByCpf(cpf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCpf", reflect.TypeOf((*MockClientRepository)(nil).GetByCpf), cpf)
}

// GetByCpfOrEmail mocks base method.
func (m *MockClientRepository) GetByCpfOrEmail(cpf int, email string) (*entity.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCpfOrEmail", cpf, email)
	ret0, _ := ret[0].(*entity.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCpfOrEmail indicates an expected call of GetByCpfOrEmail.
func (mr *MockClientRepositoryMockRecorder) GetByCpfOrEmail(cpf, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCpfOrEmail", reflect.TypeOf((*MockClientRepository)(nil).GetByCpfOrEmail), cpf, email)
}

// GetById mocks base method.
func (m *MockClientRepository) GetById(id *int) (*entity.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*entity.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockClientRepositoryMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockClientRepository)(nil).GetById), id)
}

// MockOrderRepository is a mock of OrderRepository interface.
type MockOrderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryMockRecorder
}

// MockOrderRepositoryMockRecorder is the mock recorder for MockOrderRepository.
type MockOrderRepositoryMockRecorder struct {
	mock *MockOrderRepository
}

// NewMockOrderRepository creates a new mock instance.
func NewMockOrderRepository(ctrl *gomock.Controller) *MockOrderRepository {
	mock := &MockOrderRepository{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepository) EXPECT() *MockOrderRepositoryMockRecorder {
	return m.recorder
}

// AnonymizeOrderClient mocks base method.
func (m *MockOrderRepository) AnonymizeOrderClient(clientID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnonymizeOrderClient", clientID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AnonymizeOrderClient indicates an expected call of AnonymizeOrderClient.
func (mr *MockOrderRepositoryMockRecorder) AnonymizeOrderClient(clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnonymizeOrderClient", reflect.TypeOf((*MockOrderRepository)(nil).AnonymizeOrderClient), clientID)
}

// CancelExpiredOrders mocks base method.
func (m *MockOrderRepository) CancelExpiredOrders(threshold int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelExpiredOrders", threshold)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelExpiredOrders indicates an expected call of CancelExpiredOrders.
func (mr *MockOrderRepositoryMockRecorder) CancelExpiredOrders(threshold interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelExpiredOrders", reflect.TypeOf((*MockOrderRepository)(nil).CancelExpiredOrders), threshold)
}

// Create mocks base method.
func (m *MockOrderRepository) Create(order entity.Order) (entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", order)
	ret0, _ := ret[0].(entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOrderRepositoryMockRecorder) Create(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrderRepository)(nil).Create), order)
}

// GetAll mocks base method.
func (m *MockOrderRepository) GetAll() ([]entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockOrderRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockOrderRepository)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockOrderRepository) GetById(id int) (*entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockOrderRepositoryMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockOrderRepository)(nil).GetById), id)
}

// GetByStatus mocks base method.
func (m *MockOrderRepository) GetByStatus(status enum.StatusOrder) ([]*entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByStatus", status)
	ret0, _ := ret[0].([]*entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByStatus indicates an expected call of GetByStatus.
func (mr *MockOrderRepositoryMockRecorder) GetByStatus(status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByStatus", reflect.TypeOf((*MockOrderRepository)(nil).GetByStatus), status)
}

// UpdateStatusById mocks base method.
func (m *MockOrderRepository) UpdateStatusById(id int, status enum.StatusOrder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusById", id, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatusById indicates an expected call of UpdateStatusById.
func (mr *MockOrderRepositoryMockRecorder) UpdateStatusById(id, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusById", reflect.TypeOf((*MockOrderRepository)(nil).UpdateStatusById), id, status)
}

// MockProductRepository is a mock of ProductRepository interface.
type MockProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryMockRecorder
}

// MockProductRepositoryMockRecorder is the mock recorder for MockProductRepository.
type MockProductRepositoryMockRecorder struct {
	mock *MockProductRepository
}

// NewMockProductRepository creates a new mock instance.
func NewMockProductRepository(ctrl *gomock.Controller) *MockProductRepository {
	mock := &MockProductRepository{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepository) EXPECT() *MockProductRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductRepository) Create(product entity.Product) (entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", product)
	ret0, _ := ret[0].(entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductRepositoryMockRecorder) Create(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductRepository)(nil).Create), product)
}

// Delete mocks base method.
func (m *MockProductRepository) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProductRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductRepository)(nil).Delete), id)
}

// GetById mocks base method.
func (m *MockProductRepository) GetById(id int) (*entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockProductRepositoryMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockProductRepository)(nil).GetById), id)
}

// GetProductByCategory mocks base method.
func (m *MockProductRepository) GetProductByCategory(categoryId int) ([]entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductByCategory", categoryId)
	ret0, _ := ret[0].([]entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductByCategory indicates an expected call of GetProductByCategory.
func (mr *MockProductRepositoryMockRecorder) GetProductByCategory(categoryId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductByCategory", reflect.TypeOf((*MockProductRepository)(nil).GetProductByCategory), categoryId)
}

// Update mocks base method.
func (m *MockProductRepository) Update(product entity.Product) (*entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", product)
	ret0, _ := ret[0].(*entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockProductRepositoryMockRecorder) Update(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductRepository)(nil).Update), product)
}
