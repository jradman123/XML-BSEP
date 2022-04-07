package com.example.PKI.service.impl;

import com.example.PKI.service.*;
import org.springframework.beans.factory.annotation.*;
import org.springframework.stereotype.*;

import java.io.*;
import java.security.*;
import java.security.cert.*;
import java.security.cert.Certificate;
import java.util.*;

@Service
public class KeyStoreServiceImpl implements KeyStoreService {

    @Autowired
    private KeyService keyService;

    @Override
    public void store(String keyStorePassword, String keyPassword, Certificate[] chain, PrivateKey privateKey, String alias, String keyStorePath) throws CertificateException, NoSuchAlgorithmException, KeyStoreException, IOException {
        char[] keyStorePasswordChars=keyStorePassword.toCharArray();
        char[] keyPasswordChars=keyPassword.toCharArray();

        KeyStore keyStore=getKeyStore(keyStorePath,keyStorePassword);
        KeyStore.PrivateKeyEntry privKeyEntry = new KeyStore.PrivateKeyEntry(privateKey,
                chain);
        keyStore.setEntry(alias, privKeyEntry, new KeyStore.PasswordProtection(keyPasswordChars));
        keyStore.store(new FileOutputStream(keyStorePath),keyStorePasswordChars);
    }

    @Override
    public KeyStore getKeyStore(String keyStorePath, String keyStorePassword) throws IOException, KeyStoreException, CertificateException, NoSuchAlgorithmException {
        char[] keyStorePasswordChars=keyStorePassword.toCharArray();
        KeyStore keyStore=KeyStore.getInstance("PKCS12");
        try{
            keyStore.load(new FileInputStream(keyStorePath),keyStorePasswordChars);
        }catch (FileNotFoundException e){
            keyStore.load(null,keyStorePasswordChars);
        }
        return keyStore;
    }

    //metodu treba popraviti jer se treba proslijediti tip sertifikata koji nam trebaju
    @Override
    public List<X509Certificate> getCertificates(String keyStorePass) throws CertificateException, NoSuchAlgorithmException, KeyStoreException, IOException {
        String keyStorePath=keyService.getKeyStorePath("ROOT");
        KeyStore keyStore=getKeyStore(keyStorePath,keyStorePass);
        List<X509Certificate>certificateList=new ArrayList<>();
        Enumeration<String> aliass= keyStore.aliases();

        while(aliass.hasMoreElements()){
            String alias=aliass.nextElement();
            X509Certificate certificate=(X509Certificate)keyStore.getCertificate(alias);
            certificateList.add(certificate);
        }
        return certificateList;
    }
}
