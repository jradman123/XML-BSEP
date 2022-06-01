export interface NewUser {

    username : string;
    password : string;
    email : string;
    recoveryEmail : string;
    phoneNumber : string;
    firstName : string;
    lastName : string;
    dateOfBirth : string | null;
    gender : string;
}
