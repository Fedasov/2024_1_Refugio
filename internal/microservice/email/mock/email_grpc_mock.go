// Code generated by MockGen. DO NOT EDIT.
// Source: ./email_grpc.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	proto "mail/internal/microservice/email/proto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockEmailServiceClient is a mock of EmailServiceClient interface.
type MockEmailServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockEmailServiceClientMockRecorder
}

// MockEmailServiceClientMockRecorder is the mock recorder for MockEmailServiceClient.
type MockEmailServiceClientMockRecorder struct {
	mock *MockEmailServiceClient
}

// NewMockEmailServiceClient creates a new mock instance.
func NewMockEmailServiceClient(ctrl *gomock.Controller) *MockEmailServiceClient {
	mock := &MockEmailServiceClient{ctrl: ctrl}
	mock.recorder = &MockEmailServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailServiceClient) EXPECT() *MockEmailServiceClientMockRecorder {
	return m.recorder
}

// CheckRecipientEmail mocks base method.
func (m *MockEmailServiceClient) CheckRecipientEmail(ctx context.Context, in *proto.Recipient, opts ...grpc.CallOption) (*proto.EmptyEmail, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckRecipientEmail", varargs...)
	ret0, _ := ret[0].(*proto.EmptyEmail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckRecipientEmail indicates an expected call of CheckRecipientEmail.
func (mr *MockEmailServiceClientMockRecorder) CheckRecipientEmail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRecipientEmail", reflect.TypeOf((*MockEmailServiceClient)(nil).CheckRecipientEmail), varargs...)
}

// CreateEmail mocks base method.
func (m *MockEmailServiceClient) CreateEmail(ctx context.Context, in *proto.Email, opts ...grpc.CallOption) (*proto.EmailWithID, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateEmail", varargs...)
	ret0, _ := ret[0].(*proto.EmailWithID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEmail indicates an expected call of CreateEmail.
func (mr *MockEmailServiceClientMockRecorder) CreateEmail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmail", reflect.TypeOf((*MockEmailServiceClient)(nil).CreateEmail), varargs...)
}

// CreateProfileEmail mocks base method.
func (m *MockEmailServiceClient) CreateProfileEmail(ctx context.Context, in *proto.IdSenderRecipient, opts ...grpc.CallOption) (*proto.EmptyEmail, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateProfileEmail", varargs...)
	ret0, _ := ret[0].(*proto.EmptyEmail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProfileEmail indicates an expected call of CreateProfileEmail.
func (mr *MockEmailServiceClientMockRecorder) CreateProfileEmail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProfileEmail", reflect.TypeOf((*MockEmailServiceClient)(nil).CreateProfileEmail), varargs...)
}

// DeleteEmail mocks base method.
func (m *MockEmailServiceClient) DeleteEmail(ctx context.Context, in *proto.LoginWithID, opts ...grpc.CallOption) (*proto.StatusEmail, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteEmail", varargs...)
	ret0, _ := ret[0].(*proto.StatusEmail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEmail indicates an expected call of DeleteEmail.
func (mr *MockEmailServiceClientMockRecorder) DeleteEmail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEmail", reflect.TypeOf((*MockEmailServiceClient)(nil).DeleteEmail), varargs...)
}

// GetAllIncoming mocks base method.
func (m *MockEmailServiceClient) GetAllIncoming(ctx context.Context, in *proto.LoginOffsetLimit, opts ...grpc.CallOption) (*proto.Emails, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllIncoming", varargs...)
	ret0, _ := ret[0].(*proto.Emails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllIncoming indicates an expected call of GetAllIncoming.
func (mr *MockEmailServiceClientMockRecorder) GetAllIncoming(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllIncoming", reflect.TypeOf((*MockEmailServiceClient)(nil).GetAllIncoming), varargs...)
}

// GetAllSent mocks base method.
func (m *MockEmailServiceClient) GetAllSent(ctx context.Context, in *proto.LoginOffsetLimit, opts ...grpc.CallOption) (*proto.Emails, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllSent", varargs...)
	ret0, _ := ret[0].(*proto.Emails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSent indicates an expected call of GetAllSent.
func (mr *MockEmailServiceClientMockRecorder) GetAllSent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSent", reflect.TypeOf((*MockEmailServiceClient)(nil).GetAllSent), varargs...)
}

// GetDraftEmails mocks base method.
func (m *MockEmailServiceClient) GetDraftEmails(ctx context.Context, in *proto.LoginOffsetLimit, opts ...grpc.CallOption) (*proto.Emails, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDraftEmails", varargs...)
	ret0, _ := ret[0].(*proto.Emails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDraftEmails indicates an expected call of GetDraftEmails.
func (mr *MockEmailServiceClientMockRecorder) GetDraftEmails(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDraftEmails", reflect.TypeOf((*MockEmailServiceClient)(nil).GetDraftEmails), varargs...)
}

// GetEmailByID mocks base method.
func (m *MockEmailServiceClient) GetEmailByID(ctx context.Context, in *proto.EmailIdAndLogin, opts ...grpc.CallOption) (*proto.Email, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEmailByID", varargs...)
	ret0, _ := ret[0].(*proto.Email)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmailByID indicates an expected call of GetEmailByID.
func (mr *MockEmailServiceClientMockRecorder) GetEmailByID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmailByID", reflect.TypeOf((*MockEmailServiceClient)(nil).GetEmailByID), varargs...)
}

// GetSpamEmails mocks base method.
func (m *MockEmailServiceClient) GetSpamEmails(ctx context.Context, in *proto.LoginOffsetLimit, opts ...grpc.CallOption) (*proto.Emails, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSpamEmails", varargs...)
	ret0, _ := ret[0].(*proto.Emails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpamEmails indicates an expected call of GetSpamEmails.
func (mr *MockEmailServiceClientMockRecorder) GetSpamEmails(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpamEmails", reflect.TypeOf((*MockEmailServiceClient)(nil).GetSpamEmails), varargs...)
}

// UpdateEmail mocks base method.
func (m *MockEmailServiceClient) UpdateEmail(ctx context.Context, in *proto.Email, opts ...grpc.CallOption) (*proto.StatusEmail, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateEmail", varargs...)
	ret0, _ := ret[0].(*proto.StatusEmail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEmail indicates an expected call of UpdateEmail.
func (mr *MockEmailServiceClientMockRecorder) UpdateEmail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmail", reflect.TypeOf((*MockEmailServiceClient)(nil).UpdateEmail), varargs...)
}

// MockEmailServiceServer is a mock of EmailServiceServer interface.
type MockEmailServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockEmailServiceServerMockRecorder
}

// MockEmailServiceServerMockRecorder is the mock recorder for MockEmailServiceServer.
type MockEmailServiceServerMockRecorder struct {
	mock *MockEmailServiceServer
}

// NewMockEmailServiceServer creates a new mock instance.
func NewMockEmailServiceServer(ctrl *gomock.Controller) *MockEmailServiceServer {
	mock := &MockEmailServiceServer{ctrl: ctrl}
	mock.recorder = &MockEmailServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailServiceServer) EXPECT() *MockEmailServiceServerMockRecorder {
	return m.recorder
}

// CheckRecipientEmail mocks base method.
func (m *MockEmailServiceServer) CheckRecipientEmail(arg0 context.Context, arg1 *proto.Recipient) (*proto.EmptyEmail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRecipientEmail", arg0, arg1)
	ret0, _ := ret[0].(*proto.EmptyEmail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckRecipientEmail indicates an expected call of CheckRecipientEmail.
func (mr *MockEmailServiceServerMockRecorder) CheckRecipientEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRecipientEmail", reflect.TypeOf((*MockEmailServiceServer)(nil).CheckRecipientEmail), arg0, arg1)
}

// CreateEmail mocks base method.
func (m *MockEmailServiceServer) CreateEmail(arg0 context.Context, arg1 *proto.Email) (*proto.EmailWithID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmail", arg0, arg1)
	ret0, _ := ret[0].(*proto.EmailWithID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEmail indicates an expected call of CreateEmail.
func (mr *MockEmailServiceServerMockRecorder) CreateEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmail", reflect.TypeOf((*MockEmailServiceServer)(nil).CreateEmail), arg0, arg1)
}

// CreateProfileEmail mocks base method.
func (m *MockEmailServiceServer) CreateProfileEmail(arg0 context.Context, arg1 *proto.IdSenderRecipient) (*proto.EmptyEmail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProfileEmail", arg0, arg1)
	ret0, _ := ret[0].(*proto.EmptyEmail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProfileEmail indicates an expected call of CreateProfileEmail.
func (mr *MockEmailServiceServerMockRecorder) CreateProfileEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProfileEmail", reflect.TypeOf((*MockEmailServiceServer)(nil).CreateProfileEmail), arg0, arg1)
}

// DeleteEmail mocks base method.
func (m *MockEmailServiceServer) DeleteEmail(arg0 context.Context, arg1 *proto.LoginWithID) (*proto.StatusEmail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEmail", arg0, arg1)
	ret0, _ := ret[0].(*proto.StatusEmail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEmail indicates an expected call of DeleteEmail.
func (mr *MockEmailServiceServerMockRecorder) DeleteEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEmail", reflect.TypeOf((*MockEmailServiceServer)(nil).DeleteEmail), arg0, arg1)
}

// GetAllIncoming mocks base method.
func (m *MockEmailServiceServer) GetAllIncoming(arg0 context.Context, arg1 *proto.LoginOffsetLimit) (*proto.Emails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllIncoming", arg0, arg1)
	ret0, _ := ret[0].(*proto.Emails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllIncoming indicates an expected call of GetAllIncoming.
func (mr *MockEmailServiceServerMockRecorder) GetAllIncoming(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllIncoming", reflect.TypeOf((*MockEmailServiceServer)(nil).GetAllIncoming), arg0, arg1)
}

// GetAllSent mocks base method.
func (m *MockEmailServiceServer) GetAllSent(arg0 context.Context, arg1 *proto.LoginOffsetLimit) (*proto.Emails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSent", arg0, arg1)
	ret0, _ := ret[0].(*proto.Emails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSent indicates an expected call of GetAllSent.
func (mr *MockEmailServiceServerMockRecorder) GetAllSent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSent", reflect.TypeOf((*MockEmailServiceServer)(nil).GetAllSent), arg0, arg1)
}

// GetDraftEmails mocks base method.
func (m *MockEmailServiceServer) GetDraftEmails(arg0 context.Context, arg1 *proto.LoginOffsetLimit) (*proto.Emails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDraftEmails", arg0, arg1)
	ret0, _ := ret[0].(*proto.Emails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDraftEmails indicates an expected call of GetDraftEmails.
func (mr *MockEmailServiceServerMockRecorder) GetDraftEmails(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDraftEmails", reflect.TypeOf((*MockEmailServiceServer)(nil).GetDraftEmails), arg0, arg1)
}

// GetEmailByID mocks base method.
func (m *MockEmailServiceServer) GetEmailByID(arg0 context.Context, arg1 *proto.EmailIdAndLogin) (*proto.Email, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmailByID", arg0, arg1)
	ret0, _ := ret[0].(*proto.Email)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmailByID indicates an expected call of GetEmailByID.
func (mr *MockEmailServiceServerMockRecorder) GetEmailByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmailByID", reflect.TypeOf((*MockEmailServiceServer)(nil).GetEmailByID), arg0, arg1)
}

// GetSpamEmails mocks base method.
func (m *MockEmailServiceServer) GetSpamEmails(arg0 context.Context, arg1 *proto.LoginOffsetLimit) (*proto.Emails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpamEmails", arg0, arg1)
	ret0, _ := ret[0].(*proto.Emails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpamEmails indicates an expected call of GetSpamEmails.
func (mr *MockEmailServiceServerMockRecorder) GetSpamEmails(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpamEmails", reflect.TypeOf((*MockEmailServiceServer)(nil).GetSpamEmails), arg0, arg1)
}

// UpdateEmail mocks base method.
func (m *MockEmailServiceServer) UpdateEmail(arg0 context.Context, arg1 *proto.Email) (*proto.StatusEmail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEmail", arg0, arg1)
	ret0, _ := ret[0].(*proto.StatusEmail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEmail indicates an expected call of UpdateEmail.
func (mr *MockEmailServiceServerMockRecorder) UpdateEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmail", reflect.TypeOf((*MockEmailServiceServer)(nil).UpdateEmail), arg0, arg1)
}

// mustEmbedUnimplementedEmailServiceServer mocks base method.
func (m *MockEmailServiceServer) mustEmbedUnimplementedEmailServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedEmailServiceServer")
}

// mustEmbedUnimplementedEmailServiceServer indicates an expected call of mustEmbedUnimplementedEmailServiceServer.
func (mr *MockEmailServiceServerMockRecorder) mustEmbedUnimplementedEmailServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedEmailServiceServer", reflect.TypeOf((*MockEmailServiceServer)(nil).mustEmbedUnimplementedEmailServiceServer))
}

// MockUnsafeEmailServiceServer is a mock of UnsafeEmailServiceServer interface.
type MockUnsafeEmailServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeEmailServiceServerMockRecorder
}

// MockUnsafeEmailServiceServerMockRecorder is the mock recorder for MockUnsafeEmailServiceServer.
type MockUnsafeEmailServiceServerMockRecorder struct {
	mock *MockUnsafeEmailServiceServer
}

// NewMockUnsafeEmailServiceServer creates a new mock instance.
func NewMockUnsafeEmailServiceServer(ctrl *gomock.Controller) *MockUnsafeEmailServiceServer {
	mock := &MockUnsafeEmailServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeEmailServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeEmailServiceServer) EXPECT() *MockUnsafeEmailServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedEmailServiceServer mocks base method.
func (m *MockUnsafeEmailServiceServer) mustEmbedUnimplementedEmailServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedEmailServiceServer")
}

// mustEmbedUnimplementedEmailServiceServer indicates an expected call of mustEmbedUnimplementedEmailServiceServer.
func (mr *MockUnsafeEmailServiceServerMockRecorder) mustEmbedUnimplementedEmailServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedEmailServiceServer", reflect.TypeOf((*MockUnsafeEmailServiceServer)(nil).mustEmbedUnimplementedEmailServiceServer))
}
