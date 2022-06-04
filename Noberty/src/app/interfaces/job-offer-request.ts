export interface IJobOfferRequest{
    companyId: number,
    requirements: string[],
    position : string,
    jobDescription : string,
    dueDate : string | null 
}