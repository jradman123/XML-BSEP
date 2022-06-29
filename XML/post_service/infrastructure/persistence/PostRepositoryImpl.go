package persistence

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"post/module/domain/model"
	"post/module/domain/repositories"
)

const (
	DATABASE           = "posts_service"
	CollectionPost     = "postsData"
	CollectionJobOffer = "jobOffersData"
)

type PostRepositoryImpl struct {
	posts     *mongo.Collection
	jobOffers *mongo.Collection
}

func NewPostRepositoryImpl(client *mongo.Client) repositories.PostRepository {
	posts := client.Database(DATABASE).Collection(CollectionPost)
	jobOffers := client.Database(DATABASE).Collection(CollectionJobOffer)

	return &PostRepositoryImpl{
		posts:     posts,
		jobOffers: jobOffers,
	}
}

func (p PostRepositoryImpl) Get(id primitive.ObjectID) (post *model.Post, err error) {
	filter := bson.M{"_id": id}
	return p.filterOne(filter)
}

func (p PostRepositoryImpl) GetAll() ([]*model.Post, error) {
	filter := bson.D{}
	return p.filter(filter)
}

func (p PostRepositoryImpl) Create(post *model.Post) error {
	result, err := p.posts.InsertOne(context.TODO(), post)
	if err != nil {
		return err
	}
	post.Id = result.InsertedID.(primitive.ObjectID)

	return nil
}

func (p PostRepositoryImpl) GetAllByUsername(username string) ([]*model.Post, error) {
	filter := bson.M{"username": username}
	return p.filter(filter)
}

func (p PostRepositoryImpl) CreateComment(post *model.Post, comment *model.Comment) error {
	comments := append(post.Comments, *comment)

	_, err := p.posts.UpdateOne(context.TODO(), bson.M{"_id": post.Id}, bson.D{
		{"$set", bson.D{{"comments", comments}}},
	},
	)
	if err != nil {
		return err
	}

	return nil
}

func (p PostRepositoryImpl) CreateJobOffer(offer *model.JobOffer) error {
	result, err := p.jobOffers.InsertOne(context.TODO(), offer)
	if err != nil {
		return err
	}
	offer.Id = result.InsertedID.(primitive.ObjectID)

	return nil
}

func (p PostRepositoryImpl) GetAllJobOffers() ([]*model.JobOffer, error) {
	filter := bson.D{}
	return p.filterJobOffers(filter)
}
func (p PostRepositoryImpl) UpdateUserPosts(user *model.User) error {
	//filter := bson.M{"user_id": user.UserId}
	//posts, err := p.filter(filter)
	//if err != nil {
	//	panic(err)
	//}
	//for _, post := range posts {
	//	_, err = p.posts.UpdateOne(context.TODO(), bson.M{"_id": post.Id}, bson.D{
	//		{"$set", bson.D{{"username", user.Username}}},
	//	})
	//}
	p.posts.UpdateMany(context.TODO(), bson.M{"user_id": user.UserId},
		bson.D{
			{"$set", bson.D{{"username", user.Username}}}})
	return nil
}
func (p PostRepositoryImpl) LikePost(post *model.Post, userId uuid.UUID) error {
	var reactions []model.Reaction

	reactionExists := false
	for _, reaction := range post.Reactions {
		if reaction.UserId != userId.String() {
			reactions = append(reactions, reaction)
		} else {
			if reaction.Reaction != model.LIKED {
				reaction.Reaction = model.LIKED
				reactions = append(reactions, reaction)
			}
			reactionExists = true
		}

	}
	if !reactionExists {
		reaction := model.Reaction{
			UserId:   userId.String(),
			Reaction: model.LIKED,
		}
		reactions = append(reactions, reaction)
	}

	_, err := p.posts.UpdateOne(context.TODO(), bson.M{"_id": post.Id}, bson.D{
		{"$set", bson.D{{"reactions", reactions}}},
	},
	)
	if err != nil {
		return err
	}

	return nil
}

func (p PostRepositoryImpl) DislikePost(post *model.Post, userId uuid.UUID) error {
	var reactions []model.Reaction

	reactionExists := false
	for _, reaction := range post.Reactions {
		if reaction.UserId != userId.String() {
			reactions = append(reactions, reaction)
		} else {
			if reaction.Reaction != model.DISLIKED {
				reaction.Reaction = model.DISLIKED
				reactions = append(reactions, reaction)
			}
			reactionExists = true
		}

	}
	if !reactionExists {
		reaction := model.Reaction{
			UserId:   userId.String(),
			Reaction: model.DISLIKED,
		}
		reactions = append(reactions, reaction)
	}

	_, err := p.posts.UpdateOne(context.TODO(), bson.M{"_id": post.Id}, bson.D{
		{"$set", bson.D{{"reactions", reactions}}},
	},
	)
	if err != nil {
		return err
	}

	return nil
}

func (p PostRepositoryImpl) filterOne(filter interface{}) (post *model.Post, err error) {
	result := p.posts.FindOne(context.TODO(), filter)
	err = result.Decode(&post)
	return
}

func (p PostRepositoryImpl) filter(filter interface{}) ([]*model.Post, error) {
	cursor, err := p.posts.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, err
	}

	return decode(cursor)
}

func (p PostRepositoryImpl) filterJobOffers(filter bson.D) ([]*model.JobOffer, error) {
	cursor, err := p.jobOffers.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeJobOffers(cursor)
}

func decodeJobOffers(cursor *mongo.Cursor) (offers []*model.JobOffer, err error) {
	for cursor.Next(context.TODO()) {
		var offer model.JobOffer
		err = cursor.Decode(&offer)
		if err != nil {
			return
		}
		offers = append(offers, &offer)
	}
	err = cursor.Err()
	return
}

func decode(cursor *mongo.Cursor) (posts []*model.Post, err error) {
	for cursor.Next(context.TODO()) {
		var post model.Post
		err = cursor.Decode(&post)
		if err != nil {
			return
		}
		posts = append(posts, &post)
	}
	err = cursor.Err()
	return
}
