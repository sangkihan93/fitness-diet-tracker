// Use the configured API URL when available, otherwise default to the local Go API.
const API_BASE_URL =
    import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

function buildUrl(path: string): string {
    const normalizedPath = path.startsWith('/') ? path : `/${path}`;

    return `${API_BASE_URL}${normalizedPath}`;
}

function createHeaders(): HeadersInit {
    return {
        Accept: 'application/json',
        // Future auth support: add Authorization header here when Cognito is implemented.
    };
}

async function get<T>(path: string): Promise<T> {
    const response = await fetch(buildUrl(path), {
        method: 'GET',
        headers: createHeaders(),
    });

    if (!response.ok) {
        throw new Error(`GET ${path} failed with status ${response.status}`);
    }

    return response.json() as Promise<T>;
}

async function post<TResponse, TRequest>(
        path: string,
        body: TRequest
    ): Promise<TResponse> {
        const response = await fetch(buildUrl(path), {
            method: 'POST',
            headers: {
                ...createHeaders(),
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(body),
        });

        if (!response.ok) {
            throw new Error(`POST ${path} failed with status ${response.status}`);
        }

        return response.json() as Promise<TResponse>;
}

export const apiClient = {
    get,
    post,
};
