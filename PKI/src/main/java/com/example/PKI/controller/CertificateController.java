package com.example.PKI.controller;

import com.example.PKI.dto.CertificateDto;
import com.example.PKI.dto.DownloadCertificateDto;
import com.example.PKI.model.*;
import com.example.PKI.model.Certificate;
import com.example.PKI.model.Subject;
import com.example.PKI.model.User;
import com.example.PKI.repository.*;
import com.example.PKI.service.Base64Encoder;
import com.example.PKI.service.KeyService;
import com.example.PKI.service.cert.CertificateService;
import com.example.PKI.util.keyStoreUtils.KeyStoreReader;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import java.io.IOException;
import java.security.KeyStoreException;
import java.security.NoSuchAlgorithmException;
import java.security.NoSuchProviderException;
import java.security.cert.CertificateException;
import java.util.ArrayList;
import java.util.Collection;

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


    @Autowired
    private CertificateRepository certificateRepository;

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


    @GetMapping("/api/certificate/getCAsForSigningClientsCertificatesInDateRange")
    public ResponseEntity<?> getCAsForSigningInDateRange(@RequestParam("email") String email,@RequestParam("startDate") String startDate, @RequestParam("endDate") String endDate) throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException {
        ArrayList<Certificate> cs = certificateService.getAllValidSignersForUser(email,startDate,endDate);
        return new ResponseEntity<ArrayList<Certificate>>(cs, HttpStatus.OK);
    }

    @PostMapping("/api/certificate/generateByClient")
    public ResponseEntity<String> generateCertificateByClient(@RequestBody CertificateDto certificateDto) throws Exception {
        Subject generatedSubjectData = certificateService.generateSubjectData(certificateDto.getSubjectId());
        certificateService.generateCertificateByUser(certificateDto, generatedSubjectData);
        Certificate certificate = certificateService.saveCertificateDB(certificateDto, certificateDto.getSubjectId());
        if (certificate != null) {
            return new ResponseEntity<String>("Success!", HttpStatus.OK);
        } else {
            return new ResponseEntity<String>("Error!", HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }
    @GetMapping("/api/certificate/getCAsForSigning")
    public ResponseEntity<?> getCAsForSigning(@RequestParam("startDate") String startDate, @RequestParam("endDate") String endDate) throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException {
        return new ResponseEntity<ArrayList<User>>(certificateService.getAllValidSignersForDateRange(startDate, endDate), HttpStatus.OK);

    }
}
