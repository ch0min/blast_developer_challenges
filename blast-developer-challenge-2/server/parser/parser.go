package parser

import (
	"blast_developer_challenges/models"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

func ensurePlayer(match *models.Match, name string) *models.Player {
	// Check player exist
	if _, exists := match.Players[name]; !exists {
		// If player don't exist, create a new player with 0 stats
		match.Players[name] = &models.Player{
			Name: name,
			Kills: 0,
			Deaths: 0,
		}
	}
	return match.Players[name]
}

func skipLastRound(rounds []models.Round) []models.Round {
	skipRounds := []models.Round{}
	
	for _, round := range rounds {
		if round.Duration == 0 {
			continue
		}
		skipRounds = append(skipRounds, round)
	}
	return skipRounds
}

func ParseMatchLog(filePath string) (*models.Match, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	match := &models.Match{
		Players: make(map[string]*models.Player),
		Rounds: []models.Round{},
	}
	scanner := bufio.NewScanner(file)

	warmUpRounds := 3
	totalRoundsStarted := 0
	totalRoundsEnded := 0

	timestampRegex := regexp.MustCompile(`(\d{2}/\d{2}/\d{4} - \d{2}:\d{2}:\d{2})`)
	roundStartRegex := regexp.MustCompile(`World triggered "Round_Start"`)
	roundEndRegex := regexp.MustCompile(`World triggered "Round_End"`)
	roundNumberRegex := regexp.MustCompile(`RoundsPlayed: (\d+)`)
	killEventRegex := regexp.MustCompile(`"([^"]+)<\d+><STEAM_[^>]+><[^>]+>" \[-?\d+ -?\d+ -?\d+\] killed "([^"]+)<\d+><STEAM_[^>]+><[^>]+>"`)


	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println("Processing line:", line)

		// Detect and parse timestamp from the line
		timestampMatch := timestampRegex.FindStringSubmatch(line)
		if len(timestampMatch) > 0 {
			timestampStr := timestampMatch[1]
			eventTime, err := time.Parse("01/02/2006 - 15:04:05", timestampStr)
			if err != nil {
				return nil, fmt.Errorf("failed to parse timestamp: %v", err)
			}
		
			// Detect and handle round number
			if roundNumberMatch := roundNumberRegex.FindStringSubmatch(line); len(roundNumberMatch) > 1 {
				roundNumber := roundNumberMatch[1]
				fmt.Println("Detected round number update:", roundNumber)

				// Update Round Number
				if len(match.Rounds) > 0 {
					lastRound := &match.Rounds[len(match.Rounds)-1]
					lastRound.RoundNumber = match.TotalRounds
				}
			}

			// Detect and handle round start
			if roundStartRegex.MatchString(line) {
				totalRoundsStarted++
				// Skip Warm-up Rounds
				if totalRoundsStarted > warmUpRounds {
					match.TotalRounds++
					match.Rounds = append(match.Rounds, models.Round{
						RoundNumber: match.TotalRounds,
						StartTime: eventTime.Unix(),
					})
					fmt.Println("Detected round start:", match.TotalRounds)
				} else {
					fmt.Println("Skipping warm-up round:", totalRoundsStarted)
				}
			}

			if roundEndRegex.MatchString(line) {
				totalRoundsEnded++
				if len(match.Rounds) > 0 {
					lastRound := &match.Rounds[len(match.Rounds)-1]
					lastRound.EndTime = eventTime.Unix()
					lastRound.Duration = float64(lastRound.EndTime - lastRound.StartTime)
					fmt.Println("Detected round end:", lastRound.RoundNumber)
				}
			}

			// Detect and handle kill events
			if killEventRegex.MatchString(line) && totalRoundsStarted > warmUpRounds {
				fmt.Println("Kill event detected:", line)
				matches := killEventRegex.FindStringSubmatch(line)
				if len(matches) < 3 {
					fmt.Println("Unexpected kill event format:", matches)
					continue
				}

				killer, victim := matches[1], matches[2]
				fmt.Printf("Detected kill: %s killed %s\n", killer, victim)

				killerPlayer := ensurePlayer(match, killer)
				victimPlayer := ensurePlayer(match, victim)

				killerPlayer.Kills++
				victimPlayer.Deaths++

			} /*else {
				fmt.Println("No kill event matched for line:", line)
				fmt.Println(">")
			}*/
		}
	}
	match.Rounds = skipLastRound(match.Rounds)
	
	return match, scanner.Err()
}