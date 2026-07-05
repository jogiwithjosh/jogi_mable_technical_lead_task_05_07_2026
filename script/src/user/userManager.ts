import { User } from "../types/user";

export class UserManager {

    private user?: User;

    identify(user: User): void {
        this.user = user;
    }

    current(): User | undefined {
        return this.user;
    }

    reset(): void {
        this.user = undefined;
    }

    isAuthenticated(): boolean {
        return this.user !== undefined;
    }
}