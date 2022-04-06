package com.example.PKI.dto;

import org.bouncycastle.asn1.x500.*;

import java.security.*;
import java.util.*;

public class IssuerDto {
    private PrivateKey privateKey;
    private X500Name x500Name;
    private String alias;
}
