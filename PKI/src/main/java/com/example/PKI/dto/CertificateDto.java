package com.example.PKI.dto;

import lombok.*;
import org.bouncycastle.asn1.x500.*;

import java.security.*;
import java.util.*;

@Data
public class CertificateDto {

    private Integer subjectId;
    private Integer issuerId;
    private String startDate;
    private String endDate;
    private String issuerSerialNumber;
    private String type;

    public CertificateDto() {
    }

    public CertificateDto(Integer subjectId, Integer issuerId, String startDate, String endDate, String issuerSerialNumber, String type) {
        this.subjectId = subjectId;
        this.issuerId = issuerId;
        this.startDate = startDate;
        this.endDate = endDate;
        this.issuerSerialNumber = issuerSerialNumber;
        this.type = type;
    }

    public Integer getSubjectId() {
        return subjectId;
    }

    public void setSubjectId(Integer subjectId) {
        this.subjectId = subjectId;
    }

    public Integer getIssuerId() {
        return issuerId;
    }

    public void setIssuerId(Integer issuerId) {
        this.issuerId = issuerId;
    }

    public String getStartDate() {
        return startDate;
    }

    public void setStartDate(String startDate) {
        this.startDate = startDate;
    }

    public String getEndDate() {
        return endDate;
    }

    public void setEndDate(String endDate) {
        this.endDate = endDate;
    }

    public String getIssuerSerialNumber() {
        return issuerSerialNumber;
    }

    public void setIssuerSerialNumber(String issuerSerialNumber) {
        this.issuerSerialNumber = issuerSerialNumber;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }
}
