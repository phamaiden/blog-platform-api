package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)
	postId, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "missing post id parameter",
		})
		return
	}

	// read in updated post from http request
	var request models.UpdatePost
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	// check if post exists
	_, err = bh.blogService.ReadPostById(r.Context(), postId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "post not found",
		})
		return
	}

	// update post
	updated, err := bh.blogService.UpdatePost(r.Context(), postId, &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}

func (bh *BlogHandler) GetBlogById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "missing post id parameter",
		})
		return
	}

	post, err := bh.blogService.ReadPostById(r.Context(), postId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "post not found",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func (bh *BlogHandler) GetBlogByFilter(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("term")
	posts, err := bh.blogService.ReadPostByFilter(r.Context(), term)
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

func (bh *BlogHandler) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "missing post id parameter",
		})
		return
	}

	_, err := bh.blogService.ReadPostById(r.Context(), postId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = bh.blogService.DeleteBlog(r.Context(), postId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
