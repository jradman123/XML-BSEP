package com.example.AgentApp.model;

import com.example.AgentApp.enums.Gender;
import com.example.AgentApp.enums.UserRole;
import lombok.Data;
import lombok.ToString;
import org.apache.commons.codec.binary.Base32;

import javax.persistence.*;
import java.security.SecureRandom;
import java.util.Date;
import java.util.List;

@Entity
@Data
@Table(name = "agent_user")
public class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(unique = true, nullable = false)
    private String username;

    @Column(unique = true, nullable = false)
    private String password;

    @Column(unique = true, nullable = false)
    private String email;

    @Column(nullable = false)
    private String recoveryEmail;

    @Column(nullable = false)
    private String phoneNumber;

    @Column(nullable = false)
    private String firstName;

    @Column(nullable = false)
    private String lastName;

    @Column(nullable = false)
    private Date dateOfBirth;

    @Column(nullable = false)
    private Gender gender;

    @Column(nullable = false)
    private UserRole role;

    @Column(nullable = false)
    private boolean isConfirmed;
    private boolean isUsing2FA;
    private String secret;
//
//    @Column(nullable = false)
//    @OneToMany(mappedBy = "owner", fetch = FetchType.LAZY)
//    @ToString.Exclude
//    private List<Company> companies;

//    @Column(nullable = false)
//    @OneToMany(mappedBy = "user", fetch = FetchType.LAZY)
//    @ToString.Exclude
//    private List<SalaryComment> salaryComments;
//
//    @Column(nullable = false)
//    @OneToMany(mappedBy = "user", fetch = FetchType.LAZY)
//    @ToString.Exclude
//    private List<Interview> interviews;
//
//    @Column(nullable = false)
//    @OneToMany(mappedBy = "user", fetch = FetchType.LAZY)
//    @ToString.Exclude
//    private List<Comment> comments;

    public User(String username, String password, String email, String recoveryEmail, String phoneNumber, String firstName, String lastName, Date dateOfBirth, Gender gender, UserRole role, boolean isConfirmed) {
        this.username = username;
        this.password = password;
        this.email = email;
        this.recoveryEmail = recoveryEmail;
        this.phoneNumber = phoneNumber;
        this.firstName = firstName;
        this.lastName = lastName;
        this.dateOfBirth = dateOfBirth;
        this.gender = gender;
        this.role = role;
        this.isConfirmed = isConfirmed;
        this.isUsing2FA = false;
        this.secret = generateSecretKey();
    }

    private static String generateSecretKey() {
        SecureRandom random = new SecureRandom();
        byte[] bytes = new byte[20];
        random.nextBytes(bytes);
        Base32 base32 = new Base32();
        return base32.encodeToString(bytes);
    }

    public void setSecret() {
        this.secret = generateSecretKey();
    }

    public User() {
    }
}
