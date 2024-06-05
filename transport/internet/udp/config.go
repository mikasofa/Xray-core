package udp

import (
	"github.com/mikasofa/xray-core/common"
	"github.com/mikasofa/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
