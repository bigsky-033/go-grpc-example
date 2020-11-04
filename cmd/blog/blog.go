package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"

	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"

	pb "github.com/bigsky-park/go-grpc-example/api/blog"
	"github.com/bigsky-park/go-grpc-example/internal/blog"
	"github.com/bigsky-park/go-grpc-example/internal/blog/storage"
)

var (
	grpcPort            = pflag.Int32("grpc_port", 18080, "Grpc service port. Default: 8080.")
	httpHealthCheckPort = pflag.Int32("http_health_check_port", 8081, "Health check port. Default: 8081.")
)

func main() {
	klog.InitFlags(nil)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.WriteHeader(200)
		}
	})
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", *httpHealthCheckPort), nil); err != nil {
			klog.Fatalf("Failed to start HTTP health check. Details: %v", err)
		}
	}()

	posts := map[string][]*pb.Post{}
	blogService := blog.NewBlogService(storage.NewLocalStorage(posts))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		klog.Fatalf("Failed to listen grpc port. Details: %v", err)
	} else {
		klog.Infof("Listen grpc port: %d", *grpcPort)
	}

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, blogService)

	if err := s.Serve(lis); err != nil {
		klog.Fatalf("Failed to serve grpc service. Details: %v", err)
	}
}
