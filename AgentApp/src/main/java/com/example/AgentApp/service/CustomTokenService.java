package com.example.AgentApp.service;

import com.example.AgentApp.model.CustomToken;
import com.example.AgentApp.model.User;

public interface CustomTokenService {
    void sendVerificationToken(User user);
    CustomToken findByToken(String token);
    CustomToken findByUser(User user);
    void deleteById(Long id);
    void sendResetPasswordToken(User user);
    boolean checkResetPasswordCode(String sentCode,String codeDb);
}
