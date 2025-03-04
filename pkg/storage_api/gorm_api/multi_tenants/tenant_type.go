package multi_tenants

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/spf13/cast"
	"gorm.io/gorm/clause"
	gormschema "gorm.io/gorm/schema"
	"time"
)

type TenantID sql.Null[uint]

func (t TenantID) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.V)
	}
	return json.Marshal(nil)
}

func (t *TenantID) UnmarshalJSON(b []byte) error {
	if util.Bytes2StringNoCopy(b) == "null" {
		t.Valid = false
		return nil
	} else if err := json.Unmarshal(b, &t); err == nil {
		t.Valid = true
		return nil
	} else {
		return err
	}
}

func (t TenantID) QueryClauses(f *gormschema.Field) []clause.Interface {
	return []clause.Interface{&Tenant{Field: Field{
		Name:    f.Name,
		DBTable: f.Schema.Table,
		DBName:  f.DBName,
		Value:   t.V,
		Values:  nil,
	}}}
}

func (t TenantID) UpdateClauses(f *gormschema.Field) []clause.Interface {
	return []clause.Interface{&Tenant{Field: Field{
		Name:    f.Name,
		DBTable: f.Schema.Table,
		DBName:  f.DBName,
		Value:   t.V,
		Values:  nil,
	}}}
}

func (t TenantID) DeleteClauses(f *gormschema.Field) []clause.Interface {
	return []clause.Interface{&Tenant{Field: Field{
		Name:    f.Name,
		DBTable: f.Schema.Table,
		DBName:  f.DBName,
		Value:   t.V,
		Values:  nil,
	}}}
}

type Time sql.NullTime

var ErrTimeNotValid = errors.New("value is not null or time.Time or sql.NullTime or timestamp")

func (t Time) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Time)
	}
	return json.Marshal(nil)
}

func (t *Time) UnmarshalJSON(b []byte) error {
	if bStr := util.Bytes2StringNoCopy(b); bStr == "null" {
		t.Valid = false
		return nil
	} else if err := t.Time.UnmarshalJSON(b); err == nil {
		t.Valid = true
		return nil
	} else if bInt64 := cast.ToInt64(bStr); bInt64 > 0 {
		if unixTime := time.Unix(bInt64, 0); !unixTime.IsZero() {
			t.Valid = true
			t.Time = unixTime
			return nil
		}
	}
	t.Valid = false
	return ErrTimeNotValid
}
