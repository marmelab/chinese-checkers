import { create } from "zustand";
import { CellIdentifier, isSameCell } from "../model/cell";

export interface MovesHint {
	bestMove?: CellIdentifier[];
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
