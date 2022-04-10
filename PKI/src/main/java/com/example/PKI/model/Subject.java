package com.example.PKI.model;

import java.math.*;
import java.security.KeyPair;
import java.security.PublicKey;
import java.util.Date;

import org.bouncycastle.asn1.x500.X500Name;

import lombok.Data;

import javax.security.auth.x500.*;

@Data
public class Subject {
    private X500Name x500Name;
    private KeyPair keyPair;

    public Subject() {
    }

    public Subject(X500Name x500Name, KeyPair keyPair) {
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

