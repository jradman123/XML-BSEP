package com.example.AgentApp.dto;

import lombok.Data;

@Data
public class UserInformationResponseDto {
    private String username;
    private String firstName;
    private String lastName;
    private String email;
    private String phoneNumber;
    private String dateOfBirth;
    private String recoveryEmail;
    private String gender;
}
