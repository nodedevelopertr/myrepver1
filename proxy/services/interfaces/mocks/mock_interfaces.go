// Code generated by MockGen. DO NOT EDIT.
// Source: proxy/services/interfaces/interfaces.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	context "context"
	io "io"
	reflect "reflect"

	driver "github.com/distribution/distribution/v3/registry/storage/driver"
	interfaces "github.com/forta-network/disco/proxy/services/interfaces"
	gomock "github.com/golang/mock/gomock"
	shell "github.com/ipfs/go-ipfs-api"
)

// MockIPFSClient is a mock of IPFSClient interface.
type MockIPFSClient struct {
	ctrl     *gomock.Controller
	recorder *MockIPFSClientMockRecorder
}

// MockIPFSClientMockRecorder is the mock recorder for MockIPFSClient.
type MockIPFSClientMockRecorder struct {
	mock *MockIPFSClient
}

// NewMockIPFSClient creates a new mock instance.
func NewMockIPFSClient(ctrl *gomock.Controller) *MockIPFSClient {
	mock := &MockIPFSClient{ctrl: ctrl}
	mock.recorder = &MockIPFSClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPFSClient) EXPECT() *MockIPFSClientMockRecorder {
	return m.recorder
}

// FilesCp mocks base method.
func (m *MockIPFSClient) FilesCp(ctx context.Context, src, dest string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilesCp", ctx, src, dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// FilesCp indicates an expected call of FilesCp.
func (mr *MockIPFSClientMockRecorder) FilesCp(ctx, src, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesCp", reflect.TypeOf((*MockIPFSClient)(nil).FilesCp), ctx, src, dest)
}

// FilesLs mocks base method.
func (m *MockIPFSClient) FilesLs(ctx context.Context, path string, options ...shell.FilesOpt) ([]*shell.MfsLsEntry, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, path}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FilesLs", varargs...)
	ret0, _ := ret[0].([]*shell.MfsLsEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilesLs indicates an expected call of FilesLs.
func (mr *MockIPFSClientMockRecorder) FilesLs(ctx, path interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, path}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesLs", reflect.TypeOf((*MockIPFSClient)(nil).FilesLs), varargs...)
}

// FilesMkdir mocks base method.
func (m *MockIPFSClient) FilesMkdir(ctx context.Context, path string, options ...shell.FilesOpt) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, path}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FilesMkdir", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// FilesMkdir indicates an expected call of FilesMkdir.
func (mr *MockIPFSClientMockRecorder) FilesMkdir(ctx, path interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, path}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesMkdir", reflect.TypeOf((*MockIPFSClient)(nil).FilesMkdir), varargs...)
}

// FilesMv mocks base method.
func (m *MockIPFSClient) FilesMv(ctx context.Context, src, dest string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilesMv", ctx, src, dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// FilesMv indicates an expected call of FilesMv.
func (mr *MockIPFSClientMockRecorder) FilesMv(ctx, src, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesMv", reflect.TypeOf((*MockIPFSClient)(nil).FilesMv), ctx, src, dest)
}

// FilesRead mocks base method.
func (m *MockIPFSClient) FilesRead(ctx context.Context, path string, options ...shell.FilesOpt) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, path}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FilesRead", varargs...)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilesRead indicates an expected call of FilesRead.
func (mr *MockIPFSClientMockRecorder) FilesRead(ctx, path interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, path}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesRead", reflect.TypeOf((*MockIPFSClient)(nil).FilesRead), varargs...)
}

// FilesRm mocks base method.
func (m *MockIPFSClient) FilesRm(ctx context.Context, path string, force bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilesRm", ctx, path, force)
	ret0, _ := ret[0].(error)
	return ret0
}

// FilesRm indicates an expected call of FilesRm.
func (mr *MockIPFSClientMockRecorder) FilesRm(ctx, path, force interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesRm", reflect.TypeOf((*MockIPFSClient)(nil).FilesRm), ctx, path, force)
}

// FilesStat mocks base method.
func (m *MockIPFSClient) FilesStat(ctx context.Context, path string, options ...shell.FilesOpt) (*shell.FilesStatObject, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, path}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FilesStat", varargs...)
	ret0, _ := ret[0].(*shell.FilesStatObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilesStat indicates an expected call of FilesStat.
func (mr *MockIPFSClientMockRecorder) FilesStat(ctx, path interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, path}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesStat", reflect.TypeOf((*MockIPFSClient)(nil).FilesStat), varargs...)
}

// FilesWrite mocks base method.
func (m *MockIPFSClient) FilesWrite(ctx context.Context, path string, data io.Reader, options ...shell.FilesOpt) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, path, data}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FilesWrite", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// FilesWrite indicates an expected call of FilesWrite.
func (mr *MockIPFSClientMockRecorder) FilesWrite(ctx, path, data interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, path, data}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesWrite", reflect.TypeOf((*MockIPFSClient)(nil).FilesWrite), varargs...)
}

// GetClientFor mocks base method.
func (m *MockIPFSClient) GetClientFor(ctx context.Context, path string) (interfaces.IPFSFilesAPI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientFor", ctx, path)
	ret0, _ := ret[0].(interfaces.IPFSFilesAPI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientFor indicates an expected call of GetClientFor.
func (mr *MockIPFSClientMockRecorder) GetClientFor(ctx, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientFor", reflect.TypeOf((*MockIPFSClient)(nil).GetClientFor), ctx, path)
}

// MockIPFSFilesAPI is a mock of IPFSFilesAPI interface.
type MockIPFSFilesAPI struct {
	ctrl     *gomock.Controller
	recorder *MockIPFSFilesAPIMockRecorder
}

// MockIPFSFilesAPIMockRecorder is the mock recorder for MockIPFSFilesAPI.
type MockIPFSFilesAPIMockRecorder struct {
	mock *MockIPFSFilesAPI
}

// NewMockIPFSFilesAPI creates a new mock instance.
func NewMockIPFSFilesAPI(ctrl *gomock.Controller) *MockIPFSFilesAPI {
	mock := &MockIPFSFilesAPI{ctrl: ctrl}
	mock.recorder = &MockIPFSFilesAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPFSFilesAPI) EXPECT() *MockIPFSFilesAPIMockRecorder {
	return m.recorder
}

// FilesCp mocks base method.
func (m *MockIPFSFilesAPI) FilesCp(ctx context.Context, src, dest string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilesCp", ctx, src, dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// FilesCp indicates an expected call of FilesCp.
func (mr *MockIPFSFilesAPIMockRecorder) FilesCp(ctx, src, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesCp", reflect.TypeOf((*MockIPFSFilesAPI)(nil).FilesCp), ctx, src, dest)
}

// FilesLs mocks base method.
func (m *MockIPFSFilesAPI) FilesLs(ctx context.Context, path string, options ...shell.FilesOpt) ([]*shell.MfsLsEntry, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, path}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FilesLs", varargs...)
	ret0, _ := ret[0].([]*shell.MfsLsEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilesLs indicates an expected call of FilesLs.
func (mr *MockIPFSFilesAPIMockRecorder) FilesLs(ctx, path interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, path}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesLs", reflect.TypeOf((*MockIPFSFilesAPI)(nil).FilesLs), varargs...)
}

// FilesMkdir mocks base method.
func (m *MockIPFSFilesAPI) FilesMkdir(ctx context.Context, path string, options ...shell.FilesOpt) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, path}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FilesMkdir", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// FilesMkdir indicates an expected call of FilesMkdir.
func (mr *MockIPFSFilesAPIMockRecorder) FilesMkdir(ctx, path interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, path}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesMkdir", reflect.TypeOf((*MockIPFSFilesAPI)(nil).FilesMkdir), varargs...)
}

// FilesMv mocks base method.
func (m *MockIPFSFilesAPI) FilesMv(ctx context.Context, src, dest string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilesMv", ctx, src, dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// FilesMv indicates an expected call of FilesMv.
func (mr *MockIPFSFilesAPIMockRecorder) FilesMv(ctx, src, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesMv", reflect.TypeOf((*MockIPFSFilesAPI)(nil).FilesMv), ctx, src, dest)
}

// FilesRead mocks base method.
func (m *MockIPFSFilesAPI) FilesRead(ctx context.Context, path string, options ...shell.FilesOpt) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, path}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FilesRead", varargs...)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilesRead indicates an expected call of FilesRead.
func (mr *MockIPFSFilesAPIMockRecorder) FilesRead(ctx, path interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, path}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesRead", reflect.TypeOf((*MockIPFSFilesAPI)(nil).FilesRead), varargs...)
}

// FilesRm mocks base method.
func (m *MockIPFSFilesAPI) FilesRm(ctx context.Context, path string, force bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilesRm", ctx, path, force)
	ret0, _ := ret[0].(error)
	return ret0
}

// FilesRm indicates an expected call of FilesRm.
func (mr *MockIPFSFilesAPIMockRecorder) FilesRm(ctx, path, force interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesRm", reflect.TypeOf((*MockIPFSFilesAPI)(nil).FilesRm), ctx, path, force)
}

// FilesStat mocks base method.
func (m *MockIPFSFilesAPI) FilesStat(ctx context.Context, path string, options ...shell.FilesOpt) (*shell.FilesStatObject, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, path}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FilesStat", varargs...)
	ret0, _ := ret[0].(*shell.FilesStatObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilesStat indicates an expected call of FilesStat.
func (mr *MockIPFSFilesAPIMockRecorder) FilesStat(ctx, path interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, path}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesStat", reflect.TypeOf((*MockIPFSFilesAPI)(nil).FilesStat), varargs...)
}

// FilesWrite mocks base method.
func (m *MockIPFSFilesAPI) FilesWrite(ctx context.Context, path string, data io.Reader, options ...shell.FilesOpt) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, path, data}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FilesWrite", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// FilesWrite indicates an expected call of FilesWrite.
func (mr *MockIPFSFilesAPIMockRecorder) FilesWrite(ctx, path, data interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, path, data}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesWrite", reflect.TypeOf((*MockIPFSFilesAPI)(nil).FilesWrite), varargs...)
}

// MockStorageDriver is a mock of StorageDriver interface.
type MockStorageDriver struct {
	ctrl     *gomock.Controller
	recorder *MockStorageDriverMockRecorder
}

// MockStorageDriverMockRecorder is the mock recorder for MockStorageDriver.
type MockStorageDriverMockRecorder struct {
	mock *MockStorageDriver
}

// NewMockStorageDriver creates a new mock instance.
func NewMockStorageDriver(ctrl *gomock.Controller) *MockStorageDriver {
	mock := &MockStorageDriver{ctrl: ctrl}
	mock.recorder = &MockStorageDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageDriver) EXPECT() *MockStorageDriverMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockStorageDriver) Delete(ctx context.Context, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockStorageDriverMockRecorder) Delete(ctx, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStorageDriver)(nil).Delete), ctx, path)
}

// GetContent mocks base method.
func (m *MockStorageDriver) GetContent(ctx context.Context, path string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContent", ctx, path)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContent indicates an expected call of GetContent.
func (mr *MockStorageDriverMockRecorder) GetContent(ctx, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContent", reflect.TypeOf((*MockStorageDriver)(nil).GetContent), ctx, path)
}

// List mocks base method.
func (m *MockStorageDriver) List(ctx context.Context, path string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, path)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockStorageDriverMockRecorder) List(ctx, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStorageDriver)(nil).List), ctx, path)
}

// Move mocks base method.
func (m *MockStorageDriver) Move(ctx context.Context, sourcePath, destPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Move", ctx, sourcePath, destPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// Move indicates an expected call of Move.
func (mr *MockStorageDriverMockRecorder) Move(ctx, sourcePath, destPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Move", reflect.TypeOf((*MockStorageDriver)(nil).Move), ctx, sourcePath, destPath)
}

// Name mocks base method.
func (m *MockStorageDriver) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockStorageDriverMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockStorageDriver)(nil).Name))
}

// PutContent mocks base method.
func (m *MockStorageDriver) PutContent(ctx context.Context, path string, content []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutContent", ctx, path, content)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutContent indicates an expected call of PutContent.
func (mr *MockStorageDriverMockRecorder) PutContent(ctx, path, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutContent", reflect.TypeOf((*MockStorageDriver)(nil).PutContent), ctx, path, content)
}

// Reader mocks base method.
func (m *MockStorageDriver) Reader(ctx context.Context, path string, offset int64) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reader", ctx, path, offset)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reader indicates an expected call of Reader.
func (mr *MockStorageDriverMockRecorder) Reader(ctx, path, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reader", reflect.TypeOf((*MockStorageDriver)(nil).Reader), ctx, path, offset)
}

// Stat mocks base method.
func (m *MockStorageDriver) Stat(ctx context.Context, path string) (driver.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat", ctx, path)
	ret0, _ := ret[0].(driver.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat.
func (mr *MockStorageDriverMockRecorder) Stat(ctx, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockStorageDriver)(nil).Stat), ctx, path)
}

// URLFor mocks base method.
func (m *MockStorageDriver) URLFor(ctx context.Context, path string, options map[string]interface{}) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URLFor", ctx, path, options)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// URLFor indicates an expected call of URLFor.
func (mr *MockStorageDriverMockRecorder) URLFor(ctx, path, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URLFor", reflect.TypeOf((*MockStorageDriver)(nil).URLFor), ctx, path, options)
}

// Walk mocks base method.
func (m *MockStorageDriver) Walk(ctx context.Context, path string, f driver.WalkFn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Walk", ctx, path, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// Walk indicates an expected call of Walk.
func (mr *MockStorageDriverMockRecorder) Walk(ctx, path, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Walk", reflect.TypeOf((*MockStorageDriver)(nil).Walk), ctx, path, f)
}

// Writer mocks base method.
func (m *MockStorageDriver) Writer(ctx context.Context, path string, append bool) (driver.FileWriter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Writer", ctx, path, append)
	ret0, _ := ret[0].(driver.FileWriter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Writer indicates an expected call of Writer.
func (mr *MockStorageDriverMockRecorder) Writer(ctx, path, append interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Writer", reflect.TypeOf((*MockStorageDriver)(nil).Writer), ctx, path, append)
}
