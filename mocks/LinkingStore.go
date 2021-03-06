// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	urls "url-short/internal/repos/urls"

	mock "github.com/stretchr/testify/mock"
)

// LinkingStore is an autogenerated mock type for the LinkingStore type
type LinkingStore struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, hash
func (_m *LinkingStore) Create(ctx context.Context, hash urls.UrlHash) error {
	ret := _m.Called(ctx, hash)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, urls.UrlHash) error); ok {
		r0 = rf(ctx, hash)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Read provides a mock function with given fields: ctx, hash
func (_m *LinkingStore) Read(ctx context.Context, hash urls.UrlHash) chan string {
	ret := _m.Called(ctx, hash)

	var r0 chan string
	if rf, ok := ret.Get(0).(func(context.Context, urls.UrlHash) chan string); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan string)
		}
	}

	return r0
}
