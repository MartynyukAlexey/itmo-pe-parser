package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	cfg := AuthConfig{
		AccessToken:  "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICIwSVliSmNVLW1wbEdBdzhFMzNSNkNKTUdWa3hZdUQ2eUItdWt3RlBJOXV3In0.eyJleHAiOjE3MjU1NTE5NTcsImlhdCI6MTcyNTU1MDE1NywiYXV0aF90aW1lIjoxNzI0Njc2MjM2LCJqdGkiOiI5N2E4ZjAyMy0zY2I1LTQzNGMtYWViZC02MjA2NmJmY2NiODMiLCJpc3MiOiJodHRwczovL2lkLml0bW8ucnUvYXV0aC9yZWFsbXMvaXRtbyIsImF1ZCI6InlhbmRleCIsInN1YiI6ImQ3MjFmYzQ4LTk2YWYtNDllMy1hY2JiLTQ1NWE5Njg0NTBhMiIsInR5cCI6IkJlYXJlciIsImF6cCI6InN0dWRlbnQtcGVyc29uYWwtY2FiaW5ldCIsInNlc3Npb25fc3RhdGUiOiI2NTUyNWZmZC04Y2Q0LTQ4YTktYjI4OC04YzQ1MzljMTU1NmYiLCJhbGxvd2VkLW9yaWdpbnMiOlsiaHR0cHM6Ly9teS5pdG1vLnJ1IiwiaHR0cHM6Ly9pc3UuaWZtby5ydSJdLCJyZXNvdXJjZV9hY2Nlc3MiOnsieWFuZGV4Ijp7InJvbGVzIjpbImVkaXQ6YWNjb3VudCJdfX0sInNjb3BlIjoib3BlbmlkIHByb2ZpbGUgZWR1IHdvcmsiLCJzaWQiOiI2NTUyNWZmZC04Y2Q0LTQ4YTktYjI4OC04YzQ1MzljMTU1NmYiLCJpc3UiOjM2ODQ5MywicHJlZmVycmVkX3VzZXJuYW1lIjoiaWQzNjg0OTMifQ.DVMA92giNzIEPYLaV8tz1oMzhp6J2bVBRg5q6l8_44Tx9EKQDwJsuZz-yZWus1vBsLR_av3l8e6kIHkBFI1fFPYqZXi6LtOnKtHKgPNhPffvA8hlTaYyP8FlA7wP7nCifM0zgkF8Qy76ADrGtRwPmHgOD46SBy_OyAAUarLc1Lt15LOxk10NoIzQ32EHH9m7Kr6kU028UJEsh2o0cf3mla5y1WsIwIn2_R7Oa1EX01pkE0stNSeghzDPDP7TSJyTYYRldGvc1nXRQ6WGQovELSGvHSvM8pu_hxqC3Wt_BfZFTeYEXQLy6errLdMtMjmNhq7v8QVUG-z_UoHDLe2rrg",
		RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJjNTcxMzVjYy01ZjEwLTQ4ZTAtYTU5ZS1lYTYwMmY3ZTcxYzAifQ.eyJleHAiOjE3MjgxNDAzNTgsImlhdCI6MTcyNTU0ODM1OCwianRpIjoiMTAyYmIyMzEtNDBlOS00Y2Y1LTg1MDgtODJjYjk5MTkxNzE2IiwiaXNzIjoiaHR0cHM6Ly9pZC5pdG1vLnJ1L2F1dGgvcmVhbG1zL2l0bW8iLCJhdWQiOiJodHRwczovL2lkLml0bW8ucnUvYXV0aC9yZWFsbXMvaXRtbyIsInN1YiI6ImQ3MjFmYzQ4LTk2YWYtNDllMy1hY2JiLTQ1NWE5Njg0NTBhMiIsInR5cCI6IlJlZnJlc2giLCJhenAiOiJzdHVkZW50LXBlcnNvbmFsLWNhYmluZXQiLCJzZXNzaW9uX3N0YXRlIjoiNjU1MjVmZmQtOGNkNC00OGE5LWIyODgtOGM0NTM5YzE1NTZmIiwic2NvcGUiOiJvcGVuaWQgcHJvZmlsZSBlZHUgd29yayIsInNpZCI6IjY1NTI1ZmZkLThjZDQtNDhhOS1iMjg4LThjNDUzOWMxNTU2ZiJ9.S3Pi4xSu9NbOzEI_ro2Z6wLi6H3cJEQ_pxAUGLjDFik",
		IdToken:      "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICIwSVliSmNVLW1wbEdBdzhFMzNSNkNKTUdWa3hZdUQ2eUItdWt3RlBJOXV3In0.eyJleHAiOjE3MjU1NTAxNTgsImlhdCI6MTcyNTU0ODM1OCwiYXV0aF90aW1lIjoxNzI0Njc2MjM2LCJqdGkiOiI2M2JmMDIyMi05ODhjLTQ2NmItYmU4Zi0zNTZhYTg1ZmU2YWEiLCJpc3MiOiJodHRwczovL2lkLml0bW8ucnUvYXV0aC9yZWFsbXMvaXRtbyIsImF1ZCI6InN0dWRlbnQtcGVyc29uYWwtY2FiaW5ldCIsInN1YiI6ImQ3MjFmYzQ4LTk2YWYtNDllMy1hY2JiLTQ1NWE5Njg0NTBhMiIsInR5cCI6IklEIiwiYXpwIjoic3R1ZGVudC1wZXJzb25hbC1jYWJpbmV0Iiwic2Vzc2lvbl9zdGF0ZSI6IjY1NTI1ZmZkLThjZDQtNDhhOS1iMjg4LThjNDUzOWMxNTU2ZiIsImF0X2hhc2giOiJEMVdHWU1QamliaWFYbHRFMlIzcjdnIiwic2lkIjoiNjU1MjVmZmQtOGNkNC00OGE5LWIyODgtOGM0NTM5YzE1NTZmIiwiem9uZWluZm8iOiJFdXJvcGUvTW9zY293IiwiYmlydGhkYXRlIjoiMjAwNS0wMi0wNyIsImdlbmRlciI6Im1hbGUiLCJuYW1lIjoi0JDQu9C10LrRgdC10Lkg0JzQsNGA0YLRi9C90Y7QuiIsImdyb3VwcyI6W3sicXVhbGlmaWNhdGlvbiI6eyJjb2RlIjo2MiwibmFtZSI6ItCR0LDQutCw0LvQsNCy0YAifSwibmFtZSI6IkszMzIwIiwiY291cnNlIjozLCJmYWN1bHR5Ijp7Im5hbWUiOiLRhNCw0LrRg9C70YzRgtC10YIg0LjQvdGE0L7QutC-0LzQvNGD0L3QuNC60LDRhtC40L7QvdC90YvRhSDRgtC10YXQvdC-0LvQvtCz0LjQuSIsInNob3J0X25hbWUiOiLQpNCY0JrQoiIsImlkIjo3MjV9fSx7InF1YWxpZmljYXRpb24iOnsiY29kZSI6NjIsIm5hbWUiOiLQkdCw0LrQsNC70LDQstGAIn0sIm5hbWUiOiJLMzIyMCIsImNvdXJzZSI6MiwiZmFjdWx0eSI6eyJuYW1lIjoi0YTQsNC60YPQu9GM0YLQtdGCINC40L3RhNC-0LrQvtC80LzRg9C90LjQutCw0YbQuNC-0L3QvdGL0YUg0YLQtdGF0L3QvtC70L7Qs9C40LkiLCJzaG9ydF9uYW1lIjoi0KTQmNCa0KIiLCJpZCI6NzI1fX1dLCJpc3UiOjM2ODQ5MywicHJlZmVycmVkX3VzZXJuYW1lIjoiaWQzNjg0OTMiLCJnaXZlbl9uYW1lIjoi0JDQu9C10LrRgdC10LkiLCJsb2NhbGUiOiJydSIsIm1pZGRsZV9uYW1lIjoi0J_QtdGC0YDQvtCy0LjRhyIsImlzX3N0dWRlbnQiOnRydWUsImZhbWlseV9uYW1lIjoi0JzQsNGA0YLRi9C90Y7QuiJ9.DjBd5UscuRas4d09C79T95sxHCuzBOlby6ThZZwQgbGOtwK7ZHWDuoF-HnxLWdRdWnP6b7yLhvuTWb-sWCCV7TzsF9y9cygDH_dj3Qagj79dxpsu9vsyNAK-yOHOdqgzREhmFC4LrGPJW7SbQn9ndjVyP95kY0vo0s6OhQsxJMbvmFN_HbEeAzdrLNx74SrhQP9Fx7MFMgLr67icI9w_JMe1OvOwx0jmqu0BB-4lyiyPXmFLJcMHiK6uq-fA3atMkitmu_AUKUJ7mtXwIzIBORuK2qqmrMjNmdeFkB2p3KJYwv8u3xMPyC_jtvbj8hOTqWptSczIuojgTGDdm1VCUw",
		Scope:        "openid profile edu work",
		TokenType:    "Bearer",
	}

	// in seconds!
	sleepTime := 3
	workingTime := 0
	for {
		daySchedules, err := cfg.getSchedule()
		if err != nil {
			fmt.Println(err)
			break
		}

		for _, schedule := range daySchedules {
			for _, lesson := range schedule.Lessons {
				if lesson.TypeId == 5 && (lesson.Date.Day() == 21 || lesson.Date.Day() == 22) && lesson.SectionName == "Фитнес" {
					available := lesson.Available > 0

					if available {
						req, err := http.NewRequest("GET", "https://api.telegram.org/bot7416355085:AAHPgy8M5qf-U0_nfXOCYTrxSgfeaSA1dJo/sendMessage?chat_id=1286487696&text=%D0%95%D1%81%D1%82%D1%8C%20%D0%BC%D0%B5%D1%81%D1%82%D0%B0!", nil)
						if err != nil {
							fmt.Println(err)
							continue
						}
						resp, err := http.DefaultClient.Do(req)
						if err != nil {
							fmt.Println(err)
							continue
						}

						resp.Body.Close()
					}

					if workingTime%60 == 0 {
						fmt.Printf(
							"Lesson: %s, Date: %s, Signed: %t, Reason: %s\n",
							lesson.SectionName,
							lesson.Date,
							lesson.Signed,
							lesson.CanSignIn.UnavailableReasons,
						)
					}
				}
			}
		}

		if workingTime%60 == 0 {
			fmt.Print(time.Now().Format("15:04:05") + "\t")
			fmt.Printf("Working for %d minutes\n", workingTime/60)
		}

		workingTime += sleepTime
		time.Sleep(time.Duration(sleepTime) * time.Second)
	}
}
