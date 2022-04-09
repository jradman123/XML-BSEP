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
    private CertificateRepository repository;
    @Autowired
    private UserRepository userRepository;
    private BigInteger serialNumber;

    @Override
    public X509Certificate generateCertificate(CertificateDto certificateDto,Subject generatedSubjectData) {
        try{
            System.out.println(certificateDto.getSubjectId().toString());
            SimpleDateFormat iso8601Formater = new SimpleDateFormat("yyyy-MM-dd");
            /*System.out.println(subjectDto.getStartDate().toString());
            Date startDate = iso8601Formater.parse(subjectDto.getStartDate().toString());
            Date endDate = iso8601Formater.parse(subjectDto.getEndDate().toString());*/
            Date startDate = iso8601Formater.parse("2021-12-31");
            Date endDate = iso8601Formater.parse("2022-12-31");
            serialNumber = new BigInteger(keyService.getSerialNumber().toString());
            //ovo istraziti : on koristi SHA256withECDSA
            JcaContentSignerBuilder builder = new JcaContentSignerBuilder("SHA256WithRSAEncryption");
            Security.addProvider(new BouncyCastleProvider());
            builder = builder.setProvider("BC");
            KeyStore keyStore=keyStoreService.getKeyStore(keyService.getKeyStorePath(certificateDto.getType()),keyService.getKeyStorePass());
            
            X500Name issuer = null;
            PrivateKey privateKey = null;
            if(certificateDto.getType().equals("ROOT")){
                issuer = generatedSubjectData.getX500Name();
                privateKey = generatedSubjectData.getKeyPair().getPrivate();
            }
            else {
                User issuerData = userRepository.findById(certificateDto.getIssuerId()).get();
                X509Certificate issuerCert= (X509Certificate) keyStore.getCertificate(issuerData.getEmail() + issuerData.getId());
                issuer=new JcaX509CertificateHolder(issuerCert).getSubject();
                privateKey = (PrivateKey) keyStore.getKey(certificateDto.getIssuerSerialNumber(), "key".toCharArray());
            }

            ContentSigner contentSigner = builder.build(privateKey);

            X509v3CertificateBuilder certGen = new JcaX509v3CertificateBuilder(issuer,
                    serialNumber,
                    startDate,
                    endDate,
                    generatedSubjectData.getX500Name(),
                    generatedSubjectData.getKeyPair().getPublic());

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
        } catch (ParseException e) {
            e.printStackTrace();
        }

        return null;
    }

    @Override
    public com.example.PKI.model.Certificate saveCertificateDB(CertificateDto subjectDto, User subject) {
        CertificateType cerType;
        if (subjectDto.getType().equals("ROOT")) {
            cerType = CertificateType.ROOT;
        } else if (subjectDto.getType().equals("INTERMEDIATE")) {
            cerType = CertificateType.INTERMEDIATE;
        } else {
            cerType = CertificateType.CLIENT;
        }

        if (subjectDto.getStartDate().compareTo(subjectDto.getEndDate()) < 0) {
            com.example.PKI.model.Certificate certificate = new com.example.PKI.model.Certificate();
            certificate.setSerialNumber(serialNumber.toString());
            certificate.setType(cerType);
            certificate.setValid(true);
            certificate.setSubjectCommonName(subject.getCommonName()+" "+subject.getOrganization());
            certificate.setValidFrom(subjectDto.getStartDate());
            certificate.setValidTo(subjectDto.getEndDate());
            return repository.save(certificate);
        }
        return null;
    }

    @Override
    public void createCertificate(CertificateDto certificateDto, Subject generatedSubjectData) throws CertificateException, IOException, KeyStoreException, NoSuchAlgorithmException {
        User user = userRepository.findById(certificateDto.getSubjectId()).get();
        String alias = user.getEmail() + user.getId();
        X509Certificate certificate = generateCertificate(certificateDto,generatedSubjectData);
        KeyStore keyStore = keyStoreService.getKeyStore(keyService.getKeyStorePath(certificateDto.getType()),keyService.getKeyStorePass());
        keyStoreService.store(keyService.getKeyStorePass(),keyService.getKeyPass(),new Certificate[]{certificate},generatedSubjectData.getKeyPair().getPrivate(), alias,keyService.getKeyStorePath(certificateDto.getType()));

    }

    @Override
    public X509Certificate getCertificateByAlias(String alias, KeyStore keystore) throws KeyStoreException {
        X509Certificate certificate = (X509Certificate) keystore.getCertificate(alias);
        return certificate;
    }

    @Override
    public Subject generateSubjectData(User subject) {
        KeyPair keyPairSubject = keyService.generateKeyPair();

        X500NameBuilder builder = new X500NameBuilder(BCStyle.INSTANCE);
        builder.addRDN(BCStyle.CN, subject.getCommonName());
        builder.addRDN(BCStyle.O, subject.getOrganization());
        builder.addRDN(BCStyle.OU, subject.getOrganizationUnit());
        builder.addRDN(BCStyle.C, subject.getCountry());
        builder.addRDN(BCStyle.E, subject.getEmail());

        return new Subject(builder.build(), keyPairSubject);
    }


   /* private boolean validate(String alias) throws CertificateException, NoSuchAlgorithmException, KeyStoreException, IOException {
        KeyStore keyStore=keyStoreService.getKeyStore(keyService.getKeyStorePath(),keyService.getKeyStorePass());
        X509Certificate certificate= (X509Certificate) keyStore.getCertificate(alias);
        Certificate[] chain=keyStore.getCertificateChain(alias);
        return false;
    }*/
}
