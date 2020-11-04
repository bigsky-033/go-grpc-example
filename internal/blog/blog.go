package blog

import (
	"context"

	"google.golang.org/grpc/peer"
	"k8s.io/klog/v2"

	pb "github.com/bigsky-park/go-grpc-example/api/blog"
)

type (
	BlogService struct {
		storage Storage
	}
	Storage interface {
		CreateBlog(request *pb.CreatePostRequest) (*pb.Post, error)
		GetBlog(request *pb.GetPostRequest) (*pb.Post, error)
		ListPosts(request *pb.ListPostsRequest) ([]*pb.Post, error)
	}
)

func NewBlogService(storage Storage) *BlogService {
	return &BlogService{
		storage: storage,
	}
}

func (b *BlogService) CreatePost(ctx context.Context, in *pb.CreatePostRequest) (resp *pb.CreatePostResponse, err error) {
	p, _ := peer.FromContext(ctx)
	klog.Infof("CreatePostRequest from %s, payload: (%#v)", p.Addr.String(), in)

	post, err := b.storage.CreateBlog(in)
	if err != nil {
		return nil, err
	}
	return &pb.CreatePostResponse{
		Post: post,
	}, nil
}

func (b *BlogService) GetPost(ctx context.Context, in *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	p, _ := peer.FromContext(ctx)
	klog.Infof("GetPost from %s, payload: (%#v)", p.Addr.String(), in)

	post, err := b.storage.GetBlog(in)
	if err != nil {
		return nil, err
	}
	return &pb.GetPostResponse{
		Post: post,
	}, nil
}

func (b *BlogService) ListPosts(ctx context.Context, in *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	p, _ := peer.FromContext(ctx)
	klog.Infof("ListPosts from %s, payload: (%#v)", p.Addr.String(), in)

	posts, err := b.storage.ListPosts(in)
	if err != nil {
		return nil, err
	}
	return &pb.ListPostsResponse{
		Posts: posts,
	}, nil
}
