package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nomikuh8/bossabox-challenge/src/db"
	"github.com/nomikuh8/bossabox-challenge/src/models"
)

func NewToolsController() ToolsController {
	return &toolsController{}
}

type ToolsController interface {
	GetAllTools(*gin.Context)
	GetOneTool(*gin.Context)
	InsertNewTool(*gin.Context)
}

type toolsController struct{}

func (tc *toolsController) GetAllTools(c *gin.Context) {
	db, err := db.GetDatabase()
	tc.checkError(err, c, "Error while trying to connect to database")
	defer db.Close()

	toolModel := models.NewToolsModel(db)

	toolRows, err := toolModel.GetAllTools()
	tc.checkError(err, c, "Error while to get all tools")

	tagRows, err := toolModel.GetAllToolTags()
	tc.checkError(err, c, "Error while to get all tool tags")

	tools := []models.Tool{}
	for toolRows.Next() {
		var tool models.Tool

		err := toolRows.Scan(&tool.Id, &tool.Title, &tool.Description, &tool.Link)
		tc.checkError(err, c, "Error while trying to scan tool row")

		tools = append(tools, tool)
	}

	var tags []models.ToolTag
	for tagRows.Next() {
		var toolTag models.ToolTag

		err := tagRows.Scan(&toolTag.Id, &toolTag.ToolId, &toolTag.TagName)
		tc.checkError(err, c, "Error while trying to scan tool tag row")

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

func (tc *toolsController) GetOneTool(c *gin.Context) {
	db, err := db.GetDatabase()
	tc.checkError(err, c, "Error while trying to connect to database")

	id, err := strconv.Atoi(c.Param("id"))
	tc.checkError(err, c, "Error while trying to parse id")

	toolsModel := models.NewToolsModel(db)
	toolRow, err := toolsModel.GetOneTool(id)
	tc.checkError(err, c, "Error while trying to get tool row")

	tool := models.Tool{}
	if toolRow.Next() {
		err = toolRow.Scan(&tool.Id, &tool.Title, &tool.Description, &tool.Link)
		tc.checkError(err, c, "Error while setting toolRow")
	} else {
		c.JSON(400, gin.H{
			"message": "Tool doesn't exist",
		})
		return
	}

	tags := []string{}
	tagRows, err := toolsModel.GetOneToolTags(id)
	tc.checkError(err, c, "Error while getting tags")

	for tagRows.Next() {
		tag := models.ToolTag{}
		err := tagRows.Scan(&tag.Id, &tag.ToolId, &tag.TagName)
		tc.checkError(err, c, "Error while assigning tagRows to tag")

		tags = append(tags, tag.TagName)
	}

	tool.Tags = tags
	c.JSON(200, tool)
}

func (tc *toolsController) InsertNewTool(c *gin.Context) {
	var tool models.Tool

	err := c.BindJSON(&tool)
	tc.checkError(err, c, "Body invalid")

	db, err := db.GetDatabase()
	tc.checkError(err, c, "Couldn't connect to database")

	toolsModel := models.NewToolsModel(db)
	_, err = toolsModel.InsertTool(tool)
	tc.checkError(err, c, "Couldn't insert new tool")

	tools, err := toolsModel.GetLastTool()
	tc.checkError(err, c, "Error trying to get last tool")

	if tools.Next() {
		err := tools.Scan(&tool.Id, &tool.Title, &tool.Link, &tool.Description)
		tc.checkError(err, c, "Error binding last tool")
	} else {
		c.JSON(500, gin.H{
			"message": "Tool wasn't inserted",
		})
		return
	}

	for _, tag := range tool.Tags {
		_, err := toolsModel.InsertToolTag(tool.Id, tag)
		tc.checkError(err, c, "Couldn't insert some tags")
	}

	c.JSON(201, tool)
}

func (tc *toolsController) checkError(err error, c *gin.Context, message string) {
	if err != nil {
		c.JSON(500, gin.H{
			"message": message,
		})
		log.Panic(err)
	}
}
