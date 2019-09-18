// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import context "context"
import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-incubator/compass/components/director/internal/model"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, tenant, apiID, runtimeID
func (_m *Repository) Delete(ctx context.Context, tenant string, apiID string, runtimeID string) error {
	ret := _m.Called(ctx, tenant, apiID, runtimeID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, tenant, apiID, runtimeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, tenant, apiID, runtimeID
func (_m *Repository) Get(ctx context.Context, tenant string, apiID string, runtimeID string) (*model.RuntimeAuth, error) {
	ret := _m.Called(ctx, tenant, apiID, runtimeID)

	var r0 *model.RuntimeAuth
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *model.RuntimeAuth); ok {
		r0 = rf(ctx, tenant, apiID, runtimeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RuntimeAuth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, tenant, apiID, runtimeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrDefault provides a mock function with given fields: ctx, tenant, apiID, runtimeID
func (_m *Repository) GetOrDefault(ctx context.Context, tenant string, apiID string, runtimeID string) (*model.RuntimeAuth, error) {
	ret := _m.Called(ctx, tenant, apiID, runtimeID)

	var r0 *model.RuntimeAuth
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *model.RuntimeAuth); ok {
		r0 = rf(ctx, tenant, apiID, runtimeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RuntimeAuth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, tenant, apiID, runtimeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListForAllRuntimes provides a mock function with given fields: ctx, tenant, apiID
func (_m *Repository) ListForAllRuntimes(ctx context.Context, tenant string, apiID string) ([]model.RuntimeAuth, error) {
	ret := _m.Called(ctx, tenant, apiID)

	var r0 []model.RuntimeAuth
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []model.RuntimeAuth); ok {
		r0 = rf(ctx, tenant, apiID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.RuntimeAuth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, tenant, apiID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Upsert provides a mock function with given fields: ctx, item
func (_m *Repository) Upsert(ctx context.Context, item model.RuntimeAuth) error {
	ret := _m.Called(ctx, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.RuntimeAuth) error); ok {
		r0 = rf(ctx, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}