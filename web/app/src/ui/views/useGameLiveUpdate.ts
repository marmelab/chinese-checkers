import { Game, zGame } from "../../model/game";
import { useEffect } from "react";

export function useGameLiveUpdate(
	gameUuid: string,
	updateGame: (game: Game) => void,
	serverName: string = "",
): void {
	useEffect(() => {
		const eventSource = new EventSource(
			`${serverName}/.well-known/mercure?topic=${gameUuid}`,
		);

		eventSource.addEventListener("message", (event) => {
			updateGame(zGame.parse(JSON.parse(event.data)));
		});

		return () => {
			eventSource.close();
		};
	}, [gameUuid]);
}
