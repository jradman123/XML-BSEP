package com.example.PKI.service.cert.impl;

import com.example.PKI.dto.*;
import com.example.PKI.model.Issuer;
import com.example.PKI.model.Subject;
import com.example.PKI.service.cert.CertificateService;
import org.bouncycastle.asn1.x500.*;
import org.bouncycastle.asn1.x500.style.*;

import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;

public class CertificateServiceImpl implements CertificateService {

    @Override
    public Subject generateSubjectData(SubjectDto subjectDto) {
        try {
            KeyPair keyPairSubject = generateKeyPair();

            //Datumi od kad do kad vazi sertifikat
            SimpleDateFormat iso8601Formater = new SimpleDateFormat("yyyy-MM-dd");
            Date startDate = iso8601Formater.parse(subjectDto.getStartDate().toString());
            Date endDate = iso8601Formater.parse(subjectDto.getEndDate().toString());

            //Serijski broj sertifikata
            subjectDto.setAlias((new BigInteger(64,new SecureRandom())).toString());
            BigInteger serialNumber=new BigInteger(subjectDto.getAlias());

            //klasa X500NameBuilder pravi X500Name objekat koji predstavlja podatke o vlasniku
            X500NameBuilder builder = new X500NameBuilder(BCStyle.INSTANCE);
            builder.addRDN(BCStyle.CN, subjectDto.getCommonName());
            //builder.addRDN(BCStyle.SURNAME, "Sladic");
            //builder.addRDN(BCStyle.GIVENNAME, "Goran");
            builder.addRDN(BCStyle.O, subjectDto.getOrganization());
            builder.addRDN(BCStyle.OU, subjectDto.getOrganizationUnit());
            builder.addRDN(BCStyle.C, subjectDto.getCountry());
            builder.addRDN(BCStyle.E, subjectDto.getEmail());
            //UID (USER ID) je ID korisnika
            //NE ZNAM STA SA OVIM
            //IZVUCI USERA IZ BAZE KOJI IMA ISTI EMAIL I NJEGOV ID DODJELITI KAO UID?
            builder.addRDN(BCStyle.UID, "123456");

            //Kreiraju se podaci za sertifikat, sto ukljucuje:
            // - javni kljuc koji se vezuje za sertifikat
            // - podatke o vlasniku
            // - serijski broj sertifikata
            // - od kada do kada vazi sertifikat
            return new Subject(keyPairSubject.getPublic(), builder.build(), serialNumber, startDate, endDate);
        } catch (ParseException e) {
            e.printStackTrace();
        }
        return null;
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
