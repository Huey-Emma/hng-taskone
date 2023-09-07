package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type PersonalInfo struct {
		SlackName 		string 	   `json:"slack_name"`
		CurrentDay 		string 	   `json:"current_day"`
		UTCTime 		string     `json:"utc_time"`
		Track           string     `json:"track"`
		GithubFileURL 	string 	   `json:"github_file_url"`
		GithubRepoURL 	string 	   `json:"github_repo_url"`
		StatusCode 		int 	   `json:"status_code"`
}	

func day(t time.Time) string {
		fmt := t.Format("2006-01-02 15:04:05 Monday")
		parts := strings.Split(fmt, " ")
		return parts[len(parts) - 1]
}

func validstring(s string) bool {
		return len(strings.TrimSpace(s)) > 0
}

func queryParam(v url.Values, key string) string {
		return v.Get(key)
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(code)
		return json.NewEncoder(w).Encode(v)
}

type check struct {
		field string 
		cond  bool 
		msg   string
}

type validationError struct {
		Field  string `json:"field"`
		ErrMsg string `json:"errmsg"`
}

func (e validationError) Error() string {
		return e.ErrMsg
}

func validate(checks ...check) []validationError {
		errs := make([]validationError, 0)

		for _, chk := range checks {
				if !chk.cond {
						errs = append(errs, validationError{
								Field:  chk.field,
								ErrMsg: chk.msg,
						})
				}
		}

		if len(errs) == 0 {
				return nil
		}

		return errs
} 

func homeHandler(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, 200, map[string]string{
				"message": `visit /api?slack_name=<hueyemma>&track=backend>`,
		})
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query() 

		slackname := queryParam(query, "slack_name")
		track := queryParam(query, "track")

		checks := []check{
				{"slack_name", validstring(slackname), "slack_name cannot be blank"},
				{"example_name", validstring(track), "track cannot be blank"},
		}

		if errs := validate(checks...); errs != nil {
				writeJSON(w, http.StatusUnprocessableEntity, errs)
				return 
		}

		personalInfo := PersonalInfo{
				SlackName:     slackname,
				CurrentDay:    day(time.Now()),
				UTCTime:       time.Now().UTC().Format(time.RFC3339),
				Track:         track,
				GithubFileURL: "https://github.com/Huey-Emma/hng-taskone/blob/main/app.go",
				GithubRepoURL: "https://github.com/Huey-Emma/hng-taskone",
				StatusCode:    http.StatusOK,
		}

		writeJSON(w, http.StatusOK, personalInfo)
}

func logMiddleware(n http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
				log.Println(r.Method, r.RequestURI)
				n(w, r)
		}
}

func main() {
		mux := http.NewServeMux()

		mux.HandleFunc("/", logMiddleware(homeHandler))
		mux.HandleFunc("/api", logMiddleware(infoHandler))

		log.Println("app is listening on port 8000")

		err := http.ListenAndServe(":8080", mux)

		if err != nil {
				log.Fatal(err)
		}
}

