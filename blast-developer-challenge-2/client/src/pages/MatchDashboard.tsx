import React, { useState, useEffect } from "react";
import { fetchMatchData } from "../api/fetchMatchData";
import Players from "../components/Players";
import Rounds from "../components/Rounds";

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
				console.log("Match data received:", data);
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
			{matchData && <Players players={matchData.players} />}

			<h2 className="mt-10 text-xl font-semibold">Rounds</h2>
			{matchData && <Rounds rounds={matchData.rounds} />}
		</div>
	);
}

export default MatchDashboard;
