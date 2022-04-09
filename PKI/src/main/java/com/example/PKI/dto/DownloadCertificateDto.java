package com.example.PKI.dto;

public class DownloadCertificateDto {
    private String serialNumber;
    private String commonName;

    public DownloadCertificateDto() {
    }

    public DownloadCertificateDto(String serialNumber, String commonName) {
        this.serialNumber = serialNumber;
        this.commonName = commonName;
    }

    public String getSerialNumber() {
        return serialNumber;
    }

    public void setSerialNumber(String serialNumber) {
        this.serialNumber = serialNumber;
    }

    public String getCommonName() {
        return commonName;
    }

    public void setCommonName(String commonName) {
        this.commonName = commonName;
    }
}
