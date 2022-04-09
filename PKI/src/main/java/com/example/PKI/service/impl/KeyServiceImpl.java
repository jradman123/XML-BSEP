package com.example.PKI.service.impl;

import com.example.PKI.service.KeyService;
import org.springframework.stereotype.Service;

import java.math.BigInteger;
import java.security.*;

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
    public String getKeyStorePath(String type) {
        if (type.equals("ROOT")) {
            return "roots";
        } else if (type.equals("INTERMEDIATE")) {
            return "intermediates";
        } else {
            return "clients";
        }
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

    @Override
    public String getKeyPass() {
        return "key";
    }
}

