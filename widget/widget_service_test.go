package widget

import (
	"errors"
	"testing"

	"github.com/dakiva/widget-sample/widget/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateWidget(t *testing.T) {
	// given
	mockWidgetRepo := new(mocks.WidgetRepository)
	mockWidgetRepo.On("CreateWidget", mock.AnythingOfType("*domain.Widget")).Return(int64(1), nil)

	// when
	createdWidget, err := new(WidgetService).createWidget("A Widget", mockWidgetRepo)

	// then
	assert.NoError(t, err)
	assert.Equal(t, "A Widget", createdWidget.Name)
}

func TestFailedCreateWidget(t *testing.T) {
	// given
	mockWidgetRepo := new(mocks.WidgetRepository)
	mockWidgetRepo.On("CreateWidget", mock.AnythingOfType("*domain.Widget")).Return(int64(0), errors.New("Some Error"))

	// when
	createdWidget, err := new(WidgetService).createWidget("A Widget", mockWidgetRepo)

	// then
	assert.Error(t, err)
	assert.Nil(t, createdWidget)
}
