import { toast } from "react-toastify";
import { ApiError } from "../api/api";

export function showErrorToast(error: any): void {
	if (!(error instanceof ApiError)) {
		toast.error("Unknown error.");
		return;
	}

	toast.error(error.formattedErrorMessage);
}
