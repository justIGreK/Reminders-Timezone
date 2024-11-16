package models

type UserTimezone struct{
	UserID string `bson:"user_id"`
	Latitude float64 `bson:"lat"`
	Longitude float64 `bson:"long"`
	DiffHour int `bson:"diff_hour"`
}