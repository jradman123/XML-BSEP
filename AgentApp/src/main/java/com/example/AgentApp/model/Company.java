package com.example.AgentApp.model;

import com.example.AgentApp.enums.CompanyStatus;
import lombok.Data;
import lombok.ToString;

import javax.persistence.*;
import java.util.List;

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

}
