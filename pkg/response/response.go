package response

import (
	"time"

	"github.com/beego/beego/v2/server/web/context"
)

type Bee struct {
	Ctx *context.Context
}
type DataResponse struct {
	Responsemessage   string        `json:"message"`
	Responsetimestamp time.Time     `json:"time"`
	RunTime           time.Duration `json:"runTime"`
	Data              interface{}   `json:"data"`
}

func (g *Bee) Response(httpCode int, message string, duration time.Duration, data interface{}) {
	g.Ctx.Output.SetStatus(httpCode)
	g.Ctx.Output.JSON(DataResponse{
		Responsemessage:   message,
		Responsetimestamp: time.Now(),
		Data:              data,
		RunTime:           duration,
	}, true, true)
	return

}
