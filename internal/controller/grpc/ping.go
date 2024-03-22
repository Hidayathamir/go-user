package grpc

import (
	"context"

	"github.com/Hidayathamir/go-user/pkg/gousergrpc"
)

// Ping is controller GRPC for ping related.
type Ping struct {
	gousergrpc.UnimplementedPingServer
}

var _ gousergrpc.PingServer = &Ping{}

// Ping implements pb.PingServer.
func (p *Ping) Ping(context.Context, *gousergrpc.PingEmpty) (*gousergrpc.ResPing, error) {
	return &gousergrpc.ResPing{Message: "ping success"}, nil
}
