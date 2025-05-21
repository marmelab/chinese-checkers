import { FunctionField } from "react-admin";
import { Chip } from "@mui/material";
import { GamePlayer } from "../../../web/app/src/model/game-player";

export function PlayerTeamField() {
	return (
		<FunctionField
			render={(record) =>
				record.game_player == GamePlayer.Green.valueOf() ? (
					<Chip color="success" variant="outlined" label={"Green"} />
				) : (
					<Chip color="error" variant="outlined" label={"Red"} />
				)
			}
		/>
	);
}
