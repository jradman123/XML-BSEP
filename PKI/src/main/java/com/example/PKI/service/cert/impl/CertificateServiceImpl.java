package com.example.PKI.service.cert.impl;

import com.example.PKI.dto.SubjectDto;
import com.example.PKI.dto.SubjectWithPkDto;
import com.example.PKI.model.CertificateType;
import com.example.PKI.model.Subject;
import com.example.PKI.repository.CertificateRepository;
import com.example.PKI.service.KeyService;
import com.example.PKI.service.KeyStoreService;
import com.example.PKI.service.cert.CertificateService;
import org.bouncycastle.asn1.x500.X500Name;
import org.bouncycastle.asn1.x500.X500NameBuilder;
import org.bouncycastle.asn1.x500.style.BCStyle;
import org.bouncycastle.cert.X509CertificateHolder;
import org.bouncycastle.cert.X509v3CertificateBuilder;
import org.bouncycastle.cert.jcajce.JcaX509CertificateConverter;
import org.bouncycastle.cert.jcajce.JcaX509CertificateHolder;
import org.bouncycastle.cert.jcajce.JcaX509v3CertificateBuilder;
import org.bouncycastle.jce.provider.BouncyCastleProvider;
import org.bouncycastle.operator.ContentSigner;
import org.bouncycastle.operator.OperatorCreationException;
import org.bouncycastle.operator.jcajce.JcaContentSignerBuilder;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.math.BigInteger;
import java.security.*;
import java.security.cert.Certificate;
import java.security.cert.CertificateEncodingException;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.Enumeration;

@Service
public class CertificateServiceImpl implements CertificateService {

    @Autowired
    private KeyService keyService;
    @Autowired
    private KeyStoreService keyStoreService;
    @Autowired
    CertificateRepository repository;

    @Override
    public X509Certificate generateCertificate(SubjectWithPkDto subjectWithPKDto, String issuerSerialNumber) {
        try {
            //ovo istraziti : on koristi SHA256withECDSA
            JcaContentSignerBuilder builder = new JcaContentSignerBuilder("SHA256WithRSAEncryption");
            Security.addProvider(new BouncyCastleProvider());
            builder.setProvider("BC");
            KeyStore keyStore = keyStoreService.getKeyStore(keyService.getKeyStorePath("ROOT"), keyService.getKeyStorePass());
            Enumeration<String> aliases = keyStore.aliases();
            X500Name issuer;
            PrivateKey privateKey;
            if (subjectWithPKDto.getSubject().getType().equals("ROOT")) {
                issuer = subjectWithPKDto.getSubject().getX500Name();
                privateKey = subjectWithPKDto.getPrivateKey();
            } else {
                //TODO: Ne dobavi ga
                X509Certificate issuerCert = (X509Certificate) keyStore.getCertificate(issuerSerialNumber);
                issuer = new JcaX509CertificateHolder(issuerCert).getSubject();
                privateKey = (PrivateKey) keyStore.getKey(issuerSerialNumber, "key".toCharArray());
                //boolean valid = validate(issuerSerialNumber); //popraviti metodu da zaista validira issuera
            }


            ContentSigner contentSigner = builder.build(privateKey);

            X509v3CertificateBuilder certGen = new JcaX509v3CertificateBuilder(issuer,
                    subjectWithPKDto.getSubject().getSerialNumber(),
                    subjectWithPKDto.getSubject().getStartDate(),
                    subjectWithPKDto.getSubject().getEndDate(),
                    subjectWithPKDto.getSubject().getX500Name(),
                    subjectWithPKDto.getSubject().getPublicKey());

            X509CertificateHolder certHolder = certGen.build(contentSigner);

            JcaX509CertificateConverter certConverter = new JcaX509CertificateConverter();
            BouncyCastleProvider bcp = new BouncyCastleProvider();
            certConverter.setProvider(bcp);

            //Konvertuje objekat u sertifikat
            return certConverter.getCertificate(certHolder);

        } catch (CertificateEncodingException e) {
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
    public com.example.PKI.model.Certificate saveCertificateDb(SubjectWithPkDto subjectWithPK) {
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
            certificate.setSubjectCommonName(subjectWithPK.getSubject().getCommonName() + " organizacija: " + subjectWithPK.getSubject().getOrganization());

            return repository.save(certificate);
        }
        return null;
    }

    @Override
    public void createCertificate(SubjectWithPkDto subjectWithPkDto, String issuerSerialNumber) throws CertificateException, IOException, KeyStoreException, NoSuchAlgorithmException, SignatureException, InvalidKeyException, NoSuchProviderException, UnrecoverableKeyException {
        X509Certificate certificate = generateCertificate(subjectWithPkDto, issuerSerialNumber);
        KeyStore keyStore = keyStoreService.getKeyStore(keyService.getKeyStorePath(subjectWithPkDto.getSubject().getType()), keyService.getKeyStorePass());
        //TODO: Verify
        verifyCertificate(issuerSerialNumber);

        keyStoreService.store(keyService.getKeyStorePass(), keyService.getKeyPass(), new Certificate[]{certificate}, subjectWithPkDto.getPrivateKey(), subjectWithPkDto.getSubject().getAlias(), keyService.getKeyStorePath(subjectWithPkDto.getSubject().getType()));

    }

    private void verifyCertificate(String issuerSerialNumber) throws CertificateException, KeyStoreException, NoSuchAlgorithmException, IOException, SignatureException, InvalidKeyException, NoSuchProviderException, UnrecoverableKeyException {
//        KeyStore keyStore = keyStoreService.getKeyStore(keyService.getKeyStorePath("ROOT"), keyService.getKeyStorePass());
//        X509Certificate certificate= (X509Certificate) keyStore.getCertificate(issuerSerialNumber);
//        Certificate[] chain=keyStore.getCertificateChain(issuerSerialNumber);
//
//        certificate.verify(certificate.getPublicKey());
        KeyStore keyStore = KeyStore.getInstance("JSK");
        Key key = keyStore.getKey(issuerSerialNumber, "password".toCharArray());
        Certificate issuer = keyStore.getCertificate(issuerSerialNumber);

    }

    @Override
    public SubjectWithPkDto generateSubjectData(SubjectDto subjectDto) {
        try {
            KeyPair keyPairSubject = keyService.generateKeyPair();

            //Datumi od kad do kad vazi sertifikat
            SimpleDateFormat iso8601Formater = new SimpleDateFormat("yyyy-MM-dd");
            /*System.out.println(subjectDto.getStartDate().toString());
            Date startDate = iso8601Formater.parse(subjectDto.getStartDate().toString());
            Date endDate = iso8601Formater.parse(subjectDto.getEndDate().toString());*/
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

            builder.addRDN(BCStyle.UID, subjectDto.getEmail());

            //Kreiraju se podaci za sertifikat, sto ukljucuje:
            // - javni kljuc koji se vezuje za sertifikat
            // - podatke o vlasniku
            // - serijski broj sertifikata
            // - od kada do kada vazi sertifikat
            //msm da u subject treba i private key
            Subject subject = new Subject(keyPairSubject.getPublic(), builder.build(), serialNumber, startDate, endDate, subjectDto.getType(), subjectDto.getAlias(), subjectDto.getCommonName(), subjectDto.getOrganization());
            return new SubjectWithPkDto(subject, keyPairSubject.getPrivate());
        } catch (ParseException e) {
            e.printStackTrace();
        }
        return null;
    }


   /* private boolean validate(String alias) throws CertificateException, NoSuchAlgorithmException, KeyStoreException, IOException {
        KeyStore keyStore=keyStoreService.getKeyStore(keyService.getKeyStorePath(),keyService.getKeyStorePass());
        X509Certificate certificate= (X509Certificate) keyStore.getCertificate(alias);
        Certificate[] chain=keyStore.getCertificateChain(alias);
        return false;
    }*/
}