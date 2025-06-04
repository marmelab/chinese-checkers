import { z } from "zod";

export const zAccount = z.object({
	id: z.number().int(),
	name: z.string(),
	roles: z.array(z.string()),
});

export type Account = z.infer<typeof zAccount>;
