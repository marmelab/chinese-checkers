import React from "react";
import {CheckCircle, TrashSimple} from "@phosphor-icons/react";

export function MoveActionsBar() {
	return (
		<div className={"move-actions-bar"}>
			<button>
				<TrashSimple /> Cancel
			</button>
			<button>
				<CheckCircle /> Submit
			</button>
		</div>
	);
}
