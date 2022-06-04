package api

import (
	pb "common/module/proto/posts_service"
	b64 "encoding/base64"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"post/module/domain/model"
	"time"
)

func MapPost(post *model.Post) *pb.Post {
	id := post.Id.Hex()

	links := &pb.Links{
		Comment: "/post/" + id + "/comment",
		Like:    "/post/" + id + "/like",
		Dislike: "/post/" + id + "/dislike",
		User:    "/user/" + post.UserId,
	}

	likesNum, dislikesNum := findNumberOfReactions(post)

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

func findNumberOfReactions(post *model.Post) (int, int) {
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
		UserId:      commentPb.UserId,
		CommentText: commentPb.CommentText,
	}
	return comment
}

func MapNewJobOffer(offerPb *pb.JobOffer) *model.JobOffer {
	duration, _ := time.ParseDuration(offerPb.Duration)

	offer := &model.JobOffer{
		Id:             primitive.NewObjectID(),
		Publisher:      offerPb.Publisher,
		Position:       offerPb.Position,
		JobDescription: offerPb.JobDescription,
		Requirements:   offerPb.Requirements,
		DatePosted:     offerPb.DatePosted.AsTime(),
		Duration:       duration,
	}

	return offer
}

func MapJobOffer(offer *model.JobOffer) *pb.JobOffer {
	id := offer.Id.Hex()

	offerPb := &pb.JobOffer{
		Id:             id,
		Publisher:      offer.Publisher,
		Position:       offer.Position,
		JobDescription: offer.JobDescription,
		Requirements:   offer.Requirements,
		DatePosted:     timestamppb.New(offer.DatePosted),
		Duration:       offer.Duration.String(),
	}

	return offerPb
}
