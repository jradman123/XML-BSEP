package com.example.AgentApp.dto;

import lombok.Data;

@Data
public class JwtAuthenticationRequest {
    private String username;
    private String password;

}
