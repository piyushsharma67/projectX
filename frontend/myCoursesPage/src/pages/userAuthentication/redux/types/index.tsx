export type User = {
    name: string,
    email: string,
    mobile?: string,
}

export type initialStateAuth = {
    user: User,
    token: string,
    isSigningUp: boolean
}

