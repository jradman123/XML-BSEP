package com.example.PKI.service.cert;

import com.example.PKI.dto.*;
import com.example.PKI.model.Certificate;
import com.example.PKI.model.Subject;
import com.example.PKI.model.User;
import java.io.*;
import java.security.*;
import java.security.cert.*;

public interface CertificateService {
    Subject generateSubjectData(User subject);
    X509Certificate generateCertificate(CertificateDto certificateDto,Subject generatedSubjectData);
    Certificate saveCertificateDB(CertificateDto certificateDto, User subject);
    void createCertificate(CertificateDto certificateDto, Subject generatedSubjectData) throws CertificateException, IOException, KeyStoreException, NoSuchAlgorithmException;
    X509Certificate getCertificateByAlias(String alias, KeyStore keystore) throws KeyStoreException;
}
