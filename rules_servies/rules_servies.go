package rules_servies

import (
	"github.com/piexlmax/gva-plug-rules/rules-model"
	"github.com/jinzhu/gorm"
)

func RegistApi(db *gorm.DB,apis []rules_model.GvaPlugApi){
	for _,v:= range apis{
		db.Create(&v)
	}
}