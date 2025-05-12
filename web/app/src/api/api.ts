export async function fetchApi(
	input: string | URL | Request,
	init?: RequestInit,
): Promise<any> {
	const response = await fetch(input, init);

	if (response.ok) return response.json();
	else throw new ApiError(response);
}

export class ApiError extends Error {
	constructor(public response: Response) {
		super("Invalid API call.");
	}

	async getApiMessage(): Promise<string | null> {
		return (await this.response.json())?.error ?? null;
	}
}
