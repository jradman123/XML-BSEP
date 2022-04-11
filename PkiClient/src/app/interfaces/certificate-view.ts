
export enum CertificateType {
    ROOT,
    INTERMEDIATE,
    CLIENT
  }

export interface CertificateView {
    id : number;
    type : CertificateType;
    serialNumber : string;
    validFrom : string;
    validTo : string;
    isRevoked : boolean;


}

