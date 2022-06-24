package com.example.PKI.service;

import org.springframework.stereotype.Service;

@Service
public interface LoggerService {
    void loginFailed(String email,String ipAddress);
    void loginSuccess(String email);
    void signUpFailed(String ipAddress);
    void signUpSuccess(String email,String ipAddress);
    void confirmAccountFailed(String email,String ipAddress);
    void confirmAccountSuccess(String email,String ipAddress);
    void changePasswordSuccess(String email,String ipAddress);
    void sendCodeFailed(String email);
    void sendCodeSuccess(String email);
    void checkCodeFailed(String email,String ipAddress);
    void checkCodeSuccess(String email,String ipAddress);
    void resetPasswordSuccess(String email,String ipAddress);
    void checkRecoveryEmailFailed(String email,String ipAddress);
    void checkRecoveryEmailSuccess(String email,String ipAddress);
    void generateCertificateFailed(String ipAddress);
    void generateCertificateSuccess(String ipAddress);
    void downloadCertificateSuccess(String serialNumber);
}
