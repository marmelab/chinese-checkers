import { toast } from "react-toastify";
import { ApiError, formatErrorMessage } from "../api/api";

export async function handleCallbackError(error: any) {
	if (!(error instanceof ApiError)) {
		toast.error("Unknown error.");
		return;
	}

	toast.error(formatErrorMessage(await error.getApiMessage()));
}
