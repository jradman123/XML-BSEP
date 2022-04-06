package com.example.PKI.service;

import java.math.*;
import java.security.*;

public interface KeyService {
    public BigInteger getSerialNumber();
    public String getKeyStorePass();
    public String getKeyStorePath();
    public KeyPair generateKeyPair();
}
