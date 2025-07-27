package db

import (
	"github.com/ljcnh/flow/internal/pkg/consts"
	"gorm.io/gorm"
)

func beforeCreate(db *gorm.DB) {
	if db.Statement == nil || db.Statement.Schema == nil {
		return
	}
	username := db.Statement.Context.Value(consts.CtxUserName)
	if username == nil {
		return
	}
	if db.Statement.Schema.LookUpField(consts.CreatedBy) != nil {
		db.Statement.SetColumn(consts.CreatedBy, username)
	}
	if db.Statement.Schema.LookUpField(consts.UpdatedBy) != nil {
		db.Statement.SetColumn(consts.UpdatedBy, username)
	}
}

func beforeUpdate(db *gorm.DB) {
	if db.Statement == nil || db.Statement.Schema == nil {
		return
	}
	username := db.Statement.Context.Value(consts.CtxUserName)
	if username == nil {
		return
	}
	if db.Statement.Schema.LookUpField(consts.UpdatedBy) != nil {
		db.Statement.SetColumn(consts.UpdatedBy, username)
	}
}
