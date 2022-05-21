package com.example.PKI.service;

import com.example.PKI.dto.LoginDto;
import com.example.PKI.dto.UserDto;
import com.example.PKI.model.User;

public interface UserService {
    boolean login(LoginDto loginDTO);

    User findByEmail(String email);
    UserDto createUser(UserDto userDto);
    User activateAccount(User user);
}
