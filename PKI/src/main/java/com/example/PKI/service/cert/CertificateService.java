package com.example.PKI.service.cert;

import com.example.PKI.dto.*;
import com.example.PKI.model.Issuer;
import com.example.PKI.model.Subject;

import java.io.*;
import java.security.*;
import java.security.cert.*;


public interface CertificateService {
    public SubjectWithPKDto generateSubjectData(SubjectDto subjectDto);
    public X509Certificate generateCertificate(SubjectWithPKDto subjectWithPKDto, String issuerSerialNumber, String type);
    public void saveCertificateDB(SubjectWithPKDto subjectWithPK);
    void createCertificate(SubjectWithPKDto subjectWithPK, String type,String issuerSerialNumber) throws CertificateException, IOException, KeyStoreException, NoSuchAlgorithmException;
}
