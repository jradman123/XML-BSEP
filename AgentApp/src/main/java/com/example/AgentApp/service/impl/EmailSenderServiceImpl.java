package com.example.AgentApp.service.impl;

import com.example.AgentApp.service.EmailSenderService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.mail.SimpleMailMessage;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Service;

@Service
public class EmailSenderServiceImpl implements EmailSenderService {
    @Autowired
    private JavaMailSender mailSender;

    @Override
    @Async
    public void sendEmail(String toEmail, String subject, String body) {
        SimpleMailMessage message = new SimpleMailMessage();

        message.setFrom("testsemailpsw@gmail.com");
        message.setSubject(subject);
        message.setText(body);
        message.setTo(toEmail);

        try {
            mailSender.send(message);
        }catch( Exception e ){
            e.printStackTrace();
        }

    }
}
