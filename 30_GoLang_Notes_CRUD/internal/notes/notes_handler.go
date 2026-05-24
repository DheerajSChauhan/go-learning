package notes

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	repo *Repo
}

func NewHandler(repo *Repo) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) CreateNote(c *gin.Context){
	//per request context from gin
	var req CreateNoteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	now := time.Now().UTC()
	note := Note{
		ID : primitive.NewObjectID(),
		Title: req.Title,
		Content: req.Content,
		Pinned: req.Pinned,
		CreatedAt: now,
		UpdatedAt: now,
	}
	created, err := h.repo.Create(c.Request.Context(), note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *Handler) ListNotes(c *gin.Context){
	noteList, err := h.repo.List(c.Request.Context())
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": " failed to list notes: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"notes" : noteList,
	})
}

func (h *Handler) GetNoteByID(c *gin.Context){
	idStr := c.Param("id")

	//24 character hex string to ObjectID
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Invalid ID",
		})
		return
	}

	note, err := h.repo.GetByID(c.Request.Context(), objID)
	if err != nil{
		if errors.Is(err, mongo.ErrNoDocuments){
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Note not found",
			})
			return  
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get note",
		})
		return
	}
	c.JSON(http.StatusOK, note)
}

func (h *Handler) UpdateNote(c *gin.Context){
	idStr := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}
	var req UpdateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	updatedNote, err := h.repo.Update(c.Request.Context(), objID, req)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Note not found",
			})
			return
		}
	
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, updatedNote)
}

func (h *Handler) DeleteNoteByID(c *gin.Context){
	idStr := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	deleted, err := h.repo.DeleteByID(c.Request.Context(), objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Note not found",
			})
			return
		}
	
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Note not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
		"message": "Note deleted successfully",
	})

}