import {z} from "zod";

/**
 * Game player team.
 */
export enum CellContent
{
	Empty = 0,
	GreenPawn = 1,
	RedPawn = 2,
}

export const zCellContent = z.nativeEnum(CellContent);
