package handlers

import (
	common "common/module"
	pb "common/module/proto/posts_service"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
	"strings"
)

func (p PostHandler) sanitizeJobOffer(request *pb.CreateJobOfferRequest) *pb.CreateJobOfferRequest {
	policy := bluemonday.UGCPolicy()
	request.JobOffer.Publisher = strings.TrimSpace(policy.Sanitize(request.JobOffer.Publisher))
	request.JobOffer.Position = strings.TrimSpace(policy.Sanitize(request.JobOffer.Position))
	request.JobOffer.JobDescription = strings.TrimSpace(policy.Sanitize(request.JobOffer.JobDescription))
	for i := range request.JobOffer.Requirements {
		request.JobOffer.Requirements[i] = strings.TrimSpace(policy.Sanitize(request.JobOffer.Requirements[i]))
	}
	request.JobOffer.DatePosted = strings.TrimSpace(policy.Sanitize(request.JobOffer.DatePosted))
	request.JobOffer.Duration = strings.TrimSpace(policy.Sanitize(request.JobOffer.Duration))

	p2 := common.BadText(request.JobOffer.Publisher)
	p3 := common.BadText(request.JobOffer.Position)
	p4 := common.BadText(request.JobOffer.JobDescription)
	p5 := common.BadTexts(request.JobOffer.Requirements)
	p6 := common.BadDate(request.JobOffer.DatePosted)
	p7 := common.BadDate(request.JobOffer.Duration)

	if request.JobOffer.Publisher == "" || request.JobOffer.Position == "" || request.JobOffer.JobDescription == "" ||
		request.JobOffer.DatePosted == "" || request.JobOffer.Duration == "" {
		p.logError.Logger.Errorf("ERR:XSS")
	} else if p2 || p3 || p4 || p5 || p6 || p7 {
		p.logError.Logger.Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
	} else {
		p.logInfo.Logger.Infof("INFO:Handling ShareJobOffer/CreateJobOffer")
	}
	return request
}

func (p PostHandler) sanitizePost(request *pb.CreatePostRequest, userNameCtx string) *pb.CreatePostRequest {
	policy := bluemonday.UGCPolicy()
	request.Post.Username = strings.TrimSpace(policy.Sanitize(request.Post.Username))
	request.Post.PostText = strings.TrimSpace(policy.Sanitize(request.Post.PostText))
	//for i := range request.Post.ImagePaths {
	//	request.Post.ImagePaths[i] = strings.TrimSpace(policy.Sanitize(request.Post.ImagePaths[i]))
	//}
	request.Post.DatePosted = strings.TrimSpace(policy.Sanitize(request.Post.DatePosted))

	p1 := common.BadId(request.Post.Username)
	p2 := common.BadText(request.Post.PostText)
	p3 := common.BadDate(request.Post.DatePosted)
	//p4 := common.BadPaths(request.Post.ImagePaths)

	if request.Post.Username == "" || request.Post.PostText == "" || request.Post.DatePosted == "" {
		p.logError.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Errorf("ERR:XSS")
	} else if p1 || p2 || p3 {
		p.logError.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
	} else {
		p.logInfo.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Infof("INFO:Handling Create post")
	}
	return request
}

func (p PostHandler) sanitizeComment(request *pb.CreateCommentRequest, userNameCtx string) *pb.CreateCommentRequest {
	policy := bluemonday.UGCPolicy()
	request.PostId = strings.TrimSpace(policy.Sanitize(request.PostId))

	request.Comment.Username = strings.TrimSpace(policy.Sanitize(request.Comment.Username))
	request.Comment.FirstName = strings.TrimSpace(policy.Sanitize(request.Comment.FirstName))
	request.Comment.LastName = strings.TrimSpace(policy.Sanitize(request.Comment.LastName))
	request.Comment.CommentText = strings.TrimSpace(policy.Sanitize(request.Comment.CommentText))

	p1 := common.BadId(request.PostId)

	p3 := common.BadUsername(request.Comment.Username)
	p4 := common.BadName(request.Comment.FirstName)
	p5 := common.BadName(request.Comment.LastName)
	p6 := common.BadText(request.Comment.CommentText)

	if request.PostId == "" || request.Comment.Username == "" || request.Comment.FirstName == "" ||
		request.Comment.LastName == "" || request.Comment.CommentText == "" {
		p.logError.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Errorf("ERR:XSS")
	} else if p1 || p3 || p4 || p5 || p6 {
		p.logError.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
	} else {
		p.logInfo.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Infof("INFO:Handling CreateComment")
	}
	return request
}

func (p PostHandler) sanitizeGetRequest(request *pb.GetRequest) *pb.GetRequest {
	policy := bluemonday.UGCPolicy()
	request.Id = strings.TrimSpace(policy.Sanitize(request.Id))
	sqlInj := common.BadId(request.Id)
	if request.Id == "" {
		p.logError.Logger.Errorf("XSS")
	} else if sqlInj {
		p.logError.Logger.Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
	} else {
		p.logInfo.Logger.Infof("INFO:Handling GetAllByUserId posts")
	}
	return request
}

func (p PostHandler) sanitizeReactionRequest(request *pb.ReactionRequest, userNameCtx string) *pb.ReactionRequest {
	policy := bluemonday.UGCPolicy()
	request.PostId = strings.TrimSpace(policy.Sanitize(request.PostId))
	request.Username = strings.TrimSpace(policy.Sanitize(request.Username))
	p1 := common.BadId(request.PostId)
	p2 := common.BadId(request.Username)

	if request.PostId == "" || request.Username == "" {
		p.logError.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Errorf("ERR:XSS")
	} else if p1 || p2 {
		p.logError.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
	} else {
		p.logInfo.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Infof("INFO:Handling LikePost")
	}
	return request
}
