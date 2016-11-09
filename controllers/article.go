package controllers

import (
	"blog/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about article
type ArticleController struct {
	beego.Controller
}

// @Title Create
// @Description create article
// @Param	body		body 	models.Article	true		"The article content"
// @Success 200 {string} models.Article.Id
// @Failure 403 body is empty
// @router / [post]
func (o *ArticleController) Post() {
	var ob models.Article
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if ob.Title == "" || ob.Content == "" {
		o.Ctx.Output.SetStatus(403)
		return
	}
	articleid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"ArticleId": articleid}
	o.ServeJSON()
}

// @Title Get
// @Description find article by articleid
// @Param	articleId		path 	string	true		"the articleid you want to get"
// @Success 200 {article} models.Article
// @Failure 403 :articleId is empty
// @router /:articleId [get]
func (o *ArticleController) Get() {
	articleId := o.Ctx.Input.Param(":articleId")
	if articleId != "" {
		ob, err := models.GetOne(articleId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all articles
// @Success 200 {article} models.Article
// @Failure 403 :articleId is empty
// @router / [get]
func (o *ArticleController) GetAll() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the article
// @Param	articleId		path 	string	true		"The articleid you want to update"
// @Param	body		body 	models.Article	true		"The body"
// @Success 200 {article} models.Article
// @Failure 403 :articleId is empty
// @router /:articleId [put]
func (o *ArticleController) Put() {
	articleId := o.Ctx.Input.Param(":articleId")
	var ob models.Article
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Update(articleId, ob.Title, ob.Content)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the article
// @Param	articleId		path 	string	true		"The articleId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 articleId is empty
// @router /:articleId [delete]
func (o *ArticleController) Delete() {
	articleId := o.Ctx.Input.Param(":articleId")
	models.Delete(articleId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}
