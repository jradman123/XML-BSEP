import { UserTokenState } from "./user-token-state";

export interface LogedUser {
    email : string;
    role : string;
    token : UserTokenState;
}
