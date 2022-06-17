package com.example.PKI.dto;

import lombok.Data;

import javax.validation.constraints.Email;

@Data
public class CheckCodeDto {
    @Email
    private String email;
    private String code;
}
