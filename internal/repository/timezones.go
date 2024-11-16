package repository

import (
	"context"
	"errors"

	"github.com/justIGreK/Reminders-Timezone/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TimezoneRepo struct {
	collection *mongo.Collection
}

func NewTimezoneRepository(db *mongo.Client) *TimezoneRepo {
	return &TimezoneRepo{
		collection: db.Database(dbname).Collection(transactionCollection),
	}
}

func (r *TimezoneRepo) GetTimezone(ctx context.Context, userID string) (*models.UserTimezone, error) {
	var tz models.UserTimezone
	err := r.collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&tz)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &tz, nil
}
func (r *TimezoneRepo) AddTimezone(ctx context.Context, userTZ models.UserTimezone) (string, error) {
	result, err := r.collection.InsertOne(ctx, userTZ)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
func (r *TimezoneRepo) UpdateTimezone(ctx context.Context, userTZ models.UserTimezone) error {
	updateTZ := bson.M{
		"$set": bson.M{
			"lat":       userTZ.Latitude,
			"long":      userTZ.Latitude,
			"diff_hour": userTZ.DiffHour,
		},
	}
	filter := bson.M{"user_id": userTZ.UserID}
	_, err := r.collection.UpdateOne(ctx, filter, updateTZ)
	if err != nil {
		return err
	}
	return nil
}

func (r *TimezoneRepo) DeleteTimezone(ctx context.Context, userID string) error {
	filter := bson.M{"user_id": userID}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("timezone was not deleted")
	}
	return nil
}
