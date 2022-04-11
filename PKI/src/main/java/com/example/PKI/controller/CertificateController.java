package com.example.PKI.controller;

import com.example.PKI.dto.CertificateDto;
import com.example.PKI.dto.DownloadCertificateDto;
import com.example.PKI.dto.IssuerDto;
import com.example.PKI.model.Certificate;
import com.example.PKI.model.Subject;
import com.example.PKI.model.User;
import com.example.PKI.repository.CertificateRepository;
import com.example.PKI.repository.UserRepository;
import com.example.PKI.service.Base64Encoder;
import com.example.PKI.service.KeyService;
import com.example.PKI.service.cert.CertificateService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.io.IOException;
import java.security.KeyStoreException;
import java.security.NoSuchAlgorithmException;
import java.security.NoSuchProviderException;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;
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
    private CertificateRepository certificateRepository;

    @PostMapping("/api/certificate/generate")
    public ResponseEntity<String> generateCertificate(@RequestBody CertificateDto certificateDto) throws Exception {
        Subject generatedSubjectData = certificateService.generateSubjectData(certificateDto.getSubjectId());
        com.example.PKI.model.Certificate certificate = certificateService.createCertificate(certificateDto, generatedSubjectData);
        //Certificate certificate = certificateService.saveCertificateDB(certificateDto, certificateDto.getSubjectId());
        if (certificate != null) {
            return new ResponseEntity<String>("Success!", HttpStatus.OK);
        } else {
            return new ResponseEntity<String>("Error!", HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @PostMapping("/api/certificate/revoke")
    public ResponseEntity<?> revokeCertificate(@RequestBody String serialNumber) throws Exception {
        certificateService.revokeCertificate(serialNumber);
        return new ResponseEntity<>("All good bbyyyyyyyyyyyyyyy", HttpStatus.OK);
    }

    @GetMapping("/api/certificate/")
    public ResponseEntity<ArrayList<com.example.PKI.model.Certificate>> getAll() throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException {
        return new ResponseEntity<ArrayList<Certificate>>(certificateService.getAllCertificates(), HttpStatus.OK);
    }

    @PostMapping("/api/certificate/downloadCertificate")
    public ResponseEntity<?> downloadCertificate(@RequestBody DownloadCertificateDto dto) throws Exception {
        base64Encoder.downloadCertificate(dto);
        return new ResponseEntity<>(HttpStatus.OK);
    }

    @GetMapping("/api/certificate/getCAsForSigning")
    public ResponseEntity<?> getCAsForSigning(@RequestParam("startDate") String startDate, @RequestParam("endDate") String endDate) throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException {
        return new ResponseEntity<ArrayList<IssuerDto>>(certificateService.getAllValidSignersForDateRange(startDate, endDate), HttpStatus.OK);
    }
}
