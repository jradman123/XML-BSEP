package com.example.AgentApp.dto;

import lombok.Data;

public class LoggedUserDto {
    private String username;
    private String role;
    private UserTokenState token;

    public LoggedUserDto() {
    }

    public LoggedUserDto(String username, String role, UserTokenState token) {
        this.username = username;
        this.role = role;
        this.token = token;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getRole() {
        return role;
    }

    public void setRole(String role) {
        this.role = role;
    }

    public UserTokenState getToken() {
        return token;
    }

    public void setToken(UserTokenState token) {
        this.token = token;
    }

}


