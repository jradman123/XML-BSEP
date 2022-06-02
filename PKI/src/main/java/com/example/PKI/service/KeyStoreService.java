package com.example.PKI.service;

import java.io.IOException;
import java.security.KeyStore;
import java.security.KeyStoreException;
import java.security.NoSuchAlgorithmException;
import java.security.PrivateKey;
import java.security.cert.Certificate;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;
import java.util.List;

public interface KeyStoreService {
    List<X509Certificate> getCertificates(String keyStorePass) throws CertificateException, NoSuchAlgorithmException, KeyStoreException, IOException;

    KeyStore getKeyStore(String keyStorePath, String keyStorePassword) throws IOException, KeyStoreException, CertificateException, NoSuchAlgorithmException;

    void store(String keyStorePassword, String keyPassword, Certificate[] chain, PrivateKey privateKey, String alias, String keyStorePath) throws CertificateException, NoSuchAlgorithmException, KeyStoreException, IOException;
}
