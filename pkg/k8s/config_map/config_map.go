package config_map

import (
	"github.com/Juminiy/kube/pkg/util"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func New(name string) corev1.ConfigMap {
	return corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Immutable: util.NewBool(true),
		Data: map[string]string{
			"MySQLAddr":         "mysql.mysql:mysqlPort",
			"SQLiteAddr":        "sqlite3.sqlite3:sqlitePath",
			"RedisAddr":         "redis.redis:redisPort",
			"ElasticSearchAddr": "es.es:esPort",
		},
		BinaryData: map[string][]byte{
			"ChineseKey":  []byte("非UTF8中文"),
			"JapaneseKey": []byte("日本語"),
		},
	}
}
