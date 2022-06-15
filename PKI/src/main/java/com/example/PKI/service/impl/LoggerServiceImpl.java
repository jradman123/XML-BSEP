package com.example.PKI.service.impl;

import com.example.PKI.service.LoggerService;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

public class LoggerServiceImpl implements LoggerService {

    private final Logger logger;

    public LoggerServiceImpl(Class<?> parentClass) {
        this.logger = LogManager.getLogger(parentClass);
    }

    @Override
    public void loginFailed(String email, String ipAddress) {
        logger.error("Login failed. Email: " + email + ". Ip address: " + ipAddress);
    }

    @Override
    public void loginSuccess(String email) {
        logger.info("Login succeed. Email: " + email);
    }

    @Override
    public void signUpFailed(String ipAddress) {
        logger.error("Sign up failed. Ip address: " + ipAddress);
    }

    @Override
    public void signUpSuccess(String email, String ipAddress) {
        logger.info("Sign up succeed. Email: " + email +
                ". Ip address: " + ipAddress);
    }

    @Override
    public void confirmAccountFailed(String email, String ipAddress) {
        logger.error("Confirmation account failed. Email: " + email + ". Ip address: "+ipAddress);
    }

    @Override
    public void confirmAccountSuccess(String email, String ipAddress) {
        logger.info("Confirmation account succeed. Email: " + email + ". Ip address: " + ipAddress);
    }

    @Override
    public void changePasswordSuccess(String email, String ipAddress) {
        logger.info("Change password succeed. Email: " + email + ". Ip address: " + ipAddress);
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
    public void checkCodeFailed(String email, String ipAddress) {
        logger.error("Check code failed. Email: " + email + ". Ip address: " + ipAddress);
    }

    @Override
    public void checkCodeSuccess(String email, String ipAddress) {
        logger.info("Check code succeed. Email: " + email + ". Ip address: " + ipAddress);
    }

    @Override
    public void resetPasswordSuccess(String email, String ipAddress) {
        logger.info("Reset password succeed. Email: " + email + ". Ip address: " + ipAddress);
    }

    @Override
    public void checkRecoveryEmailFailed(String email, String ipAddress) {
        logger.error("Check recovery email failed. Email: " + email + ". Ip address: " + ipAddress);
    }

    @Override
    public void checkRecoveryEmailSuccess(String email, String ipAddress) {
        logger.info("Check recovery email succeed. Email: " + email + ". Ip address: " + ipAddress);
    }
}
