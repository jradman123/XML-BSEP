package com.example.PKI.controller;

import com.example.PKI.dto.*;
import com.example.PKI.model.*;
import com.example.PKI.model.Certificate;
import com.example.PKI.service.*;
import com.example.PKI.service.cert.*;
import org.bouncycastle.asn1.x500.*;
import org.bouncycastle.operator.*;
import org.springframework.beans.factory.annotation.*;
import org.springframework.web.bind.annotation.*;

import java.io.*;
import java.security.*;
import java.security.cert.*;

@CrossOrigin(origins = "http://localhost:4200")
@RestController
public class CertificateController {

    @Autowired
    private KeyService keyService;
    @Autowired
    private CertificateService certificateService;


    @PostMapping("/api/certificate/generate")
    public void generateCertificate(@RequestBody SubjectDto subjectDto) throws CertificateException, NoSuchAlgorithmException, KeyStoreException, IOException, OperatorCreationException, NoSuchProviderException, InvalidAlgorithmParameterException, UnrecoverableKeyException {



      //  IssuerDto issuerDto = new IssuerDto();
      //  subjectDto.setX500Name(certificateService.getX500NameSubject(subjectDto));
        SubjectWithPKDto subjectWithPK = certificateService.generateSubjectData(subjectDto);

        //KeyPair keyPair = keyService.generateKeyPair();
        //subjectDto.setPublicKey(keyPair.getPublic());
        //subjectDto.setPrivateKey(keyPair.getPrivate());

        //zasto da mi vraca sertifikate?
        //  keystoreService.getCertificates(keyService.getKeyStorePass());
        certificateService.createCertificate(subjectWithPK, subjectDto.getType(),subjectDto.getIssuerSerialNumber());

        //sacuva u bazi podataka sertifikat zajedno sa njegovim tipom
        //  subjectDto.setAlias(keyService.getSerialNumber().toString());
        //  subjectDto.setSerialNumber(subjectDto.getAlias());
        //  CGservice.saveCertificateDB(subjectDto);
        Certificate newCertificate = new Certificate();
        certificateService.saveCertificateDB(subjectWithPK);
    }
}
