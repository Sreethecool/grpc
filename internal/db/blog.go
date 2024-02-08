package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"grpc/internal/proto"
)

var (
	ErrPostNotFound = errors.New("post not found")
)

type PostRepository interface {
	CreatePost(ctx context.Context, post *proto.Post) (*proto.Post, error)
	ReadPost(ctx context.Context, postID string) (*proto.Post, error)
	UpdatePost(ctx context.Context, post *proto.Post) (*proto.Post, error)
	DeletePost(ctx context.Context, postID string) error
}

type InMemoryPostRepository struct {
	posts map[string]*proto.Post
}

func NewPostRepository() *InMemoryPostRepository {
	return &InMemoryPostRepository{
		posts: make(map[string]*proto.Post),
	}
}

func (repo *InMemoryPostRepository) CreatePost(ctx context.Context, post *proto.Post) (*proto.Post, error) {
	post.PostId = fmt.Sprintf("P%d", time.Now().UnixNano())
	post.PublicationDate = time.Now().Format("02-01-2008 15:04:05")
	repo.posts[post.PostId] = post
	return post, nil
}

func (repo *InMemoryPostRepository) ReadPost(ctx context.Context, postID string) (*proto.Post, error) {
	post, ok := repo.posts[postID]
	if !ok {
		return nil, ErrPostNotFound
	}
	return post, nil
}

func (repo *InMemoryPostRepository) UpdatePost(ctx context.Context, post *proto.Post) (*proto.Post, error) {
	_, err := repo.ReadPost(ctx, post.PostId)
	if err != nil {
		return nil, err
	}
	repo.posts[post.PostId] = post
	return post, nil
}

func (repo *InMemoryPostRepository) DeletePost(ctx context.Context, postID string) error {
	_, ok := repo.posts[postID]
	if !ok {
		return ErrPostNotFound
	}
	delete(repo.posts, postID)
	return nil
}
