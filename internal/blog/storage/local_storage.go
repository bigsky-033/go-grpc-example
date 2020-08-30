package storage

import (
	"strconv"
	"sync/atomic"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/bigsky-park/go-grpc-example/api/blog"
)

type LocalStorage struct {
	// Store posts. Key is writer, value is list of posts
	posts   map[string][]*pb.Post
	counter *uint64
}

func NewLocalStorage(posts map[string][]*pb.Post) *LocalStorage {
	return &LocalStorage{
		posts:   posts,
		counter: new(uint64),
	}
}

func (l *LocalStorage) CreateBlog(request *pb.CreatePostRequest) (*pb.Post, error) {
	now := ptypes.TimestampNow()
	post := &pb.Post{
		Id:        strconv.FormatUint(*l.counter, 10),
		Title:     request.Title,
		Writer:    request.Writer,
		Content:   request.Content,
		CreatedAt: now,
		UpdatedAt: now,
	}
	atomic.AddUint64(l.counter, 1)
	l.posts[request.Writer] = append(l.posts[request.Writer], post)
	return post, nil
}

func (l *LocalStorage) GetBlog(request *pb.GetPostRequest) (*pb.Post, error) {
	for _, p := range l.posts[request.Writer] {
		if request.Id == p.Id {
			return p, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "Cannot find post by given request: %v", request)
}

func (l *LocalStorage) ListPosts(request *pb.ListPostsRequest) ([]*pb.Post, error) {
	return l.posts[request.Writer], nil
}
