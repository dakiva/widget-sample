package widget

import (
	"testing"

	"github.com/dakiva/widget-sample/domain"
	"github.com/stretchr/testify/assert"
)

func TestWidgetCreation(t *testing.T) {
	// given
	tx, _ := testdb.Beginx()
	defer tx.Rollback()
	repo := NewWidgetRepository(tx, testQueryMap)

	widget := &domain.Widget{
		Uuid: "516f4464-041e-446e-9ab3-d2f0f7ff0d34",
		Name: "A Widget",
	}

	// when
	widgetId, err := repo.CreateWidget(widget)

	// then
	assert.NoError(t, err)
	assert.True(t, widgetId >= 0)

	found, err := repo.FindWidget(widget.Uuid)
	assert.NoError(t, err)
	assert.Equal(t, widgetId, found.Id)
	assert.Equal(t, widget.Uuid, found.Uuid)
	assert.Equal(t, widget.Name, found.Name)
}

func TestWidgetUpdate(t *testing.T) {
	// given
	tx, _ := testdb.Beginx()
	defer tx.Rollback()
	repo := NewWidgetRepository(tx, testQueryMap)

	widget := &domain.Widget{
		Uuid: "516f4464-041e-446e-9ab3-d2f0f7ff0d34",
		Name: "A Widget",
	}
	_, err := repo.CreateWidget(widget)
	assert.NoError(t, err)

	// when
	widget.Name = "Updated Name"
	err = repo.UpdateWidget(widget)

	// then
	assert.NoError(t, err)
	found, err := repo.FindWidget(widget.Uuid)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", found.Name)
}

func TestFindAll(t *testing.T) {
	// given
	tx, _ := testdb.Beginx()
	defer tx.Rollback()
	repo := NewWidgetRepository(tx, testQueryMap)

	widget := &domain.Widget{
		Uuid: "516f4464-041e-446e-9ab3-d2f0f7ff0d34",
		Name: "A Widget",
	}

	// when
	widgetId, err := repo.CreateWidget(widget)

	// then
	assert.NoError(t, err)
	assert.True(t, widgetId >= 0)

	foundAll, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, foundAll, 1)
	assert.Equal(t, widgetId, foundAll[0].Id)
	assert.Equal(t, widget.Uuid, foundAll[0].Uuid)
	assert.Equal(t, widget.Name, foundAll[0].Name)
}
