package com.example.PKI.service.cert;

import com.example.PKI.dto.*;
import com.example.PKI.model.Certificate;
import com.example.PKI.model.Issuer;
import com.example.PKI.model.Subject;

import java.io.*;
import java.security.*;
import java.security.cert.*;


public interface CertificateService {
    SubjectWithPKDto generateSubjectData(SubjectDto subjectDto);
    X509Certificate generateCertificate(SubjectWithPKDto subjectWithPKDto, String issuerSerialNumber);
    Certificate saveCertificateDB(SubjectWithPKDto subjectWithPK);
    void createCertificate(SubjectWithPKDto subjectWithPK,String issuerSerialNumber) throws CertificateException, IOException, KeyStoreException, NoSuchAlgorithmException;
}
