package com.example.AgentApp.dto;

import javax.validation.constraints.Pattern;

public class CommentDto {

    public Long companyId;
    @Pattern(regexp= "^[a-zA-Z0-9]([._-](?![._-])|[a-zA-Z0-9]){3,18}[a-zA-Z0-9]$", message =  "Username format not valid")
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
