package com.example.PKI.dto;

import javax.validation.constraints.Email;
import javax.validation.constraints.Pattern;

public class JwtAuthenticationRequest {
    @Email
    private String email;
    @Pattern(regexp= "^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!\"#$@%&()*<>+_|~]).*$", message =  "Password format not valid")
    private String password;
    @Pattern(regexp = "^[0-9]{1,6}$", message = "Code format not valid")
    private String code;


    public JwtAuthenticationRequest() {
        super();
    }

    public JwtAuthenticationRequest(String email, String password, String code) {
        this.setEmail(email);
        this.setPassword(password);
        this.setCode(code);
    }

    public String getCode() {
        return code;
    }

    public void setCode(String code) {
        this.code = code;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getPassword() {
        return this.password;
    }

    public void setPassword(String password) {
        this.password = password;
    }
}
