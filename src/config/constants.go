package config

import "github.com/tohruyaginuma/TranscriptKeeper-Backend/src/lib"

const Port = "8080"
const FirebaseUIDKey = "firebaseUID"
const UserIDKey = "userID"
const Version = "v1"

var ClientURLWeb = lib.GetEnv("CLIENT_URL_WEB", "http://localhost:5173")
var ClientURLDesktop = lib.GetEnv("CLIENT_URL_DESKTOP", "http://localhost:5174")
var GoogleApplicationCredentials = lib.GetEnv("GOOGLE_APPLICATION_CREDENTIALS", "")
var CFAPIToken = lib.GetEnv("CF_API_TOKEN", "")
var CFAccountID = lib.GetEnv("CF_ACCOUNT_ID", "")
