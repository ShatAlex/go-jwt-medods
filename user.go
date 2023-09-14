package tokens

import "time"

type User struct {
	Username     string    `json:"username" bson:"username" binding:"required"`
	Password     string    `json:"password" bson:"password" binding:"required"`
	Guid         string    `json:"guid" bson:"guid" binding:"required"`
	RefreshToken string    `json:"refresh_token" bson:"refreshtoken"`
	ExpiresAt    time.Time `json:"expiresAt" bson:"expiresat"`
}

type SignUpUser struct {
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type SignInUser struct {
	Guid string `json:"guid" bson:"guid" binding:"required"`
}
