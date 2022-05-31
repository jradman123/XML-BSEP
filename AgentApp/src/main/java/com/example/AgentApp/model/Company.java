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

    @OneToOne(cascade = CascadeType.ALL)
    @JoinColumn(name = "company_info_id", referencedColumnName = "id")
    private CompanyInfo companyInfo;

    @Column(nullable = false)
    private String companyPolicy;

    @Column(nullable = false)
    @OneToMany(mappedBy = "company", fetch = FetchType.LAZY)
    @ToString.Exclude
    private List<SalaryComment> salaryComments;

    @Column(nullable = false)
    @OneToMany(mappedBy = "company", fetch = FetchType.LAZY)
    @ToString.Exclude
    private List<Interview> interviews;

    @Column(nullable = false)
    @OneToMany(mappedBy = "company", fetch = FetchType.LAZY)
    @ToString.Exclude
    private List<Comment> comments;

    
}
