package api

import (
	pb "common/module/proto/posts_service"
	events "common/module/saga/user_events"
	b64 "encoding/base64"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"image"
	"image/jpeg"
	"log"
	"os"
	"post/module/domain/model"
	"strings"
	"time"
)

func MapNewPost(postPb *pb.Post, user *model.User) *model.Post {
	post := &model.Post{
		Id:         primitive.NewObjectID(),
		Username:   user.Username,
		UserId:     user.UserId,
		PostText:   postPb.PostText,
		DatePosted: time.Now(),
		IsDeleted:  false,
	}
	base64toJpg(postPb.ImagePaths)
	post.ImagePaths = convertBase64ToByte(postPb.ImagePaths)
	return post
}
func base64toJpg(img string) {
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

func convertBase64ToByte(image string) []byte {

	fmt.Println("convertBase64ToByte")
	imageDec := image[strings.IndexByte(image, ',')+1:]
	dec, err := b64.StdEncoding.DecodeString(imageDec)
	if err != nil {
		panic(err)
	}
	return dec

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
		User:    "/user/" + post.UserId.String(),
	}

	likesNum, dislikesNum := FindNumberOfReactions(post)

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
func convertByteToBase64(image []byte) string {
	imageEnc := b64.StdEncoding.EncodeToString(image)
	return imageEnc
}
func MapUserCommentsForPost(user *model.User, commentText string) *pb.Comment {

	commentPb := &pb.Comment{
		Username:    user.Username,
		CommentText: commentText,
	}

	return commentPb
}
