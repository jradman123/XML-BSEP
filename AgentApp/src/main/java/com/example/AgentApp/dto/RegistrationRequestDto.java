package com.example.AgentApp.dto;

import com.example.AgentApp.enums.Gender;
import com.example.AgentApp.enums.UserRole;
import lombok.Data;

import javax.persistence.Column;
import javax.validation.constraints.Email;
import javax.validation.constraints.Pattern;
import java.util.Date;

@Data
public class RegistrationRequestDto {
    private String username;
    @Pattern(regexp= "^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!\"#$@%&()*<>+_|~]).*$", message =  "Password format not valid")
    private String password;
    @Email
    private String email;
    @Email
    private String recoveryEmail;
    private String phoneNumber;
    @Pattern(regexp = "[a-zA-Z]+")
    private String firstName;
    @Pattern(regexp = "[a-zA-Z]+")
    private String lastName;
    private String dateOfBirth;
    @Pattern(regexp = "[a-zA-Z]+")
    private String gender;

}
