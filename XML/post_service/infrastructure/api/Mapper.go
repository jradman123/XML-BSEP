package api

import (
	pb "common/module/proto/posts_service"
	events "common/module/saga/user_events"
	"context"
	b64 "encoding/base64"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"image"
	"image/jpeg"
	"log"
	tracer "monitoring/module"
	"os"
	"post/module/domain/model"
	"strings"
	"time"
)

func MapNewPost(postPb *pb.Post, user *model.User, ctx context.Context) *model.Post {
	span := tracer.StartSpanFromContext(ctx, "mapNewPost")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	post := &model.Post{
		Id:         primitive.NewObjectID(),
		Username:   user.Username,
		UserId:     user.UserId,
		PostText:   postPb.PostText,
		DatePosted: time.Now(),
		IsDeleted:  false,
	}
	base64toJpg(postPb.ImagePaths, ctx)
	post.ImagePaths = convertBase64ToByte(postPb.ImagePaths, ctx)
	return post
}
func base64toJpg(img string, ctx context.Context) {
	span := tracer.StartSpanFromContext(ctx, "base64ToJpg")
	defer span.Finish()

	data := img[strings.IndexByte(img, ',')+1:]
	reader := b64.NewDecoder(b64.StdEncoding, strings.NewReader(data))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	fmt.Println("base64toJpg", bounds, formatString)

	//Encode from image format to writer
	pngFilename := "test"
	f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = jpeg.Encode(f, m, &jpeg.Options{Quality: 75})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Jpg file", pngFilename, "created")
	err = f.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

}

func convertBase64ToByte(image string, ctx context.Context) []byte {
	span := tracer.StartSpanFromContext(ctx, "convertBase64ToByte")
	defer span.Finish()

	fmt.Println("convertBase64ToByte")
	imageDec := image[strings.IndexByte(image, ',')+1:]
	dec, err := b64.StdEncoding.DecodeString(imageDec)
	if err != nil {
		panic(err)
	}
	return dec

}
func MapNewComment(commentPb *pb.Comment, ctx context.Context) *model.Comment {
	span := tracer.StartSpanFromContext(ctx, "mapNewComment")
	defer span.Finish()

	comment := &model.Comment{
		Id:          primitive.NewObjectID(),
		Username:    commentPb.Username,
		CommentText: commentPb.CommentText,
	}
	return comment
}

func MapNewJobOffer(offerPb *pb.JobOffer, ctx context.Context) *model.JobOffer {
	span := tracer.StartSpanFromContext(ctx, "mapNewJobOffer")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	offer := &model.JobOffer{
		Id:             primitive.NewObjectID(),
		Publisher:      offerPb.Publisher,
		Position:       offerPb.Position,
		JobDescription: offerPb.JobDescription,
		Requirements:   offerPb.Requirements,
		DatePosted:     mapToDate(offerPb.DatePosted, ctx),
		Duration:       mapToDate(offerPb.Duration, ctx),
	}

	return offer
}
func MapNewUser(command *events.UserCommand, ctx context.Context) *model.User {
	span := tracer.StartSpanFromContext(ctx, "mapNewUser")
	defer span.Finish()

	user := &model.User{
		Id:        primitive.NewObjectID(),
		UserId:    command.User.UserId,
		Username:  command.User.Username,
		FirstName: command.User.FirstName,
		LastName:  command.User.LastName,
		Email:     command.User.Email,
		Active:    false,
	}
	return user
}
func MapUserReply(user *model.User, replyType events.UserReplyType, ctx context.Context) (reply *events.UserReply) {
	span := tracer.StartSpanFromContext(ctx, "mapUserReply")
	defer span.Finish()
	reply = &events.UserReply{
		Type: replyType,
		PostUser: events.PostUser{
			Id:        user.Id,
			UserId:    user.UserId,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		},
	}
	return reply
}
func mapToDate(birth string, ctx context.Context) time.Time {
	span := tracer.StartSpanFromContext(ctx, "mapToDate")
	defer span.Finish()

	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth, _ := time.Parse(layout, birth)
	return dateOfBirth

}

func MapJobOfferReply(offer *model.JobOffer, ctx context.Context) *pb.JobOffer {
	span := tracer.StartSpanFromContext(ctx, "mapJobOfferReply")
	defer span.Finish()

	id := offer.Id.Hex()

	offerPb := &pb.JobOffer{
		Id:             id,
		Publisher:      offer.Publisher,
		Position:       offer.Position,
		JobDescription: offer.JobDescription,
		Requirements:   offer.Requirements,
		DatePosted:     offer.DatePosted.String(),
		Duration:       offer.Duration.String(),
	}

	return offerPb
}

func MapPostReply(post *model.Post, ctx context.Context) *pb.Post {
	span := tracer.StartSpanFromContext(ctx, "mapPostReply")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	id := post.Id.Hex()

	links := &pb.Links{
		Comment: "/post/" + id + "/comment",
		Like:    "/post/" + id + "/like",
		Dislike: "/post/" + id + "/dislike",
		User:    "/user/" + post.UserId.String(),
	}

	likesNum, dislikesNum := FindNumberOfReactions(post, ctx)

	postPb := &pb.Post{
		Id:             id,
		Username:       post.Username,
		PostText:       post.PostText,
		DatePosted:     post.DatePosted.String(),
		Links:          links,
		LikesNumber:    int32(likesNum),
		DislikesNumber: int32(dislikesNum),
		CommentsNumber: int32(len(post.Comments)),
	}
	postPb.ImagePaths = convertByteToBase64(post.ImagePaths, ctx)

	return postPb
}

func FindNumberOfReactions(post *model.Post, ctx context.Context) (int, int) {
	span := tracer.StartSpanFromContext(ctx, "findNumberOfReactions")
	defer span.Finish()

	likesNum := 0
	dislikesNum := 0

	for _, reaction := range post.Reactions {
		if reaction.Reaction == model.LIKED {
			likesNum++
		} else if reaction.Reaction == model.DISLIKED {
			dislikesNum++
		}
	}
	return likesNum, dislikesNum
}
func convertByteToBase64(image []byte, ctx context.Context) string {
	span := tracer.StartSpanFromContext(ctx, "convertByteToBase64")
	defer span.Finish()

	imageEnc := b64.StdEncoding.EncodeToString(image)
	return imageEnc
}
func MapUserCommentsForPost(user *model.User, commentText string, ctx context.Context) *pb.Comment {

	span := tracer.StartSpanFromContext(ctx, "mapUserCommentsForPost")
	defer span.Finish()

	commentPb := &pb.Comment{
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		CommentText: commentText,
	}

	return commentPb
}
