package com.example.AgentApp.service;

import com.example.AgentApp.dto.ChangePasswordDto;
import com.example.AgentApp.dto.RegistrationRequestDto;
import com.example.AgentApp.model.User;

import java.text.ParseException;

public interface UserService {
    User findByUsername(String username);
    User findByEmail(String email);
    User addUser(RegistrationRequestDto registrationRequestDto) throws ParseException;
    User activateAccount(User user);
    void changePassword(String username, ChangePasswordDto changePasswordDto);
    void resetPassword(String username,String newPassword);

    Long getByUsername(String username);
}
