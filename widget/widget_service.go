package widget

import (
	"time"

	"github.com/dakiva/widget-sample/domain"
	"github.com/pborman/uuid"
)

type WidgetService struct {
}

// Finds and returns all widgets
func (this *WidgetService) FindWidgets() ([]*domain.Widget, error) {
	widgetRepo := NewWidgetRepository(GetDB(), GetQueryMap())
	return this.findWidgets(widgetRepo)
}

// Creates a new widget
func (this *WidgetService) CreateWidget(widgetName string) (*domain.Widget, error) {
	tx, err := GetDB().Beginx()
	if err != nil {
		return nil, err
	}
	var serviceErr error
	defer func() {
		if serviceErr != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	widgetRepo := NewWidgetRepository(tx, GetQueryMap())
	return this.createWidget(widgetName, widgetRepo)
}

func (this *WidgetService) findWidgets(widgetRepo WidgetRepository) ([]*domain.Widget, error) {
	widgets, err := widgetRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return widgets, nil
}

func (this *WidgetService) createWidget(widgetName string, widgetRepo WidgetRepository) (*domain.Widget, error) {
	widget := &domain.Widget{
		Uuid: uuid.New(),
		Name: widgetName,
	}
	widgetId, err := widgetRepo.CreateWidget(widget)
	if err != nil {
		return nil, err
	}
	widget.Id = widgetId
	now := time.Now()
	widget.CreatedOn = now
	widget.ModifiedOn = now
	return widget, nil
}
