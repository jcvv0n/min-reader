package service

import (
	"fmt"
	"html/template"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type StoryContent struct {
	PageNo   uint64
	PageDesc string
	Content  string
}

type PageInfo struct {
	MinPage uint64
	MaxPage uint64
}

func Reader(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	storyIdStr := ctx.Param("storyId")
	pageStr := ctx.Query("p")
	storyId, err1 := strconv.ParseUint(storyIdStr, 10, 64)
	page, err2 := strconv.ParseUint(pageStr, 10, 64)
	if err1 != nil || err2 != nil {
		ctx.JSON(404, nil)
		return
	}
	si := GetStoryInfo(namespace, storyId)
	if si == nil {
		ctx.JSON(404, nil)
		return
	}
	sc, pg := getStroyContent(si, page)
	if sc == nil {
		ctx.JSON(404, nil)
		return
	}
	ctx.HTML(200, "reader.tmpl", gin.H{
		"title":    si.StoryName,
		"desc":     sc.PageDesc,
		"content":  template.HTML(sc.Content),
		"curPage":  page,
		"prePage":  prePage(page, pg.MinPage),
		"nextPage": nextPage(page, pg.MaxPage),
	})
}

func getStroyContent(si *StoryInfo, page uint64) (*StoryContent, *PageInfo) {
	db, err := gorm.Open(sqlite.Open(si.DbPath), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("数据库链接失败")
		return nil, nil
	}

	var sc StoryContent
	result := db.Where("page_no = ?", page).Take(&sc)
	if result.RowsAffected > 0 {
		var pageInfo PageInfo
		pageRes := db.Raw("SELECT min(page_no) min_page,max(page_no) max_page FROM story_content").Scan(&pageInfo)
		if pageRes.RowsAffected > 0 {
			return &sc, &pageInfo
		}
	}
	return nil, nil
}

func prePage(page, minPage uint64) uint64 {
	if page > 1 && page > minPage {
		return page - 1
	}
	return 0
}

func nextPage(page, maxPage uint64) uint64 {
	if page < maxPage {
		return page + 1
	}
	return 0
}
