package com.example.AgentApp.dto;

public class CommentDto {

    public Long companyId;
   public String userUsername;
    public String comment;

    public CommentDto(Long companyId, String userUsername, String comment) {
        this.companyId = companyId;
        this.userUsername = userUsername;
        this.comment = comment;
    }

    public CommentDto() {
    }

    public Long getCompanyId() {
        return companyId;
    }

    public void setCompanyId(Long companyId) {
        this.companyId = companyId;
    }

    public String getUserUsername() {
        return userUsername;
    }

    public void setUserUsername(String userUsername) {
        this.userUsername = userUsername;
    }

    public String getComment() {
        return comment;
    }

    public void setComment(String comment) {
        this.comment = comment;
    }
}
