package handlers

import (
	"encoding/json"
	"net/http"

	_ "github.com/phamaiden/blog-platform-api/internal/db"
	"github.com/phamaiden/blog-platform-api/internal/models"
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

func (bh *BlogHandler) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	posts, err := bh.blogService.ReadAllPosts(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (bh *BlogHandler) PostBlog(w http.ResponseWriter, r *http.Request) {
	var request models.CreatePost
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	post, err := bh.blogService.CreatePost(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func (bh *BlogHandler) PutBlog(w http.ResponseWriter, r *http.Request) {

}

func (bh *BlogHandler) GetBlogById(w http.ResponseWriter, r *http.Request) {

}

func (bh *BlogHandler) DeleteBlog(w http.ResponseWriter, r *http.Request) {}
