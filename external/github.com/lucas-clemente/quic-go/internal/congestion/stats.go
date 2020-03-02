package congestion

import "github.com/gitamenet/v2ray-core/external/github.com/lucas-clemente/quic-go/internal/protocol"

type connectionStats struct {
	slowstartPacketsLost protocol.PacketNumber
	slowstartBytesLost   protocol.ByteCount
}
