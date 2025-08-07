package services

import (
	"context"

	"github.com/phamaiden/blog-platform-api/internal/db"
	"github.com/phamaiden/blog-platform-api/internal/models"
)

type BlogService interface {
	CreatePost(ctx context.Context, post *models.CreatePost) (*db.Post, error)
	ReadAllPosts(ctx context.Context) (*[]db.Post, error)
	ReadPostById(ctx context.Context) error
	ReadPostByFilter(ctx context.Context) error
	UpdatePost(ctx context.Context) error
	DeleteBlog(ctx context.Context) error
}

type blogService struct {
	dbQueries *db.Queries
}

func NewBlogService(q *db.Queries) *blogService {
	return &blogService{
		dbQueries: q,
	}
}

func (bs *blogService) CreatePost(ctx context.Context, post *models.CreatePost) (*db.Post, error) {
	newPost := db.CreatePostParams{
		Title:    post.Title,
		Content:  post.Content,
		Category: post.Category,
		Tags:     post.Tags,
	}

	resp, err := bs.dbQueries.CreatePost(ctx, newPost)
	return &resp, err
}

func (bs *blogService) ReadAllPosts(ctx context.Context) (*[]db.Post, error) {
	resp, err := bs.dbQueries.ListPosts(ctx)
	return &resp, err
}

func (bs *blogService) ReadPostById(ctx context.Context) error {
	return nil
}

func (bs *blogService) ReadPostByFilter(ctx context.Context) error {
	return nil
}

func (bs *blogService) UpdatePost(ctx context.Context) error {
	return nil
}

func (bs *blogService) DeleteBlog(ctx context.Context) error {
	return nil
}
