package com.example.PKI.service.impl;

import com.example.PKI.model.CustomToken;
import com.example.PKI.model.TokenType;
import com.example.PKI.model.User;
import com.example.PKI.repository.CustomTokenRepository;
import com.example.PKI.service.EmailSenderService;
import com.example.PKI.service.CustomTokenService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.UUID;

@Service
public class CustomTokenServiceImpl implements CustomTokenService {

    @Autowired
    private CustomTokenRepository customTokenRepository;

    @Autowired
    private EmailSenderService emailSenderService;
    

    @Override
    public CustomToken createToken(User user, TokenType type) {
        CustomToken token = new CustomToken(UUID.randomUUID().toString(),user,type);
        customTokenRepository.save(token);
        return token;

    }

    @Override
    public void sendVerificationToken(User user) {
        String confirmationLink = "http://localhost:8443/api/confirmAccount/" + createToken(user,TokenType.Confirmation).getToken();
        emailSenderService.sendEmail(user.getEmail(),"Confirm account", "Click on following link to confirm " +
                "your account \n" + confirmationLink);

    }

    @Override
    public CustomToken findByToken(String token) {
        return customTokenRepository.findByToken(token);
    }

    @Override
    public void deleteById(Long id) {
        customTokenRepository.deleteById(id);
    }


}
