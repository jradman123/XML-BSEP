package com.example.PKI.service.impl;

import com.example.PKI.service.*;
import org.springframework.stereotype.*;

import java.math.*;
import java.security.*;
import java.security.spec.*;

@Service
public class KeyServiceImpl implements KeyService {

    @Override
    public BigInteger getSerialNumber() {
        return new BigInteger(64, new SecureRandom());
    }


    @Override
    public String getKeyStorePass() {
        return "password";
    }

    @Override
    public String getKeyStorePath() {
        return "test.pks";
    }

    @Override
    public KeyPair generateKeyPair() {
        try {
            KeyPairGenerator keyGen = KeyPairGenerator.getInstance("RSA");
            SecureRandom random = SecureRandom.getInstance("SHA1PRNG", "SUN");
            keyGen.initialize(2048, random);
            return keyGen.generateKeyPair();
        } catch (NoSuchAlgorithmException e) {
            e.printStackTrace();
        } catch (NoSuchProviderException e) {
            e.printStackTrace();
        }
        return null;
    }
    }

