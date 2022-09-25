package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// Region 区域模型
type Region struct {
	PosX   int `json:"pos_x"`  // x坐标
	PosY   int `json:"pos_y"`  // y坐标
	Width  int `json:"width"`  // 宽度
	Height int `json:"height"` // 长度
}

// Scan 实现sql.Scanner接口
func (region *Region) Scan(pSrc any) error {
	bytes, ok := pSrc.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", pSrc))
	}
	return json.Unmarshal(bytes, region)
}

// Value 实现sql.Valuer接口
func (region Region) Value() (driver.Value, error) {
	return json.Marshal(&region)
}
