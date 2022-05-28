import { UserTokenState } from "./user-token-state";

export interface LoggedUser {
    email : string;
    role : string;
    token : UserTokenState;
}
