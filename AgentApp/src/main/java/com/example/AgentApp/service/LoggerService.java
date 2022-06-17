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
    void createCompanyRequestFailed(String companyId);
    void createCompanyRequestSuccess(String companyId);
    void approveCompanyRequestFailed(String companyId);
    void approveCompanyRequestSuccess(String companyId,String username);
    void rejectCompanyRequestFailed(String companyId);
    void rejectCompanyRequestSuccess(String companyId,String username);
    void editCompanyFailed(String companyId);
    void editCompanySuccess(String companyId,String username);
    void createJobOfferFailed(String companyId);
    void createJobOfferSuccess(String companyId,String username);
    void leaveCommentFailed(String companyId,String username);
    void leaveCommentSuccess(String companyId,String username);
    void leaveSalaryCommentFailed(String companyId,String username);
    void leaveSalaryCommentSuccess(String companyId,String username);
    void leaveInterviewCommentFailed(String companyId,String username);
    void leaveInterviewCommentSuccess(String companyId,String username);
    void sendLinkForPasswordlessFailed(String email);
    void sendLinkForPasswordlessSuccess(String email);
    void passwordlessLoginFailed(String username,String ipAddress);
    void passwordlessLoginSuccess(String username);
    void changeTwoFactorStatus(String username,String ipAddress);


}
