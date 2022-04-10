package com.example.PKI.controller;

import com.example.PKI.dto.*;
import com.example.PKI.model.*;
import com.example.PKI.model.Certificate;
import com.example.PKI.repository.UserRepository;
import com.example.PKI.service.*;
import com.example.PKI.service.cert.*;
import com.example.PKI.util.keyStoreUtils.KeyStoreReader;
import org.bouncycastle.operator.*;
import org.springframework.beans.factory.annotation.*;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.io.*;
import java.security.*;
import java.security.cert.*;
import java.util.ArrayList;

@CrossOrigin(origins = "http://localhost:4200")
@RestController
public class CertificateController {

    @Autowired
    private KeyService keyService;
    @Autowired
    private CertificateService certificateService;
    @Autowired
    private UserRepository userRepository;
    @Autowired
    private Base64Encoder base64Encoder;
    @Autowired
    private KeyStoreReader keyStoreReader;


    @PostMapping("/api/certificate/generate")
    public ResponseEntity<String> generateCertificate(@RequestBody CertificateDto certificateDto) throws Exception {
        Subject generatedSubjectData = certificateService.generateSubjectData(certificateDto.getSubjectId());
        certificateService.createCertificate(certificateDto, generatedSubjectData);
        Certificate certificate = certificateService.saveCertificateDB(certificateDto, certificateDto.getSubjectId());
        if (certificate != null) {
            return new ResponseEntity<String>("Success!", HttpStatus.OK);
        } else {
            return new ResponseEntity<String>("Error!", HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @PostMapping("/api/certificate/revoke")
    public ResponseEntity<ArrayList<com.example.PKI.model.Certificate>> revokeCertificate(@RequestBody String serialNumber) throws Exception {
        certificateService.revokeCertificate(serialNumber);
        return new ResponseEntity<>(certificateService.getAllCertificates(), HttpStatus.OK);
    }

    @GetMapping("/api/certificate/")
    public ResponseEntity<ArrayList<com.example.PKI.model.Certificate>> getAll() throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException {
        return new ResponseEntity<ArrayList<Certificate>>(certificateService.getAllCertificates(), HttpStatus.OK);
    }

    @PostMapping("/api/certificate/downloadCertificate")
    public ResponseEntity<?> downloadCertificate(@RequestBody String serialNumber) throws Exception {
        base64Encoder.downloadCertificate(serialNumber);
        return new ResponseEntity<>(HttpStatus.OK);
    }

    @GetMapping("/api/certificate/getAllUsersCertificates/{email}")
    public ResponseEntity<ArrayList<com.example.PKI.model.Certificate>> getAllUsersCertificates(@PathVariable String email) throws Exception {
        return new ResponseEntity<ArrayList<com.example.PKI.model.Certificate>>(certificateService.getAllUsersCertificates(email),HttpStatus.OK);
    }


}
