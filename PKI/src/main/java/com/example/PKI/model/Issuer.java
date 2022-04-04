package com.example.PKI.model;

import java.security.PrivateKey;

import org.bouncycastle.asn1.x500.X500Name;

import lombok.Data;

@Data
public class Issuer {
    private X500Name x500Name;
    private PrivateKey privateKey;
}