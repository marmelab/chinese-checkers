import React, {useContext} from "react";
import {create} from "zustand";
import "./Modals.css";
import {Modal} from "./Modal";

const useModalsStore = create<{
	modals: React.ReactElement[];
}>(() => ({
	modals: [],
}));

/**
 * Open the provided modal.
 * @param element The React element of the modal.
 */
export function openModal(element: React.ReactElement): void {
	useModalsStore.setState({
		modals: [...useModalsStore.getState().modals, element],
	});
}

/**
 * Close a modal of the provided index.
 * @param index Index of the modal to close.
 */
export function closeModal(index: number): void {
	useModalsStore.setState({
		modals: useModalsStore.getState().modals.toSpliced(index, 1),
	});
}

/**
 * Get the modal close function of the current modal.
 */
export function useModalClose(): () => void {
	const {index} = useContext(ModalContext);
	return () => {
		closeModal(index);
	};
}

const ModalContext = React.createContext({
	index: 0,
});

export function Modals() {
	const modals = useModalsStore((state) => state.modals);

	if (modals.length <= 0) return null;

	return (
		<div className="modals">
			{modals.map((modal, index) => (
				<ModalContext.Provider
					key={index}
					value={{
						index,
					}}
				>
					{modal}
				</ModalContext.Provider>
			))}
		</div>
	);
}
