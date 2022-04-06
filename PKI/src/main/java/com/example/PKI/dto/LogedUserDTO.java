package com.example.PKI.dto;

public class LogedUserDTO {
    private String email;
    private String role;

    public LogedUserDTO() {
    }

    public LogedUserDTO( String email, String role) {
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
