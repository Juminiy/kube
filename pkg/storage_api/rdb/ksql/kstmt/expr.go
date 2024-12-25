package kstmt

import "github.com/Juminiy/kube/pkg/storage_api/rdb/kinternal"

type Expr interface {
	kinternal.WriteBuilder
}
