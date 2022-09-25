package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/saisesai/ancient-database-backend/model"
	"net/http"
	"time"
)

type PageOutput struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ArtworkId uint      `json:"artwork_id"`
	Chars     []uint    `json:"chars"`
}

type PageAddInput struct {
	ArtworkId *uint `json:"artwork_id" binding:"required"`
}

type PageAddOutput struct {
	Msg  string     `json:"msg"`
	Data PageOutput `json:"data"`
}

func PageAddHandler(ctx *gin.Context) {
	var err error
	input := PageAddInput{}
	output := PageAddOutput{}

	// 绑定输入参数
	err = ctx.BindJSON(&input)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, &output)
		return
	}

	// 创建对象
	pageNew, err := model.NewPage(*input.ArtworkId)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusInternalServerError, &output)
		return
	}

	// 整理返回结果
	output.Data.ID = pageNew.ID
	output.Data.CreatedAt = pageNew.CreatedAt
	output.Data.UpdatedAt = pageNew.UpdatedAt
	output.Data.ArtworkId = pageNew.ArtworkId
	output.Data.ArtworkId = pageNew.ArtworkId
	output.Data.Chars = pageNew.Chars

	// 返回
	output.Msg = "ok"
	ctx.JSON(http.StatusOK, &output)
}

type PageDeleteInput struct {
	ID uint `json:"id" binding:"required"`
}

type PageDeleteOutput struct {
	Msg string `json:"msg"`
}

func PageDeleteHandler(ctx *gin.Context) {
	var err error
	input := PageDeleteInput{}
	output := PageDeleteOutput{}

	// 绑定输入
	err = ctx.BindJSON(&input)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, &output)
		return
	}

	// 查找对象
	qr, err := model.FindPages(input.ID)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, &output)
		return
	}
	q := qr[0]

	// 删除对象
	err = q.Delete()
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusInternalServerError, &output)
		return
	}

	// 返回
	output.Msg = "ok"
	ctx.JSON(http.StatusOK, &output)
}

type PageGetInput struct {
	Condition []interface{} `json:"condition" binding:"required"`
}

type PageGetOutput struct {
	Msg  string       `json:"msg"`
	Data []PageOutput `json:"data"`
}

func PageGetHandler(ctx *gin.Context) {
	var err error
	input := PageGetInput{}
	output := PageGetOutput{}

	// 绑定参数
	err = ctx.BindJSON(&input)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, &output)
		return
	}

	// 查询
	qr, err := model.FindPages(input.Condition...)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, &output)
		return
	}

	// 整理结果
	for _, v := range qr {
		out := PageOutput{}
		out.ID = v.ID
		out.CreatedAt = v.CreatedAt
		out.UpdatedAt = v.UpdatedAt
		out.ArtworkId = v.ArtworkId
		out.Chars = v.Chars
		output.Data = append(output.Data, out)
	}

	// 输出结果
	output.Msg = "ok"
	ctx.JSON(http.StatusOK, &output)
}

type PageModifyInput struct {
	ID uint `json:"id" binding:"required"`
}

type PageModifyOutput struct {
	Msg string `json:"msg"`
}

func PageModifyHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &PageModifyOutput{Msg: "接口未实现"})
}
