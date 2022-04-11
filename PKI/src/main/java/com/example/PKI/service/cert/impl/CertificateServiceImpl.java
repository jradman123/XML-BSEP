package com.example.PKI.service.cert.impl;
import com.example.PKI.dto.CertificateDto;
import com.example.PKI.dto.IssuerDto;
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
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.math.BigInteger;
import java.security.*;
import java.security.cert.Certificate;
import java.security.cert.*;
import java.text.*;
import java.time.*;
import java.util.*;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Collections;
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
    public com.example.PKI.model.Certificate generateCertificate(CertificateDto certificateDto, Subject generatedSubjectData) throws Exception {
        try {

            // TODO: provjeri je l ima usera i issuera u bazii brrrr

            SimpleDateFormat sdf = new SimpleDateFormat("MM/dd/yyyy");
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

            if (certificateDto.getType().equalsIgnoreCase("ROOT")) {
                issuerAlias = serialNumber.toString();
                issuer = generatedSubjectData.getX500Name();
                issuerKeyStore = keyStoreService.getKeyStore(keyService.getKeyStorePath(CertificateType.ROOT.toString()),
                        keyService.getKeyStorePass());
                privateKey = generatedSubjectData.getKeyPair().getPrivate();
            } else {
                issuerAlias = certificateDto.getIssuerSerialNumber();

                issuerKeyStore = keyStoreReader.getKeyStore(keyService.getKeyStorePath(certificateRepository
                        .findTypeBySerialNumber(certificateDto.getIssuerSerialNumber()).toString()), keyService.getKeyStorePass());

                issuer = new JcaX509CertificateHolder((X509Certificate) issuerKeyStore.getCertificate(issuerAlias)).getSubject();
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

            return saveCertificateDB(certificateDto, certificateDto.getSubjectId());

            //return certConverter.getCertificate(certHolder);

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
        if (subjectDto.getType().equalsIgnoreCase("ROOT")) {
            cerType = CertificateType.ROOT;
        } else if (subjectDto.getType().equalsIgnoreCase("INTERMEDIATE")) {
            cerType = CertificateType.INTERMEDIATE;
        } else {
            cerType = CertificateType.CLIENT;
        }

        User subject = userRepository.findById(subjectId).get();

        if (subjectDto.getStartDate().compareTo(subjectDto.getEndDate()) < 0) {
            com.example.PKI.model.Certificate certificate = new com.example.PKI.model.Certificate();
            certificate.setSerialNumber(serialNumber.toString());
            certificate.setType(cerType);
            certificate.setIsRevoked(false);
            certificate.setSubjectEmail(subject.getEmail());
            certificate.setValidFrom(subjectDto.getStartDate());
            certificate.setValidTo(subjectDto.getEndDate());
            return certificateRepository.save(certificate);
        }
        return null;
    }

    @Override
    public com.example.PKI.model.Certificate createCertificate(CertificateDto certificateDto, Subject generatedSubjectData) throws Exception {
        //User user = userRepository.findById(certificateDto.getSubjectId()).get();
        com.example.PKI.model.Certificate certificate = generateCertificate(certificateDto, generatedSubjectData);
        return certificate;
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

    @Override
    public void revokeCertificate(String serialNumber) throws Exception {

        com.example.PKI.model.Certificate certificate = certificateRepository.findBySerialNumber(serialNumber);
        if (certificate != null)
            if (certificate.isRevoked()) throw new Exception("Certificate is already revoked!");

        revokeAllBelow(serialNumber);
    }

    private void revokeAllBelow(String serialNumber) throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException {
        KeyStore[] keyStores = getKeyStores();
        for (KeyStore ks : keyStores) revokeAllByKS(ks, serialNumber);
    }

    private KeyStore[] getKeyStores() throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException {
        KeyStore rootKS = keyStoreReader.getKeyStore("roots", "password");
        KeyStore intermediateKS = keyStoreReader.getKeyStore("intermediates", "password");
        KeyStore clientKS = keyStoreReader.getKeyStore("clients", "password");
        return new KeyStore[]{rootKS, intermediateKS, clientKS};
    }

    private void revokeAllByKS(KeyStore keyStore, String serialNumber) throws KeyStoreException {
        ArrayList<String> aliases = Collections.list(keyStore.aliases());
        for (String alias : aliases)
            checkCertificateChain(keyStore.getCertificateChain(alias), serialNumber);
    }

    private void checkCertificateChain(Certificate[] certificates, String serialNumber) {
        int idx = -1;
        for (int i = 0; i < certificates.length; i++) {
            if (((X509Certificate) certificates[i]).getSerialNumber().toString().equals(serialNumber)) {
                idx = i;
            }
        }

        if (idx != -1) {
            for (int i = 0; i <= idx; i++) {
                X509Certificate certificate = (X509Certificate) certificates[i];
                com.example.PKI.model.Certificate dbCertificate = certificateRepository.findBySerialNumber(certificate.getSerialNumber().toString());
                dbCertificate.setIsRevoked(true);
                certificateRepository.save(dbCertificate);
            }
        }
    }

    @Override
    public ArrayList<com.example.PKI.model.Certificate> getAllCertificates() throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException {
        ArrayList<com.example.PKI.model.Certificate> certificates = new ArrayList<com.example.PKI.model.Certificate>();
        for ( com.example.PKI.model.Certificate c : certificateRepository.findAll()) {
            if (isCertificateValid(getKeyStoreByAlias(c.getSerialNumber()), c.getSerialNumber()))
                certificates.add(c);
        }
        return certificates;
    }

    @Override
    public ArrayList<com.example.PKI.model.Certificate> getAllUsersCertificates(String email) {
        return (ArrayList<com.example.PKI.model.Certificate>) certificateRepository.findAllBySubjectEmail(email);
    }

    @Override
    public ArrayList<IssuerDto> getAllValidSignersForUser(String email,String startDate,String endDate) throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException {
        ArrayList<IssuerDto> users = new ArrayList<IssuerDto>();
        for( com.example.PKI.model.Certificate c : certificateRepository.findAllSignersCertByUser(email,startDate,endDate)){
            if (isCertificateValid(getKeyStoreByAlias(c.getSerialNumber()), c.getSerialNumber())){
                users.add(new IssuerDto(userRepository.findByEmail(email), c.getSerialNumber()));
            }
        }
        return users;
    }

    @Override
    public boolean isCertificateValid(KeyStore keyStore, String alias) throws KeyStoreException, CertificateException, NoSuchAlgorithmException, NoSuchProviderException {
        Certificate[] certificates = keyStore.getCertificateChain(alias);

        for (int i = certificates.length - 1; i >= 0; i--) {
            String currentAlias = ((X509Certificate) certificates[i]).getSerialNumber().toString();
            //KeyStore currentKS = getKeyStoreByAlias(currentAlias);

            // is it revoked
            if (certificateRepository.findBySerialNumber(currentAlias).isRevoked()) return false;

            // did it expire
            X509Certificate certificate = (X509Certificate) certificates[i];
            try {
                certificate.checkValidity();
            }
            catch (CertificateExpiredException e) {
                return false;
            }
            catch (CertificateNotYetValidException e) {
                return false;
            }

            X509Certificate parent = (X509Certificate) certificates[i];

            if( i > 0) {
                X509Certificate child = (X509Certificate) certificates[i - 1];

                try {
                    child.verify(parent.getPublicKey()); // da li je potpis validan
                }
                catch (SignatureException e) {
                    return false;
                }
                catch (InvalidKeyException e) {
                    return false;
                }
            }
        }
        return true;
    }

    @Override
    public KeyStore getKeyStoreByAlias(String alias) throws KeyStoreException, CertificateException, IOException, NoSuchAlgorithmException, NoSuchProviderException {
        KeyStore[] keyStores = getKeyStores();
        for (KeyStore ks : keyStores) {
            if (ks.containsAlias(alias)) return ks;
        }
        return null;
    }

    @Override
    public ArrayList<IssuerDto> getAllValidSignersForDateRange(String startDate, String endDate) throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException {
        ArrayList<IssuerDto> users = new ArrayList<IssuerDto>();
        for( com.example.PKI.model.Certificate c : certificateRepository.findCertificatesValidForDateRange(startDate, endDate)){
            String mejlic = c.getSubjectEmail();
            if (isCertificateValid(getKeyStoreByAlias(c.getSerialNumber()), c.getSerialNumber())){
                users.add(new IssuerDto(userRepository.findByEmail(c.getSubjectEmail()), c.getSerialNumber()));
            }
        }
        return users;
    }
   /* private boolean validate(String alias) throws CertificateException, NoSuchAlgorithmException, KeyStoreException, IOException {
        KeyStore keyStore=keyStoreService.getKeyStore(keyService.getKeyStorePath(),keyService.getKeyStorePass());
        X509Certificate certificate= (X509Certificate) keyStore.getCertificate(alias);
        Certificate[] chain=keyStore.getCertificateChain(alias);
        return false;
    }*/

    @Override
    public com.example.PKI.model.Certificate generateCertificateByUser(CertificateDto certificateDto, Subject generatedSubjectData) {
        try {

//            if (LocalDate.parse(certificateDto.getEndDate()).compareTo(LocalDate.parse(certificateDto.getStartDate())) < 0) {
//                throw new Exception("END_DATE_BEFORE_START");
//            }
//            if (certificateDto.getType() == "ROOT") {
//                throw new Exception("CAN_NOT_GENERATE_ROOT");
//            }
//            // TODO: provjeri je l ima usera i issuera u bazii brrrr
//
//            SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd");
//            Date startDate = sdf.parse(certificateDto.getStartDate());
//            Date endDate = sdf.parse(certificateDto.getEndDate());
//
//            //ovo istraziti : on koristi SHA256withECDSA
//            JcaContentSignerBuilder builder = new JcaContentSignerBuilder("SHA256WithRSAEncryption");
//            Security.addProvider(new BouncyCastleProvider());
//            builder.setProvider("BC");
            SimpleDateFormat sdf = new SimpleDateFormat("MM/dd/yyyy");
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
            issuerAlias = certificateDto.getIssuerSerialNumber();
            issuerKeyStore = keyStoreReader.getKeyStore(keyService.getKeyStorePath("INTERMEDIATE"), keyService.getKeyStorePass());
            issuer = new JcaX509CertificateHolder((X509Certificate) issuerKeyStore.getCertificate(issuerAlias)).getSubject();
            privateKey = (PrivateKey) issuerKeyStore.getKey(issuerAlias, keyService.getKeyPass().toCharArray());

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

            return saveCertificateDB(certificateDto, certificateDto.getSubjectId());


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
        } catch (NoSuchProviderException e) {
            e.printStackTrace();
        } catch (Exception e) {
            e.printStackTrace();
        }
        return null;
    }

    @Override
    public ArrayList<IssuerDto> getAllValidSignersForDateRangeByUser(String email,String startDate, String endDate) throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException {
        ArrayList<IssuerDto> users = new ArrayList<IssuerDto>();

        ArrayList<com.example.PKI.model.Certificate> cers = new ArrayList<>();
        for (com.example.PKI.model.Certificate c : certificateRepository.findAllBySubjectEmail(email)) {
            if(c.getType() == CertificateType.INTERMEDIATE ){
                try {
                    SimpleDateFormat sdf = new SimpleDateFormat("MM/dd/yyyy", Locale.ENGLISH);
                    Date start = sdf.parse(startDate);
                    Date ende = sdf.parse(endDate);

                    if( sdf.parse(c.getValidFrom()).compareTo(start) < 0 && sdf.parse(c.getValidTo()).compareTo(ende) > 0)
                        cers.add(c);

                } catch (ParseException e) {
                    e.printStackTrace();
                }
            }
        }

        //for( com.example.PKI.model.Certificate c : certificateRepository.findCertificatesValidForDateRange(startDate, endDate)){
        for( com.example.PKI.model.Certificate c : cers){
            if (isCertificateValid(getKeyStoreByAlias(c.getSerialNumber()), c.getSerialNumber())){
                users.add(new IssuerDto(userRepository.findByEmail(c.getSubjectEmail()), c.getSerialNumber()));
            }
        }
        return users;
    }

}
