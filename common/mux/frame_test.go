package mux_test

import (
	"testing"

	"github.com/gitamenet/v2ray-core/common"
	"github.com/gitamenet/v2ray-core/common/buf"
	"github.com/gitamenet/v2ray-core/common/mux"
	"github.com/gitamenet/v2ray-core/common/net"
)

func BenchmarkFrameWrite(b *testing.B) {
	frame := mux.FrameMetadata{
		Target:        net.TCPDestination(net.DomainAddress("www.v2ray.com"), net.Port(80)),
		SessionID:     1,
		SessionStatus: mux.SessionStatusNew,
	}
	writer := buf.New()
	defer writer.Release()

	for i := 0; i < b.N; i++ {
		common.Must(frame.WriteTo(writer))
		writer.Clear()
	}
}
