package main

import (
	"fmt"
	"io"
	"sync"

	examples "github.com/gengo/grpc-gateway/examples/examplepb"
	sub "github.com/gengo/grpc-gateway/examples/sub"
	"github.com/golang/glog"
	"github.com/rogpeppe/fastuuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Implements of ABitOfEverythingServiceServer

var uuidgen = fastuuid.MustNewGenerator()

type _ABitOfEverythingServer struct {
	v map[string]*examples.ABitOfEverything
	m sync.Mutex
}

func newABitOfEverythingServer() examples.ABitOfEverythingServiceServer {
	return &_ABitOfEverythingServer{
		v: make(map[string]*examples.ABitOfEverything),
	}
}

func (s *_ABitOfEverythingServer) Create(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	s.m.Lock()
	defer s.m.Unlock()

	glog.Info(msg)
	var uuid string
	for {
		uuid = fmt.Sprintf("%x", uuidgen.Next())
		if _, ok := s.v[uuid]; !ok {
			break
		}
	}
	s.v[uuid] = msg
	s.v[uuid].Uuid = uuid
	glog.Infof("%v", s.v[uuid])
	return s.v[uuid], nil
}

func (s *_ABitOfEverythingServer) CreateBody(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	return s.Create(ctx, msg)
}

func (s *_ABitOfEverythingServer) BulkCreate(stream examples.ABitOfEverythingService_BulkCreateServer) error {
	ctx := stream.Context()
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		glog.Error(msg)
		if _, err = s.Create(ctx, msg); err != nil {
			return err
		}
	}
	return stream.SendAndClose(new(examples.EmptyMessage))
}

func (s *_ABitOfEverythingServer) Lookup(ctx context.Context, msg *examples.IdMessage) (*examples.ABitOfEverything, error) {
	s.m.Lock()
	defer s.m.Unlock()

	glog.Info(msg)
	if a, ok := s.v[msg.Uuid]; ok {
		return a, nil
	}
	return nil, grpc.Errorf(codes.NotFound, "not found")
}

func (s *_ABitOfEverythingServer) List(_ *examples.EmptyMessage, stream examples.ABitOfEverythingService_ListServer) error {
	s.m.Lock()
	defer s.m.Unlock()
	for _, msg := range s.v {
		if err := stream.Send(msg); err != nil {
			return err
		}
	}
	return nil
}

func (s *_ABitOfEverythingServer) Update(ctx context.Context, msg *examples.ABitOfEverything) (*examples.EmptyMessage, error) {
	s.m.Lock()
	defer s.m.Unlock()

	glog.Info(msg)
	if _, ok := s.v[msg.Uuid]; ok {
		s.v[msg.Uuid] = msg
	} else {
		return nil, grpc.Errorf(codes.NotFound, "not found")
	}
	return new(examples.EmptyMessage), nil
}

func (s *_ABitOfEverythingServer) Delete(ctx context.Context, msg *examples.IdMessage) (*examples.EmptyMessage, error) {
	s.m.Lock()
	defer s.m.Unlock()

	glog.Info(msg)
	if _, ok := s.v[msg.Uuid]; ok {
		delete(s.v, msg.Uuid)
	} else {
		return nil, grpc.Errorf(codes.NotFound, "not found")
	}
	return new(examples.EmptyMessage), nil
}

func (s *_ABitOfEverythingServer) Echo(ctx context.Context, msg *sub.StringMessage) (*sub.StringMessage, error) {
	glog.Info(msg)
	return msg, nil
}

func (s *_ABitOfEverythingServer) BulkEcho(stream examples.ABitOfEverythingService_BulkEchoServer) error {
	var msgs []*sub.StringMessage
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		msgs = append(msgs, msg)
	}
	for _, msg := range msgs {
		glog.Info(msg)
		if err := stream.Send(msg); err != nil {
			return err
		}
	}
	return nil
}

func (s *_ABitOfEverythingServer) EchoABEProto2(ctx context.Context, msg *examples.ABitOfEverythingProto2) (*examples.ABitOfEverythingProto2, error) {
	glog.Info(msg)
	return msg, nil
}

func (s *_ABitOfEverythingServer) BulkEchoABEProto2(stream examples.ABitOfEverythingService_BulkEchoABEProto2Server) error {
	var msgs []*examples.ABitOfEverythingProto2
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		msgs = append(msgs, msg)
	}
	for _, msg := range msgs {
		glog.Info(msg)
		if err := stream.Send(msg); err != nil {
			return err
		}
	}
	return nil
}
