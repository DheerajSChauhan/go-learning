package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo struct {
	coll *mongo.Collection
}

func NewRepo(db *mongo.Database) *Repo {
	return &Repo{
		coll: db.Collection("notes"),
	}
}

func (r *Repo) Create(ctx context.Context, note Note) (Note, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.coll.InsertOne(opCtx, note)
	if err != nil {
		return Note{}, fmt.Errorf("failed to create note: %w", err)
	}

	return note, nil
}

//Repository Pattern, where all database operations are kept inside a Repo type so the rest of the application doesn't need to know MongoDB details.

func(r *Repo) List(ctx context.Context)([]Note, error){
	//list all notes from db
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{} //empty filter to get all documents
	cursor, err := r.coll.Find(opCtx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to list notes: %w", err)
	}

	//cursor must be closed after use, we can use defer to ensure it happens
	//to avoid resource leaks
	defer cursor.Close(opCtx)

	var notes []Note
	if err := cursor.All(opCtx, &notes); err != nil {
		return nil, fmt.Errorf("failed to decode notes: %w", err)
	}
	return notes, nil
}
func (r *Repo) GetByID(ctx context.Context, id primitive.ObjectID) (Note, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	var note Note

	err := r.coll.FindOne(opCtx, filter).Decode(&note)
	if err != nil {
		return Note{}, fmt.Errorf("failed to get note by ID: %w", err)
	}
	return note, nil
}

func (r *Repo) Update(ctx context.Context, id primitive.ObjectID, req UpdateNoteRequest) (Note, error){
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{
			"title": req.Title,
			"content": req.Content,
			"pinned": req.Pinned,
			"updatedAt": time.Now().UTC(),
		},
	}
	after := options.After
	opts := options.FindOneAndUpdate().SetReturnDocument(after)

	var updatedNote Note
	err := r.coll.FindOneAndUpdate(opCtx, filter, update, opts).Decode(&updatedNote)
	if err != nil {
		return Note{}, fmt.Errorf("failed to update note: %w", err)
	}
	return updatedNote, nil
}

func (r *Repo) DeleteByID(ctx context.Context, id primitive.ObjectID) (bool, error){
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}
	res, err := r.coll.DeleteOne(opCtx, filter)
	if err != nil {
		return false, fmt.Errorf("failed to delete note: %w", err)
	}
	if res.DeletedCount == 0 {
		return false, fmt.Errorf("note not found")
	}
	return true, nil
}	