package com.example.PKI.dto;

import com.example.PKI.model.*;
import lombok.*;

import java.security.*;

@Data
public class SubjectWithPKDto {
    private Subject subject;
    private PrivateKey privateKey;

    public SubjectWithPKDto(Subject subject, PrivateKey privateKey) {
        this.subject = subject;
        this.privateKey = privateKey;
    }

    public Subject getSubject() {
        return subject;
    }

    public void setSubject(Subject subject) {
        this.subject = subject;
    }

    public PrivateKey getPrivateKey() {
        return privateKey;
    }

    public void setPrivateKey(PrivateKey privateKey) {
        this.privateKey = privateKey;
    }
}
