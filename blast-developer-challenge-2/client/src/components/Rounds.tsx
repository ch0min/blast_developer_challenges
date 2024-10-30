import React from "react";

interface Round {
	round_number: number;
	start_time: number;
	end_time: number;
	duration: number;
}

interface RoundsProps {
	rounds: Round[];
}

function Rounds({ rounds }: RoundsProps) {
	return (
		<div className="p-4">
			<ul>
				{rounds.map((round) => (
					<li key={round.round_number} className="border-b border-gray-200 py-2">
						<span className="font-semibold">Round {round.round_number}</span>: Duration {round.duration} seconds
					</li>
				))}
			</ul>
		</div>
	);
}

export default Rounds;
