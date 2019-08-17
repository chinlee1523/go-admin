package controller

import (
	"github.com/chinlee1523/go-admin/context"
	"github.com/chinlee1523/go-admin/modules/auth"
	"github.com/chinlee1523/go-admin/modules/menu"
	"github.com/chinlee1523/go-admin/plugins/admin/models"
	"github.com/chinlee1523/go-admin/plugins/admin/modules/file"
	"github.com/chinlee1523/go-admin/template"
	"github.com/chinlee1523/go-admin/template/types"
	"net/http"
	"strings"
)

// 显示新建表单
func ShowNewForm(ctx *context.Context) {
	prefix := ctx.Query("prefix")
	panel := models.TableList[prefix]
	if !panel.GetCanAdd() {
		ctx.Html(http.StatusNotFound, "page not found")
		return
	}
	params := models.GetParam(ctx.Request.URL.Query())

	user := auth.Auth(ctx)

	formList := models.GetNewFormList(panel.GetForm().FormList)
	for i := 0; i < len(formList); i++ {
		formList[i].Editable = true
	}
	tmpl, tmplName := template.Get(Config.THEME).GetTemplate(ctx.Headers("X-PJAX") == "true")
	buf := template.Excecute(tmpl, tmplName, user, types.Panel{
		Content: template.Get(Config.THEME).Form().
			SetPrefix(Config.PREFIX).
			SetContent(formList).
			SetUrl(Config.PREFIX + "/new/" + prefix).
			SetToken(auth.TokenHelper.AddToken()).
			SetTitle("New").
			SetInfoUrl(Config.PREFIX + "/info/" + prefix + params.GetRouteParamStr()).
			SetHeader(panel.GetForm().HeaderHtml).
			SetFooter(panel.GetForm().FooterHtml).
			GetContent(),
		Description: panel.GetForm().Description,
		Title:       panel.GetForm().Title,
	}, Config, menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(ctx.Path(), Config.PREFIX, "", 1)))
	ctx.Html(http.StatusOK, buf.String())
}

// 新建数据
func NewForm(ctx *context.Context) {
	prefix := ctx.Query("prefix")
	panel := models.TableList[prefix]
	if !panel.GetCanAdd() {
		ctx.Html(http.StatusNotFound, "page not found")
		return
	}

	token := ctx.FormValue("_t")

	if !auth.TokenHelper.CheckToken(token) {
		ctx.Json(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  "新增失败",
		})
		return
	}

	form := ctx.Request.MultipartForm

	// 处理上传文件，目前仅仅支持传本地
	if len((*form).File) > 0 {
		_, _ = file.GetFileEngine("local").Upload(form)
	}

	if prefix == "manager" { // 管理员管理新建
		NewManager((*form).Value)
	} else if prefix == "roles" { // 管理员角色管理新建
		NewRole((*form).Value)
	} else {
		panel.InsertDataFromDatabase((*form).Value)
	}

	models.RefreshTableList()

	previous := ctx.FormValue("_previous_")
	prevUrlArr := strings.Split(previous, "?")
	params := models.GetParamFromUrl(previous)

	panelInfo := panel.GetDataFromDatabase(prevUrlArr[0], params)

	editUrl := Config.PREFIX + "/info/" + prefix + "/edit" + params.GetRouteParamStr()
	newUrl := Config.PREFIX + "/info/" + prefix + "/new" + params.GetRouteParamStr()
	deleteUrl := Config.PREFIX + "/delete/" + prefix

	dataTable := template.Get(Config.THEME).
		DataTable().
		SetInfoList(panelInfo.InfoList).
		SetThead(panelInfo.Thead).
		SetEditUrl(editUrl).
		SetNewUrl(newUrl).
		SetDeleteUrl(deleteUrl)

	table := dataTable.GetContent()

	box := template.Get(Config.THEME).Box().
		SetBody(table).
		SetHeader(dataTable.GetDataTableHeader() + panel.GetInfo().HeaderHtml).
		WithHeadBorder(false).
		SetFooter(panel.GetInfo().FooterHtml + panelInfo.Paginator.GetContent()).
		GetContent()

	user := auth.Auth(ctx)

	tmpl, tmplName := template.Get(Config.THEME).GetTemplate(true)
	buffer := template.Excecute(tmpl, tmplName, user, types.Panel{
		Content:     box,
		Description: panelInfo.Description,
		Title:       panelInfo.Title,
	}, Config, menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(previous, Config.PREFIX, "", 1)))

	ctx.Html(http.StatusOK, buffer.String())
	ctx.AddHeader("X-PJAX-URL", previous)
}
