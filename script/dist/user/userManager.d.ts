import { User } from '../types/user';
export declare class UserManager {
    private user?;
    identify(user: User): void;
    current(): User | undefined;
    reset(): void;
    isAuthenticated(): boolean;
}
//# sourceMappingURL=userManager.d.ts.map