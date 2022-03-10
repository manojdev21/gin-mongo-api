package services

import (
	"context"
	"errors"
	"gin-mongo-api/configs"
	"gin-mongo-api/exceptions"
	"gin-mongo-api/logger"
	"gin-mongo-api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	USERS = "users"
)

func GetAllUsers() ([]models.User, int64, error) {
	sort := bson.M{"createdDate": -1}
	limit := configs.EnvPageLimit()

	filter := bson.M{}
	opts := &options.FindOptions{Sort: sort, Limit: &limit}

	cur, err := configs.GetCollection(configs.Client, USERS).Find(context.Background(), filter, opts)
	if err != nil {
		return nil, 0, err
	}

	defer cur.Close(context.Background())

	var users []models.User
	for cur.Next(context.Background()) {
		var user models.User

		err := cur.Decode(&user)
		if err != nil {
			return nil, 0, err
		}

		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		return nil, 0, err
	}

	if drawCount := int64(len(users)); drawCount == 0 {
		return nil, 0, errors.New(exceptions.NO_USER_FOUND)
	}
	totalRecords, _ := configs.GetCollection(configs.Client, USERS).CountDocuments(context.Background(), filter)

	logger.InfoLogger.Println("Calculated draw count and total records")

	return users, totalRecords, nil
}

func GetUser(id primitive.ObjectID) (models.User, error) {
	found, user := getUser(id)
	if !found {
		return user, errors.New(exceptions.INVALID_ID)
	}
	return user, nil
}

func CreateUser(user models.User) (models.User, error) {
	duplicate, _ := checkEmailDuplicate(user.Email)

	if duplicate {
		return user, errors.New(exceptions.DUPLICATE_EMAIL)
	}
	user.CreatedDate = primitive.NewDateTimeFromTime(time.Now())
	logger.InfoLogger.Println("Set created date to registed user")

	result, err := configs.GetCollection(configs.Client, USERS).InsertOne(context.Background(), user)
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, err
}

func UpdateUser(user models.User) error {
	if found, _ := getUser(user.ID); !found {
		return errors.New(exceptions.INVALID_ID)
	}

	if duplicate, existUserID := checkEmailDuplicate(user.Email); duplicate && (existUserID != user.ID) {
		return errors.New(exceptions.DUPLICATE_EMAIL)
	}

	_, err := configs.GetCollection(configs.Client, USERS).UpdateByID(context.Background(), user.ID, bson.M{"$set": user})
	return err
}

func DeleteUser(id primitive.ObjectID) error {
	found, _ := getUser(id)

	if !found {
		return errors.New(exceptions.INVALID_ID)
	}

	filter := bson.M{"_id": id}

	_, err := configs.GetCollection(configs.Client, USERS).DeleteOne(context.Background(), filter)

	return err
}

func getUser(id primitive.ObjectID) (bool, models.User) {
	var user models.User

	filter := bson.M{"_id": id}
	configs.GetCollection(configs.Client, USERS).FindOne(context.Background(), filter).Decode(&user)

	logger.InfoLogger.Println("Getting user by ID")

	if !user.ID.IsZero() && user.ID == id {
		return true, user
	}
	return false, user
}

func checkEmailDuplicate(email string) (bool, primitive.ObjectID) {
	var user models.User

	filter := bson.M{"email": email}
	configs.GetCollection(configs.Client, USERS).FindOne(context.Background(), filter).Decode(&user)

	logger.InfoLogger.Println("Checking email duplication")

	if user.ID.IsZero() {
		return false, primitive.NilObjectID
	}

	return true, user.ID
}
