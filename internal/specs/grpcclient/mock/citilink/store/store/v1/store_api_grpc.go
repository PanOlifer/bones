// Code generated by MockGen. DO NOT EDIT.
// Source: ./gen/citilink/store/store/v1/store_api_grpc.pb.go

// Package mock_storev1 is a generated GoMock package.
package mock_storev1

import (
	context "context"
	reflect "reflect"

	storev1 "go.citilink.cloud/grpc-skeleton/internal/specs/grpcclient/gen/citilink/store/store/v1"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockStoreAPIClient is a mock of StoreAPIClient interface.
type MockStoreAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockStoreAPIClientMockRecorder
}

// MockStoreAPIClientMockRecorder is the mock recorder for MockStoreAPIClient.
type MockStoreAPIClientMockRecorder struct {
	mock *MockStoreAPIClient
}

// NewMockStoreAPIClient creates a new mock instance.
func NewMockStoreAPIClient(ctrl *gomock.Controller) *MockStoreAPIClient {
	mock := &MockStoreAPIClient{ctrl: ctrl}
	mock.recorder = &MockStoreAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoreAPIClient) EXPECT() *MockStoreAPIClientMockRecorder {
	return m.recorder
}

// Filter mocks base method.
func (m *MockStoreAPIClient) Filter(ctx context.Context, in *storev1.FilterRequest, opts ...grpc.CallOption) (*storev1.FilterResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Filter", varargs...)
	ret0, _ := ret[0].(*storev1.FilterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter.
func (mr *MockStoreAPIClientMockRecorder) Filter(ctx, in interface{}, opts ...interface{}) *StoreAPIClientFilterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockStoreAPIClient)(nil).Filter), varargs...)
	return &StoreAPIClientFilterCall{Call: call}
}

// StoreAPIClientFilterCall wrap *gomock.Call
type StoreAPIClientFilterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StoreAPIClientFilterCall) Return(arg0 *storev1.FilterResponse, arg1 error) *StoreAPIClientFilterCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StoreAPIClientFilterCall) Do(f func(context.Context, *storev1.FilterRequest, ...grpc.CallOption) (*storev1.FilterResponse, error)) *StoreAPIClientFilterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StoreAPIClientFilterCall) DoAndReturn(f func(context.Context, *storev1.FilterRequest, ...grpc.CallOption) (*storev1.FilterResponse, error)) *StoreAPIClientFilterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// FindAll mocks base method.
func (m *MockStoreAPIClient) FindAll(ctx context.Context, in *storev1.FindAllRequest, opts ...grpc.CallOption) (*storev1.FindAllResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindAll", varargs...)
	ret0, _ := ret[0].(*storev1.FindAllResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockStoreAPIClientMockRecorder) FindAll(ctx, in interface{}, opts ...interface{}) *StoreAPIClientFindAllCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockStoreAPIClient)(nil).FindAll), varargs...)
	return &StoreAPIClientFindAllCall{Call: call}
}

// StoreAPIClientFindAllCall wrap *gomock.Call
type StoreAPIClientFindAllCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StoreAPIClientFindAllCall) Return(arg0 *storev1.FindAllResponse, arg1 error) *StoreAPIClientFindAllCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StoreAPIClientFindAllCall) Do(f func(context.Context, *storev1.FindAllRequest, ...grpc.CallOption) (*storev1.FindAllResponse, error)) *StoreAPIClientFindAllCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StoreAPIClientFindAllCall) DoAndReturn(f func(context.Context, *storev1.FindAllRequest, ...grpc.CallOption) (*storev1.FindAllResponse, error)) *StoreAPIClientFindAllCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSpaceIds mocks base method.
func (m *MockStoreAPIClient) GetSpaceIds(ctx context.Context, in *storev1.GetSpaceIdsRequest, opts ...grpc.CallOption) (*storev1.GetSpaceIdsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSpaceIds", varargs...)
	ret0, _ := ret[0].(*storev1.GetSpaceIdsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpaceIds indicates an expected call of GetSpaceIds.
func (mr *MockStoreAPIClientMockRecorder) GetSpaceIds(ctx, in interface{}, opts ...interface{}) *StoreAPIClientGetSpaceIdsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpaceIds", reflect.TypeOf((*MockStoreAPIClient)(nil).GetSpaceIds), varargs...)
	return &StoreAPIClientGetSpaceIdsCall{Call: call}
}

// StoreAPIClientGetSpaceIdsCall wrap *gomock.Call
type StoreAPIClientGetSpaceIdsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StoreAPIClientGetSpaceIdsCall) Return(arg0 *storev1.GetSpaceIdsResponse, arg1 error) *StoreAPIClientGetSpaceIdsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StoreAPIClientGetSpaceIdsCall) Do(f func(context.Context, *storev1.GetSpaceIdsRequest, ...grpc.CallOption) (*storev1.GetSpaceIdsResponse, error)) *StoreAPIClientGetSpaceIdsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StoreAPIClientGetSpaceIdsCall) DoAndReturn(f func(context.Context, *storev1.GetSpaceIdsRequest, ...grpc.CallOption) (*storev1.GetSpaceIdsResponse, error)) *StoreAPIClientGetSpaceIdsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetTerminalSpaceIds mocks base method.
func (m *MockStoreAPIClient) GetTerminalSpaceIds(ctx context.Context, in *storev1.GetTerminalSpaceIdsRequest, opts ...grpc.CallOption) (*storev1.GetTerminalSpaceIdsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTerminalSpaceIds", varargs...)
	ret0, _ := ret[0].(*storev1.GetTerminalSpaceIdsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTerminalSpaceIds indicates an expected call of GetTerminalSpaceIds.
func (mr *MockStoreAPIClientMockRecorder) GetTerminalSpaceIds(ctx, in interface{}, opts ...interface{}) *StoreAPIClientGetTerminalSpaceIdsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTerminalSpaceIds", reflect.TypeOf((*MockStoreAPIClient)(nil).GetTerminalSpaceIds), varargs...)
	return &StoreAPIClientGetTerminalSpaceIdsCall{Call: call}
}

// StoreAPIClientGetTerminalSpaceIdsCall wrap *gomock.Call
type StoreAPIClientGetTerminalSpaceIdsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StoreAPIClientGetTerminalSpaceIdsCall) Return(arg0 *storev1.GetTerminalSpaceIdsResponse, arg1 error) *StoreAPIClientGetTerminalSpaceIdsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StoreAPIClientGetTerminalSpaceIdsCall) Do(f func(context.Context, *storev1.GetTerminalSpaceIdsRequest, ...grpc.CallOption) (*storev1.GetTerminalSpaceIdsResponse, error)) *StoreAPIClientGetTerminalSpaceIdsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StoreAPIClientGetTerminalSpaceIdsCall) DoAndReturn(f func(context.Context, *storev1.GetTerminalSpaceIdsRequest, ...grpc.CallOption) (*storev1.GetTerminalSpaceIdsResponse, error)) *StoreAPIClientGetTerminalSpaceIdsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockStoreAPIServer is a mock of StoreAPIServer interface.
type MockStoreAPIServer struct {
	ctrl     *gomock.Controller
	recorder *MockStoreAPIServerMockRecorder
}

// MockStoreAPIServerMockRecorder is the mock recorder for MockStoreAPIServer.
type MockStoreAPIServerMockRecorder struct {
	mock *MockStoreAPIServer
}

// NewMockStoreAPIServer creates a new mock instance.
func NewMockStoreAPIServer(ctrl *gomock.Controller) *MockStoreAPIServer {
	mock := &MockStoreAPIServer{ctrl: ctrl}
	mock.recorder = &MockStoreAPIServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoreAPIServer) EXPECT() *MockStoreAPIServerMockRecorder {
	return m.recorder
}

// Filter mocks base method.
func (m *MockStoreAPIServer) Filter(arg0 context.Context, arg1 *storev1.FilterRequest) (*storev1.FilterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filter", arg0, arg1)
	ret0, _ := ret[0].(*storev1.FilterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter.
func (mr *MockStoreAPIServerMockRecorder) Filter(arg0, arg1 interface{}) *StoreAPIServerFilterCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockStoreAPIServer)(nil).Filter), arg0, arg1)
	return &StoreAPIServerFilterCall{Call: call}
}

// StoreAPIServerFilterCall wrap *gomock.Call
type StoreAPIServerFilterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StoreAPIServerFilterCall) Return(arg0 *storev1.FilterResponse, arg1 error) *StoreAPIServerFilterCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StoreAPIServerFilterCall) Do(f func(context.Context, *storev1.FilterRequest) (*storev1.FilterResponse, error)) *StoreAPIServerFilterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StoreAPIServerFilterCall) DoAndReturn(f func(context.Context, *storev1.FilterRequest) (*storev1.FilterResponse, error)) *StoreAPIServerFilterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// FindAll mocks base method.
func (m *MockStoreAPIServer) FindAll(arg0 context.Context, arg1 *storev1.FindAllRequest) (*storev1.FindAllResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0, arg1)
	ret0, _ := ret[0].(*storev1.FindAllResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockStoreAPIServerMockRecorder) FindAll(arg0, arg1 interface{}) *StoreAPIServerFindAllCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockStoreAPIServer)(nil).FindAll), arg0, arg1)
	return &StoreAPIServerFindAllCall{Call: call}
}

// StoreAPIServerFindAllCall wrap *gomock.Call
type StoreAPIServerFindAllCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StoreAPIServerFindAllCall) Return(arg0 *storev1.FindAllResponse, arg1 error) *StoreAPIServerFindAllCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StoreAPIServerFindAllCall) Do(f func(context.Context, *storev1.FindAllRequest) (*storev1.FindAllResponse, error)) *StoreAPIServerFindAllCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StoreAPIServerFindAllCall) DoAndReturn(f func(context.Context, *storev1.FindAllRequest) (*storev1.FindAllResponse, error)) *StoreAPIServerFindAllCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSpaceIds mocks base method.
func (m *MockStoreAPIServer) GetSpaceIds(arg0 context.Context, arg1 *storev1.GetSpaceIdsRequest) (*storev1.GetSpaceIdsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpaceIds", arg0, arg1)
	ret0, _ := ret[0].(*storev1.GetSpaceIdsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpaceIds indicates an expected call of GetSpaceIds.
func (mr *MockStoreAPIServerMockRecorder) GetSpaceIds(arg0, arg1 interface{}) *StoreAPIServerGetSpaceIdsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpaceIds", reflect.TypeOf((*MockStoreAPIServer)(nil).GetSpaceIds), arg0, arg1)
	return &StoreAPIServerGetSpaceIdsCall{Call: call}
}

// StoreAPIServerGetSpaceIdsCall wrap *gomock.Call
type StoreAPIServerGetSpaceIdsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StoreAPIServerGetSpaceIdsCall) Return(arg0 *storev1.GetSpaceIdsResponse, arg1 error) *StoreAPIServerGetSpaceIdsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StoreAPIServerGetSpaceIdsCall) Do(f func(context.Context, *storev1.GetSpaceIdsRequest) (*storev1.GetSpaceIdsResponse, error)) *StoreAPIServerGetSpaceIdsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StoreAPIServerGetSpaceIdsCall) DoAndReturn(f func(context.Context, *storev1.GetSpaceIdsRequest) (*storev1.GetSpaceIdsResponse, error)) *StoreAPIServerGetSpaceIdsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetTerminalSpaceIds mocks base method.
func (m *MockStoreAPIServer) GetTerminalSpaceIds(arg0 context.Context, arg1 *storev1.GetTerminalSpaceIdsRequest) (*storev1.GetTerminalSpaceIdsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTerminalSpaceIds", arg0, arg1)
	ret0, _ := ret[0].(*storev1.GetTerminalSpaceIdsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTerminalSpaceIds indicates an expected call of GetTerminalSpaceIds.
func (mr *MockStoreAPIServerMockRecorder) GetTerminalSpaceIds(arg0, arg1 interface{}) *StoreAPIServerGetTerminalSpaceIdsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTerminalSpaceIds", reflect.TypeOf((*MockStoreAPIServer)(nil).GetTerminalSpaceIds), arg0, arg1)
	return &StoreAPIServerGetTerminalSpaceIdsCall{Call: call}
}

// StoreAPIServerGetTerminalSpaceIdsCall wrap *gomock.Call
type StoreAPIServerGetTerminalSpaceIdsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StoreAPIServerGetTerminalSpaceIdsCall) Return(arg0 *storev1.GetTerminalSpaceIdsResponse, arg1 error) *StoreAPIServerGetTerminalSpaceIdsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StoreAPIServerGetTerminalSpaceIdsCall) Do(f func(context.Context, *storev1.GetTerminalSpaceIdsRequest) (*storev1.GetTerminalSpaceIdsResponse, error)) *StoreAPIServerGetTerminalSpaceIdsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StoreAPIServerGetTerminalSpaceIdsCall) DoAndReturn(f func(context.Context, *storev1.GetTerminalSpaceIdsRequest) (*storev1.GetTerminalSpaceIdsResponse, error)) *StoreAPIServerGetTerminalSpaceIdsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// mustEmbedUnimplementedStoreAPIServer mocks base method.
func (m *MockStoreAPIServer) mustEmbedUnimplementedStoreAPIServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedStoreAPIServer")
}

// mustEmbedUnimplementedStoreAPIServer indicates an expected call of mustEmbedUnimplementedStoreAPIServer.
func (mr *MockStoreAPIServerMockRecorder) mustEmbedUnimplementedStoreAPIServer() *StoreAPIServermustEmbedUnimplementedStoreAPIServerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedStoreAPIServer", reflect.TypeOf((*MockStoreAPIServer)(nil).mustEmbedUnimplementedStoreAPIServer))
	return &StoreAPIServermustEmbedUnimplementedStoreAPIServerCall{Call: call}
}

// StoreAPIServermustEmbedUnimplementedStoreAPIServerCall wrap *gomock.Call
type StoreAPIServermustEmbedUnimplementedStoreAPIServerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StoreAPIServermustEmbedUnimplementedStoreAPIServerCall) Return() *StoreAPIServermustEmbedUnimplementedStoreAPIServerCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StoreAPIServermustEmbedUnimplementedStoreAPIServerCall) Do(f func()) *StoreAPIServermustEmbedUnimplementedStoreAPIServerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StoreAPIServermustEmbedUnimplementedStoreAPIServerCall) DoAndReturn(f func()) *StoreAPIServermustEmbedUnimplementedStoreAPIServerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockUnsafeStoreAPIServer is a mock of UnsafeStoreAPIServer interface.
type MockUnsafeStoreAPIServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeStoreAPIServerMockRecorder
}

// MockUnsafeStoreAPIServerMockRecorder is the mock recorder for MockUnsafeStoreAPIServer.
type MockUnsafeStoreAPIServerMockRecorder struct {
	mock *MockUnsafeStoreAPIServer
}

// NewMockUnsafeStoreAPIServer creates a new mock instance.
func NewMockUnsafeStoreAPIServer(ctrl *gomock.Controller) *MockUnsafeStoreAPIServer {
	mock := &MockUnsafeStoreAPIServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeStoreAPIServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeStoreAPIServer) EXPECT() *MockUnsafeStoreAPIServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedStoreAPIServer mocks base method.
func (m *MockUnsafeStoreAPIServer) mustEmbedUnimplementedStoreAPIServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedStoreAPIServer")
}

// mustEmbedUnimplementedStoreAPIServer indicates an expected call of mustEmbedUnimplementedStoreAPIServer.
func (mr *MockUnsafeStoreAPIServerMockRecorder) mustEmbedUnimplementedStoreAPIServer() *UnsafeStoreAPIServermustEmbedUnimplementedStoreAPIServerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedStoreAPIServer", reflect.TypeOf((*MockUnsafeStoreAPIServer)(nil).mustEmbedUnimplementedStoreAPIServer))
	return &UnsafeStoreAPIServermustEmbedUnimplementedStoreAPIServerCall{Call: call}
}

// UnsafeStoreAPIServermustEmbedUnimplementedStoreAPIServerCall wrap *gomock.Call
type UnsafeStoreAPIServermustEmbedUnimplementedStoreAPIServerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *UnsafeStoreAPIServermustEmbedUnimplementedStoreAPIServerCall) Return() *UnsafeStoreAPIServermustEmbedUnimplementedStoreAPIServerCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *UnsafeStoreAPIServermustEmbedUnimplementedStoreAPIServerCall) Do(f func()) *UnsafeStoreAPIServermustEmbedUnimplementedStoreAPIServerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *UnsafeStoreAPIServermustEmbedUnimplementedStoreAPIServerCall) DoAndReturn(f func()) *UnsafeStoreAPIServermustEmbedUnimplementedStoreAPIServerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
