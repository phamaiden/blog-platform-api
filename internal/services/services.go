package services

import (
	"context"
	"strconv"

	"github.com/phamaiden/blog-platform-api/internal/db"
	"github.com/phamaiden/blog-platform-api/internal/models"
)

type BlogService interface {
	CreatePost(ctx context.Context, post *models.CreatePost) (*db.Post, error)
	ReadAllPosts(ctx context.Context) (*[]db.Post, error)
	ReadPostById(ctx context.Context, id string) (*db.Post, error)
	ReadPostByFilter(ctx context.Context) error
	UpdatePost(ctx context.Context, id string, post *models.UpdatePost) (*db.Post, error)
	DeleteBlog(ctx context.Context, id string) error
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

func (bs *blogService) ReadPostById(ctx context.Context, id string) (*db.Post, error) {
	postId, err := strconv.Atoi(id)
	if err != nil {
		return &db.Post{}, err
	}
	resp, err := bs.dbQueries.GetPost(ctx, int32(postId))
	return &resp, err
}

func (bs *blogService) ReadPostByFilter(ctx context.Context) error {
	return nil
}

func (bs *blogService) UpdatePost(ctx context.Context, id string, post *models.UpdatePost) (*db.Post, error) {
	postId, err := strconv.Atoi(id)
	if err != nil {
		return &db.Post{}, err
	}
	updatedPost := db.UpdatePostParams{
		ID:       int32(postId),
		Title:    post.Title,
		Content:  post.Content,
		Category: post.Category,
		Tags:     post.Tags,
	}
	resp, err := bs.dbQueries.UpdatePost(ctx, updatedPost)
	if err != nil {
		return &db.Post{}, err
	}
	return &resp, err
}

func (bs *blogService) DeleteBlog(ctx context.Context, id string) error {
	postId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return bs.dbQueries.DeletePost(ctx, int32(postId))
}
