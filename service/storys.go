package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type StoryOverview struct {
	StoryId   uint64
	StoryName string
}

type StoryInfo struct {
	StoryName string
	DbPath    string
}

func Storys(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	sov := getStroyList(&namespace)
	if sov == nil {
		ctx.JSON(404, nil)
		return
	}
	ctx.HTML(200, "storys.tmpl", gin.H{
		"storys": sov,
	})
}

func GetStoryInfo(namespace string, storyId uint64) *StoryInfo {
	db, err := gorm.Open(sqlite.Open("db/story_overview"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("数据库链接失败")
		return nil
	}
	var si StoryInfo
	result := db.Table("story_overview").Select("story_name, db_path").Where("namespace = ? AND story_id = ?", namespace, storyId).Take(&si)
	if result.RowsAffected > 0 {
		return &si
	}
	return nil
}

func getStroyList(namespace *string) *[]StoryOverview {
	db, err := gorm.Open(sqlite.Open("db/story_overview"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("数据库链接失败")
		return nil
	}

	var sov []StoryOverview
	result := db.Where("namespace = ?", *namespace).Select("story_id, story_name").Find(&sov)
	if result.RowsAffected > 0 {
		return &sov
	}
	return nil
}
