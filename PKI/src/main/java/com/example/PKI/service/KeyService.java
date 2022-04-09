package com.example.PKI.service;

import java.math.BigInteger;
import java.security.KeyPair;

public interface KeyService {

    BigInteger getSerialNumber();

    String getKeyStorePass();

    String getKeyStorePath(String type);

    KeyPair generateKeyPair();

    String getKeyPass();
}
