package controllers

import (
	gvalid "ThingsPanel-Go/initialize/validate"
	"ThingsPanel-Go/services"
	response "ThingsPanel-Go/utils"
	valid "ThingsPanel-Go/validate"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	context2 "github.com/beego/beego/v2/server/web/context"
)

type HomeController struct {
	beego.Controller
}

type HomeList struct {
	CpuUsage string `json:"cpu_usage"`
	MemUsage string `json:"mem_usage"`
	Device   int64  `json:"device"`
	Msg      int64  `json:"msg"`
}

type HomeDevice struct {
	Business   int64 `json:"business"`
	Assets     int64 `json:"assets"`
	Equipment  int64 `json:"equipment"`
	Dashboard  int64 `json:"dashboard"`
	Conditions int64 `json:"conditions"`
}

// 首页数据统计
func (this *HomeController) List() {
	var ResourcesService services.ResourcesService
	r := ResourcesService.GetNew()
	var DeviceService services.DeviceService
	_, dc := DeviceService.All()
	var TSKVService services.TSKVService
	_, tc := TSKVService.All()
	u := HomeList{
		CpuUsage: r.CPU,
		MemUsage: r.MEM,
		Device:   dc,
		Msg:      tc,
	}
	response.SuccessWithDetailed(200, "success", u, map[string]string{}, (*context2.Context)(this.Ctx))
	return
}

// 首页报表 chart
func (this *HomeController) Chart() {
	var ResourcesService services.ResourcesService
	nr := ResourcesService.GetNewResource("cpu")
	response.SuccessWithDetailed(200, "success", nr, map[string]string{}, (*context2.Context)(this.Ctx))
	return
}

// 首页展示设备 show
func (this *HomeController) Show() {
	homeShowValidate := valid.HomeShowValidate{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &homeShowValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(homeShowValidate)
	if !status {
		for _, err := range v.Errors {
			// 获取字段别称
			alias := gvalid.GetAlias(homeShowValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var DeviceService services.DeviceService
	d, _ := DeviceService.GetDeviceByID(homeShowValidate.ID)
	response.SuccessWithDetailed(200, "success", d, map[string]string{}, (*context2.Context)(this.Ctx))
	return
}

// Device
func (this *HomeController) Device() {
	var BusinessService services.BusinessService
	_, bc := BusinessService.All()
	var AssetService services.AssetService
	_, ac := AssetService.All()
	var DeviceService services.DeviceService
	_, dc := DeviceService.All()
	var DashBoardService services.DashBoardService
	_, dac := DashBoardService.All()
	var ConditionsService services.ConditionsService
	_, cc := ConditionsService.All()
	d := HomeDevice{
		Business:   bc,
		Assets:     ac,
		Equipment:  dc,
		Dashboard:  dac,
		Conditions: cc,
	}
	response.SuccessWithDetailed(200, "success", d, map[string]string{}, (*context2.Context)(this.Ctx))
	return
}
