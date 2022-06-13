package com.example.AgentApp.config;

import com.example.AgentApp.security.RestAuthenticationEntryPoint;
import com.example.AgentApp.security.TokenAuthenticationFilter;
import com.example.AgentApp.security.TokenUtils;
import com.example.AgentApp.service.impl.UserDetailsService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.HttpMethod;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder;
import org.springframework.security.config.annotation.method.configuration.EnableGlobalMethodSecurity;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.builders.WebSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;
import org.springframework.security.config.http.SessionCreationPolicy;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.security.web.authentication.www.BasicAuthenticationFilter;

@Configuration
@EnableGlobalMethodSecurity(prePostEnabled = true)
@EnableWebSecurity
public class WebSecurityConfig extends WebSecurityConfigurerAdapter {

    private final UserDetailsService customUserDetailsService;

    private final RestAuthenticationEntryPoint restAuthenticationEntryPoint;

    private final PasswordEncoder passwordEncoder;

    private final TokenUtils tokenUtils;


    public WebSecurityConfig(UserDetailsService customUserDetailsService, RestAuthenticationEntryPoint restAuthenticationEntryPoint, PasswordEncoder passwordEncoder, TokenUtils tokenUtils) {
        this.customUserDetailsService = customUserDetailsService;
        this.restAuthenticationEntryPoint = restAuthenticationEntryPoint;
        this.passwordEncoder = passwordEncoder;
        this.tokenUtils = tokenUtils;
    }


    @Bean
    @Override
    public AuthenticationManager authenticationManagerBean() throws Exception {
        return super.authenticationManagerBean();
    }

    @Autowired
    public void configureGlobal(AuthenticationManagerBuilder auth) throws Exception {
        auth.userDetailsService(customUserDetailsService).passwordEncoder(passwordEncoder);
    }

    @Override
    protected void configure(HttpSecurity http) throws Exception {
        http
                .sessionManagement().sessionCreationPolicy(SessionCreationPolicy.STATELESS).and()
                .exceptionHandling().authenticationEntryPoint(restAuthenticationEntryPoint).and()

                .authorizeRequests().antMatchers("/api/auth/**").permitAll()
                .anyRequest().authenticated().and()
                .cors().and()
                .addFilterBefore(new TokenAuthenticationFilter(tokenUtils, authenticationManager(), customUserDetailsService), BasicAuthenticationFilter.class)
                .headers()
                // xss protection can detect what looks like xss and blocks it - browser wont render if it sees it
                .xssProtection()
                .and()
                .contentSecurityPolicy("script-src 'self'")
                .and();
        http.csrf().disable();
    }

    // Definisanje konfiguracije koja utice na generalnu bezbednost aplikacije
    @Override
    public void configure(WebSecurity web) throws Exception {
        web.ignoring().antMatchers(HttpMethod.POST, "/api/login");
        web.ignoring().antMatchers(HttpMethod.GET, "/", "/webjars/**", "/*.html", "favicon.ico", "/**/*.html",
                "/**/*.css", "/**/*.js");
    }

}



