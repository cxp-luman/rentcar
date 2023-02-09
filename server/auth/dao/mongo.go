package dao

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo define a mongo dao
type Mongo struct {
	col *mongo.Collection
}

// NewMongo creates a new mongo dao
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("account"),
	}
}

// ResolveOpenId resolve an account id from open id.
func (m *Mongo) ResolveOpenId(ctx context.Context, opneId string) (string, error) {
	res := m.col.FindOneAndUpdate(ctx, 
		bson.M{
			"open_id": opneId,
		},
		bson.M{
			"$set": bson.M{
				"open_id": opneId,
			},
		}, options.FindOneAndUpdate().
			SetUpsert(true).
			SetReturnDocument(options.After)
	)
	if err := res.Err(); err != nil {
		return "", fmt.Errorf("cannot FindOneAndUpdate: %v", err)
	}
	var row struct {
		ID primitive.ObjectID `bson:"_id"`
		OpenID string `bson:"open_id"`
	}
	err := res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("cannot Decode: %v", err)
	}
	return row.ID.Hex(), nil
}
