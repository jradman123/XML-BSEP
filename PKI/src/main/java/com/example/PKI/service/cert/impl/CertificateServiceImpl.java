package com.example.PKI.service.cert.impl;

import com.example.PKI.dto.CertificateDto;
import com.example.PKI.model.CertificateType;
import com.example.PKI.model.Subject;
import com.example.PKI.model.User;
import com.example.PKI.repository.CertificateRepository;
import com.example.PKI.repository.UserRepository;
import com.example.PKI.service.KeyService;
import com.example.PKI.service.KeyStoreService;
import com.example.PKI.service.cert.CertificateService;
import com.example.PKI.util.keyStoreUtils.KeyStoreReader;
import com.example.PKI.util.keyStoreUtils.KeyStoreWriter;
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
import org.springframework.context.annotation.Bean;
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
import java.time.LocalDate;
import java.util.Date;

@Service
public class CertificateServiceImpl implements CertificateService {

    @Autowired
    private KeyService keyService;
    @Autowired
    private KeyStoreService keyStoreService;
    @Autowired
    private CertificateRepository certificateRepository;
    @Autowired
    private UserRepository userRepository;

    private BigInteger serialNumber;

    @Autowired
    private KeyStoreWriter keyStoreWriter;

    @Autowired
    private KeyStoreReader keyStoreReader;


    @Override
    public X509Certificate generateCertificate(CertificateDto certificateDto, Subject generatedSubjectData) throws Exception {
        try {

            if (LocalDate.parse(certificateDto.getEndDate()).compareTo(LocalDate.parse(certificateDto.getStartDate())) < 0) {
                throw new Exception("END_DATE_BEFORE_START");
            }

            // TODO: provjeri je l ima usera i issuera u bazii brrrr

            SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd");
            Date startDate = sdf.parse(certificateDto.getStartDate());
            Date endDate = sdf.parse(certificateDto.getEndDate());

            //ovo istraziti : on koristi SHA256withECDSA
            JcaContentSignerBuilder builder = new JcaContentSignerBuilder("SHA256WithRSAEncryption");
            Security.addProvider(new BouncyCastleProvider());
            builder.setProvider("BC");
            //TODO: zamijeniti sa readerom eyyy
            KeyStore keyStore = keyStoreService.getKeyStore(keyService.getKeyStorePath(certificateDto.getType()), keyService.getKeyStorePass());

            X500Name issuer;
            PrivateKey privateKey;
            KeyStore issuerKeyStore;
            String issuerAlias;
            serialNumber = keyService.getSerialNumber();

            if (certificateDto.getType().equals("ROOT")) {
                issuerAlias = serialNumber.toString();
                issuer = generatedSubjectData.getX500Name();
                issuerKeyStore = keyStoreService.getKeyStore(keyService.getKeyStorePath(CertificateType.ROOT.toString()),
                        keyService.getKeyStorePass());
                privateKey = generatedSubjectData.getKeyPair().getPrivate();
            } else {
                issuerAlias = certificateDto.getIssuerSerialNumber();

                CertificateType typebyy = certificateRepository.findTypeBySerialNumber(certificateDto.getIssuerSerialNumber());
                String kezstorpaff = keyService.getKeyStorePath(typebyy.toString());
                //issuerKeyStore = keyStoreService.getKeyStore(kezstorpaff, keyService.getKeyStorePass());

                issuerKeyStore = keyStoreReader.getKeyStore(kezstorpaff, keyService.getKeyStorePass());

                X509Certificate ceee = (X509Certificate) issuerKeyStore.getCertificate(issuerAlias);
                issuer = new JcaX509CertificateHolder(ceee).getSubject();
                privateKey = (PrivateKey) issuerKeyStore.getKey(issuerAlias, keyService.getKeyPass().toCharArray());
            }
            // signer
            ContentSigner contentSigner = builder.build(privateKey);

            X509v3CertificateBuilder certGen = new JcaX509v3CertificateBuilder(issuer,
                    serialNumber,
                    startDate,
                    endDate,
                    generatedSubjectData.getX500Name(),
                    generatedSubjectData.getKeyPair().getPublic());

            X509CertificateHolder certHolder = certGen.build(contentSigner);

            JcaX509CertificateConverter certConverter = new JcaX509CertificateConverter();
            certConverter.setProvider("BC");

            X509Certificate cert = certConverter.getCertificate(certHolder);

            Certificate[] certificateChain;
            Certificate[] certificates = issuerKeyStore.getCertificateChain(issuerAlias);
            if (certificates != null) {
                certificateChain = new Certificate[certificates.length + 1];
                certificateChain[0] = cert;
                for (int i = 0; i < certificates.length; i++)
                    certificateChain[i + 1] = certificates[i];
            } else
                certificateChain = new Certificate[]{cert};

            // ja bih ovo uradila on init !!! ovde mi je ruzno TODO
            keyStoreWriter.loadKeyStore(keyService.getKeyStorePath(certificateDto.getType()), keyService.getKeyStorePass().toCharArray());
            keyStoreWriter.write(cert.getSerialNumber().toString(), generatedSubjectData.getKeyPair().getPrivate(), keyService.getKeyPass(),
                    certificateChain);
            keyStoreWriter.saveKeyStore(keyService.getKeyStorePath(certificateDto.getType()), keyService.getKeyStorePass().toCharArray());
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
        } catch (ParseException e) {
            e.printStackTrace();
        }

        return null;
    }

    @Override
    public com.example.PKI.model.Certificate saveCertificateDB(CertificateDto subjectDto, Integer subjectId) {
        CertificateType cerType;
        if (subjectDto.getType().equals("ROOT")) {
            cerType = CertificateType.ROOT;
        } else if (subjectDto.getType().equals("INTERMEDIATE")) {
            cerType = CertificateType.INTERMEDIATE;
        } else {
            cerType = CertificateType.CLIENT;
        }

        User subject = userRepository.findById(subjectId).get();

        if (subjectDto.getStartDate().compareTo(subjectDto.getEndDate()) < 0) {
            com.example.PKI.model.Certificate certificate = new com.example.PKI.model.Certificate();
            certificate.setSerialNumber(serialNumber.toString());
            certificate.setType(cerType);
            certificate.setValid(true);
            certificate.setSubjectCommonName(subject.getCommonName() + " " + subject.getOrganization());
            certificate.setValidFrom(subjectDto.getStartDate());
            certificate.setValidTo(subjectDto.getEndDate());
            return certificateRepository.save(certificate);
        }
        return null;
    }

    @Override
    public void createCertificate(CertificateDto certificateDto, Subject generatedSubjectData) throws Exception {
        //User user = userRepository.findById(certificateDto.getSubjectId()).get();
        X509Certificate certificate = generateCertificate(certificateDto, generatedSubjectData);
    }

    @Override
    public X509Certificate getCertificateByAlias(String alias, KeyStore keystore) throws KeyStoreException {
        X509Certificate certificate = (X509Certificate) keystore.getCertificate(alias);
        return certificate;
    }

    @Override
    public Subject generateSubjectData(Integer subjectId) {

        User subject = userRepository.findById(subjectId).get();

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
