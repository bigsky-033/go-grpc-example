package main

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"

	pb "github.com/bigsky-park/go-grpc-example/api/blog"
)

var (
	endpoint = pflag.String("endpoint", "localhost", "Grpc service endpoint. Default: localhost.")
	grpcPort = pflag.Int32("grpc_port", 8080, "Grpc service port. Default: 8080.")
)

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *endpoint, *grpcPort), grpc.WithInsecure())
	if err != nil {
		klog.Fatalf("Failed to connect. Details: %v", err)
	}
	defer conn.Close()
	client := pb.NewBlogServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	writer := "bigsky-park"
	for i := 0; i < 10; i++ {
		// Create post
		createPostResp, err := client.CreatePost(ctx, &pb.CreatePostRequest{
			Writer:  writer,
			Title:   fmt.Sprintf("post-%d", i),
			Content: fmt.Sprintf("This is post %d.", i),
		})
		if err != nil {
			klog.Fatalf("Failed to create post. Details: %v", err)
		}
		klog.Infof("Created Post: %v", createPostResp.Post)
	}

	// List Posts
	listPostsResp, err := client.ListPosts(ctx, &pb.ListPostsRequest{
		Writer: writer,
	})
	if err != nil {
		klog.Fatalf("Failed to list post. Details: %v", err)
	}
	klog.Infof("List Posts: %v", listPostsResp.Posts)

	// Get Post
	for _, i := range []int{1, 999} {
		getPostResp, err := client.GetPost(ctx, &pb.GetPostRequest{
			Id:     fmt.Sprintf("%d", i),
			Writer: writer,
		})
		if err != nil {
			// Get post with Id: 999 will fail
			klog.Warningf("Failed to get post. Details: %v", err)
		}
		klog.Infof("Get Post: %v", getPostResp.Post)
	}
}
