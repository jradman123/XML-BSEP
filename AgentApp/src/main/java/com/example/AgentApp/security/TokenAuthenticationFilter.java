package com.example.AgentApp.security;

import com.example.AgentApp.model.UserDetails;
import com.example.AgentApp.service.impl.UserDetailsService;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.web.authentication.www.BasicAuthenticationFilter;
import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

public class TokenAuthenticationFilter extends BasicAuthenticationFilter {

    private TokenUtils tokenUtils;
    private UserDetailsService userDetailsService;

    public TokenAuthenticationFilter(TokenUtils tokenHelper, AuthenticationManager authManager, UserDetailsService userDetailsService) {
        super(authManager);
        this.tokenUtils= tokenHelper;
        this.userDetailsService = userDetailsService;
    }

    @Override
    protected void doFilterInternal(HttpServletRequest req,
                                    HttpServletResponse res,
                                    FilterChain chain) throws IOException, ServletException {
        String header = req.getHeader("Authorization");

        if (header == null || !header.startsWith("Bearer ")) {
            chain.doFilter(req, res);
            return;
        }

        UsernamePasswordAuthenticationToken authentication = getAuthentication(req);

        SecurityContextHolder.getContext().setAuthentication(authentication);
        chain.doFilter(req, res);
    }

    private UsernamePasswordAuthenticationToken getAuthentication(HttpServletRequest request) {
        String token = tokenUtils.getToken(request);
        System.out.println("TOKEN " + token);
        if (token != null) {
            // parse the token.
            String username = tokenUtils.getUsernameFromToken(token);

            if (username != null) {
                UserDetails userDetails = userDetailsService.loadUserByUsername(username);
                return new UsernamePasswordAuthenticationToken(username, null, userDetails.getAuthorities());
            }

            return null;
        }
        return null;
    }
}

