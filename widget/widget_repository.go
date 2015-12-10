package widget

import (
	"github.com/dakiva/dbx"
	"github.com/dakiva/widget-sample/domain"
)

// abstraction that manages widget data
type WidgetRepository interface {
	// finds and returns all widgets. Returns an empty array if no widgets could be found.
	FindAll() ([]*domain.Widget, error)
	// finds and returns a widget for the given for the given uuid. Returns nil if a widget could not be found.
	FindWidget(widgetUuid string) (*domain.Widget, error)
	// creates a new widget, returning its internal id.
	CreateWidget(widget *domain.Widget) (int64, error)
	// updates a widget
	UpdateWidget(widget *domain.Widget) error
}

// implementation that is backed by a database
type dbBackedWidgetRepository struct {
	ctx      dbx.DBContext
	queryMap dbx.QueryMap
}

// Construct a WidgetRepository
func NewWidgetRepository(ctx dbx.DBContext, queryMap dbx.QueryMap) WidgetRepository {
	return &dbBackedWidgetRepository{ctx: ctx, queryMap: queryMap}
}

func (this *dbBackedWidgetRepository) FindAll() ([]*domain.Widget, error) {
	rows, err := this.ctx.NamedQuery(this.queryMap.Q("FindAllWidgets"), map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	widgets := make([]*domain.Widget, 0)
	defer rows.Close()
	for rows.Next() {
		widget := &domain.Widget{}
		err = rows.StructScan(widget)
		if err != nil {
			return nil, err
		}
		widgets = append(widgets, widget)
	}
	return widgets, nil
}

func (this *dbBackedWidgetRepository) FindWidget(widgetUuid string) (*domain.Widget, error) {
	rows, err := this.ctx.NamedQuery(this.queryMap.Q("FindWidget"), map[string]interface{}{"widget_uuid": widgetUuid})
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		widget := &domain.Widget{}
		err = rows.StructScan(widget)
		if err != nil {
			return nil, err
		}
		return widget, nil
	}
	return nil, nil
}

func (this *dbBackedWidgetRepository) CreateWidget(widget *domain.Widget) (int64, error) {
	rows, err := this.ctx.NamedQuery(this.queryMap.Q("InsertWidget"), widget)
	if err != nil {
		return -1, err
	}
	id, err := GetNextId(rows)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (this *dbBackedWidgetRepository) UpdateWidget(widget *domain.Widget) error {
	_, err := this.ctx.NamedExec(this.queryMap.Q("UpdateWidget"), widget)
	if err != nil {
		return err
	}
	return nil
}
