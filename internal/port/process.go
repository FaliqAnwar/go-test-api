package port

import (
	"context"
	"net"
)

type ProcessStartStopper interface {
	Start(ctx context.Context, l net.Listener)
	Stop(ctx context.Context)
}
