package config

import(
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"  // the _ in front is a must
)

var (
  db * gorm.DB
)

func Connect() {
  d, err := gorm.Open("mysql", "Snow-00:)
}