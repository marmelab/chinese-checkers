import React from "react";
import "./BotMaxTimeInput.css";

export function BotMaxTimeInput({
	value,
	onChange,
}: {
	value: number;
	onChange: (value: number) => void;
}) {
	return (
		<div className="max-time-input">
			<label htmlFor="max-time">Robot maximum thinking time (in seconds)</label>
			<input
				type="number"
				id="max-time"
				name="max-time"
				min={1}
				max={60}
				value={value}
				onChange={(event) => onChange(event.currentTarget.valueAsNumber)}
			/>
			<input
				type="range"
				min={1}
				max={60}
				value={value}
				onChange={(event) => onChange(event.currentTarget.valueAsNumber)}
			/>
		</div>
	);
}
