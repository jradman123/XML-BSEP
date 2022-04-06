package com.example.PKI.model;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import lombok.Data;

@Data
@Entity
public class Certificate {

    @Id
    private int id;

    @Column
    private String serialNumber;

    @Column
    private CertificateType type;
    
    @Column
    private String validFrom;
    
    @Column
    private String validTo;
 
    @Column
    private boolean valid;

    @Column
    private String subjectCommonName;

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

    public boolean isValid() {
        return valid;
    }

    public void setValid(boolean valid) {
        this.valid = valid;
    }

    public String getSubjectCommonName() {
        return subjectCommonName;
    }

    public void setSubjectCommonName(String subjectCommonName) {
        this.subjectCommonName = subjectCommonName;
    }
}
