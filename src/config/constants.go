package config

import "github.com/tohruyaginuma/TranscriptKeeper-Backend/src/lib"

const Port = "8080"
const FirebaseUIDKey = "firebaseUID"
const UserIDKey = "userID"
const Version = "v1"

var ClientURLWeb = lib.GetEnv("CLIENT_URL_WEB", "http://localhost:5173")
var GoogleApplicationCredentials = lib.GetEnv("GOOGLE_APPLICATION_CREDENTIALS", "")
