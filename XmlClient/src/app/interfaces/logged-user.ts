import { UserTokenState } from "./user-token-state";

export interface LoggedUser {
    username : string;
    email : string;
    role : string;
    token : UserTokenState;
}
