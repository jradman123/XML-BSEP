package com.example.PKI.service;

import com.example.PKI.model.CustomToken;
import com.example.PKI.model.TokenType;
import com.example.PKI.model.User;

public interface CustomTokenService {
    void sendVerificationToken(User user);
    CustomToken findByToken(String token);
    CustomToken findByUser(User user);
    void deleteById(Long id);
    void sendResetPasswordToken(User user);
    boolean checkResetPasswordCode(String sentCode,String codeDb);
    void sendMagicLink(User user);
}
