package date_agent

import (
	"context"
	pb "github.com/Shanghai-Lunara/date-agent/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"k8s.io/klog/v2"
	"net"
	"net/http"
	"time"
)

var kaep = keepalive.EnforcementPolicy{
	MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true,            // Allow pings even when there are no active streams
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
	MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
}

type Server struct {
	ret  *Return
	hub  *Hub
	http *http.Server
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	if err := s.hub.Register(req.Hostname); err != nil {
		klog.V(2).Info(err)
		return nil, err
	}
	return &pb.RegisterReply{}, nil
}

func (s *Server) PullTask(ctx context.Context, req *pb.PullTaskRequest) (*pb.PullTaskReply, error) {
	task := s.hub.PullTask(req.Hostname)
	return &pb.PullTaskReply{Task: &pb.Task{TaskId: task.Id, Command: task.Command}}, nil
}

func (s *Server) CompleteTask(ctx context.Context, req *pb.CompleteTaskRequest) (*pb.CompleteTaskReply, error) {
	if err := s.hub.CompleteTask(req.Hostname, req.TaskId, req.OutPut); err != nil {
		klog.V(2).Info(err)
		return nil, err
	}
	return &pb.CompleteTaskReply{}, nil
}

func (s *Server) Close() {
	if err := s.http.Shutdown(context.Background()); err != nil {
		klog.V(2).Info(err)
	}
}

func (s *Server) NewTask(commands []string) {
	s.hub.NewTask(commands)
}

func NewServer(grpcAddr string, httpAddr string) *Server {
	hub := NewHub(10)
	s := &Server{
		hub:  hub,
		http: InitHttp(httpAddr, hub),
	}
	svr := grpc.NewServer()
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		klog.Fatal("Failed to listen on addr:", grpcAddr)
	}
	pb.RegisterDateAgentServer(svr, s)
	go func() {
		if err := svr.Serve(lis); err != nil {
			klog.Fatal(err)
		}
	}()
	go func() {
		klog.Info("new task")
		s.NewTask([]string{"date"})
	}()
	return s
}
