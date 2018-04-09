package plugin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func sendTo(method, url string, payload map[string]interface{}) (int, string, error) {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return -1, "", err
	}

	var req *http.Request

	if payload != nil {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(payloadJSON))
		if err != nil {
			return -1, "", err
		}
	} else {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return -1, "", err
		}
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return -1, "", err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, string(body), nil
}

func validateResp(statusCode int, body string, err error) {
	if err != nil {
		panic(err)
	}
	if statusCode != 200 {
		fmt.Printf("Request failed with status code %d\n", statusCode)
		panic(body)
	}
}

func buildURL(pathname string) string {
	// todo:
	config := viper.New()
	url := config.Get("podium.url")
	return fmt.Sprintf("%s%s", url, pathname)
}

func buildDeleteLeaderboardURL(leaderboard string) string {
	var pathname = fmt.Sprintf("/l/%s", leaderboard)
	return buildURL(pathname)
}

func buildGetTopPercentURL(leaderboard string, percentage int) string {
	var pathname = fmt.Sprintf("/l/%s/top-percent/%d", leaderboard, percentage)
	return buildURL(pathname)
}

func buildUpdateScoreURL(leaderboard string, playerID string) string {
	var pathname = fmt.Sprintf("/l/%s/members/%s/score", leaderboard, playerID)
	return buildURL(pathname)
}

func buildIncrementScoreURL(leaderboard string, playerID string) string {
	return buildUpdateScoreURL(leaderboard, playerID)
}

func buildUpdateScoresURL(playerID string) string {
	var pathname = fmt.Sprintf("/m/%s/scores", playerID)
	return buildURL(pathname)
}

func buildRemoveMemberFromLeaderboardURL(leaderboard string, member string) string {
	var pathname = fmt.Sprintf("/l/%s/members/%s", leaderboard, member)
	return buildURL(pathname)
}

// page is 1-based
func buildGetTopURL(leaderboard string, page int, pageSize int) string {
	var pathname = fmt.Sprintf("/l/%s/top/%d?pageSize=%d", leaderboard, page, pageSize)
	return buildURL(pathname)
}

func buildGetPlayerURL(leaderboard string, playerID string) string {
	var pathname = fmt.Sprintf("/l/%s/members/%s", leaderboard, playerID)
	return buildURL(pathname)
}

func buildHealthcheckURL() string {
	var pathname = "/healthcheck"
	return buildURL(pathname)
}
