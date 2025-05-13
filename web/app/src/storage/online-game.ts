import Cookies from "js-cookie";

export function getOnlineGamePlayerId(gameId: string): string | null {
	const onlineGames = JSON.parse(Cookies.get("online-games") ?? "{}");
	return onlineGames?.[gameId] ?? null;
}
