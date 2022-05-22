package com.example.PKI.dto;

import lombok.Data;
import java.util.*;

@Data
public class UserDto {

    public Integer id;

    public String email;

    public String password;

    public String commonName;

    public String organization;

    public String organizationUnit;

    public String locality;

    public String country;

    public String recoveryMail;

    public boolean isPawned;
}
