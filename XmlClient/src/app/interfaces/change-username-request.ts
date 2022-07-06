import { UsernameRequest } from "./username-request";

export interface ChangeUsernameRequest {
    userId : string | null,
    username : UsernameRequest,
}
