package rules_servies

import (
	"github.com/jinzhu/gorm"
	"github.com/piexlmax/gva-plug-rules/rules-model"
)

func RegisterApi(db *gorm.DB, apis []rules_model.GvaPlugApi) {
	for _, v := range apis {
		if NeedRegister := db.First(&rules_model.GvaPlugApi{}, "path = ? AND method = ?", v.Path, v.Method).RecordNotFound(); NeedRegister {
			db.Create(&v)
		}
	}
}
