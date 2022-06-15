package com.example.AgentApp.service;

import org.springframework.stereotype.Service;

@Service
public interface LoggerService {
    void loginFailed(String username,String ipAddress);
    void loginSuccess(String username);
    void signUpFailed(String ipAddress);
    void signUpSuccess(String username,String ipAddress);
    void confirmAccountFailed(String username,String ipAddress);
    void confirmAccountSuccess(String username,String ipAddress);
    void sendCodeFailed(String email);
    void sendCodeSuccess(String email);
    void checkCodeFailed(String username,String ipAddress);
    void checkCodeSuccess(String username,String ipAddress);
    void resetPasswordSuccess(String username,String ipAddress);
    void changePasswordSuccess(String username,String ipAddress);
    void userEditInfo(String username);


}
