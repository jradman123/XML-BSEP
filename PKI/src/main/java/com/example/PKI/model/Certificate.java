package com.example.PKI.model;

import javax.persistence.*;

import lombok.Data;

@Data
@Entity
public class Certificate {

    @Id
    @GeneratedValue(strategy= GenerationType.IDENTITY)
    private int id;

    @Column(unique = true)
    private String serialNumber;

    @Column
    private CertificateType type;
    
    @Column
    private String validFrom;
    
    @Column
    private String validTo;
 
    @Column
    private boolean isRevoked;

    @Column
    private String subjectEmail;

    public Certificate() {
    }

    public String getSerialNumber() {
        return serialNumber;
    }

    public void setSerialNumber(String serialNumber) {
        this.serialNumber = serialNumber;
    }

    public CertificateType getType() {
        return type;
    }

    public void setType(CertificateType type) {
        this.type = type;
    }

    public boolean isRevoked() {
        return isRevoked;
    }

    public void setIsRevoked(boolean isRevoked) {
        this.isRevoked = isRevoked;
    }

    public String getSubjectEmail() {
        return subjectEmail;
    }

    public void setSubjectEmail(String subjectEmail) {
        this.subjectEmail = subjectEmail;
    }
}
