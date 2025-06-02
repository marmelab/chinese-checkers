import { z } from "zod";
import { zCellIdentifier } from "./cell";

export const zPathsTree = z.object({
	cell: zCellIdentifier,
	move: z.array(zCellIdentifier),
	get paths(): z.ZodArray<any> {
		return z.array(zPathsTree);
	},
});

export type PathsTree = z.infer<typeof zPathsTree>;
