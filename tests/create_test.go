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
			ACOL: "11",
			BCOL: "11",
			//CCOL: 11,
			DCOL: "11",
			ECOL: a,
			FCOL: `{"nzx": 1}`,
		},
		{
			ACOL: "12",
			BCOL: "12",
			CCOL: 12,
			DCOL: "12",
			ECOL: a,
			FCOL: `{"nzx": 2}`,
		},
	}
	//tmps := &T1{
	//	ACOL: "11",
	//	BCOL: "11",
	//	//CCOL: 11,
	//	DCOL: "11",
	//	ECOL: a,
	//	FCOL: `{"nzx": 1}`,
	//}
	res := db.Create(&tmps)
	//res := db.Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "A_COL"}},
	//	DoUpdates: clause.AssignmentColumns([]string{"B_COL", "C_COL", "D_COL"}),
	//}).Create(&tmps)
	//fmt.Println(sql)
	assert.Nil(t, res.Error)
}
