package com.example.PKI.service.cert;

import com.example.PKI.dto.*;
import com.example.PKI.model.Certificate;
import com.example.PKI.model.Subject;
import com.example.PKI.model.User;

import java.io.*;
import java.security.*;
import java.security.cert.*;
import java.util.ArrayList;

public interface CertificateService {
    Subject generateSubjectData(Integer subjectId);

    X509Certificate generateCertificate(CertificateDto certificateDto, Subject generatedSubjectData) throws Exception;

    Certificate saveCertificateDB(CertificateDto certificateDto, Integer subjectId);

    void createCertificate(CertificateDto certificateDto, Subject generatedSubjectData) throws Exception;

    X509Certificate getCertificateByAlias(String alias, KeyStore keystore) throws KeyStoreException;

    void revokeCertificate(String serialNumber) throws Exception;

    boolean isCertificateValid(KeyStore keyStore, String alias) throws KeyStoreException, CertificateException, IOException, NoSuchAlgorithmException, NoSuchProviderException, SignatureException, InvalidKeyException;

    KeyStore getKeyStoreByAlias(String alias) throws KeyStoreException, CertificateException, IOException, NoSuchAlgorithmException, NoSuchProviderException;

    ArrayList<Certificate> getAllCertificates() throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException;

    ArrayList<Certificate> getAllUsersCertificates(String email);

    ArrayList<Certificate> getAllValidSignersForUser(String email,String startDate, String endDate) throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException;

    void generateCertificateByUser(CertificateDto certificateDto, Subject generatedSubjectData);

    ArrayList<User> getAllValidSignersForDateRange(String startDate, String endDate) throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException;

    ArrayList[] getAllCertificateChains(String email) throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException;
}
