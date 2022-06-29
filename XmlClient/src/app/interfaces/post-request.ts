
export interface IPost {
    CommentsNumber: number
    DatePosted: Date
    DislikesNumber: number
    Id: string
    ImagePaths: string
    LikesNumber: number
    Links: {
        Comment: string, 
        Like: string, 
        Dislike: string, 
        User: string
    }
    PostText: string
    Username: string
}
export interface IPostRequest {
    Username: string
    PostText:string
    ImagePaths:string
}
export interface IPosts {
    Posts: IPost[]
}