package server

import (
	"fmt"
	"time"

	"io"
	"log"

	"github.com/gopiesy/project-protos/policies"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PolicyServiceServer struct {
	policies.UnimplementedPolicyServiceServer
}

func (p PolicyServiceServer) StreamSnapshots(stream policies.PolicyService_StreamSnapshotsServer) error {

	streamCpy := stream
	go func(stream policies.PolicyService_StreamSnapshotsServer) {
		for {
			status, err := stream.Recv()
			if err == io.EOF {
				log.Println("EOF received and client exited")
				return
			} else if err != nil {
				log.Panic(err)
			} else {
				// received status
				log.Println("Recv Status: ", status.GetSnapshotName())
			}
		}
	}(streamCpy)

	for {
		now := timestamppb.Now()
		name := fmt.Sprintf("NewSnapshot.%d", now.Seconds)
		if err := stream.Send(&policies.Snapshot{Name: name, Time: now, Data: nil}); err != nil {
			return err
		}
		fmt.Println("Send NewSnapshot: ", name)
		<-time.After(time.Duration(time.Second * 2))
	}
}

func NewPolicyService() PolicyServiceServer {
	return PolicyServiceServer{}
}
