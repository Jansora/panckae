package routers

import (
	pg "github.com/Jansora/pancake/backend/postgres"
	"github.com/Jansora/pancake/backend/postgres/article"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func InitArticle(r *gin.Engine) {
	Article(r)
	GetTagList(r)
}

func Article(r *gin.Engine) {
	r.GET("/Golang/Article/:Url", func(c *gin.Context) {

		if a, err := article.Select(pg.Client, c.Param("Url"), !ValidateLoginStatus(c)); err == nil {
			c.JSON(http.StatusOK, gin.H{"ret": true, "res": a})
		} else {
			c.JSON(http.StatusOK, gin.H{"ret": false, "res": "获取Article信息失败！" + err.Error()})
		}
	})

	r.POST("/Golang/Article/Insert", func(c *gin.Context) {
		var j InsertArticleType

		if err := c.BindJSON(&j); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"ret": false,
				"res": "Decode json error！" + err.Error(),
			})
			return
		}
		if !ValidateLoginStatus(c) {
			return
		}
		a := article.Article{
			Author:      j.Author,
			Create_time: time.Now(),
			Modify_time: time.Now(),
			Site:        j.Site,
			Read_num:    0,
			Like_num:    0,
			Tags:        j.Tags,
			Url:         j.Url,
			Is_public:   j.IsPublic,
			Logo_url:    j.LogoUrl,
			Title:       j.Title,
			Summary:     j.Summary,
			Content:     j.Content,
			Html:        j.Html,
		}
		if err := article.Insert(pg.Client, a); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"ret": false, "res": "参数不能为空！" + err.Error(),
			})
			return
		}
		println(a.Title)
		c.JSON(http.StatusOK, gin.H{
			"ret": true, "res": "文章发表成功！",
		})
	})
	r.POST("/Golang/Article/UpdateDoc/:Url", func(c *gin.Context) {
		var j InsertArticleType

		if err := c.BindJSON(&j); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"ret": false,
				"res": "Decode json error！" + err.Error(),
			})
			return
		}
		if !ValidateLoginStatus(c) {
			c.JSON(http.StatusOK, gin.H{"ret": false, "res": "没有操作权限"})
			return
		}
		if a, err := article.Select(pg.Client, c.Param("Url"), false); err == nil {
			a.Content = j.Content
			if err := article.Update(pg.Client, a, c.Param("Url")); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"ret": false, "res": "更新失敗！" + err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"ret": true, "res": "文章正文更新成功！",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{"ret": false, "res": "获取Article信息失败！" + err.Error()})
		}
		return

	})

	r.POST("/Golang/Article/Update/:Url", func(c *gin.Context) {
		var j InsertArticleType
		if err := c.BindJSON(&j); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"ret": false,
				"res": "Decode json error！" + err.Error(),
			})
			return
		}
		if !ValidateLoginStatus(c) {
			c.JSON(http.StatusOK, gin.H{"ret": false, "res": "没有操作权限"})
			return
		}

		if a, err := article.Select(pg.Client, c.Param("Url"), false); err == nil {
			a.Title = j.Title
			a.Url = j.Url
			a.Author = j.Author
			a.Modify_time = time.Now()
			a.Site = j.Site
			a.Tags = j.Tags
			a.Is_public = j.IsPublic
			a.Logo_url = j.LogoUrl
			a.Summary = j.Summary
			a.Content = j.Content
			a.Html = j.Html
			if err := article.Update(pg.Client, a, c.Param("Url")); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"ret": false, "res": "更新失敗！" + err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"ret": true, "res": "文章更新成功！",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{"ret": false, "res": "获取Article信息失败！" + err.Error()})
		}
		return

	})

	r.POST("/Golang/Article/UpdateLike/:Url", func(c *gin.Context) {
		if a, err := article.Select(pg.Client, c.Param("Url"), !ValidateLoginStatus(c)); err == nil {
			a.Like_num += 1
			if err := article.Update(pg.Client, a, c.Param("Url")); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"ret": false, "res": "更新失敗！" + err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"ret": true, "res": "Like更新成功！",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{"ret": false, "res": "获取Article信息失败！" + err.Error()})
		}
		return

	})

	r.DELETE("/Golang/Article/:Url", func(c *gin.Context) {
		if !ValidateLoginStatus(c) {
			c.JSON(http.StatusOK, gin.H{"ret": false, "res": "没有操作权限"})
			return
		}
		if _, err := article.Delete(pg.Client, c.Param("Url")); err == nil {
			c.JSON(http.StatusOK, gin.H{"ret": true, "res": ""})
		} else {
			c.JSON(http.StatusOK, gin.H{"ret": false, "res": "获取Article信息失败！" + err.Error()})
		}

	})

	r.GET("/Golang/Article", func(c *gin.Context) {

		var con article.Condition
		con.Init(c)
		if as, err := article.Selects(pg.Client, con, !ValidateLoginStatus(c)); err == nil {
			length, _ := article.SelectsLength(pg.Client, con, !ValidateLoginStatus(c))
			c.JSON(http.StatusOK, gin.H{"ret": true, "res": as, "total": length})
		} else {
			Return(c, false, "get Article failed！"+err.Error())
		}
	})

	r.GET("/Golang/ArticleLength", func(c *gin.Context) {

		var con article.Condition
		con.Init(c)
		if length, err := article.SelectsLength(pg.Client, con, !ValidateLoginStatus(c)); err == nil {
			Return(c, true, length)
		} else {
			Return(c, false, "get Article length failed！"+err.Error())
		}
	})

	r.GET("/Golang/ArticleByIds", func(c *gin.Context) {

		arrStr := strings.Split(c.DefaultQuery("ids", ""), ",")
		if arrStr[0] == "" {
			Return(c, false, "ArticleByIds is null！")
			return
		}

		if as, err := article.SelectsByIds(pg.Client, arrStr, !ValidateLoginStatus(c)); err == nil {
			Return(c, true, as)
		} else {
			Return(c, false, "get Article failed！"+err.Error())
		}
	})

}

func GetTagList(r *gin.Engine) {

	r.GET("/Golang/Tags", func(c *gin.Context) {

		if as, err := article.SelectTags(pg.Client, !ValidateLoginStatus(c)); err == nil {
			Return(c, true, as)
		} else {
			Return(c, false, "获取 tag 失败！"+err.Error())
		}

	})
}
