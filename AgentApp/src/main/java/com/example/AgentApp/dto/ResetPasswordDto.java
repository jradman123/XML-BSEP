package com.example.AgentApp.dto;

import lombok.Data;

import javax.validation.constraints.Email;
import javax.validation.constraints.Pattern;

@Data
public class ResetPasswordDto {

    private String username;

    @Pattern(regexp= "^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!\"#$@%&()*<>+_|~]).*$", message =  "Password format not valid")
    private String newPassword;
}
