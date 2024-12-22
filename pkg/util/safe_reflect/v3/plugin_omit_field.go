package safe_reflectv3

import (
	"github.com/Juminiy/kube/pkg/util"
	"k8s.io/apimachinery/pkg/util/sets"
)

type FieldOmit struct {
	read  sets.Set[string]
	write sets.Set[string]
}

func NewFieldOmit() *FieldOmit {
	return &FieldOmit{
		read:  sets.New[string](),
		write: sets.New[string](),
	}
}

func (etv *ETv) OmitReadByTag2(tagKey, valKey, valVal string) *ETv {
	tagValName := etv.Tag2VName(tagKey, valKey)
	if name, ok := util.MapElemOk(tagValName, valVal); ok {
		return etv.OmitRead(name)
	}
	return etv
}

func (etv *ETv) OmitRead(s ...string) *ETv {
	if etv.FieldOmit == nil {
		etv.FieldOmit = NewFieldOmit()
	}
	etv.read.Insert(s...)
	return etv
}

func (etv *ETv) OmitWrite(s ...string) *ETv {
	if etv.FieldOmit == nil {
		etv.FieldOmit = NewFieldOmit()
	}
	etv.write.Insert(s...)
	return etv
}

func (etv *ETv) StructFields() Fields {
	fields := etv.Tv.StructFields()
	util.MapDelete(fields, etv.read.UnsortedList()...)
	return fields
}
