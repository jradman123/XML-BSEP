package com.example.AgentApp.model;

import lombok.Data;

import javax.persistence.*;
import java.time.LocalDate;
import java.util.Set;

@Entity
@Data
public class JobOffer {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    @ManyToOne
    @JoinColumn(name = "company_id")
    private Company company;
    private String position;
    @Column(columnDefinition = "TEXT")
    private String jobDescription;
    @ElementCollection(targetClass = String.class, fetch = FetchType.EAGER)
    private Set<String> requirements;
    private LocalDate dateCreated;
    private LocalDate dueDate;

}
