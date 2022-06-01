package com.example.AgentApp.dto;

import lombok.Data;

import javax.validation.constraints.Email;

@Data
public class CheckCodeDto {

    private String username;
    private String code;
}

