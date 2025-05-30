import { z } from "zod";

export const zGameEvaluation = z.object({
	evaluation: z.object({
		green: z.number().int(),
		red: z.number().int(),
	}),
});

export type GameEvaluation = z.infer<typeof zGameEvaluation>;
