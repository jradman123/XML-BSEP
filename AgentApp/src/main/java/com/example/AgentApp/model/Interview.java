package com.example.AgentApp.model;

import com.example.AgentApp.enums.InterviewDifficulty;
import lombok.Data;

import javax.persistence.*;


@Entity
@Data
public class Interview {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @ManyToOne
    @JoinColumn(name = "user_id")
    private User user;

    @ManyToOne
    @JoinColumn(name = "company_id")
    private Company company;

    @Column(nullable = false)
    private String comment;

    @Column(nullable = false)
    private int rating;

    @Column(nullable = false)
    private InterviewDifficulty difficulty;

}
