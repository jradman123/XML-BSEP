package com.example.PKI.service;

import com.example.PKI.dto.LoginDto;
import com.example.PKI.dto.UserDto;
import com.example.PKI.model.User;

public interface UserService {
    boolean login(LoginDto loginDTO);

    void logout(String email);

    User findByEmail(String email);

    UserDto createUser(UserDto userDto);
}
