package notes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(r *gin.Engine, db *mongo.Database) {
	//create repo and handler once at start

	repo := NewRepo(db)
	h := NewHandler(repo)

	notesGroup := r.Group("/notes")
	{
		notesGroup.POST("", h.CreateNote)
		//other routes like GET, PUT, DELETE can be added here
		notesGroup.GET("", h.ListNotes)
		notesGroup.GET("/:id", h.GetNoteByID)
		notesGroup.PUT("/:id", h.UpdateNote)
		notesGroup.DELETE("/:id", h.DeleteNoteByID)

	}
}

