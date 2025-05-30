import React, { useState } from "react";
import { Lightbulb } from "@phosphor-icons/react";
import { Loader } from "../kit/Loader";
import { getHint } from "../../api/games";
import { Game } from "../../model/game";
import { setBestMoveHint } from "../../storage/moves-hint";
import { showErrorToast } from "../showErrorToast";

export function GetMoveHint({ game }: { game: Game }) {
	const [loading, setLoading] = useState(false);

	const findBestMove = async () => {
		setLoading(true);
		try {
			setBestMoveHint(await getHint(game));
		} catch (error) {
			showErrorToast(error);
		}
		setLoading(false);
	};

	return loading ? (
		<Loader />
	) : (
		<button type="button" onClick={findBestMove}>
			<Lightbulb /> Hint
		</button>
	);
}
