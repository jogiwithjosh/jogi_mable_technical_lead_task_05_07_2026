import { Api } from "../lib/api";
import { Storage } from "../lib/storage";

import {
    LoginRequest,
    LoginResponse,
    User
} from "../types/auth";

export async function login(
    request: LoginRequest
): Promise<LoginResponse> {

    // -------------------------------------------------------
    // Temporary mock for frontend development.
    // Replace with the Api.post() call once the Go API is ready.
    // -------------------------------------------------------

    // const response: LoginResponse = {
    //     token: "demo-token",
    //     user: {
    //         id: "1",
    //         name: "Demo User",
    //         email: request.email
    //     }
    // };

    
    const response = await Api.post<LoginResponse>(
        "http://localhost:8080/uapi/login",
        request
    );
    

    Storage.setToken(response.token);
    Storage.setUser(response.user);

    return response;
}

export function logout(): void {

    Storage.removeToken();
    Storage.removeUser();
    Storage.clearCart();
}

export function currentUser(): User | null {

    return Storage.getUser();
}

export function isAuthenticated(): boolean {

    return Storage.getToken() !== null;
}

export function getToken(): string | null {

    return Storage.getToken();
}

export interface SignupRequest {

    name: string;

    email: string;

    password: string;
}

export async function signup(
    request: SignupRequest,
) {

    const response = await Api.post<LoginResponse>(
        "http://localhost:8080/uapi/signup",
        request
    );

    return response;
}