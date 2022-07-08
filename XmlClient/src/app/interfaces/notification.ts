export interface INotification {
    id : string,
    content : string,
    from : string,
    to : string,
    read : boolean,
    redirectPath : string,
    notificationType : string,
    time : Date
}
