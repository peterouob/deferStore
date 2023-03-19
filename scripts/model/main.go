package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	dsn := "peter:password@tcp(localhost:3306)/deferStore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("%s", "connect to database have error"))
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:      "/Users/peter/Desktop/deferStore/source/logic/orm/dal",
		ModelPkgPath: "/Users/peter/Desktop/deferStore/source/logic/orm/model",
		Mode:         gen.WithDefaultQuery | gen.WithoutContext,
	})
	g.UseDB(db)
	g.ApplyBasic(g.GenerateModel("Account"), g.GenerateModel("block"), g.GenerateModel("goods", gen.FieldType("price", "decimal.Decimal"), gen.FieldType("original_price", "decimal.Decimal")))
	g.Execute()
}
