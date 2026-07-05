import {
    createContext,
    ReactNode,
    useContext,
    useMemo,
    useState
} from "react";

import { Analytics } from "../lib/tracker";
import * as AuthService from "../services/auth";

import {
    LoginResponse,
    User
} from "../types/auth";

interface AuthContextValue {

    user: User | null;

    authenticated: boolean;

    loading: boolean;

    signIn(
        email: string,
        password: string
    ): Promise<void>;

    signOut(): void;
}

const AuthContext =
    createContext<AuthContextValue | null>(null);

interface Props {
    children: ReactNode;
}

export function AuthProvider({
    children
}: Props) {

   const [user, setUser] = useState<User | null>(() => AuthService.currentUser());
    const [loading] = useState(false);

    async function signIn(
        email: string,
        password: string
    ) {

        const response: LoginResponse =
            await AuthService.login({
                email,
                password
            });

        setUser(response.user);

        Analytics.identify(
            response.user.id,
            response.user.email
        );
    }

    function signOut() {

        Analytics.logout();

        AuthService.logout();

        setUser(null);

    }

    const value = useMemo(() => ({

        user,

        authenticated:
            user !== null,

        loading,

        signIn,

        signOut

    }), [user, loading]);

    return (

        <AuthContext.Provider value={value}>

            {children}

        </AuthContext.Provider>

    );

}

export function useAuthContext() {

    const context =
        useContext(AuthContext);

    if (!context) {

        throw new Error(
            "useAuthContext must be used inside AuthProvider"
        );

    }

    return context;

}