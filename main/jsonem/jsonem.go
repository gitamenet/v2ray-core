package jsonem

import (
	"github.com/gitamenet/v2ray-core"
	"github.com/gitamenet/v2ray-core/common"
	"github.com/gitamenet/v2ray-core/infra/conf/serial"
)

func init() {
	common.Must(core.RegisterConfigLoader(&core.ConfigFormat{
		Name:      "JSON",
		Extension: []string{"json"},
		Loader:    serial.LoadJSONConfig,
	}))
}
