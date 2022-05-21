package com.example.PKI.service.impl;

import com.example.PKI.model.User;
import com.example.PKI.model.VerificationToken;
import com.example.PKI.repository.VerificationTokenRepository;
import com.example.PKI.service.EmailSenderService;
import com.example.PKI.service.VerificationTokenService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.UUID;

@Service
public class VerificationTokenServiceImpl implements VerificationTokenService {

    @Autowired
    private VerificationTokenRepository verificationTokenRepository;

    @Autowired
    private EmailSenderService emailSenderService;

    @Override
    public VerificationToken createToken(User user) {
        VerificationToken token = new VerificationToken(UUID.randomUUID().toString(),user);
        verificationTokenRepository.save(token);
        return token;

    }

    @Override
    public void sendVerificationToken(User user) {
        String confirmationLink = "http://localhost:8443/api/confirmAccount/" + createToken(user).getToken();
        emailSenderService.sendEmail(user.getEmail(),"Confirm account", "Click on following link to confirm " +
                "your account \n" + confirmationLink);

    }

    @Override
    public VerificationToken findByToken(String token) {
        return verificationTokenRepository.findByToken(token);
    }

    @Override
    public void deleteById(Long id) {
        verificationTokenRepository.deleteById(id);
    }


}
