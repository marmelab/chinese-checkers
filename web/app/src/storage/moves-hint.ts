import { create } from "zustand";
import { CellIdentifier, isSameCell } from "../model/cell";
import { PathsTree } from "../model/paths";

export interface MovesHint {
	bestMove?: CellIdentifier[];
	validMoves?: PathsTree;
}

export const useMoveHint = create<MovesHint>(() => ({}));

/**
 * Get the best move hint, if there is one.
 */
export function useBestMoveHint(): CellIdentifier[] | undefined {
	return useMoveHint((state) => state.bestMove);
}

export function useIsBestMoveHintStart(cell: CellIdentifier): boolean {
	const bestMoveHint = useBestMoveHint();
	return isSameCell(cell, bestMoveHint?.[0]);
}

export function useIsInBestMoveHint(cell: CellIdentifier): boolean {
	const bestMoveHint = useBestMoveHint();
	return bestMoveHint?.some(
		(currentCell) =>
			currentCell.row == cell.row && currentCell.column == cell.column,
	);
}

export function useIsBestMoveHintEnd(cell: CellIdentifier): boolean {
	const bestMoveHint = useBestMoveHint();
	return isSameCell(cell, bestMoveHint?.[bestMoveHint.length - 1]);
}

export function setBestMoveHint(bestMove: CellIdentifier[]): void {
	useMoveHint.setState({ bestMove });
}

export function resetMovesHint(): void {
	useMoveHint.setState({
		bestMove: undefined,
	});
}

export function useValidMoves() {
	return useMoveHint((state) => state.validMoves);
}

/**
 * Get valid moves from a cell reached using the provided move.
 * @param validMoves
 * @param move
 */
export function findValidMovesFromCell(
	validMoves: PathsTree | null,
	move: CellIdentifier[],
): PathsTree | null {
	const currentCell = move?.[0];

	if (!currentCell) return null;
	if (!validMoves) return null;

	if (isSameCell(validMoves.cell, currentCell)) {
		if (move.length == 1) return validMoves;

		for (const child of validMoves.paths) {
			const childValidMoves = findValidMovesFromCell(child, move.slice(1));
			if (childValidMoves) return childValidMoves;
		}
	}

	return null;
}

/**
 * Find a valid move to the target cell.
 * @param validMoves
 * @param target
 */
export function findValidMoveToCell(
	validMoves: PathsTree | null,
	target: CellIdentifier,
): CellIdentifier[] | null {
	if (!validMoves) return null;

	if (isSameCell(validMoves.cell, target)) {
		return validMoves.move;
	} else {
		for (const child of validMoves.paths) {
			const move = findValidMoveToCell(child, target);
			if (move) return move;
		}
		return null; // No valid move to the target cell.
	}
}

export function useIsCellReachable(
	move: CellIdentifier[],
	cell: CellIdentifier,
): boolean {
	const allValidMoves = useValidMoves();

	if (move?.length == 0) return true;

	// Try to find a valid move to the target cell from the current cell.
	return !!findValidMoveToCell(
		findValidMovesFromCell(allValidMoves, move),
		cell,
	);
}

export function setValidMoves(paths: PathsTree): void {
	useMoveHint.setState({
		validMoves: paths,
	});
}

export function resetValidMoves(): void {
	useMoveHint.setState({
		validMoves: undefined,
	});
}
