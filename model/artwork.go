package model

import (
	"errors"
	"gorm.io/gorm"
)

// Artwork 作品模型
type Artwork struct {
	gorm.Model           // GORM基础模型
	Pages      UintSlice `json:"pages"`        // 页面ID
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

// NewArtwork 新建作品
func NewArtwork() (*Artwork, error) {
	artworkNew := &Artwork{}
	return artworkNew, DB.Create(artworkNew).Error
}

// FindArtworks 依据条件查询作品
func FindArtworks(pCondition ...interface{}) ([]Artwork, error) {
	var artworks []Artwork
	err := DB.Find(&artworks, pCondition...).Error
	if len(artworks) == 0 {
		err = errors.New("artwork not found")
	}
	return artworks, err
}

// Save 保存作品信息
func (artwork *Artwork) Save() error {
	return DB.Save(artwork).Error
}

// Delete 删除作品及其相关内容
func (artwork *Artwork) Delete() error {
	// 重新从数据库加载
	err := artwork.Reload()
	if err != nil {
		return err
	}
	// 删除相关页面
	if len(artwork.Pages) != 0 {
		pages, err := FindPages(artwork.Pages.ToIntSlice())
		if err != nil {
			return err
		}
		for _, page := range pages {
			err = page.Delete()
			if err != nil {
				return err
			}
		}
	}
	// 删除作品自身
	return DB.Delete(artwork).Error
}

// Reload 重新从数据库加载数值
func (artwork *Artwork) Reload() error {
	as, err := FindArtworks(artwork.ID)
	if err != nil {
		return err
	}
	*artwork = as[0]
	return nil
}
