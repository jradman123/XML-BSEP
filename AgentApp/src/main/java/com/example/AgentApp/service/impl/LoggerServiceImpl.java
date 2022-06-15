package com.example.AgentApp.service.impl;

import com.example.AgentApp.service.LoggerService;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

public class LoggerServiceImpl implements LoggerService {

    private final Logger logger;

    public LoggerServiceImpl(Class<?> parentClass) {
        this.logger = LogManager.getLogger(parentClass);
    }

    @Override
    public void loginFailed(String username,String ipAddress) {
        logger.error("Login failed. Username: " + username + ". Ip address: " + ipAddress);
    }

    @Override
    public void loginSuccess(String username) {logger.info("Login succeed. Username: " + username);}

    @Override
    public void signUpFailed(String ipAddress) {logger.error("Sign up failed. Ip address: " + ipAddress);}

    @Override
    public void signUpSuccess(String username, String ipAddress) {logger.info("Sign up succeed. Username: " + username +
            ". Ip address: " + ipAddress);}

    @Override
    public void confirmAccountFailed(String username,String ipAddress) {
        logger.error("Confirmation account failed. Username: " + username + ". Ip address: "+ipAddress);
    }

    @Override
    public void confirmAccountSuccess(String username,String ipAddress) {
        logger.info("Confirmation account succeed. Username: " + username + ". Ip address: " + ipAddress);
    }

    @Override
    public void sendCodeFailed(String email) {
        logger.error("Send code failed. Email: " + email );
    }

    @Override
    public void sendCodeSuccess(String email) {
        logger.info("Send code succeed. Email: " + email );
    }

    @Override
    public void checkCodeFailed(String username, String ipAddress) {
        logger.error("Check code failed. Username: " + username + ". Ip address: " + ipAddress);
    }

    @Override
    public void checkCodeSuccess(String username, String ipAddress) {
        logger.info("Check code succeed. Username: " + username + ". Ip address: " + ipAddress);
    }

    @Override
    public void resetPasswordSuccess(String username, String ipAddress) {
        logger.info("Reset password succeed. Username: " + username + ". Ip address: " + ipAddress);
    }

    @Override
    public void changePasswordSuccess(String username, String ipAddress) {
        logger.info("Change password succeed. Username: " + username + ". Ip address: " + ipAddress);
    }

    @Override
    public void userEditInfo(String username) {
        logger.info("User edits info. Username: " + username);
    }
}
