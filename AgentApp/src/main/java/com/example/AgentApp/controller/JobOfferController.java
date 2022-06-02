package com.example.AgentApp.controller;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.mapper.*;
import com.example.AgentApp.model.*;
import com.example.AgentApp.service.*;
import org.springframework.beans.factory.annotation.*;
import org.springframework.http.*;
import org.springframework.web.bind.annotation.*;

import java.util.*;

@CrossOrigin(origins = "http://localhost:4200")
@RequestMapping("/offer")
@RestController
public class JobOfferController {

    @Autowired
    private JobOfferService jobOfferService;
    @Autowired
    private JobOfferMapper jobOfferMapper;

    //NE RADI
    //mzd svi sta znam
    @GetMapping("all/{companyId}")
    public ResponseEntity<?> allOffersForCompany(@PathVariable Long companyId){
        Set<JobOffer> offers = jobOfferService.getAllOffersForCompany(companyId);
        if (offers != null){
            return new ResponseEntity<List<JobOfferResponseDto>>(jobOfferMapper.mapToDtos(offers), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to get all job offers for company!", HttpStatus.CONFLICT);
    }

    //SVI MZD
    @GetMapping("all")
    public ResponseEntity<?> allJobOffers(){
        List<JobOffer> offers = jobOfferService.getAllJobOffers();
        if (offers != null){
            return new ResponseEntity<List<JobOfferWithCompanyResponseDto>>(jobOfferMapper.mapToDtosWithCompany(offers), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to get all job offers!", HttpStatus.CONFLICT);
    }
}
