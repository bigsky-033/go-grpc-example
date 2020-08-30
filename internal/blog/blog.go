package blog

import (
	"context"

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
	post, err := b.storage.CreateBlog(in)
	if err != nil {
		return nil, err
	}
	return &pb.CreatePostResponse{
		Post: post,
	}, nil
}

func (b *BlogService) GetPost(ctx context.Context, in *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	post, err := b.storage.GetBlog(in)
	if err != nil {
		return nil, err
	}
	return &pb.GetPostResponse{
		Post: post,
	}, nil
}

func (b *BlogService) ListPosts(ctx context.Context, in *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	posts, err := b.storage.ListPosts(in)
	if err != nil {
		return nil, err
	}
	return &pb.ListPostsResponse{
		Posts: posts,
	}, nil
}
