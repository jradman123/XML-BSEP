package com.example.PKI.model;

import lombok.Data;
import org.bouncycastle.asn1.x500.X500Name;

import java.math.BigInteger;
import java.security.PublicKey;
import java.util.Date;

@Data
public class Subject {
    //msm da u subject treba i private key

    private PublicKey publicKey;
    private X500Name x500Name;
    private BigInteger serialNumber;
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

    public Subject(PublicKey publicKey, X500Name x500Name, BigInteger serialNumber, Date startDate, Date endDate, String type, String alias, String commonName, String organization) {
        this.publicKey = publicKey;
        this.x500Name = x500Name;
        this.serialNumber = serialNumber;
        this.startDate = startDate;
        this.endDate = endDate;
        this.type = type;
        this.alias = alias;
        this.commonName = commonName;
        this.organization = organization;
    }

    public PublicKey getPublicKey() {
        return publicKey;
    }

    public void setPublicKey(PublicKey publicKey) {
        this.publicKey = publicKey;
    }

    public X500Name getX500Name() {
        return x500Name;
    }

    public void setX500Name(X500Name x500Name) {
        this.x500Name = x500Name;
    }

    public BigInteger getSerialNumber() {
        return serialNumber;
    }

    public void setSerialNumber(BigInteger serialNumber) {
        this.serialNumber = serialNumber;
    }

    public Date getStartDate() {
        return startDate;
    }

    public void setStartDate(Date startDate) {
        this.startDate = startDate;
    }

    public Date getEndDate() {
        return endDate;
    }

    public void setEndDate(Date endDate) {
        this.endDate = endDate;
    }

    public String getOrganization() {
        return organization;
    }

    public void setOrganization(String organization) {
        this.organization = organization;
    }

    public String getOrganizationUnit() {
        return organizationUnit;
    }

    public void setOrganizationUnit(String organizationUnit) {
        this.organizationUnit = organizationUnit;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getCountry() {
        return country;
    }

    public void setCountry(String country) {
        this.country = country;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public String getAlias() {
        return alias;
    }

    public void setAlias(String alias) {
        this.alias = alias;
    }

    public String getCommonName() {
        return commonName;
    }

    public void setCommonName(String commonName) {
        this.commonName = commonName;
    }

    public String getLocality() {
        return locality;
    }

    public void setLocality(String locality) {
        this.locality = locality;
    }
}

