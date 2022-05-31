package com.example.AgentApp.service.impl;

import com.example.AgentApp.enums.TokenType;
import com.example.AgentApp.model.CustomToken;
import com.example.AgentApp.model.User;
import com.example.AgentApp.repository.CustomTokenRepository;
import com.example.AgentApp.service.CustomTokenService;
import com.example.AgentApp.service.EmailSenderService;
import net.bytebuddy.utility.RandomString;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.UUID;

@Service
public class CustomTokenServiceImpl implements CustomTokenService {

    @Autowired
    private CustomTokenRepository customTokenRepository;

    @Autowired
    private EmailSenderService emailSenderService;

    @Autowired
    private PasswordEncoder passwordEncoder;


    private CustomToken createConfirmationToken(User user) {
        CustomToken token = new CustomToken(UUID.randomUUID().toString(),user, TokenType.CONFIRMATION);
        customTokenRepository.save(token);
        return token;

    }

    private CustomToken createResetPasswordToken(User user) {
        CustomToken token = new CustomToken(RandomString.make(8),user,TokenType.RESET_PASSWORD);
        return token;

    }

    private void saveToken(CustomToken customToken){
        String valueOfToken = customToken.getToken();
        customToken.setToken(passwordEncoder.encode(valueOfToken));
        customTokenRepository.save(customToken);
    }

    @Override
    public void sendVerificationToken(User user) {
        String confirmationLink = "http://localhost:8443/api/confirmAccount/" + createConfirmationToken(user).getToken();
        emailSenderService.sendEmail(user.getEmail(),"Confirm account", "Click on following link to confirm " +
                "your account \n" + confirmationLink);

    }

    @Override
    public CustomToken findByToken(String token) {
        return customTokenRepository.findByToken(token);
    }

    @Override
    public CustomToken findByUser(User user) {
        return customTokenRepository.findByUser(user);
    }

    @Override
    public void deleteById(Long id) {
        customTokenRepository.deleteById(id);
    }

    @Override
    public void sendResetPasswordToken(User user) {
        CustomToken customToken = createResetPasswordToken(user);
        String passwordCode = customToken.getToken();
        saveToken(customToken);
        emailSenderService.sendEmail(user.getRecoveryEmail(),"Reset password", "Following code is your new temporary " +
                "password \nCode : " + passwordCode);


    }

    @Override
    public boolean checkResetPasswordCode(String sentCode, String codeDb) {
        return passwordEncoder.matches(sentCode,codeDb);
    }


}

