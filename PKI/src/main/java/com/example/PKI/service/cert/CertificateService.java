package com.example.PKI.service.cert;

import com.example.PKI.dto.SubjectDto;
import com.example.PKI.dto.SubjectWithPkDto;
import com.example.PKI.model.Certificate;

import java.io.IOException;
import java.security.*;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;

public interface CertificateService {

    SubjectWithPkDto generateSubjectData(SubjectDto subjectDto);

    X509Certificate generateCertificate(SubjectWithPkDto subjectWithPKDto, String issuerSerialNumber);

    Certificate saveCertificateDb(SubjectWithPkDto subjectWithPK);

    void createCertificate(SubjectWithPkDto subjectWithPK, String issuerSerialNumber) throws CertificateException, IOException, KeyStoreException, NoSuchAlgorithmException, SignatureException, InvalidKeyException, NoSuchProviderException, UnrecoverableKeyException;
}
