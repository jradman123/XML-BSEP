export enum Gender{
    MALE,
    FEMALE

}

export interface UserData {

    username:    string,
    password:    string,
    email:       string,
    phoneNumber: string,
    firstName:   string,
    lastName:    string,
    gender:      Gender,
    dateOfBirth : Date,
    recoveryEmail : string,
}
