package com.example.AgentApp.service.impl;

import com.example.AgentApp.dto.ChangePasswordDto;
import com.example.AgentApp.dto.RegistrationRequestDto;
import com.example.AgentApp.mapper.UserMapper;
import com.example.AgentApp.model.User;
import com.example.AgentApp.repository.UserRepository;
import com.example.AgentApp.service.CustomTokenService;
import com.example.AgentApp.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import java.text.ParseException;

@Service
public class UserServiceImpl implements UserService {

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private CustomTokenService customTokenService;

    @Autowired
    private PasswordEncoder passwordEncoder;

    @Override
    public User findByUsername(String username) {
        return userRepository.findByUsername(username);
    }

    @Override
    public User findByEmail(String email) {
        return userRepository.findByEmail(email);
    }

    @Override
    public User addUser(RegistrationRequestDto registrationRequestDto) throws ParseException {
        User newUser = UserMapper.mapToUser(registrationRequestDto);
        User created = userRepository.save(newUser);
        customTokenService.sendVerificationToken(created);
        return created;
    }

    @Override
    public User activateAccount(User user) {
        User userDb = findByUsername(user.getUsername());
        userDb.setConfirmed(true);
        User saved = userRepository.save(userDb);
        return saved;
    }

    @Override
    public void changePassword(String username, ChangePasswordDto changePasswordDto) {
        User user = findByUsername(username);
        if (passwordEncoder.matches(changePasswordDto.getCurrentPassword(),user.getPassword())) {
            user.setPassword(passwordEncoder.encode(changePasswordDto.getNewPassword()));
        }

        userRepository.save(user);
    }

    @Override
    public void resetPassword(String username, String newPassword) {
        User user = findByUsername(username);
        user.setPassword(passwordEncoder.encode(newPassword));
        userRepository.save(user);
    }

    @Override
    public Long getByUsername(String username) {
        return userRepository.findByUsername(username).getId();
    }

    @Override
    public String change2FAStatus(String username, Boolean status) {
        User user = findByUsername(username);
        user.setUsing2FA(status);
        user.setSecret();
        userRepository.save(user);
        return status ? user.getSecret() : "";
    }

    @Override
    public boolean check2FAStatus(String username) {
        return findByUsername(username).isUsing2FA();
    }

}
