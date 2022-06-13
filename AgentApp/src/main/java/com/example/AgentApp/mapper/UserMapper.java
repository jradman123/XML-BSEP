package com.example.AgentApp.mapper;


import com.example.AgentApp.dto.RegistrationRequestDto;
import com.example.AgentApp.dto.UserInformationResponseDto;
import com.example.AgentApp.enums.Gender;
import com.example.AgentApp.enums.UserRole;
import com.example.AgentApp.model.User;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Component;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;

@Component
public class UserMapper {
    private static PasswordEncoder passwordEncoder;
    public UserMapper(PasswordEncoder passwordEncoder) {
        this.passwordEncoder = passwordEncoder;
    }

    public static User mapToUser(RegistrationRequestDto dto) throws ParseException {
        User user = new User();
        user.setConfirmed(false);
        user.setFirstName(dto.getFirstName());
        user.setLastName(dto.getLastName());
        user.setDateOfBirth(getDateOfBirthFromRequest(dto.getDateOfBirth()));
        user.setGender(getGenderFromRequest(dto.getGender()));
        user.setPassword(passwordEncoder.encode(dto.getPassword()));
        user.setEmail(dto.getEmail());
        user.setRecoveryEmail(dto.getRecoveryEmail());
        user.setUsername(dto.getUsername());
        user.setPhoneNumber(dto.getPhoneNumber());
        user.setRole(UserRole.REGISTERED_USER);
        return user;
    }

    private static Date getDateOfBirthFromRequest(String dateOfBirth) throws ParseException {
        return new SimpleDateFormat("MM/dd/yyyy").parse(dateOfBirth);
    }

    private static Gender getGenderFromRequest(String gender) {
        if (gender.equals(Gender.FEMALE.toString())) {
            return Gender.FEMALE;
        } else {
            return Gender.MALE;
        }
    }

    public static UserInformationResponseDto mapToDto(User user) {
        UserInformationResponseDto dto = new UserInformationResponseDto();
        dto.setEmail(user.getEmail());
        dto.setUsername(user.getUsername());
        dto.setLastName(user.getLastName());
        dto.setFirstName(user.getFirstName());
        dto.setPhoneNumber(user.getPhoneNumber());
        dto.setRecoveryEmail(user.getRecoveryEmail());
        dto.setGender(user.getGender().toString());
        dto.setDateOfBirth(convertDateToString(user.getDateOfBirth()));
        return dto;
    }

    private static String convertDateToString(Date dateOfBirth) {
        SimpleDateFormat dateFormat = new SimpleDateFormat("MM/dd/yyyy");
        return dateFormat.format(dateOfBirth);
    }
}
