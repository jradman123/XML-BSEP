package com.example.PKI.dto;

import com.example.PKI.model.User;
import org.bouncycastle.asn1.x500.*;

import java.security.*;
import java.util.*;

public class IssuerDto {
    public Integer id;

    public String email;

    public String password;

    public String commonName;

    public String organization;

    public String organizationUnit;

    public String locality;

    public String country;

    public String serialNumber;

    public IssuerDto(User user, String serialNumber){
        this.id = user.getId();
        this.email = user.getEmail();
        this.password = user.getPassword();
        this.commonName = user.getCommonName();
        this.organization = user.getOrganization();
        this.organizationUnit = user.getOrganizationUnit();
        this.locality = user.getLocality();
        this.country = user.getCountry();
        this.serialNumber = serialNumber;
    }
}
