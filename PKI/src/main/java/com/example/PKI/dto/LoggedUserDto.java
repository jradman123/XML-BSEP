package com.example.PKI.dto;

public class LoggedUserDto {
    private String email;
    private String role;
    private UserTokenState token;

    public LoggedUserDto() {
    }

    public LoggedUserDto(String email, String role, UserTokenState token) {
        this.email = email;
        this.role = role;
        this.token = token;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
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
