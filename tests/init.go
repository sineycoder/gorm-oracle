package tests

import (
	"time"

	gorm_oracle "github.com/sineycoder/gorm-oracle"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	url := gorm_oracle.BuildDSN("10.174.240.182", 32619, "xe", "siney", "siney", nil)
	//url := oracle.BuildUrl("127.0.0.1", 1521, "pdb1", "system", "000000", nil)
	dba, err := gorm.Open(gorm_oracle.Open(url))
	if err != nil {
		panic(err)
	}
	db = dba.Debug()
}

type T1 struct {
	ID   int64     // NUMBER
	ACOL string    `gorm:"column:A_COL"`                           // NVARCHAR2
	BCOL string    `gorm:"column:B_COL"`                           // NCLOB
	CCOL int64     `gorm:"column:C_COL"`                           // NUMBER
	DCOL string    `gorm:"column:D_COL"`                           // NVARCHAR2
	ECOL time.Time `gorm:"column:E_COL;default:CURRENT_TIMESTAMP"` // timestamp
	FCOL string    `gorm:"column:F_COL"`
	GCOL bool      `gorm:"column:G_COL"`
}

type User struct {
	ID      int64
	Name    string
	Age     int
	OtherID int64
	T1      T1 `gorm:"foreignKey:OtherID;references:ID"`
}
