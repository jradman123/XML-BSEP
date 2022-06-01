package com.example.AgentApp.service;

import com.example.AgentApp.model.*;
import org.springframework.stereotype.*;

import java.util.*;

@Service
public interface JobOfferService {
    List<JobOffer> getAllOffersForCompany(Long companyId);

    List<JobOffer> getAllJobOffers();
}
