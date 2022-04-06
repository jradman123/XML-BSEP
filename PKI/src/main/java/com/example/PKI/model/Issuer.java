package com.example.PKI.model;

import java.security.PrivateKey;

import org.bouncycastle.asn1.x500.X500Name;

import lombok.Data;

@Data
public class Issuer {
    public Issuer(PrivateKey privKey, X500Name issuerName) {
    }

    public X500Name getX500Name() {
        return x500Name;
    }

    public void setX500Name(X500Name x500Name) {
        this.x500Name = x500Name;
    }

    public PrivateKey getPrivateKey() {
        return privateKey;
    }

    public void setPrivateKey(PrivateKey privateKey) {
        this.privateKey = privateKey;
    }

    private X500Name x500Name;
    private PrivateKey privateKey;
}