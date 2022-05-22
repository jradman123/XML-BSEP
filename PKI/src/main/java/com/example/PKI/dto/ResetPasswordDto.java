package com.example.PKI.dto;

import lombok.Data;

@Data
public class ResetPasswordDto {
    private String email;
    private String newPassword;
}
