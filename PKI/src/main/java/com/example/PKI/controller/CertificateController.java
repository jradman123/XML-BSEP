package com.example.PKI.controller;

import com.example.PKI.dto.*;
import com.example.PKI.model.*;
import com.example.PKI.model.Certificate;
import com.example.PKI.service.*;
import com.example.PKI.service.cert.*;
import org.bouncycastle.asn1.x500.*;
import org.bouncycastle.operator.*;
import org.springframework.beans.factory.annotation.*;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
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
    public ResponseEntity<String> generateCertificate(@RequestBody SubjectDto subjectDto) throws CertificateException, NoSuchAlgorithmException, KeyStoreException, IOException, OperatorCreationException, NoSuchProviderException, InvalidAlgorithmParameterException, UnrecoverableKeyException {
        SubjectWithPKDto subjectWithPK = certificateService.generateSubjectData(subjectDto);
        certificateService.createCertificate(subjectWithPK, subjectDto.getIssuerSerialNumber());
        Certificate certificate= certificateService.saveCertificateDB(subjectWithPK);
        if(certificate != null){
            return new ResponseEntity<String>("Success!", HttpStatus.OK);
        }else{
            return new ResponseEntity<String>("Error!",HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }
}
