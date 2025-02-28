package service

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type StoryCatalog struct {
	// StoryName string
	PageNo   uint64
	PageDesc string
}

func Catalog(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	storyIdStr := ctx.Param("storyId")
	storyId, err1 := strconv.ParseUint(storyIdStr, 10, 64)
	if err1 != nil {
		ctx.JSON(404, nil)
		return
	}
	si := GetStoryInfo(namespace, storyId)
	if si == nil {
		ctx.JSON(404, nil)
		return
	}
	ct := getStroyCatalog(si, storyId)
	if ct == nil {
		ctx.JSON(404, nil)
		return
	}
	ctx.HTML(200, "catalog.tmpl", gin.H{
		"storyId":  storyId,
		"stNames":  si.StoryName,
		"storyCat": ct,
	})
}

func getStroyCatalog(si *StoryInfo, storyId uint64) *[]StoryCatalog {
	db, err := gorm.Open(sqlite.Open(si.DbPath), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("数据库链接失败")
		return nil
	}

	var sc []StoryCatalog
	result := db.Table("story_content").Select("page_no, page_desc").Find(&sc)
	if result.RowsAffected > 0 {
		return &sc
	}
	return nil
}

// func groupByName(ct *[]StoryCatalog) (*[]string, *map[string][]StoryCatalog) {
// 	catMap := make(map[string][]StoryCatalog)
// 	var stNames []string
// 	for _, c := range *ct {
// 		mval, ok := catMap[c.StoryName]
// 		if ok {
// 			mval = append(mval, c)
// 		} else {
// 			stNames = append(stNames, c.StoryName)
// 			var newC []StoryCatalog
// 			mval = append(newC, c)
// 		}
// 		catMap[c.StoryName] = mval
// 	}
// 	return &stNames, &catMap
// }

// func GetStroyCat(m *map[string][]StoryCatalog, key string) (val []StoryCatalog) {
// 	return (*m)[key]
// }
