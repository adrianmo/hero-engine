package main

// GameRequest is used to interact with the game controller and get a reply back
type GameRequest struct {
  Response chan GameResponse
}

// GameResponse is used to respond to game requests
type GameResponse struct {
  success bool
  message string
}

// TokenRequest is used to access token-protected resources
type TokenRequest struct {
  GameRequest
  token string
}

// JoinRequest is used when a hero wants to join the game
type JoinRequest struct {
  TokenRequest
  firstName string
  lastName  string
  email     string
  twitter   string
  heroName  string
  heroClass string
}

// ActivateHeroRequest is used to activate a Hero
type ActivateHeroRequest struct {
  TokenRequest
  name string
}
