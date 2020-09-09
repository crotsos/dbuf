package dbuf

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type dbufService struct {
	UnimplementedDbufServiceServer
	bq *BufferQueue
}

func newDbufService(bq *BufferQueue) *dbufService {
	s := &dbufService{}
	s.bq = bq
	return s
}

func (s *dbufService) GetDbufState(
	ctx context.Context, req *GetDbufStateRequest,
) (*GetDbufStateResponse, error) {
	state := s.bq.GetState()
	return &state, nil
}

func (s *dbufService) GetQueueState(
	ctx context.Context, req *GetQueueStateRequest,
) (*GetQueueStateResponse, error) {
	state, err := s.bq.GetQueueState(req.QueueId)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

func (s *dbufService) ModifyQueue(
	ctx context.Context, req *QueueOperationRequest,
) (*QueueOperationResponse, error) {
	switch req.Action {
	case QueueOperationRequest_QUEUE_ACTION_RELEASE:
		if err := s.bq.ReleasePackets(uint32(req.BufferId), false); err != nil {
			return nil, err
		}
	case QueueOperationRequest_QUEUE_ACTION_DROP:
		if err := s.bq.ReleasePackets(uint32(req.BufferId), true); err != nil {
			return nil, err
		}
	default:
		return nil, status.Errorf(
			codes.InvalidArgument, "unknown queue operation: %v", req.Action,
		)
	}
	resp := &QueueOperationResponse{}

	return resp, nil
}

func (s *dbufService) Subscribe(req *SubscribeRequest, stream DbufService_SubscribeServer) error {
	ch := make(chan Notification, 16)
	if err := s.bq.RegisterSubscriber(ch); err != nil {
		return err
	}
	defer s.bq.UnregisterSubscriber(ch)
	ready := &Notification{MessageType: &Notification_Ready_{&Notification_Ready{}}}
	if err := stream.Send(ready); err != nil {
		return err
	}

readLoop:
	for {
		select {
		case <-stream.Context().Done():
			log.Print("Client cancelled subscription")
			break readLoop
		case n := <-ch:
			log.Printf("Sending notification %v", n)
			if err := stream.Send(&n); err != nil {
				return err
			}
		}
	}

	return nil
}