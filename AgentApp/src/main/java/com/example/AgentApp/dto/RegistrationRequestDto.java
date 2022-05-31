package com.example.AgentApp.dto;

import com.example.AgentApp.enums.Gender;
import com.example.AgentApp.enums.UserRole;
import lombok.Data;

import javax.persistence.Column;
import java.util.Date;

@Data
public class RegistrationRequestDto {
    private String username;
    private String password;
    private String email;
    private String recoveryEmail;
    private String phoneNumber;
    private String firstName;
    private String lastName;
    private String dateOfBirth;
    private String gender;
    private String isPwned;

}
