package com.example.PKI.service.cert.impl;

import com.example.PKI.dto.*;
import com.example.PKI.model.*;
import com.example.PKI.repository.*;
import com.example.PKI.service.*;
import com.example.PKI.service.cert.CertificateService;
import org.bouncycastle.asn1.x500.*;
import org.bouncycastle.asn1.x500.style.*;
import org.bouncycastle.cert.*;
import org.bouncycastle.cert.jcajce.*;
import org.bouncycastle.jce.provider.*;
import org.bouncycastle.operator.*;
import org.bouncycastle.operator.jcajce.*;
import org.springframework.beans.factory.annotation.*;
import org.springframework.stereotype.*;

import java.io.*;
import java.math.*;
import java.security.*;
import java.security.cert.*;
import java.security.cert.Certificate;
import java.text.*;
import java.util.*;

@Service
public class CertificateServiceImpl implements CertificateService {

    @Autowired
    private KeyService keyService;
    @Autowired
    private KeyStoreService keyStoreService;
    @Autowired
    CertificateRepository repository;

    @Override
    public X509Certificate generateCertificate(SubjectWithPKDto subjectWithPKDto, String issuerSerialNumber, String type) {
        try{
            //ovo istraziti : on koristi SHA256withECDSA
            JcaContentSignerBuilder builder = new JcaContentSignerBuilder("SHA256WithRSAEncryption");
            Security.addProvider(new BouncyCastleProvider());
            builder = builder.setProvider("BC");
            KeyStore keyStore=keyStoreService.getKeyStore(keyService.getKeyStorePath(),keyService.getKeyStorePass());

            X500Name issuer = null;
            PrivateKey privKey = null;
            if(type.equals("ROOT")){
                issuer=subjectWithPKDto.getSubject().getX500Name();
                privKey=subjectWithPKDto.getPrivateKey();
            }
            else {
                X509Certificate issuerCert= (X509Certificate) keyStore.getCertificate(issuerSerialNumber);
                issuer=new JcaX509CertificateHolder(issuerCert).getSubject();
                privKey = (PrivateKey) keyStore.getKey(issuerSerialNumber, "key".toCharArray());
                boolean valid = validate(issuerSerialNumber);
            }


            ContentSigner contentSigner = builder.build(privKey);

            X509v3CertificateBuilder certGen = new JcaX509v3CertificateBuilder(issuer,
                    subjectWithPKDto.getSubject().getSerialNumber(),
                    subjectWithPKDto.getSubject().getStartDate(),
                    subjectWithPKDto.getSubject().getEndDate(),
                    subjectWithPKDto.getSubject().getX500Name(),
                    subjectWithPKDto.getSubject().getPublicKey());

            X509CertificateHolder certHolder = certGen.build(contentSigner);

            JcaX509CertificateConverter certConverter = new JcaX509CertificateConverter();
            BouncyCastleProvider bcp = new BouncyCastleProvider();
            certConverter = certConverter.setProvider(bcp);

            //Konvertuje objekat u sertifikat
            return certConverter.getCertificate(certHolder);

        }catch (CertificateEncodingException e) {
            e.printStackTrace();
        } catch (IllegalArgumentException e) {
            e.printStackTrace();
        } catch (IllegalStateException e) {
            e.printStackTrace();
        } catch (OperatorCreationException e) {
            e.printStackTrace();
        } catch (CertificateException e) {
            e.printStackTrace();
        } catch (UnrecoverableKeyException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        } catch (KeyStoreException e) {
            e.printStackTrace();
        } catch (NoSuchAlgorithmException e) {
            e.printStackTrace();
        }

        return null;
    }

    @Override
    public void saveCertificateDB(SubjectWithPKDto subjectWithPK) {
        CertificateType cerType;
        if (subjectWithPK.getSubject().getType().equals("ROOT")) {
            cerType = CertificateType.ROOT;
        } else if (subjectWithPK.getSubject().getType().equals("INTERMEDIATE")) {
            cerType = CertificateType.INTERMEDIATE;
        } else {
            cerType = CertificateType.CLIENT;
        }

        if (subjectWithPK.getSubject().getStartDate().compareTo(subjectWithPK.getSubject().getEndDate()) < 0) {
            com.example.PKI.model.Certificate certificate = new com.example.PKI.model.Certificate();
            certificate.setSerialNumber(subjectWithPK.getSubject().getSerialNumber().toString());
            certificate.setType(cerType);
            certificate.setValid(true);
            certificate.setSubjectCommonName(subjectWithPK.getSubject().getCommonName()+"organizacija:"+subjectWithPK.getSubject().getOrganization());

            repository.save(certificate);
        }
    }

    @Override
    public void createCertificate(SubjectWithPKDto subjectWithPK, String type,String issuerSerialNumber) throws CertificateException, IOException, KeyStoreException, NoSuchAlgorithmException {
        X509Certificate certificate = generateCertificate(subjectWithPK,issuerSerialNumber,type);
        String keyPass="key";
        KeyStore keyStore = keyStoreService.getKeyStore(keyService.getKeyStorePath(),keyService.getKeyStorePass());
        keyStoreService.store(keyService.getKeyStorePass(),keyPass,new Certificate[]{certificate},subjectWithPK.getPrivateKey(), subjectWithPK.getSubject().getAlias(),keyService.getKeyStorePath());

    }

    @Override
    public SubjectWithPKDto generateSubjectData(SubjectDto subjectDto) {
        try {
            KeyPair keyPairSubject = keyService.generateKeyPair();

            //Datumi od kad do kad vazi sertifikat
            SimpleDateFormat iso8601Formater = new SimpleDateFormat("yyyy-MM-dd");
           // Date startDate = iso8601Formater.parse(subjectDto.getStartDate().toString());
           // Date endDate = iso8601Formater.parse(subjectDto.getEndDate().toString());
            //da se ne zezam sad s datumima
            Date startDate = iso8601Formater.parse("2021-12-31");
            Date endDate = iso8601Formater.parse("2022-12-31");


            //Serijski broj sertifikata
            subjectDto.setAlias(keyService.getSerialNumber().toString());
            BigInteger serialNumber = new BigInteger(subjectDto.getAlias());

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
            //msm da u subject treba i private key
            Subject subject = new Subject(keyPairSubject.getPublic(), builder.build(), serialNumber, startDate, endDate);
            subject.setAlias(subjectDto.getAlias());
            subject.setType(subjectDto.getType());
            return new SubjectWithPKDto(subject,keyPairSubject.getPrivate());
        } catch (ParseException e) {
            e.printStackTrace();
        }
        return null;
    }


    private boolean validate(String alias) throws CertificateException, NoSuchAlgorithmException, KeyStoreException, IOException {
        KeyStore keyStore=keyStoreService.getKeyStore(keyService.getKeyStorePath(),keyService.getKeyStorePass());
        X509Certificate certificate= (X509Certificate) keyStore.getCertificate(alias);
        Certificate[] chain=keyStore.getCertificateChain(alias);
        return false;
    }
}
