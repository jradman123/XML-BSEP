package com.example.PKI.dto;

public class LoginDto {
	
	private String password;
	private String email;
	
	public LoginDto() {
		super();
	}

	public LoginDto(String password, String email) {
		super();
		this.password = password;
		this.email = email;
	}

	public String getPassword() {
		return password;
	}

	public void setPassword(String password) {
		this.password = password;
	}

	public String getEmail() {
		return email;
	}

	public void setEmail(String email) {
		this.email = email;
	}
	

}
