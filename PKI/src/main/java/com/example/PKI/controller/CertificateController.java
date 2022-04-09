package com.example.PKI.controller;

import com.example.PKI.dto.SubjectDto;
import com.example.PKI.dto.SubjectWithPkDto;
import com.example.PKI.model.Certificate;
import com.example.PKI.service.KeyService;
import com.example.PKI.service.cert.CertificateService;
import org.bouncycastle.operator.OperatorCreationException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import java.io.IOException;
import java.security.*;
import java.security.cert.CertificateException;

@CrossOrigin(origins = "http://localhost:4200")
@RestController
public class CertificateController {

    @Autowired
    private KeyService keyService;
    @Autowired
    private CertificateService certificateService;

    @PostMapping("/api/certificate/generate")
    public ResponseEntity<String> generateCertificate(@RequestBody SubjectDto subjectDto) throws CertificateException, NoSuchAlgorithmException, KeyStoreException, IOException, OperatorCreationException, NoSuchProviderException, InvalidAlgorithmParameterException, UnrecoverableKeyException, SignatureException, InvalidKeyException {
        SubjectWithPkDto subjectWithPk = certificateService.generateSubjectData(subjectDto);
        certificateService.createCertificate(subjectWithPk, subjectDto.getIssuerSerialNumber());
        Certificate certificate = certificateService.saveCertificateDb(subjectWithPk);

        if (certificate != null) {
            return new ResponseEntity<String>("Success!", HttpStatus.OK);

        } else {
            return new ResponseEntity<String>("Error!", HttpStatus.INTERNAL_SERVER_ERROR);

        }
    }

}
