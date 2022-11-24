package persistence

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	tracer "monitoring/module"
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

func (p PostRepositoryImpl) Get(id primitive.ObjectID, ctx context.Context) (post *model.Post, err error) {
	span := tracer.StartSpanFromContext(ctx, "getRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	filter := bson.M{"_id": id}
	return p.filterOne(filter, ctx)
}

func (p PostRepositoryImpl) GetAll(ctx context.Context) ([]*model.Post, error) {

	span := tracer.StartSpanFromContext(ctx, "getAllRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	filter := bson.D{}
	return p.filter(filter, ctx)
}

func (p PostRepositoryImpl) Create(post *model.Post, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "createPostRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	result, err := p.posts.InsertOne(ctx, post)
	if err != nil {
		return err
	}
	post.Id = result.InsertedID.(primitive.ObjectID)

	return nil
}

func (p PostRepositoryImpl) GetAllByUsername(username string, ctx context.Context) ([]*model.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "getAllByUsernameRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	filter := bson.M{"username": username}
	return p.filter(filter, ctx)
}

func (p PostRepositoryImpl) CreateComment(post *model.Post, comment *model.Comment, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "createCommentService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	comments := append(post.Comments, *comment)

	_, err := p.posts.UpdateOne(ctx, bson.M{"_id": post.Id}, bson.D{
		{"$set", bson.D{{"comments", comments}}},
	},
	)
	if err != nil {
		return err
	}

	return nil
}

func (p PostRepositoryImpl) CreateJobOffer(offer *model.JobOffer, ctx context.Context) (*model.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "createJobOfferRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	result, err := p.jobOffers.InsertOne(ctx, offer)
	if err != nil {
		return nil, err
	}
	offer.Id = result.InsertedID.(primitive.ObjectID)
	createdJobOffer := p.jobOffers.FindOne(ctx, bson.M{"_id": offer.Id})
	var retVal model.JobOffer
	createdJobOffer.Decode(&retVal)
	return &retVal, nil

}

func (p PostRepositoryImpl) GetAllJobOffers(ctx context.Context) ([]*model.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "getAllJobOffersRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	filter := bson.D{}
	return p.filterJobOffers(filter, ctx)
}

func (p PostRepositoryImpl) GetUsersJobOffers(username string, ctx context.Context) ([]*model.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "getUsersJobOffersRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	filter := bson.M{"publisher": username}
	return p.filterJobOffers(filter, ctx)
}

func (p PostRepositoryImpl) UpdateUserPosts(user *model.User, ctx context.Context) error {
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
	span := tracer.StartSpanFromContext(ctx, "updateUserPostsRepository")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	p.posts.UpdateMany(ctx, bson.M{"user_id": user.UserId},
		bson.D{
			{"$set", bson.D{{"username", user.Username}}}})
	return nil
}
func (p PostRepositoryImpl) LikePost(post *model.Post, userId uuid.UUID, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "likePostRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
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

	_, err := p.posts.UpdateOne(ctx, bson.M{"_id": post.Id}, bson.D{
		{"$set", bson.D{{"reactions", reactions}}},
	},
	)
	if err != nil {
		return err
	}

	return nil
}

func (p PostRepositoryImpl) DislikePost(post *model.Post, userId uuid.UUID, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "dislikePostRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
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

	_, err := p.posts.UpdateOne(ctx, bson.M{"_id": post.Id}, bson.D{
		{"$set", bson.D{{"reactions", reactions}}},
	},
	)
	if err != nil {
		return err
	}

	return nil
}
func (p PostRepositoryImpl) CheckLikedStatus(id primitive.ObjectID, userId uuid.UUID, ctx context.Context) (model.ReactionType, error) {
	span := tracer.StartSpanFromContext(ctx, "checkLikedStatusRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	post, err := p.Get(id, ctx)
	if err != nil {
		panic(err)
	}
	for _, reaction := range post.Reactions {
		if reaction.UserId == userId.String() {
			return reaction.Reaction, nil
		}
	}
	return model.Neutral, nil
}

func (p PostRepositoryImpl) filterOne(filter interface{}, ctx context.Context) (post *model.Post, err error) {
	span := tracer.StartSpanFromContext(ctx, "filterOne")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	result := p.posts.FindOne(ctx, filter)
	err = result.Decode(&post)
	return
}

func (p PostRepositoryImpl) filter(filter interface{}, ctx context.Context) ([]*model.Post, error) {
	span := tracer.StartSpanFromContext(ctx, "filter")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	cursor, err := p.posts.Find(ctx, filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	if err != nil {
		return nil, err
	}

	return decode(cursor, ctx)
}

func (p PostRepositoryImpl) filterJobOffers(filter interface{}, ctx context.Context) ([]*model.JobOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "filterJobOffers")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	cursor, err := p.jobOffers.Find(ctx, filter)
	defer cursor.Close(ctx)

	if err != nil {
		return nil, err
	}

	return decodeJobOffers(cursor, ctx)
}

func decodeJobOffers(cursor *mongo.Cursor, ctx context.Context) (offers []*model.JobOffer, err error) {
	span := tracer.StartSpanFromContext(ctx, "decodeJobOffers")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	for cursor.Next(ctx) {
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

func decode(cursor *mongo.Cursor, ctx context.Context) (posts []*model.Post, err error) {
	span := tracer.StartSpanFromContext(ctx, "decode")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	for cursor.Next(ctx) {
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
