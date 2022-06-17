package com.example.AgentApp.dto;

import lombok.Data;

import javax.validation.constraints.Email;
import javax.validation.constraints.Pattern;

@Data
public class CheckCodeDto {

    @Pattern(regexp= "^[a-zA-Z0-9]([._-](?![._-])|[a-zA-Z0-9]){3,18}[a-zA-Z0-9]$", message =  "Username format not valid")
    private String username;
    private String code;
}

