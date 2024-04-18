package controllers

import (
	"encoding/json"
	"go-olin/models"
	"go-olin/pkg/response"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type ObjectController struct {
	beego.Controller
}

// Post @Title Create
// @Description create object
// @Param	requestBody body models.RequestSoal1Model	true	"Request body in JSON format"
// @Success 200 interfaces{}
// @Failure 400 bodies are empty
// @router /soal1 [post]
func (o *ObjectController) Post() {
	appB := response.Bee{
		Ctx: o.Ctx,
	}
	var ob models.RequestSoal1Model
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := models.Soal1Model(ob)
	if objectid.Result == nil {
		appB.Response(http.StatusBadRequest, "Sukses", objectid.Time, nil)
		return
	}
	appB.Response(http.StatusOK, "Sukses", objectid.Time, objectid.Result)
	return
}

// Post2 @Title Create
// @Description create object
// @Param	requestBody body models.RequestSoal2Model	true	"The object content"
// @Success 200 interfaces{}
// @Failure 400 bodies are empty
// @router /soal2 [post]
func (o *ObjectController) Post2() {
	appB := response.Bee{
		Ctx: o.Ctx,
	}
	var ob models.RequestSoal2Model
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := models.Soal2Model(ob)
	if objectid.Result == nil {
		appB.Response(http.StatusBadRequest, "Sukses", objectid.Time, nil)
		return
	}
	appB.Response(http.StatusOK, "Sukses", objectid.Time, objectid.Result)
	return
}

// Post3 @Title Create
// @Description create object
// @Param	requestBody body models.RequestSoal3Model	"The object content"
// @Success 200 interfaces{}
// @Failure 400 bodies are empty
// @router /soal3 [post]
func (o *ObjectController) Post3() {
	appB := response.Bee{
		Ctx: o.Ctx,
	}
	var ob models.RequestSoal3Model
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := models.Soal3Model(ob)
	if objectid.Result == nil {
		appB.Response(http.StatusBadRequest, "Sukses", objectid.Time, nil)
		return
	}
	appB.Response(http.StatusOK, "Sukses", objectid.Time, objectid.Result)
	return
}
