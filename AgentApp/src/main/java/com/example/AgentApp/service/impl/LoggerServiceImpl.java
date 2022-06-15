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
    public void createCompanyRequestFailed(String companyId) {
        logger.error("Company request creation failed. Company: " + companyId);
    }

    @Override
    public void createCompanyRequestSuccess(String companyId) {
        logger.info("Company request successfully created. Company: " + companyId);
    }

    @Override
    public void approveCompanyRequestFailed(String companyId) {
        logger.error("Company " + companyId + " approving failed.");
    }

    @Override
    public void approveCompanyRequestSuccess(String companyId,String username) {
        logger.info("Company " + companyId + " approved by " + username + ".");

    }

    @Override
    public void rejectCompanyRequestFailed(String companyId) {
        logger.error("Company " + companyId + " rejecting failed.");
    }

    @Override
    public void rejectCompanyRequestSuccess(String companyId, String username) {
        logger.info("Company " + companyId + " rejected by " + username + ".");
    }

    @Override
    public void editCompanyFailed(String companyId) {
        logger.error("Company " + companyId + " editing failed.");
    }

    @Override
    public void editCompanySuccess(String companyId, String username) {
        logger.info("Company " + companyId + "successfully edited by " + username + ".");
    }

    @Override
    public void createJobOfferFailed(String companyId) {
        logger.error("Creating job offer for company " + companyId + "failed.");
    }

    @Override
    public void createJobOfferSuccess(String companyId, String username) {
        logger.info("Creating job offer for company " + companyId + "by " + username + ".");
    }

    @Override
    public void leaveCommentFailed(String companyId, String username) {
        logger.error("Leaving comment for company " + companyId + "failed.Username " + username + " tried leave comment.");
    }

    @Override
    public void leaveCommentSuccess(String companyId, String username) {
        logger.info("Leaving comment for company " + companyId + "successfully by " + username + ".");
    }

    @Override
    public void leaveSalaryCommentFailed(String companyId, String username) {
        logger.error("Leaving salary comment for company " + companyId + "failed.Username " + username + " tried leave salary comment.");
    }

    @Override
    public void leaveSalaryCommentSuccess(String companyId, String username) {
        logger.info("Leaving salary comment for company " + companyId + "successfully by " + username + ".");
    }

    @Override
    public void leaveInterviewCommentFailed(String companyId, String username) {
        logger.error("Leaving salary comment for company " + companyId + "failed.Username " + username + " tried leave interview comment.");
    }

    @Override
    public void leaveInterviewCommentSuccess(String companyId, String username) {
        logger.info("Leaving interview comment for company " + companyId + "successfully by " + username + ".");
    }


}
