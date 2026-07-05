import { Storage } from "./storage";

async function request<T>(
    path: string,
    options: RequestInit
): Promise<T> {

    const token = Storage.getToken();

    const response = await fetch(
        `${path}`,
        {

            ...options,

            headers: {

                "Content-Type": "application/json",

                ...(token
                    ? {
                        Authorization: `Bearer ${token}`
                    }
                    : {}),

                ...(options.headers ?? {})

            }

        }
    );

    if (!response.ok) {

        throw new Error(await response.text());

    }

    return response.json();

}

export const Api = {

    get<T>(path: string) {

        return request<T>(path, {

            method: "GET"

        });

    },

    post<T>(
        path: string,
        body: unknown
    ) {

        return request<T>(path, {

            method: "POST",
            credentials: "include",
            body: JSON.stringify(body)

        });

    },

    put<T>(
        path: string,
        body: unknown
    ) {

        return request<T>(path, {

            method: "PUT",

            body: JSON.stringify(body)

        });

    },

    delete<T>(
        path: string
    ) {

        return request<T>(path, {

            method: "DELETE"

        });

    }

};