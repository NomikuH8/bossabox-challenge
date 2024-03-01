package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nomikuh8/bossabox-challenge/src/db"
	"github.com/nomikuh8/bossabox-challenge/src/models"
)

func NewToolsController() ToolsController {
	return &toolsController{}
}

type ToolsController interface {
	GetAllTools(*gin.Context)
}

type toolsController struct{}

func (tc *toolsController) GetAllTools(c *gin.Context) {
	db, err := db.GetDatabase()
	tc.checkError(err, c)
	defer db.Close()

	toolModel := models.NewToolsModel(db)

	toolRows, err := toolModel.GetAllTools()
	tc.checkError(err, c)

	tagRows, err := toolModel.GetAllToolTags()
	tc.checkError(err, c)

	tools := []models.Tool{}
	for toolRows.Next() {
		var tool models.Tool

		err := toolRows.Scan(&tool.Id, &tool.Title, &tool.Description, &tool.Link)
		tc.checkError(err, c)

		tools = append(tools, tool)
	}

	var tags []models.ToolTag
	for tagRows.Next() {
		var toolTag models.ToolTag

		err := tagRows.Scan(&toolTag.Id, &toolTag.TagName, &toolTag.ToolId)
		tc.checkError(err, c)

		tags = append(tags, toolTag)
	}

	for _, t := range tools {
		for _, tag := range tags {
			if tag.ToolId == t.Id {
				t.Tags = append(t.Tags, tag.TagName)
			}
		}
	}

	c.JSON(200, tools)
}

func (tc *toolsController) checkError(err error, c *gin.Context) {
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Something went wrong while taking all tools",
		})
		log.Panic(err)
	}
}
