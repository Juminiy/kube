package postgres

import (
	"gorm.io/gorm"
	"math"
	"testing"
)

type Sale struct {
	gorm.Model
	Int2    int16   `gorm:"type:int2"`
	Int4    int32   `gorm:"type:int4"`
	Int8    int64   `gorm:"type:int8"`
	Decimal string  `gorm:"type:decimal"`
	Numeric string  `gorm:"type:numeric"`
	Real    float32 `gorm:"type:real"`
	Double  float64 `gorm:"type:double precision"`
	Serial2 uint16  `gorm:"type:smallint"`
	Serial4 uint32  `gorm:"type:integer"`
	Serial8 uint64  `gorm:"type:bigint"`
}

func TestNumericType(t *testing.T) {
	//Err(t, txMigrate().AutoMigrate(Sale{}))
	var sale = Sale{
		Int2:    math.MinInt16,
		Int4:    math.MaxInt32,
		Int8:    math.MaxInt64,
		Decimal: "1919.810",
		Numeric: "-114.514",
		Real:    666.666,
		Double:  888.888,
		Serial2: 0,
		Serial4: 0,
		Serial8: 0,
	}
	Err(t, _txTenant().Create(&sale).Error)
	t.Log(Enc(sale))
	//time.Sleep(util.TimeSecond(60))
}

type SaleR struct {
	gorm.Model
	Varchar string `gorm:"type:varchar(8)"`
	Char    string `gorm:"type:char(1)"`
	//BpChar  string `gorm:"type:bpchar(6)"`
	Text string `gorm:"type:text"`
}

func TestCharType(t *testing.T) {
	//Err(t, txMigrate().AutoMigrate(SaleR{}))
	var saleR = SaleR{
		Varchar: "1234567",
		Char:    "1",
		//BpChar:  "whoami",
		Text: "iamhajimi",
	}
	Err(t, _txTenant().Create(&saleR).Error)
	t.Log(Enc(saleR))
}
