package com.example.PKI.dto;

public class LoggedUserDto {
    private String email;
    private String role;

    public LoggedUserDto() {
    }

    public LoggedUserDto(String email, String role) {
        this.email = email;
        this.role = role;
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
}
