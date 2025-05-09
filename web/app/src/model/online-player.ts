import {z} from "zod";
import {GamePlayer} from "./game-player";

export const zOnlinePlayer = z.object({
	uuid: z.string().uuid(),
	name: z.string(),
	gamePlayer: z.nativeEnum(GamePlayer),
});

export type OnlinePlayer = z.infer<typeof zOnlinePlayer>;
