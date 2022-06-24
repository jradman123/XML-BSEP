package com.example.PKI.dto;

import lombok.Data;

import javax.validation.constraints.Email;
import javax.validation.constraints.Pattern;
import java.util.*;

@Data
public class UserDto {

    public Integer id;

    @Email
    public String email;

    @Pattern(regexp= "^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!\"#$@%&()*<>+_|~]).*$", message =  "Password format not valid")
    public String password;

    @Pattern(regexp = "[a-zA-Z]+")
    public String commonName;

    @Pattern(regexp = "[a-zA-Z]+")
    public String organization;

    @Pattern(regexp = "[a-zA-Z]+")
    public String organizationUnit;

    @Pattern(regexp = "[a-zA-Z]+")
    public String locality;

    @Pattern(regexp = "[a-zA-Z]+")
    public String country;

    @Email
    public String recoveryMail;

    public boolean isPawned;
}
