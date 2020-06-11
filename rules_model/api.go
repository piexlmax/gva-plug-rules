package rules_model

import "github.com/jinzhu/gorm"

type GvaPlugApi struct {
	gorm.Model
	Path            string `json:"path" gorm:"comment:'api路径'"`
	Description     string `json:"description" gorm:"comment:'api中文描述'"`
	ApiGroup        string `json:"apiGroup" gorm:"comment:'api组'"`
	Method          string `json:"method" gorm:"default:'POST'" gorm:"comment:'方法'"`
	CustomTableName string `gorm:"-"`
}

func (g GvaPlugApi) TableName() string {
	if g.CustomTableName != "" {
		return g.CustomTableName
	} else {
		return "sys_apis"
	}
}
