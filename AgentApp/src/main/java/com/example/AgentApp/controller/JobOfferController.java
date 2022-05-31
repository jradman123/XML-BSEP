package com.example.AgentApp.controller;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.mapper.*;
import com.example.AgentApp.model.*;
import com.example.AgentApp.service.*;
import org.springframework.beans.factory.annotation.*;
import org.springframework.http.*;
import org.springframework.web.bind.annotation.*;

import java.util.*;

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
        List<JobOffer> offers = jobOfferService.getAllOfersForCompany(companyId);
        if (offers != null){
            return new ResponseEntity<List<JobOfferResponseDto>>(jobOfferMapper.mapToDtos(offers), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to add job offer to company!", HttpStatus.CONFLICT);
    }
}
