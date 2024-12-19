package dynmgrm_test

import (
	"log"

	"github.com/himgang/dynmgrm"
	"github.com/miyamo2/sqldav"
	"gorm.io/gorm"
)

func ExampleListAppend() {
	db, err := gorm.Open(
		dynmgrm.New(),
		&gorm.Config{
			SkipDefaultTransaction: true,
		})
	if err != nil {
		log.Fatalf("failed to open database, got error %v", err)
	}
	db.Model(&TestTable{PK: "Partition1", SK: 1}).
		Update("list_type_attr",
			dynmgrm.ListAppend(sqldav.Map{"Foo": "Bar"}))
}
