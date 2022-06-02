package com.example.AgentApp.service.impl;

import com.example.AgentApp.model.*;
import com.example.AgentApp.repository.*;
import com.example.AgentApp.service.*;
import org.springframework.beans.factory.annotation.*;
import org.springframework.stereotype.*;

import java.util.*;

@Service
public class JobOfferServiceImpl implements JobOfferService {

    @Autowired
    private JobOfferRepository jobOfferRepository;

    @Override
    public Set<JobOffer> getAllOffersForCompany(Long companyId) {
        Set<JobOffer> offers = jobOfferRepository.findAllByCompanyId(companyId);
        return offers;
    }

    @Override
    public List<JobOffer> getAllJobOffers() {
        List<JobOffer> offers = jobOfferRepository.findAll();
        return offers;
    }
}
