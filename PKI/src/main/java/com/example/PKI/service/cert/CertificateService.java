package com.example.PKI.service.cert;

import com.example.PKI.dto.*;
import com.example.PKI.model.Issuer;
import com.example.PKI.model.Subject;

import java.security.KeyPair;
import java.security.PrivateKey;

public interface CertificateService {
    public Subject generateSubjectData(SubjectDto subjectDto);
    public KeyPair generateKeyPair();
}
