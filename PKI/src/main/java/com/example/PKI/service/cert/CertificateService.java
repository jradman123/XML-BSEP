package com.example.PKI.service.cert;

import com.example.PKI.dto.*;
import com.example.PKI.model.Certificate;
import com.example.PKI.model.Subject;
import com.example.PKI.model.User;
import org.bouncycastle.jcajce.provider.asymmetric.X509;
import org.springframework.http.HttpStatus;

import java.io.*;
import java.security.*;
import java.security.cert.*;
import java.util.ArrayList;

public interface CertificateService {
    Subject generateSubjectData(Integer subjectId);

    com.example.PKI.model.Certificate generateCertificate(CertificateDto certificateDto, Subject generatedSubjectData) throws Exception;

    Certificate saveCertificateDB(CertificateDto certificateDto, Integer subjectId);

    com.example.PKI.model.Certificate createCertificate(CertificateDto certificateDto, Subject generatedSubjectData) throws Exception;

    X509Certificate getCertificateByAlias(String alias, KeyStore keystore) throws KeyStoreException;

    void revokeCertificate(String serialNumber) throws Exception;

    boolean isCertificateValid(KeyStore keyStore, String alias) throws KeyStoreException, CertificateException, IOException, NoSuchAlgorithmException, NoSuchProviderException, SignatureException, InvalidKeyException;

    KeyStore getKeyStoreByAlias(String alias) throws KeyStoreException, CertificateException, IOException, NoSuchAlgorithmException, NoSuchProviderException;

    ArrayList<Certificate> getAllCertificates() throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException;

    ArrayList<IssuerDto> getAllValidSignersForDateRange(String startDate, String endDate) throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException;
}
