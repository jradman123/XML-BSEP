package com.example.AgentApp.dto;

import lombok.Data;

import javax.validation.constraints.Email;

@Data
public class RequestCheckDto {

    @Email
    private String email;
    @Email
    private String recoveryEmail;
}
