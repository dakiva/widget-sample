package mocks

import "github.com/stretchr/testify/mock"

import "github.com/dakiva/widget-sample/domain"

type WidgetRepository struct {
	mock.Mock
}

func (_m *WidgetRepository) FindAll() ([]*domain.Widget, error) {
	ret := _m.Called()

	var r0 []*domain.Widget
	if rf, ok := ret.Get(0).(func() []*domain.Widget); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Widget)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *WidgetRepository) FindWidget(widgetUuid string) (*domain.Widget, error) {
	ret := _m.Called(widgetUuid)

	var r0 *domain.Widget
	if rf, ok := ret.Get(0).(func(string) *domain.Widget); ok {
		r0 = rf(widgetUuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Widget)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(widgetUuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *WidgetRepository) CreateWidget(widget *domain.Widget) (int64, error) {
	ret := _m.Called(widget)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*domain.Widget) int64); ok {
		r0 = rf(widget)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Widget) error); ok {
		r1 = rf(widget)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *WidgetRepository) UpdateWidget(widget *domain.Widget) error {
	ret := _m.Called(widget)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Widget) error); ok {
		r0 = rf(widget)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
