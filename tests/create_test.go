package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	a := time.Now()
	tmps := []*T1{
		{
			ACol: "11",
			BCol: "11",
			//CCol: 11,
			DCol: "11",
			ECol: a,
			FCol: `{"nzx": 1}`,
		},
		{
			ACol: "12",
			BCol: "12",
			CCol: 12,
			DCol: "12",
			ECol: a,
			FCol: `{"nzx": 2}`,
			GCol: true,
		},
	}
	//tmps := &T1{
	//	ACol: "11",
	//	BCol: "11",
	//	//CCol: 11,
	//	DCol: "11",
	//	ECol: a,
	//	FCol: `{"nzx": 1}`,
	//}
	res := db.Create(&tmps)
	//res := db.Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "A_COL"}},
	//	DoUpdates: clause.AssignmentColumns([]string{"B_COL", "C_COL", "D_COL"}),
	//}).Create(&tmps)
	//fmt.Println(sql)
	assert.Nil(t, res.Error)
}
