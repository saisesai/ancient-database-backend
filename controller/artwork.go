package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mohae/deepcopy"
	"github.com/saisesai/ancient-database-backend/model"
	"net/http"
	"time"
)

type ArtworkOutput struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Pages      []uint    `json:"pages"`        // 页面ID
	Cover      string    `json:"cover"`        // 封面
	TiMing     string    `json:"ti_ming"`      // 题名
	ZangNian   string    `json:"zang_nian"`    // 葬年
	ChaoDai    string    `json:"chao_dai"`     // 朝代
	ChuTuDi    string    `json:"chu_tu_di"`    // 出土地
	XianCangDi string    `json:"xian_cang_di"` // 现藏地
	HangZiShu  string    `json:"hang_zi_shu"`  // 行字数
	ChiCun     string    `json:"chi_cun"`      // 尺寸
	ShuoMing   string    `json:"shuo_ming"`    // 说明
	ZuNian     string    `json:"zu_nian"`      // 卒年
	NianLing   string    `json:"nian_ling"`    // 年龄
	XingBie    string    `json:"xing_bie"`     // 性别
	JiGuan     string    `json:"ji_guan"`      // 籍贯
	ZhiGai     string    `json:"zhi_gai"`      // 志盖
	MingWen    string    `json:"ming_wen"`     // 铭文
}

type ArtworkAddInput struct {
}
type ArtworkAddOutput struct {
	Msg  string `json:"msg"`
	Data struct {
		ID        uint      `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"data"`
}

func ArtworkAddHandler(ctx *gin.Context) {
	var err error
	input := ArtworkAddInput{}
	output := ArtworkAddOutput{}

	// 绑定输入参数
	err = ctx.BindJSON(&input)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, &output)
		return
	}

	// 创建对象
	artworkNew, err := model.NewArtwork()
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusInternalServerError, &output)
		return
	}

	// 整理结果
	output.Data.ID = artworkNew.ID
	output.Data.CreatedAt = artworkNew.CreatedAt
	output.Data.UpdatedAt = artworkNew.UpdatedAt

	// 返回
	output.Msg = "ok"
	ctx.JSON(http.StatusOK, &output)
}

type ArtworkDeleteInput struct {
	ID uint `json:"id"`
}
type ArtworkDeleteOutput struct {
	Msg string `json:"msg"`
}

func ArtworkDeleteHandler(ctx *gin.Context) {
	var err error
	input := ArtworkDeleteInput{}
	output := ArtworkDeleteOutput{}

	// 绑定输入
	err = ctx.BindJSON(&input)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, &output)
		return
	}

	// 查找对象
	qr, err := model.FindArtworks(input.ID)
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

type ArtworkGetInput struct {
	Condition []interface{} `json:"condition" binding:"required"`
}
type ArtworkGetOutput struct {
	Msg  string          `json:"msg"`
	Data []ArtworkOutput `json:"data"`
}

func ArtworkGetHandler(ctx *gin.Context) {
	var err error
	input := ArtworkGetInput{}
	output := ArtworkGetOutput{}

	// 绑定参数
	err = ctx.BindJSON(&input)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, &output)
		return
	}

	// 查询
	qr, err := model.FindArtworks(input.Condition...)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, &output)
		return
	}

	// 整理结果
	qrBytes, err := json.Marshal(&qr)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusInternalServerError, &output)
		return
	}
	err = json.Unmarshal(qrBytes, &output.Data)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusInternalServerError, &output)
		return
	}

	// 输出结果
	output.Msg = "ok"
	ctx.JSON(http.StatusOK, &output)
}

type ArtworkModifyInput struct {
	ID         uint   `json:"id" binging:"required"`
	Cover      string `json:"cover"`        // 封面
	TiMing     string `json:"ti_ming"`      // 题名
	ZangNian   string `json:"zang_nian"`    // 葬年
	ChaoDai    string `json:"chao_dai"`     // 朝代
	ChuTuDi    string `json:"chu_tu_di"`    // 出土地
	XianCangDi string `json:"xian_cang_di"` // 现藏地
	HangZiShu  string `json:"hang_zi_shu"`  // 行字数
	ChiCun     string `json:"chi_cun"`      // 尺寸
	ShuoMing   string `json:"shuo_ming"`    // 说明
	ZuNian     string `json:"zu_nian"`      // 卒年
	NianLing   string `json:"nian_ling"`    // 年龄
	XingBie    string `json:"xing_bie"`     // 性别
	JiGuan     string `json:"ji_guan"`      // 籍贯
	ZhiGai     string `json:"zhi_gai"`      // 志盖
	MingWen    string `json:"ming_wen"`     // 铭文
}

type ArtworkModifyOutput struct {
	Msg  string        `json:"msg"`
	Data ArtworkOutput `json:"data"`
}

func ArtworkModifyHandler(ctx *gin.Context) {
	var err error
	input := ArtworkModifyInput{}
	output := ArtworkModifyOutput{}

	// 绑定参数
	err = ctx.BindJSON(&input)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, &output)
		return
	}

	// 查找
	qr, err := model.FindArtworks(input.ID)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusBadRequest, &output)
		return
	}
	q := qr[0]

	// 备份数据
	qc := deepcopy.Copy(q).(model.Artwork)

	// 更新数据
	inputBytes, err := json.Marshal(input)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusInternalServerError, &output)
		return
	}
	err = json.Unmarshal(inputBytes, &q)
	q.ID = qc.ID
	q.CreatedAt = qc.CreatedAt
	q.UpdatedAt = qc.UpdatedAt
	q.DeletedAt = qc.DeletedAt
	q.Pages = qc.Pages
	err = q.Save()
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusInternalServerError, &output)
		return
	}

	// 整理输出数据
	qBytes, err := json.Marshal(&q)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusInternalServerError, &output)
		return
	}
	err = json.Unmarshal(qBytes, &output.Data)
	if err != nil {
		output.Msg = err.Error()
		ctx.JSON(http.StatusInternalServerError, &output)
		return
	}

	// 输出结果
	output.Msg = "ok"
	ctx.JSON(http.StatusOK, &output)
}
