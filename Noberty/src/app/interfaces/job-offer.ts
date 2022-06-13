export interface IJobOffer{
    id:string,
    requirements:string[],
    position : string,
    jobDescription : string,
    dateCreated : string,
    dueDate : string,
    companyName : string,
    companyId : number
}