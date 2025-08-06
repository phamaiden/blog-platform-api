package services

import (
	"github.com/jackc/pgx/v5"
)

type BlogService interface {
	PostBlog()
	PutBlog()
	GetBlog()
	DeleteBlog()
}

type blogService struct {
	dbConn *pgx.Conn
}

func NewBlogService(db *pgx.Conn) *blogService {
	return &blogService{
		dbConn: db,
	}
}

func (bs *blogService) PostBlog() {

}

func (bs *blogService) PutBlog() {

}

func (bs *blogService) GetBlog() {

}

func (bs *blogService) DeleteBlog() {

}
