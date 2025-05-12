import React from "react";
import { CheckCircle, TrashSimple } from "@phosphor-icons/react";
import "./MoveActionBar.css";
import { MoveState } from "../board/PlayableGameBoard";

export function MoveActionsBar({
	move,
	onCancel,
	onSubmit,
}: {
	move: MoveState;

	onCancel: () => void;
	onSubmit: () => void;
}) {
	// A move is submittable if at least 2 cells are clicked.
	const isSubmittable = move?.length >= 2;

	return (
		<div className="move-actions-bar">
			<button className="cancel" onClick={onCancel}>
				<TrashSimple /> Cancel
			</button>
			<button disabled={!isSubmittable} onClick={onSubmit}>
				<CheckCircle /> Submit
			</button>
		</div>
	);
}
