package com.example.PKI.model;

import java.security.KeyPair;
import java.security.PrivateKey;

import org.bouncycastle.asn1.x500.X500Name;

import lombok.Data;

@Data
public class Issuer {
    private X500Name x500Name;
    private KeyPair keyPair;

    public Issuer() {
    }

    public Issuer(X500Name x500Name, KeyPair keyPair) {
        this.x500Name = x500Name;
        this.keyPair = keyPair;
    }

    public X500Name getX500Name() {
        return x500Name;
    }

    public void setX500Name(X500Name x500Name) {
        this.x500Name = x500Name;
    }

    public KeyPair getKeyPair() {
        return keyPair;
    }

    public void setKeyPair(KeyPair keyPair) {
        this.keyPair = keyPair;
    }
}