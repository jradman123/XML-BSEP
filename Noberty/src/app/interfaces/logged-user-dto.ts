import { UserStateToken } from "./user-state-token";

export interface LoggedUserDto {
    username : string;
    role : string
    token : UserStateToken;
}
