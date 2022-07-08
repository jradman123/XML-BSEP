import { NotificationSettings } from "./notification-settings";

export interface ChangeSettingsRequest {
    username : string,
    settings : NotificationSettings
}
