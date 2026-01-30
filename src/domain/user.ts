export const ERR_INVALID_USER_ID = "invalid user id"
export const ERR_INVALID_GOOGLE_ID = "invalid google id"

export class UserId {
    public readonly value: number

    constructor(value: number) {
        this.validate(value)

        this.value = value
    }

    private validate(value: number) {
        if (value <= 0) {
            throw new Error(ERR_INVALID_USER_ID)
        }
    }
}

export class GoogleId {
    public readonly value: string

    constructor(value: string) {
        this.validate(value)

        this.value = value
    }

    private validate(value: string) {
        if (!value) {
            throw new Error(ERR_INVALID_GOOGLE_ID)
        }
    }
}

export class User {
    public readonly id: UserId
    public readonly googleId: GoogleId

    constructor(id: UserId, googleId: GoogleId) {
        this.id = id
        this.googleId = googleId
    }
}