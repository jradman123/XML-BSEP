package com.example.PKI.service;

public interface EmailSenderService {

    void sendEmail(String toEmail,String subject,String body);
}
