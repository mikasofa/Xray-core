package tcp

import (
	"github.com/mikasofa/xray-core/common"
	"github.com/mikasofa/xray-core/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
