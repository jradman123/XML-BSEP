package com.example.AgentApp.model;

import com.example.AgentApp.enums.CompanyStatus;
import lombok.Data;
import lombok.ToString;

import javax.persistence.*;
import java.util.*;

@Entity
@Data
public class Company {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @ManyToOne
    @JoinColumn(name = "owner_id")
    private User owner;

    @Column(nullable = false)
    @OneToMany(mappedBy = "company", fetch = FetchType.LAZY, cascade = CascadeType.ALL)
    @ToString.Exclude
    private List<JobOffer> jobOffers;

    @Column(nullable = false)
    private CompanyStatus companyStatus;

    @Column(nullable = false)
    private String contactInfo;

    @Column(nullable = false)
    private String companyPolicy;

    public Company() {
    }

    public Company(User owner, CompanyStatus companyStatus, String contactInfo, String companyPolicy) {
        this.owner = owner;
        this.jobOffers = new ArrayList<JobOffer>();
        this.companyStatus = companyStatus;
        this.contactInfo = contactInfo;
        this.companyPolicy = companyPolicy;
    }

    public Company(User owner, List<JobOffer> jobOffers, CompanyStatus companyStatus, String contactInfo, String companyPolicy) {
        this.owner = owner;
        this.jobOffers = jobOffers;
        this.companyStatus = companyStatus;
        this.contactInfo = contactInfo;
        this.companyPolicy = companyPolicy;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public User getOwner() {
        return owner;
    }

    public void setOwner(User owner) {
        this.owner = owner;
    }

    public List<JobOffer> getJobOffers() {
        return jobOffers;
    }

    public void setJobOffers(List<JobOffer> jobOffers) {
        this.jobOffers = jobOffers;
    }

    public CompanyStatus getCompanyStatus() {
        return companyStatus;
    }

    public void setCompanyStatus(CompanyStatus companyStatus) {
        this.companyStatus = companyStatus;
    }

    public String getContactInfo() {
        return contactInfo;
    }

    public void setContactInfo(String contactInfo) {
        this.contactInfo = contactInfo;
    }

    public String getCompanyPolicy() {
        return companyPolicy;
    }

    public void setCompanyPolicy(String companyPolicy) {
        this.companyPolicy = companyPolicy;
    }
    
    
}
