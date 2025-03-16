package service

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type FullText struct {
	PageDesc string
	Content  string
}

func Fulltext(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	storyIdStr := ctx.Param("storyId")
	storyId, err := strconv.ParseUint(storyIdStr, 10, 64)
	if err != nil {
		ctx.JSON(404, nil)
		return
	}
	si := GetStoryInfo(namespace, storyId)
	if si == nil {
		ctx.JSON(404, nil)
		return
	}
	ft := getFullText(si)
	if ft == nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+si.StoryName+".txt")
	ctx.Header("Content-Transfer-Encoding", "binary")
	for _, f := range *ft {
		ctx.Writer.WriteString("\n\n" + si.StoryName + " " + f.PageDesc + "\n\n")
		ctx.Writer.WriteString(strings.ReplaceAll(strings.ReplaceAll(f.Content, "<p>", "    "), "</p>", "\n"))
		ctx.Writer.WriteString("\n")
	}
	ctx.Writer.Flush()
}

func getFullText(si *StoryInfo) *[]FullText {
	db, err := gorm.Open(sqlite.Open(si.DbPath), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("database link failed")
		return nil
	}

	var ft []FullText
	result := db.Table("story_content").Select("page_desc, content").Find(&ft)
	if result.RowsAffected > 0 {
		return &ft
	}
	return nil
}
