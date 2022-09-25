package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// UintSlice 无符号整形切片模型
type UintSlice []uint

// Scan 实现sql.Scanner接口
func (us *UintSlice) Scan(pSrc any) error {
	bytes, ok := pSrc.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", pSrc))
	}
	return json.Unmarshal(bytes, us)
}

// Value 实现sql.Valuer接口
func (us UintSlice) Value() (driver.Value, error) {
	return json.Marshal(&us)
}

// FindAndDelete 查找并删除切片中的元素
func (us *UintSlice) FindAndDelete(pValue uint) {
	var tmp []uint
	for _, v := range *us {
		if v != pValue {
			tmp = append(tmp, v)
		}
	}
	*us = tmp
}

// ToIntSlice 转换为int切片来给gorm查找使用
func (us *UintSlice) ToIntSlice() []int {
	if len(*us) == 0 { // 防止返回空数组 造成全选
		return []int{0}
	}
	var is []int
	for _, v := range *us {
		is = append(is, int(v))
	}
	return is
}
