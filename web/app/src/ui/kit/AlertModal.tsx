import React from "react";
import { Warning } from "@phosphor-icons/react";
import "./AlertModal.css";
import { Modal } from "./Modal";
import { useModalClose } from "./Modals";

export function AlertModal({ children }: { children: string }) {
	const close = useModalClose();

	return (
		<Modal className="alert">
			<Warning className="alert icon" weight="duotone" />

			<p className="center">{children}</p>

			<button onClick={close}>OK</button>
		</Modal>
	);
}
