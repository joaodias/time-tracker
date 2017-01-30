package infrastructure

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"strconv"
	"time"
)

type GCalendarHandler struct {
	Logger *Logger
}

// TODO: Really comment the duration situation and maybe a blog post :D
func (gcalendar *GCalendarHandler) CreateEvent(name string, duration int, initialTimestamp string, token string) error {
	gcalendar.Logger.Log(name + "  " + token)
	ctx := context.Background()
	scope := make([]string, 1)
	scope[0] = "https://www.googleapis.com/auth/calendar"
	config := &oauth2.Config{
		ClientID:     "612927008159-ftu0ijkhk41a8coiil2psvcksei1r49h.apps.googleusercontent.com",
		ClientSecret: "VQdVHJZVsDJsy-Vq3Um3o2_-",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://thetimetracker.surge.sh",
		Scopes:       scope,
	}
	accessToken := &oauth2.Token{
		AccessToken: token,
	}
	client := config.Client(ctx, accessToken)
	fmt.Println("Client: ", client)
	srv, err := calendar.New(client)
	if err != nil {
		return err
	}

	startTime := initialTimestamp
	description := "Actual duration in seconds: " + strconv.Itoa(duration)
	endTime := time.Now().Format(time.RFC3339)
	event := &calendar.Event{
		Summary: name,
		Start: &calendar.EventDateTime{
			DateTime: startTime,
		},
		End: &calendar.EventDateTime{
			DateTime: endTime,
		},
		Description: description,
	}

	calendarId := "primary"
	event, err = srv.Events.Insert(calendarId, event).Do()
	if err != nil {
		return err
	}
	gcalendar.Logger.Log("Created Event at " + event.HtmlLink)
	return nil
}
