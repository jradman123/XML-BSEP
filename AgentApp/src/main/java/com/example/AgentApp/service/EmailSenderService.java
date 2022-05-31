package com.example.AgentApp.service;

public interface EmailSenderService {

    void sendEmail(String toEmail,String subject,String body);
}
