import React from "react";

interface Player {
	name: string;
	kills: number;
	deaths: number;
}

interface PlayersProps {
	players: { [key: string]: Player };
}

function Players({ players }: PlayersProps) {
	return (
		<div className="flex p-4 bg-gray-800">
			<ul className="mb-4">
				{Object.entries(players).map(([name, player]) => (
					<li key={name} className="text-left border-b border-gray-200 p-2">
						<span className="font-semibold text-gray-50 px-4">{name}</span>
						<div className="text-center">
							<span className="px-4 text-green-400">{player.kills} Kills</span>{" "}
							<span className="text-red-500">{player.deaths} Deaths</span>
						</div>
					</li>
				))}
			</ul>
		</div>
	);
}

export default Players;
