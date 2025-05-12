import React, { useContext } from "react";
import { create } from "zustand";
import "./Toasts.css";

const useToastsStore = create<{
	toasts: React.ReactElement[];
}>(() => ({
	toasts: [],
}));

/**
 * @param element The React element of the toast.
 */
export function openToast(element: React.ReactElement): void {
	useToastsStore.setState({
		toasts: [...useToastsStore.getState().toasts, element],
	});
}

/**
 * @param index Index of the toast to close.
 */
export function closeToast(index: number): void {
	useToastsStore.setState({
		toasts: useToastsStore.getState().toasts.toSpliced(index, 1),
	});
}

export function useToastClose(): () => void {
	const { index } = useContext(ToastContext);
	return () => {
		closeToast(index);
	};
}

const ToastContext = React.createContext({
	index: 0,
});

export function Toasts() {
	const toasts = useToastsStore((state) => state.toasts);

	if (toasts.length <= 0) return null;

	return (
		<div className="toasts">
			{toasts.map((toast, index) => (
				<ToastContext.Provider
					key={index}
					value={{
						index,
					}}
				>
					{toast}
				</ToastContext.Provider>
			))}
		</div>
	);
}
