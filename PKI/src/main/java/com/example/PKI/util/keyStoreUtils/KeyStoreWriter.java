package com.example.PKI.util.keyStoreUtils;

import org.springframework.stereotype.Component;

import java.io.*;
import java.security.*;
import java.security.cert.Certificate;
import java.security.cert.CertificateException;


@Component
public class KeyStoreWriter {

    private KeyStore keyStore;

    public KeyStoreWriter() {
        try {
            keyStore = KeyStore.getInstance("JKS", "SUN");
        } catch (KeyStoreException e) {
            e.printStackTrace();
        } catch (NoSuchProviderException e) {
            e.printStackTrace();
        }
    }

    public void loadKeyStore(String fileName, char[] password) {
        try {
            String fullPath = fileName + ".jks";
            File f = new File(fullPath);
            if (fullPath != null && f.exists()) {
                keyStore.load(new FileInputStream(fullPath), password);
            } else {
                keyStore.load(null, password);

            }
        } catch (NoSuchAlgorithmException e) {
            e.printStackTrace();
        } catch (CertificateException e) {
            e.printStackTrace();
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }

    }

    public void saveKeyStore(String fileName, char[] password) {
        try {
            keyStore.store(new FileOutputStream(fileName + ".jks"), password);
        } catch (KeyStoreException e) {
            e.printStackTrace();
        } catch (NoSuchAlgorithmException e) {
            e.printStackTrace();
        } catch (CertificateException e) {
            e.printStackTrace();
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }


    public void write(String alias, PrivateKey privateKey, String password, Certificate[] chain) {
        try {
            keyStore.setKeyEntry(alias, privateKey, password.toCharArray(), chain);
        } catch (KeyStoreException e) {
            e.printStackTrace();
        }
    }


}
