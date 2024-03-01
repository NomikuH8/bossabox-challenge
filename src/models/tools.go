package models

import (
	"database/sql"
)

func NewToolsModel(db *sql.DB) ToolsModel {
	return &toolsModel{db}
}

type ToolsModel interface {
	GetAllTools() (*sql.Rows, error)
	GetAllToolTags() (*sql.Rows, error)
	GetOneTool(int) (*sql.Rows, error)
	GetOneToolTags(int) (*sql.Rows, error)
	InsertTool(Tool) (sql.Result, error)
	InsertToolTag(int, string) (sql.Result, error)
	UpdateTool(int, Tool) (sql.Result, error)
	DeleteTool(int) (sql.Result, error)
	DeleteAllToolTags(int) (sql.Result, error)
}

type toolsModel struct {
	db *sql.DB
}

type Tool struct {
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	Link        string   `json:"link"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

type ToolTag struct {
	Id      int
	ToolId  int
	TagName string
}

func (tm *toolsModel) GetAllTools() (rows *sql.Rows, err error) {
	query := `
		SELECT *
		FROM tool
	`

	rows, err = tm.db.Query(query)
	return
}

func (tm *toolsModel) GetAllToolTags() (rows *sql.Rows, err error) {
	query := `
		SELECT *
		FROM tool_tag
	`

	rows, err = tm.db.Query(query)
	return
}

func (tm *toolsModel) GetOneTool(id int) (rows *sql.Rows, err error) {
	query := `
		SELECT *
		FROM tool
		WHERE id = $1
	`

	rows, err = tm.db.Query(query, id)
	return
}

func (tm *toolsModel) GetOneToolTags(toolId int) (rows *sql.Rows, err error) {
	query := `
		SELECT *
		FROM tool_tag
		WHERE 
	`

	rows, err = tm.db.Query(query)
	return
}

func (tm *toolsModel) InsertTool(tool Tool) (res sql.Result, err error) {
	query := `
		INSERT INTO tool (
			title, link, description, tags
		) VALUES (
			$1, $2, $3, $4
		)
	`

	res, err = tm.db.Exec(query, tool.Title, tool.Link, tool.Description, tool.Tags)
	return
}

func (tm *toolsModel) InsertToolTag(toolId int, tag string) (res sql.Result, err error) {
	query := `
		INSERT INTO tool_tag (
			tool_id, tag_name
		) VALUES (
			$1, $2
		)
	`

	res, err = tm.db.Exec(query, toolId, tag)
	return
}

func (tm *toolsModel) UpdateTool(id int, tool Tool) (res sql.Result, err error) {
	query := `
		UPDATE tool
		SET
			title = $1,
			link = $2,
			description = $3,
			tags = $4
		WHERE id = $5
	`

	res, err = tm.db.Exec(query, tool.Title, tool.Link, tool.Description, tool.Tags, id)
	return
}

func (tm *toolsModel) DeleteTool(id int) (res sql.Result, err error) {
	query := `
		DELETE FROM tool
		WHERE id = $1
	`

	res, err = tm.db.Exec(query, id)
	return
}

func (tm *toolsModel) DeleteAllToolTags(toolId int) (res sql.Result, err error) {
	query := `
		DELETE FROM tool_tag
		WHERE tool_id = $1
	`

	res, err = tm.db.Exec(query, toolId)
	return
}
