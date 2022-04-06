package com.example.PKI.service.cert;

import com.example.PKI.model.Issuer;
import com.example.PKI.model.Subject;

import java.security.cert.X509Certificate;

public interface CertificateGenerator {
    public X509Certificate generateCertificate(Subject subjectData, Issuer issuerData);
}
