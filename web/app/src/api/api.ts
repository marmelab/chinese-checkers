export async function fetchApi(
	input: string | URL | Request,
	init?: RequestInit,
): Promise<any> {
	const response = await fetch(input, init);

	if (response.ok) return response.json();
	else throw await ApiError.fromResponse(response);
}

export class ApiError extends Error {
	constructor(
		public response: Response,
		public errorMessage: string | null,
	) {
		super("Invalid API call.");
	}

	get formattedErrorMessage(): string {
		return formatErrorMessage(this.errorMessage ?? "unknown error");
	}

	static async fromResponse(response: Response): Promise<ApiError> {
		return new ApiError(response, (await response.json())?.error ?? null);
	}
}

export function formatErrorMessage(errorMessage: string): string {
	return `${errorMessage[0].toUpperCase()}${errorMessage.slice(1)}.`;
}
