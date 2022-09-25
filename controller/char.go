package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/saisesai/ancient-database-backend/model"
	"net/http"
	"time"
)

type CharOutput struct {
	ID        uint         `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	PageId    uint         `json:"page_id"`
	Value     string       `json:"value"`
	Region    model.Region `json:"region"`
}

type CharNewInput struct {
	PageId *uint  `json:"page_id" binding:"required"`
	Value  string `json:"value"   binding:"required"`
}

type CharNewOutput struct {
	Msg  string     `json:"msg"`
	Data CharOutput `json:"data"`
}

func CharAddHandler(ctx *gin.Context) {
	var err error
	input := CharNewInput{}
	output := CharNewOutput{}

	// 绑定参数
	err = ctx.BindJSON(&input)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, &output)
		return
	}

	// 创建对象
	ch, err := model.NewChar(*input.PageId, input.Value)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusInternalServerError, &output)
		return
	}

	// 填充结果
	output.Data.ID = ch.ID
	output.Data.CreatedAt = ch.CreatedAt
	output.Data.UpdatedAt = ch.UpdatedAt
	output.Data.PageId = ch.PageId
	output.Data.Value = ch.Value
	output.Data.Region = ch.Region

	// 输出结果
	output.Msg = "ok"
	ctx.JSON(http.StatusOK, &output)
}

type CharDeleteInput struct {
	ID uint `json:"id" binding:"required"`
}

type CharDeleteOutput struct {
	Msg string `json:"msg"`
}

func CharDeleteHandler(ctx *gin.Context) {
	var err error
	input := CharDeleteInput{}
	output := CharDeleteOutput{}

	// 绑定参数
	err = ctx.BindJSON(&input)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, &output)
		return
	}

	// 查找对象
	chs, err := model.FindChars("id = ?", input.ID)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, &output)
		return
	}

	// 删除对象
	err = chs[0].Delete()
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusInternalServerError, &output)
		return
	}

	// 输出结果
	output.Msg = "ok"
	ctx.JSON(http.StatusOK, &output)
}

type CharGetInput struct {
	Condition []interface{} `json:"condition" binding:"required"`
}

type CharGetOutput struct {
	Msg  string       `json:"msg"`
	Data []CharOutput `json:"data"`
}

func CharGetHandler(ctx *gin.Context) {
	var err error
	input := CharGetInput{}
	output := CharGetOutput{}

	// 绑定参数
	err = ctx.BindJSON(&input)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, &output)
		return
	}

	// 执行查询
	queryResult, err := model.FindChars(input.Condition...)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, &output)
		return
	}

	// 整理结果
	for _, charRaw := range queryResult {
		out := CharOutput{}
		out.ID = charRaw.ID
		out.CreatedAt = charRaw.CreatedAt
		out.UpdatedAt = charRaw.UpdatedAt
		out.PageId = charRaw.PageId
		out.Value = charRaw.Value
		out.Region = charRaw.Region
		output.Data = append(output.Data, out)
	}

	// 输出结果
	output.Msg = "ok"
	ctx.JSON(http.StatusOK, &output)
}

type CharModifyInput struct {
	ID     uint         `json:"id" binding:"required"`
	Value  string       `json:"value" binding:"required"`
	Region model.Region `json:"region" binding:"required"`
}

type CharModifyOutput struct {
	Msg  string     `json:"msg"`
	Data CharOutput `json:"data"`
}

func CharModifyHandler(ctx *gin.Context) {
	var err error
	input := CharModifyInput{}
	output := CharModifyOutput{}

	// 绑定参数
	err = ctx.BindJSON(&input)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, &output)
		return
	}

	// 查找
	queryResult, err := model.FindChars(input.ID)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, &output)
		return
	}

	// 更新数据
	queryResult[0].Value = input.Value
	queryResult[0].Region = input.Region
	err = queryResult[0].Save()
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusInternalServerError, &output)
		return
	}

	// 整理输出数据
	output.Data.ID = queryResult[0].ID
	output.Data.CreatedAt = queryResult[0].CreatedAt
	output.Data.UpdatedAt = queryResult[0].UpdatedAt
	output.Data.PageId = queryResult[0].PageId
	output.Data.Value = queryResult[0].Value
	output.Data.Region = queryResult[0].Region

	// 输出结果
	output.Msg = "ok"
	ctx.JSON(http.StatusOK, &output)
}
