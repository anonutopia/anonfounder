package main

import (
	"github.com/jinzhu/gorm"
)

// KeyValue model is used for storing key/values
type KeyValue struct {
	gorm.Model
	Key      string `sql:"size:255;unique_index"`
	ValueInt uint64 `sql:"type:int"`
	ValueStr string `sql:"type:string"`
}
