package external

import (
	"context"

	firebase "firebase.google.com/go/v4"
	firebaseAuth "firebase.google.com/go/v4/auth"
	"github.com/tohruyaginuma/TranscriptKeeper-Backend/src/config"
	"google.golang.org/api/option"
)

func NewFirebaseApp() (app *firebase.App, err error) {
	app, err = firebase.NewApp(context.Background(), nil, option.WithCredentialsJSON([]byte(config.GoogleApplicationCredentials)))
	if err != nil {
		return nil, err
	}

	return app, nil
}

type Firebase struct {
	auth *firebaseAuth.Client
}

func NewFirebase(app *firebase.App) (*Firebase, error) {
	auth, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}
	return &Firebase{auth: auth}, nil
}

func (f *Firebase) VerifyIDToken(ctx context.Context, idToken string) (uid string, err error) {
	token, err := f.auth.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", err
	}
	uid = token.UID

	return uid, nil
}
