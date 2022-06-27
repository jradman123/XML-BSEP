package api

import (
	pb "common/module/proto/posts_service"
	events "common/module/saga/user_events"
	b64 "encoding/base64"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post/module/domain/model"
	"time"
)

func MapNewPost(postPb *pb.Post) *model.Post {
	post := &model.Post{
		Id:         primitive.NewObjectID(),
		UserId:     postPb.UserId,
		PostText:   postPb.PostText,
		DatePosted: time.Now(),
	}
	post.ImagePaths = convertBase64ToByte(postPb.ImagePaths)

	return post
}
func convertBase64ToByte(images []string) [][]byte {
	var decodedImages [][]byte
	for _, image := range images {
		fmt.Println(image)
		imageDec, _ := b64.StdEncoding.DecodeString(image)
		fmt.Println(string(imageDec))
		decodedImages = append(decodedImages, imageDec)
	}
	return decodedImages
}
func MapNewComment(commentPb *pb.Comment) *model.Comment {
	comment := &model.Comment{
		Id:          primitive.NewObjectID(),
		Username:    commentPb.Username,
		CommentText: commentPb.CommentText,
	}
	return comment
}

func MapNewJobOffer(offerPb *pb.JobOffer) *model.JobOffer {

	offer := &model.JobOffer{
		Id:             primitive.NewObjectID(),
		Publisher:      offerPb.Publisher,
		Position:       offerPb.Position,
		JobDescription: offerPb.JobDescription,
		Requirements:   offerPb.Requirements,
		DatePosted:     mapToDate(offerPb.DatePosted),
		Duration:       mapToDate(offerPb.Duration),
	}

	return offer
}
func MapNewUser(command *events.UserCommand) *model.User {
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
func MapUserReply(user *model.User, replyType events.UserReplyType) (reply *events.UserReply) {
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
func mapToDate(birth string) time.Time {
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth, _ := time.Parse(layout, birth)
	return dateOfBirth

}

func MapJobOfferReply(offer *model.JobOffer) *pb.JobOffer {
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

func MapPostReply(post *model.Post) *pb.Post {
	id := post.Id.Hex()

	links := &pb.Links{
		Comment: "/post/" + id + "/comment",
		Like:    "/post/" + id + "/like",
		Dislike: "/post/" + id + "/dislike",
		User:    "/user/" + post.UserId,
	}

	likesNum, dislikesNum := FindNumberOfReactions(post)

	postPb := &pb.Post{
		Id:             id,
		UserId:         post.UserId,
		PostText:       post.PostText,
		DatePosted:     post.DatePosted.String(),
		Links:          links,
		LikesNumber:    int32(likesNum),
		DislikesNumber: int32(dislikesNum),
		CommentsNumber: int32(len(post.Comments)),
	}
	postPb.ImagePaths = convertByteToBase64(post.ImagePaths)

	return postPb
}

func FindNumberOfReactions(post *model.Post) (int, int) {
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
func convertByteToBase64(images [][]byte) []string {
	var encodedImages []string
	for _, image := range images {
		fmt.Println(image)
		imageEnc := b64.StdEncoding.EncodeToString(image)
		fmt.Println(string(imageEnc))
		encodedImages = append(encodedImages, imageEnc)
	}
	return encodedImages
}
