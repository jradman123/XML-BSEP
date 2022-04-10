package com.example.PKI.dto;

import lombok.Data;

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
}
