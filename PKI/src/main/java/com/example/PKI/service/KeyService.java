package com.example.PKI.service;

import java.math.*;
import java.security.*;

public interface KeyService {
    BigInteger getSerialNumber();

    String getKeyStorePass();

    String getKeyStorePath(String type);

    KeyPair generateKeyPair();

    String getKeyPass();
}
