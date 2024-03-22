package grpc

import (
	"context"

	"github.com/Hidayathamir/go-user/pkg/gouser/grpc/pb"
)

// Ping is controller GRPC for ping related.
type Ping struct {
	pb.UnimplementedPingServer
}

var _ pb.PingServer = &Ping{}

// Ping implements pb.PingServer.
func (p *Ping) Ping(context.Context, *pb.PingEmpty) (*pb.ResPing, error) {
	return &pb.ResPing{Message: "ping success"}, nil
}
