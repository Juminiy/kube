package multi_tenants

import (
	"database/sql"
	"encoding/json"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm/clause"
	gormschema "gorm.io/gorm/schema"
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

func (t TenantID) CreateClauses(f *gormschema.Field) []clause.Interface {
	return []clause.Interface{}
}

func (t TenantID) QueryClauses(f *gormschema.Field) []clause.Interface {
	return []clause.Interface{}
}

func (t TenantID) UpdateClauses(f *gormschema.Field) []clause.Interface {
	return []clause.Interface{}
}

func (t TenantID) DeleteClauses(f *gormschema.Field) []clause.Interface {
	return []clause.Interface{}
}
