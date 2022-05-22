package com.example.PKI;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.scheduling.annotation.EnableAsync;

@SpringBootApplication
@EnableAsync(proxyTargetClass = true)
public class PkiApplication {

	public static void main(String[] args) {
		SpringApplication.run(PkiApplication.class, args);
	}

}
