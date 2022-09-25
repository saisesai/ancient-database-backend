package model

import (
	"errors"
	"gorm.io/gorm"
)

// Page 页面模型
type Page struct {
	gorm.Model           // GORM基础模型
	ArtworkId  uint      `json:"artwork_id"` // 所属作品ID
	Chars      UintSlice `json:"chars"` // 字符ID
}

// NewPage 新建页面
func NewPage(pArtworkId uint) (*Page, error) {
	pageNew := &Page{
		ArtworkId: pArtworkId,
	}
	// 创建对象至数据库以获取id
	err := DB.Create(pageNew).Error
	if err != nil {
		return nil, err
	}

	// 添加作品对页面的引用
	if pageNew.ArtworkId != 0 { // ArtworkId等于0时跳过以便于测试
		parentArtwork, err := FindArtworks(pArtworkId)
		if err != nil {
			return nil, err
		}
		parentArtwork[0].Pages = append(parentArtwork[0].Pages, pageNew.ID)
		err = parentArtwork[0].Save()
		if err != nil {
			return nil, err
		}
	}

	return pageNew, nil
}

// FindPages 依据条件查询页面
func FindPages(pCondition ...interface{}) ([]Page, error) {
	var pages []Page
	err := DB.Find(&pages, pCondition...).Error
	if len(pages) == 0 {
		err = errors.New("page not found")
	}
	return pages, err
}

// Save 保存页面
func (page *Page) Save() error {
	return DB.Save(page).Error
}

// Delete 删除页面及其相关内容
func (page *Page) Delete() error {
	// 重新从数据库加载数据
	err := page.Reload()
	if err != nil {
		return err
	}

	// 删除作品对页面的引用
	if page.ArtworkId != 0 { // ArtworkId等于0时跳过以便于测试
		parentArtwork, err := FindArtworks(page.ArtworkId)
		if err != nil {
			return err
		}
		parentArtwork[0].Pages.FindAndDelete(page.ID)
		err = parentArtwork[0].Save()
		if err != nil {
			return err
		}
	}

	// 删除所有相关字符
	if len(page.Chars) != 0 {
		chars, err := FindChars(page.Chars.ToIntSlice())
		if err != nil {
			return err
		}
		for _, char := range chars {
			err = char.Delete()
			if err != nil {
				return err
			}
		}
	}
	// 删除页面自身
	return DB.Delete(page).Error
}

// Reload 重新加载数据库中的值
func (page *Page) Reload() error {
	ps, err := FindPages(page.ID)
	if err != nil {
		return err
	}
	*page = ps[0]
	return nil
}
