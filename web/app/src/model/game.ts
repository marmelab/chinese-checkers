import {z} from "zod";

const zGame = z.object({
	uuid: z.string().uuid(),
	greenPlayer: z.string(),
	redPlayer: z.string(),
});

export type Game = z.infer<typeof zGame>;
