package com.example.PKI.dto;

import lombok.*;
import org.bouncycastle.asn1.x500.*;

import java.security.*;
import java.util.*;

@Data
public class SubjectDto {
    private String serialNumber;
    private Date startDate;
    private Date endDate;
    private String commonName;
    private String organization;
    private String organizationUnit;
    private String email;
    private String locality;
    private String country;
    private String type;
    private String alias;
}
