package com.example.PKI.service;

import com.example.PKI.model.CustomToken;
import com.example.PKI.model.TokenType;
import com.example.PKI.model.User;

public interface CustomTokenService {
    CustomToken createToken(User user, TokenType type);
    void sendVerificationToken(User user);
    CustomToken findByToken(String token);
    void deleteById(Long id);
}
