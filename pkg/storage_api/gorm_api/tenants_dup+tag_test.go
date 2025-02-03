package gorm_api

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"testing"
)

/*
	BeforeWriteDupCheck
*/

var Create = func(t *testing.T, i any) {
	Err(t, _txTenant().Create(i).Error)
}

var Update = func(t *testing.T, model, dest any, conds ...any) {
	Err(t, _txTenant().Model(model).Where(conds[0], conds[1:]...).Updates(dest).Error)
}

type UniqueTest1 struct {
	gorm.Model
	Name     string `mt:"unique:name"`
	NumberID string `mt:"unique:name,birth"`
	Birth    int    `mt:"unique:birth"`
}

func TestOneFieldOverlapInGe2Groups(t *testing.T) {
	Create(t,
		&UniqueTest1{
			Name:     "Galaxy",
			NumberID: "0019527",
			Birth:    1919730,
		})
}

func TestUpdate1(t *testing.T) {
	Update(t,
		UniqueTest1{},
		UniqueTest1{
			Name:     "Galaxy",
			NumberID: "0019527",
			Birth:    1919730,
		},
		clause.Eq{Column: "id", Value: 1},
	)

	Update(t,
		UniqueTest1{},
		map[string]any{
			"name":      "Galaxy",
			"number_id": "0019527",
			"birth":     1919730,
		},
		clause.Eq{Column: "id", Value: 1},
	)
}

type UniqueTest2 struct {
	gorm.Model
	RegionCode int    `mt:"unique:region_ip,region_host,region_node"`
	IPAddr     string `mt:"unique:region_ip"`
	Hostname   string `mt:"unique:region_host,host_node"`
	MacAddr    string `mt:"unique:mac_addr"`
	NodeID     string `mt:"unique:region_node,host_node"`
}

func TestGe2FieldOverlapInGe2Groups(t *testing.T) {
	Create(t,
		&UniqueTest2{
			RegionCode: 666,
			IPAddr:     "10.101.22.10",
			Hostname:   "BJ01-HPC-0008",
			MacAddr:    "02:1A:2B:3C:4D:5E",
			NodeID:     "pja-0x8090621",
		})
}

type UniqueTest3 struct {
	gorm.Model
	Name string `mt:"unique"`
}

func TestOneFieldInOneGroup(t *testing.T) {
	Create(t,
		&UniqueTest3{
			Name: "RR",
		})

	Update(t, UniqueTest3{},
		map[string]any{"name": "RR"},
		clause.Eq{Column: "id", Value: 1})
}

type UniqueTest4 struct {
	gorm.Model
	Name string `mt:"unique"`
	RID  int    `mt:"unique"`
}

func TestEachOneFieldInEachOneGroup(t *testing.T) {
	Create(t,
		&UniqueTest4{
			Name: "RR",
			RID:  666,
		})
}

type UniqueTest5 struct {
	gorm.Model
	Name string `mt:"unique:name"`
	Desc string `mt:"unique:name"`
	Perm string `mt:"unique:name"`
}

func TestGe2FieldsInOneGroup(t *testing.T) {
	Create(t,
		&UniqueTest5{
			Name: "RR",
			Desc: "ff-60",
			Perm: "06-21",
		})
}

type UniqueTest6 struct {
	gorm.Model
	Name   string `mt:"unique:name"`
	Desc   string `mt:"unique:name"`
	Height int    `mt:"unique:BMI"`
	Weight int    `mt:"unique:BMI"`
	Char1  int    `mt:"unique:char"`
	Char2  int    `mt:"unique:char"`
	Char3  int    `mt:"unique:char"`
}

func TestGe2FieldsInGe2GroupNoOverlap(t *testing.T) {
	Create(t,
		&UniqueTest6{
			Name:   "RR",
			Desc:   "ff-06",
			Height: 183,
			Weight: 73,
			Char1:  5,
			Char2:  2,
			Char3:  1,
		})
}
