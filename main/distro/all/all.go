package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/mikasofa/xray-core/app/dispatcher"
	_ "github.com/mikasofa/xray-core/app/proxyman/inbound"
	_ "github.com/mikasofa/xray-core/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/mikasofa/xray-core/app/commander"
	_ "github.com/mikasofa/xray-core/app/log/command"
	_ "github.com/mikasofa/xray-core/app/proxyman/command"
	_ "github.com/mikasofa/xray-core/app/stats/command"

	// Developer preview services
	_ "github.com/mikasofa/xray-core/app/observatory/command"

	// Other optional features.
	_ "github.com/mikasofa/xray-core/app/dns"
	_ "github.com/mikasofa/xray-core/app/dns/fakedns"
	_ "github.com/mikasofa/xray-core/app/log"
	_ "github.com/mikasofa/xray-core/app/metrics"
	_ "github.com/mikasofa/xray-core/app/policy"
	_ "github.com/mikasofa/xray-core/app/reverse"
	_ "github.com/mikasofa/xray-core/app/router"
	_ "github.com/mikasofa/xray-core/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/mikasofa/xray-core/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/mikasofa/xray-core/app/observatory"

	// Inbound and outbound proxies.
	_ "github.com/mikasofa/xray-core/proxy/blackhole"
	_ "github.com/mikasofa/xray-core/proxy/dns"
	_ "github.com/mikasofa/xray-core/proxy/dokodemo"
	_ "github.com/mikasofa/xray-core/proxy/freedom"
	_ "github.com/mikasofa/xray-core/proxy/http"
	_ "github.com/mikasofa/xray-core/proxy/loopback"
	_ "github.com/mikasofa/xray-core/proxy/shadowsocks"
	_ "github.com/mikasofa/xray-core/proxy/socks"
	_ "github.com/mikasofa/xray-core/proxy/trojan"
	_ "github.com/mikasofa/xray-core/proxy/vless/inbound"
	_ "github.com/mikasofa/xray-core/proxy/vless/outbound"
	_ "github.com/mikasofa/xray-core/proxy/vmess/inbound"
	_ "github.com/mikasofa/xray-core/proxy/vmess/outbound"
	_ "github.com/mikasofa/xray-core/proxy/wireguard"

	// Transports
	_ "github.com/mikasofa/xray-core/transport/internet/domainsocket"
	_ "github.com/mikasofa/xray-core/transport/internet/grpc"
	_ "github.com/mikasofa/xray-core/transport/internet/http"
	_ "github.com/mikasofa/xray-core/transport/internet/httpupgrade"
	_ "github.com/mikasofa/xray-core/transport/internet/kcp"
	_ "github.com/mikasofa/xray-core/transport/internet/quic"
	_ "github.com/mikasofa/xray-core/transport/internet/reality"
	_ "github.com/mikasofa/xray-core/transport/internet/tcp"
	_ "github.com/mikasofa/xray-core/transport/internet/tls"
	_ "github.com/mikasofa/xray-core/transport/internet/udp"
	_ "github.com/mikasofa/xray-core/transport/internet/websocket"

	// Transport headers
	_ "github.com/mikasofa/xray-core/transport/internet/headers/http"
	_ "github.com/mikasofa/xray-core/transport/internet/headers/noop"
	_ "github.com/mikasofa/xray-core/transport/internet/headers/srtp"
	_ "github.com/mikasofa/xray-core/transport/internet/headers/tls"
	_ "github.com/mikasofa/xray-core/transport/internet/headers/utp"
	_ "github.com/mikasofa/xray-core/transport/internet/headers/wechat"
	_ "github.com/mikasofa/xray-core/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "github.com/mikasofa/xray-core/main/json"
	_ "github.com/mikasofa/xray-core/main/toml"
	_ "github.com/mikasofa/xray-core/main/yaml"

	// Load config from file or http(s)
	_ "github.com/mikasofa/xray-core/main/confloader/external"

	// Commands
	_ "github.com/mikasofa/xray-core/main/commands/all"
)
