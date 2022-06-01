package com.example.AgentApp.service.impl;

import com.example.AgentApp.dto.ChangePasswordDto;
import com.example.AgentApp.dto.RegistrationRequestDto;
import com.example.AgentApp.enums.Gender;
import com.example.AgentApp.enums.UserRole;
import com.example.AgentApp.model.User;
import com.example.AgentApp.repository.UserRepository;
import com.example.AgentApp.service.CustomTokenService;
import com.example.AgentApp.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;

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
        Gender gender = getGenderFromRequest(registrationRequestDto.getGender());
        Date dateOfBirth = getDateOfBirthFromRequest(registrationRequestDto.getDateOfBirth());
        User newUser = new User(registrationRequestDto.getUsername(),passwordEncoder.encode(registrationRequestDto.getPassword()),
                                registrationRequestDto.getEmail(),registrationRequestDto.getRecoveryEmail(),
                                registrationRequestDto.getPhoneNumber(),registrationRequestDto.getFirstName(),
                                registrationRequestDto.getLastName(),dateOfBirth,
                                gender, UserRole.REGISTERED_USER,false);
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

    private Date getDateOfBirthFromRequest(String dateOfBirth) throws ParseException {
        return new SimpleDateFormat("MM/dd/yyyy").parse(dateOfBirth);
    }

    private Gender getGenderFromRequest(String gender) {
        if(gender == Gender.FEMALE.toString()){
            return Gender.FEMALE;
        }else{
            return Gender.MALE;
        }

    }
}
