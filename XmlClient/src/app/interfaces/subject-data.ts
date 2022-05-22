export enum Gender{
    MALE,
    FEMALE

}

export interface UserData {

    Username:    string,
    Password:    string,
    Email:       string,
    RecoveryMail: string,
    PhoneNumber: string,
    FirstName:   string,
    LastName:    string,
    Gender:      Gender,
    DateOfBirth : Date,
    RecoveryEmail : string,
}
