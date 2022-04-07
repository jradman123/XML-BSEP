package com.example.PKI.service;

import java.io.*;
import java.security.*;
import java.security.cert.*;
import java.security.cert.Certificate;
import java.util.*;

public interface KeyStoreService {
    List<X509Certificate> getCertificates(String keyStorePass) throws CertificateException, NoSuchAlgorithmException, KeyStoreException, IOException;
    KeyStore getKeyStore(String keyStorePath, String keyStorePassword) throws IOException, KeyStoreException, CertificateException, NoSuchAlgorithmException;
    void store(String keyStorePassword, String keyPassword, Certificate[] chain, PrivateKey privateKey, String alias, String keyStorePath) throws CertificateException, NoSuchAlgorithmException, KeyStoreException, IOException;
}
