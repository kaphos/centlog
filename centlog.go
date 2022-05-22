package centlog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func Log(app, action, user string, data interface{}) {
	url := os.Getenv("LOGGING_URL")
	key := os.Getenv("LOGGING_KEY")
	env := os.Getenv("ENV")
	if env == "" {
		return // don't log
	}

	body, _ := json.Marshal(map[string]interface{}{
		"app":    app,
		"env":    env,
		"user":   user,
		"action": action,
		"data":   data,
	})

	respBody := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", url, respBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+key)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp.Body.Close()
}
