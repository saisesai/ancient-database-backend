package model

import (
	"errors"
	"gorm.io/gorm"
)

// Char 文字模型
type Char struct {
	gorm.Model        // GORM基础模型
	PageId     uint   `json:"page_id"` // 属于的页面ID
	Value      string `json:"value"`   // UTF8字符
	Region     Region `json:"region"`  // 在页面上的区域
}

// NewChar 新建字符
func NewChar(pPageId uint, pValue string) (*Char, error) {
	// 填入初始化数值
	charNew := &Char{
		PageId: pPageId,
		Value:  pValue,
		Region: Region{0, 0, 0, 0},
	}

	// 创建数据库对象
	if err := DB.Create(charNew).Error; err != nil {
		return nil, err
	}

	// 添加页对字符的引用
	if pPageId != 0 { // PageId等于0时跳过以便于测试
		parentPage, err := FindPages(pPageId)
		if err != nil {
			return nil, err
		}
		parentPage[0].Chars = append(parentPage[0].Chars, charNew.ID)
		err = parentPage[0].Save()
		if err != nil {
			return nil, err
		}
	}
	return charNew, nil
}

// FindChars 依据条件查找字符
func FindChars(pCondition ...interface{}) ([]Char, error) {
	var chars []Char
	err := DB.Find(&chars, pCondition...).Error
	if len(chars) == 0 {
		err = errors.New("char not found")
	}
	return chars, err
}

// Save 保存字符
func (char *Char) Save() error {
	return DB.Save(char).Error
}

// Delete 删除字符
func (char *Char) Delete() error {
	// 删除页对字符的引用
	if char.PageId != 0 { // PageId等于0时跳过以便于测试
		parentPage, err := FindPages(char.PageId)
		if err != nil {
			return err
		}
		parentPage[0].Chars.FindAndDelete(char.ID)
		err = parentPage[0].Save()
		if err != nil {
			return err
		}
	}
	// 删除字符自身
	return DB.Delete(char).Error
}
