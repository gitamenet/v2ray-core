package outbound_test

import (
	"testing"

	. "github.com/gitamenet/v2ray-core/app/proxyman/outbound"
	"github.com/gitamenet/v2ray-core/features/outbound"
)

func TestInterfaces(t *testing.T) {
	_ = (outbound.Handler)(new(Handler))
	_ = (outbound.Manager)(new(Manager))
}
