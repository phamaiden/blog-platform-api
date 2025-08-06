package handlers

import (
	"net/http"

	"github.com/phamaiden/blog-platform-api/internal/services"
)

type BlogHandler struct {
	blogService services.BlogService
}

func NewBlogHandler(bs services.BlogService) *BlogHandler {
	return &BlogHandler{
		blogService: bs,
	}
}

func (bh *BlogHandler) PostBlog(w http.ResponseWriter, r *http.Request) {

}

func (bh *BlogHandler) PutBlog(w http.ResponseWriter, r *http.Request) {

}

func (bh *BlogHandler) GetBlog(w http.ResponseWriter, r *http.Request) {}

func (bh *BlogHandler) DeleteBlog(w http.ResponseWriter, r *http.Request)
