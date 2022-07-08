import { EmailRequest } from "./email-request";

export interface ChangeEmailRequest {
    userId : string | null,
    email : EmailRequest,
}
