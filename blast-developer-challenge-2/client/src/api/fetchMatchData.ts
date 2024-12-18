import axios from "axios";

const API_URL = "http://localhost:8080/api/match";

export const fetchMatchData = async () => {
	try {
		const response = await axios.get(API_URL);
		console.log("Fetched match data:", response.data);
		return response.data;
	} catch (error) {
		console.error("Error fetching match data:", error);
		throw error;
	}
};
