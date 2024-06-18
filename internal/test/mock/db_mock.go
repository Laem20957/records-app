package mock

import (
	"context"
	"reflect"

	"records-app/internal/adapters/database/schemas"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
)

// MockDB is a mock of Client interface.
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB.
type MockDBMockRecorder struct {
	mock *MockDB
}

// MockNewDB creates a new mock instance.
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// GetConnection mocks base method.
func (m *MockDB) GetConnection() (*gorm.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnection")
	ret0, _ := ret[0].(*gorm.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConnection indicates an expected call of Register.
func (mr *MockDBMockRecorder) GetConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock,
		"GetConnection", reflect.TypeOf((*MockDB)(nil).GetConnection))
}

// GetAllRecordsDB mocks base method.
func (m *MockDB) GetAllRecordsDB(arg0 context.Context) ([]schemas.Records, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRecordsDB", arg0)
	ret0, _ := ret[0].([]schemas.Records)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRecordsDB indicates an expected call of GetAllRecordsDB.
func (mr *MockDBMockRecorder) GetAllRecordsDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock,
		"GetAllRecordsDB", reflect.TypeOf((*MockDB)(nil).GetAllRecordsDB), arg0)
}

// GetByIDRecordsDB mocks base method.
func (m *MockDB) GetByIDRecordsDB(ctx context.Context, arg0 int) (schemas.Records, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDRecordsDB", ctx, arg0)
	ret0, _ := ret[0].(schemas.Records)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDRecordsDB indicates an expected call of GetByIDRecordsDB.
func (mr *MockDBMockRecorder) GetByIDRecordsDB(ctx context.Context, arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock,
		"GetByIDRecordsDB", reflect.TypeOf((*MockDB)(nil).GetByIDRecordsDB), ctx, arg0)
}

// CreateRecordsDB mocks base method.
func (m *MockDB) CreateRecordsDB(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRecordsDB", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRecordsDB indicates an expected call of CreateRecordsDB.
func (mr *MockDBMockRecorder) CreateRecordsDB(ctx context.Context) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock,
		"CreateRecordsDB", reflect.TypeOf((*MockDB)(nil).CreateRecordsDB), ctx)
}

// UpdateRecordsDB mocks base method.
func (m *MockDB) UpdateRecordsDB(ctx context.Context, arg0 int, arg1, arg2 string) (schemas.Records, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRecordsDB", ctx, arg0, arg1, arg2)
	ret0, _ := ret[0].(schemas.Records)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRecordsDB indicates an expected call of UpdateRecordsDB.
func (mr *MockDBMockRecorder) UpdateRecordsDB(ctx context.Context, arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock,
		"UpdateRecordsDB", reflect.TypeOf((*MockDB)(nil).UpdateRecordsDB), ctx, arg0, arg1, arg2)
}

// DeleteRecordsDB mocks base method.
func (m *MockDB) DeleteRecordsDB(ctx context.Context, arg0 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecordsDB", ctx, arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRecordsDB indicates an expected call of DeleteRecordsDB.
func (mr *MockDBMockRecorder) DeleteRecordsDB(ctx context.Context, arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock,
		"DeleteRecordsDB", reflect.TypeOf((*MockDB)(nil).DeleteRecordsDB), ctx, arg0)
}
