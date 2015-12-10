package widget

import (
	"net/http"

	"github.com/dakiva/widget-sample/domain"
	"github.com/emicklei/go-restful"
)

type WidgetController struct {
}

func (this *WidgetController) register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/widgets").
		Doc("Widget endpoint").
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("").To(this.findWidgets).
		Doc("Retrieves widget data.").
		Operation("findWidgets").
		Writes([]domain.Widget{}))

	container.Add(ws)
}

func (this *WidgetController) findWidgets(req *restful.Request, resp *restful.Response) {
	widgets, err := new(WidgetService).FindWidgets()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.WriteEntity(widgets)
}
