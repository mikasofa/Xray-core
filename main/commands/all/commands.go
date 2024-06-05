package all

import (
	"github.com/mikasofa/xray-core/main/commands/all/api"
	"github.com/mikasofa/xray-core/main/commands/all/tls"
	"github.com/mikasofa/xray-core/main/commands/base"
)

// go:generate go run github.com/mikasofa/xray-core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
		cmdWG,
	)
}
