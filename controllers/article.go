package controllers

import (
	"golang_cms/models"
	"fmt"
	"strconv"
	"strings"
	"time"
	"os"
	"image/jpeg"
	"io"
	"github.com/nfnt/resize"
)

type ArticleController struct {
	BaseController
}

var article = models.Article{}

func (this *ArticleController)  GetArticleById() {
	id, err := this.GetInt64("id")
	articleDetail := article.GetArticleById(id)
	if (err == nil) {
		this.Data["articleDetail"] = articleDetail
	}
	this.setTplName("articleDetail.html")
}

func (this *ArticleController)  DeleteArticleById() {
	id, _ := this.GetInt64("id")
	article.DeleteArticleById(id)
	this.Redirect("/admin/article/pagelist",302)
}

func (this *ArticleController)  AddArticle() {
	articleDetail := models.Article{}
	//调用 Controller ParseForm 这个方法的时候，传入的参数必须为一个 struct 的指针，否则对 struct 的赋值不会成功并返回 xx must be a struct pointer 的错误。
	if err := this.ParseForm(&articleDetail); err != nil {
		this.Data["showMessage"] = "操作失败！"
	} else {
		_, err := article.AddArticle(articleDetail)

		if (err == nil) {
			this.Data["articleDetail"] = articleDetail
			this.Data["showMessage"] = "操作成功！"
		} else {
			fmt.Println("操作失败")
			this.Data["showMessage"] = "操作失败！"
		}
	}

	this.setTplName("admin_addSinglePage.html")
}

func (this *ArticleController) GetArticleList() {
	singlePageList := article.GetArticleList();
	this.Data["singlePageList"] = singlePageList;
	this.setTplName("admin_singlePageList.html")
}

func (this *ArticleController)  LinkToAddArticle() {
	this.setTplName("admin_addSinglePage.html")
}

func (this *ArticleController)  LinkToEditArticle() {
	id, err := this.GetInt64("id")
	articleDetail := article.GetArticleById(id)
	if (err == nil) {
		this.Data["articleDetail"] = articleDetail
		//fmt.Println("typeId=" + strconv.FormatInt(articleDetail.Typeid, 10))
		//this.Data["typeId"] = strconv.FormatInt(articleDetail.Typeid,10)
		//this.Data["typeId"] = articleDetail.Typeid
	} else {
		fmt.Println("caozuoshibai")
	}
	this.setTplName("admin_editSinglePage.html")
}

func (this *ArticleController)  UpdateArticle() {
	articleDetail := models.Article{}
	//调用 Controller ParseForm 这个方法的时候，传入的参数必须为一个 struct 的指针，否则对 struct 的赋值不会成功并返回 xx must be a struct pointer 的错误。
	if err := this.ParseForm(&articleDetail); err != nil {
		this.Data["showMessage"] = "操作失败！"
	} else {
		_, err := article.UpdateArticle(articleDetail)

		if (err == nil) {
			this.Data["articleDetail"] = &articleDetail
			this.Data["showMessage"] = "操作成功！"
		} else {
			fmt.Println("操作失败")
			this.Data["showMessage"] = "操作失败！"
		}
	}
	this.setTplName("admin_editSinglePage.html")
}

func (this *ArticleController) Upload() {
	file, header, err := this.GetFile("upfile")
	utype := this.GetString("type")
	if utype == "" {
		utype = "1"
	}
	index, _ := strconv.Atoi(utype)

	ext := strings.ToLower(header.Filename[strings.LastIndex(header.Filename, "."):])
	out := make(map[string]string)
	out["url"] = ""
	out["fileType"] = ext
	out["original"] = header.Filename
	out["state"] = "SUCCESS"
	if err != nil {
		out["state"] = err.Error()
	} else {
		savepath := pathArr[index] + time.Now().Format("20060102")
		if err = os.MkdirAll(savepath, os.ModePerm); err != nil {
			out["state"] = err.Error()
		} else {
			filename := fmt.Sprintf("%s/%d%s", savepath, time.Now().UnixNano(), ext)
			if this.GetString("type") == "2" {
				w, _ := strconv.Atoi(this.GetString("w"))
				h, _ := strconv.Atoi(this.GetString("h"))
				err = createSmallPic(file, filename, w, h)
				if err != nil {
					out["state"] = err.Error()
				}
			} else {
				if err = this.SaveToFile("upfile", filename); err != nil {
					out["state"] = err.Error()
				}
			}
			out["url"] = filename[1:]
		}
	}
	this.Data["json"] = out
	this.ServeJSON()
}

func createSmallPic(file io.Reader, fileSmall string, w, h int) error {
	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}
	b := img.Bounds()
	if w > b.Dx() {
		w = b.Dx()
	}
	if h > b.Dy() {
		h = b.Dy()
	}
	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

	out, err := os.Create(fileSmall)
	if err != nil {
		return err
	}
	defer out.Close()

	// write new image to file
	return jpeg.Encode(out, m, nil)
}