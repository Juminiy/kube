package main

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/netserver/http/stdserver"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"net/http"
	"strings"
	"sync"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.Use(gin.Recovery())
	r.GET("/", setValueGetValues)
	r.GET("/delete", deleteValue)
	r.GET("/search", searchValues)
	stdserver.ListenAndServeInfoF(false, 8000, util.IsIPv4)
	stdlog.Fatal(r.Run("0.0.0.0:8000"))
}

var _clipBoard = sync.Map{}

func setValueGetValues(c *gin.Context) {
	value := c.Query("value")
	if len(value) > 0 {
		_clipBoard.Store(value, struct{}{})
	}
	renderValues(c, "")
}

func deleteValue(c *gin.Context) {
	value := c.Query("value")
	if len(value) > 0 {
		_clipBoard.Delete(value)
	}
	renderValues(c, "")
}

func searchValues(c *gin.Context) {
	value := c.Query("value")
	renderValues(c, value)
}

func clipValues() []string {
	values := make([]string, 0, util.MagicSliceCap)
	_clipBoard.Range(func(key, value any) bool {
		values = append(values, key.(string))
		return true
	})
	return values
}

func renderValues(c *gin.Context, search string) {
	values := clipValues()
	if len(search) > 0 {
		values = lo.Filter(lo.Map(clipValues(), func(item string, index int) string {
			if strings.Contains(search, item) ||
				strings.Contains(item, search) {
				return item
			}
			return ""
		}), func(item string, index int) bool {
			return len(item) != 0
		})
	}
	c.String(http.StatusOK, safe_json.Pretty(values))
}
