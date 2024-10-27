import axios from "axios";

const API_URL = "http://localhost:8080/api/match";

export const fetchMatchData = async () => {
	try {
		const reponse = await axios.get(API_URL);
		return reponse.data;
	} catch (error) {
		console.error("Error fetching match data:", error);
		throw error;
	}
};
