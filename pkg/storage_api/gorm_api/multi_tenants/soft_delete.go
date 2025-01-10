package multi_tenants

import (
	"github.com/Juminiy/kube/pkg/util"
	gormschema "gorm.io/gorm/schema"
)

func DeletedAt(schema *gormschema.Schema) *Field { // maybe not required
	deletedAt := schema.LookUpField("DeletedAt")
	if deletedAt == nil {
		deletedAt = schema.LookUpField("deleted_at")
		if deletedAt == nil { // pkg soft_delete
			return nil
		}
	}
	return util.New(FieldFromSchema(deletedAt))
}
