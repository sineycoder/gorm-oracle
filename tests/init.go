package tests

import (
	"time"

	gorm_oracle "github.com/sineycoder/gorm-oracle"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	url := gorm_oracle.BuildDSN("10.174.252.190", 32619, "xe", "siney", "siney", nil)
	//url := oracle.BuildUrl("127.0.0.1", 1521, "pdb1", "system", "000000", nil)
	dba, err := gorm.Open(gorm_oracle.Open(url), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	db = dba
}

type T1 struct {
	ID   int64     // NUMBER
	ACol string    // NVARCHAR2
	BCol string    // NCLOB
	CCol int64     // NUMBER
	DCol string    // NVARCHAR2
	ECol time.Time `gorm:"default:SYSTIMESTAMP"` // timestamp
	FCol string
	GCol bool
}

type User struct {
	ID      int64
	Name    string
	Age     int
	OtherID int64
	T1      T1 `gorm:"foreignKey:OtherID;references:ID"`
}
