package com.example.PKI.service.impl;

import com.example.PKI.dto.ChangePasswordDto;
import com.example.PKI.dto.LoginDto;
import com.example.PKI.dto.UserDto;
import com.example.PKI.model.Permission;
import com.example.PKI.model.User;
import com.example.PKI.repository.PermissionRepository;
import com.example.PKI.repository.UserRepository;
import com.example.PKI.service.UserService;
import com.example.PKI.service.CustomTokenService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.*;

@Service
public class UserServiceImpl implements UserService {

    private PasswordEncoder passwordEncoder;
    private UserRepository userRepository;
    private PermissionRepository permissionRepository;
    private CustomTokenService verificationTokenService;

    @Autowired
    public UserServiceImpl(UserRepository userRepository, PasswordEncoder pe, PermissionRepository pr, CustomTokenService vts) {
        this.userRepository = userRepository;
        this.passwordEncoder = pe;
        this.permissionRepository = pr;
        this.verificationTokenService = vts;
    }

    @Override
    public boolean login(LoginDto user) {
        User foundUser = userRepository.findByEmail(user.getEmail());
        if (foundUser.getPassword().equals(user.getPassword())) {
            //foundUser.setLoggedIn(true);
            userRepository.save(foundUser);
            return true;
        }
        return false;
    }

    @Override
    public User findByEmail(String email) {
        return userRepository.findByEmail(email);
    }

    @Override
    public UserDto createUser(UserDto userDto) {
        User newUser = new User(userDto.getEmail(), passwordEncoder.encode(userDto.getPassword()),
                userDto.getCommonName(), userDto.getOrganization(), userDto.getOrganizationUnit(),
                userDto.getLocality(), userDto.getCountry(),userDto.getRecoveryMail());

        List<Permission> permissions = new ArrayList<Permission>();
        permissions.add(permissionRepository.findByName("user_download"));
        permissions.add(permissionRepository.findByName("user_read"));
        newUser.setPermissions(permissions);

        User created = userRepository.save(newUser);
        userDto.setId(created.getId());
        verificationTokenService.sendVerificationToken(created);

        return userDto;
    }

    @Override
    public User activateAccount(User user) {
        User userDb = findByEmail(user.getEmail());
        userDb.setActivated(true);
        User saved = userRepository.save(userDb);
        return saved;
    }

    @Override
    public void changePassword(String email, ChangePasswordDto changePasswordDto) {
        User user = findByEmail(email);
        if (passwordEncoder.matches(changePasswordDto.getCurrentPassword(),user.getPassword())) {
            user.setPassword(passwordEncoder.encode(changePasswordDto.getNewPassword()));
        }

        userRepository.save(user);
    }

    @Override
    public void resetPassword(String email, String newPassword) {
        User user = findByEmail(email);
        user.setPassword(passwordEncoder.encode(newPassword));
        userRepository.save(user);
    }


}
