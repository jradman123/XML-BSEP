package com.example.PKI.service;

import com.example.PKI.model.User;
import com.example.PKI.model.VerificationToken;

public interface VerificationTokenService {
    VerificationToken createToken(User user);
    void sendVerificationToken(User user);
    VerificationToken findByToken(String token);
    void deleteById(Long id);
}
