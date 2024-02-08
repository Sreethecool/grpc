package blog

import (
	"context"
	"log"
	"time"

	"grpc/internal/db"
	"grpc/internal/proto"
)

// Server implements the BlogServiceServer interface.
type Server struct {
	postRepo db.PostRepository
	proto.UnimplementedBlogServiceServer
}

// NewGRPCServer creates a new instance of Server with the given post repository.
func NewGRPCServer(postRepo db.PostRepository) *Server {
	return &Server{
		postRepo: postRepo,
	}
}

func (s *Server) CreatePost(ctx context.Context, req *proto.CreatePostRequest) (*proto.Post, error) {
	log.Println("CreatePost RPC invoked")

	post := &proto.Post{
		Title:   req.GetTitle(),
		Content: req.GetContent(),
		Author:  req.GetAuthor(),
		Tags:    req.GetTags(),
	}

	createdPost, err := s.postRepo.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	return createdPost, nil
}

func (s *Server) ReadPost(ctx context.Context, req *proto.ReadPostRequest) (*proto.Post, error) {
	log.Println("ReadPost RPC invoked")

	post, err := s.postRepo.ReadPost(ctx, req.GetPostId())
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *Server) UpdatePost(ctx context.Context, req *proto.UpdatePostRequest) (*proto.Post, error) {
	log.Println("UpdatePost RPC invoked")

	post := &proto.Post{
		PostId:          req.GetPostId(),
		Title:           req.GetTitle(),
		Content:         req.GetContent(),
		Author:          req.GetAuthor(),
		Tags:            req.GetTags(),
		PublicationDate: time.Now().Format("02-01-2008 15:04:05"),
	}

	updatedPost, err := s.postRepo.UpdatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	return updatedPost, nil
}

func (s *Server) DeletePost(ctx context.Context, req *proto.DeletePostRequest) (*proto.DeletePostResponse, error) {
	log.Println("DeletePost RPC invoked")

	if err := s.postRepo.DeletePost(ctx, req.GetPostId()); err != nil {
		return nil, err
	}

	return &proto.DeletePostResponse{}, nil
}
