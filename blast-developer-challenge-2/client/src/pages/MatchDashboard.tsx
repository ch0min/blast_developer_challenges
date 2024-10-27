import React, { useState, useEffect } from "react";
import { fetchMatchData } from "../api/fetchMatchData";

interface Player {
	name: string;
	kills: number;
	deaths: number;
}

interface Round {
	roundNumber: number;
	startTime: number;
	endTime: number;
	duration: number;
}

interface Match {
	rounds: Round[];
	players: { [key: string]: Player };
	totalRounds: number;
}

function MatchDashboard() {
	const [matchData, setMatchData] = useState<Match | null>(null);
	const [loading, setLoading] = useState<boolean>(true);
	const [error, setError] = useState<string | null>(null);

	useEffect(() => {
		const getMatchData = async () => {
			try {
				const data = await fetchMatchData();
				setMatchData(data);
				setLoading(false);
				// eslint-disable-next-line @typescript-eslint/no-unused-vars
			} catch (error) {
				setError("Failed to fetch match data");
				setLoading(false);
			}
		};

		getMatchData();
	}, []);

	if (loading) {
		return <div className="text-center">Loading...</div>;
	}

	if (error) {
		return <div className="text-center text-red-500">{error}</div>;
	}

	return (
		<div className="p-4">
			<h1 className="text-2xl font-bold mb-4">Match Dashboard</h1>
			<h2 className="text-xl font-semibold">Players</h2>
			<ul className="mb-4">
				{matchData &&
					Object.entries(matchData.players).map(([name, player]) => (
						<li key={name} className="border-b border-gray-200 py-2">
							<span className="font-semibold">{name}</span>: {player.kills} Kills, {player.deaths} Deaths
						</li>
					))}
			</ul>

			<h2 className="text-xl font-semibold">Rounds</h2>
			<ul>
				{matchData &&
					matchData.rounds.map((round) => (
						<li key={round.roundNumber} className="border-b border-gray-200 py-2">
							<span className="font-semibold">Round {round.roundNumber}</span>: Duration{" "}
							{round.duration.toFixed(2)} seconds
						</li>
					))}
			</ul>
		</div>
	);
}

export default MatchDashboard;
