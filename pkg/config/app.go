package config

import(
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"  // the _ in front is a must
)

var (
  db *gorm.DB
)

func Connect() {
  d, err := gorm.Open("mysql", "shadow00:ygdrassil@/simplerest?charset=utf8&parseTime=True&loc=Local")  // account:password@/db_name?
  if err != nil {
    panic(err)
  }
  db = d
}

func GetDB() *gorm.DB {
  return db
}