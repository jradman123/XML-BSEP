export interface IMesssage {
    Id: string;
    SenderUsername: string;
    ReceiverUsername: string;
    MessageText:string;
    TimeSent: string;
}
export interface IMesssages {
    Messages: IMesssage[]
}