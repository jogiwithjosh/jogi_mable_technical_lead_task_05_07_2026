import { Api } from "../lib/api";
import { Storage } from "../lib/storage";

import {
    LoginRequest,
    LoginResponse,
    User
} from "../types/auth";

const BASE_URL = import.meta.env.VITE_API_URL ?? "http://localhost:8080";

export async function login(
    request: LoginRequest
): Promise<LoginResponse> {
    
    const response = await Api.post<LoginResponse>(
        BASE_URL+"/uapi/login",
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
        BASE_URL+"/uapi/signup",
        request
    );

    return response;
}