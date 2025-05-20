import { toast } from "react-toastify";
import { ApiError, formatErrorMessage } from "../api/api";

export async function handleCallbackError(
	error: any,
	handledErrors: Record<string, () => void> = {},
) {
	if (!(error instanceof ApiError)) {
		toast.error("Unknown error.");
		return;
	}

	const errorMessage = await error.getApiMessage();

	if (handledErrors?.[errorMessage]) {
		return handledErrors[errorMessage]();
	}

	toast.error(formatErrorMessage(errorMessage));
}
