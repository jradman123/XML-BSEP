package com.example.PKI.controller;

import com.example.PKI.dto.CertificateDto;
import com.example.PKI.dto.DownloadCertificateDto;
import com.example.PKI.model.*;
import com.example.PKI.dto.IssuerDto;
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
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.*;

import java.io.IOException;
import java.security.*;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;
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
    @Autowired
    private CertificateRepository certificateRepository;

    @PreAuthorize("hasAuthority('ADMIN')")
    @PostMapping("/api/certificate/generate")
    public ResponseEntity<String> generateCertificate(@RequestBody CertificateDto certificateDto) throws Exception {
        Subject generatedSubjectData = certificateService.generateSubjectData(certificateDto.getSubjectId());
        com.example.PKI.model.Certificate certificate = certificateService.createCertificate(certificateDto, generatedSubjectData);
        if (certificate != null) {
            return new ResponseEntity<String>("Success!", HttpStatus.OK);
        } else {
            return new ResponseEntity<String>("Error!", HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @PreAuthorize("hasAuthority('ADMIN') || hasAuthority('USER_ROOT') || hasAuthority('USER_INTERMEDIATE')")
    @PostMapping("/api/certificate/revoke")
    public ResponseEntity<ArrayList<com.example.PKI.model.Certificate>> revokeCertificate(@RequestBody String serialNumber) throws Exception {
        certificateService.revokeCertificate(serialNumber);
        return new ResponseEntity<>(certificateService.getAllCertificates(), HttpStatus.OK);
    }

    @PreAuthorize("hasAuthority('ADMIN')")
    @GetMapping("/api/certificate/")
    public ResponseEntity<ArrayList<com.example.PKI.model.Certificate>> getAll() throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException {
        return new ResponseEntity<ArrayList<Certificate>>(certificateService.getAllCertificates(), HttpStatus.OK);
    }

    @PreAuthorize("hasAuthority('ADMIN') || hasAuthority('USER_ROOT') || hasAuthority('USER_INTERMEDIATE') || hasAuthority('USER_END_ENTITY')")
    @PostMapping("/api/certificate/downloadCertificate")
    public ResponseEntity<?> downloadCertificate(@RequestBody String serialNumber) throws Exception {
        base64Encoder.downloadCertificate(serialNumber);
        return new ResponseEntity<>(HttpStatus.OK);
    }
    @PreAuthorize("hasAuthority('ADMIN') || hasAuthority('USER_ROOT') || hasAuthority('USER_INTERMEDIATE') || hasAuthority('USER_END_ENTITY')")
    @GetMapping("/api/certificate/getAllUsersCertificates/{email}")
    public ResponseEntity<ArrayList<com.example.PKI.model.Certificate>> getAllUsersCertificates(@PathVariable String email) throws Exception {
        return new ResponseEntity<ArrayList<com.example.PKI.model.Certificate>>(certificateService.getAllUsersCertificates(email), HttpStatus.OK);
    }

    @PreAuthorize("hasAuthority('ADMIN') || hasAuthority('USER_ROOT') || hasAuthority('USER_INTERMEDIATE') || hasAuthority('USER_END_ENTITY')")
    @GetMapping("/api/certificate/chains/{email}")
    public ArrayList[] getAllCertificateChains(@PathVariable String email) throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException {
        ArrayList[] certificate = certificateService.getAllCertificateChains(email);

        return certificate;
    }

    @PreAuthorize("hasAuthority('ADMIN') || hasAuthority('USER_ROOT') || hasAuthority('USER_INTERMEDIATE')")
    @GetMapping("/api/certificate/getCAsForSigningClientsCertificatesInDateRange")
    public ResponseEntity<?> getCAsForSigningInDateRange(@RequestParam("email") String email, @RequestParam("startDate") String startDate, @RequestParam("endDate") String endDate) throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException {
        return new ResponseEntity<ArrayList<IssuerDto>>(certificateService.getAllValidSignersForDateRangeByUser(email, startDate, endDate), HttpStatus.OK);
    }

    @PreAuthorize("hasAuthority('USER_ROOT') || hasAuthority('USER_INTERMEDIATE')")
    @PostMapping("/api/certificate/generateByClient")
    public ResponseEntity<String> generateCertificateByClient(@RequestBody CertificateDto certificateDto) throws Exception {
        Subject generatedSubjectData = certificateService.generateSubjectData(certificateDto.getSubjectId());
        com.example.PKI.model.Certificate certificate = certificateService.generateCertificateByUser(certificateDto, generatedSubjectData);
        //Certificate certificate = certificateService.saveCertificateDB(certificateDto, certificateDto.getSubjectId());
        if (certificate != null) {
            return new ResponseEntity<String>("Success!", HttpStatus.OK);
        } else {
            return new ResponseEntity<String>("Error!", HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @PreAuthorize("hasAuthority('ADMIN') || hasAuthority('USER_ROOT') || hasAuthority('USER_INTERMEDIATE')")
    @GetMapping("/api/certificate/getCAsForSigning")
    public ResponseEntity<?> getCAsForSigning(@RequestParam("startDate") String startDate, @RequestParam("endDate") String endDate) throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException {
        return new ResponseEntity<ArrayList<IssuerDto>>(certificateService.getAllValidSignersForDateRange(startDate, endDate), HttpStatus.OK);
    }

    @PreAuthorize("hasAuthority('ADMIN') || hasAuthority('USER_ROOT') || hasAuthority('USER_INTERMEDIATE')  || hasAuthority('USER_END_ENTITY')")
    @PostMapping("/api/certificate/isCertificateValid")
    public ResponseEntity<?> isCertificateValid(@RequestBody String serialNumber) throws CertificateException, KeyStoreException, IOException, NoSuchAlgorithmException, NoSuchProviderException, SignatureException, InvalidKeyException {
        return new ResponseEntity<Boolean>(certificateService.isCertificateValid(certificateService.getKeyStoreByAlias(serialNumber), serialNumber), HttpStatus.OK);
    }

}
