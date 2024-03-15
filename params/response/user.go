package response

/**
{
  "user": {
    "email": "jake@jake.jake",
    "token": "jwt.token.here",
    "username": "jake",
    "bio": "I work at statefarm",
    "image": null
  }
}
*/

type UserAuthenticationResponse struct {
	User UserAuthenticationBody `json:"user"`
}

type UserAuthenticationBody struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

/*
*

	{
	    "profile": {
	        "username": "xxxx123123",
	        "bio": null,
	        "image": "https://api.realworld.io/images/smiley-cyrus.jpeg",
	        "following": false
	    }
	}
*/
type UserProfileResponse struct {
	UserProfile UserProfile `json:"profile"`
}

type UserProfile struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"imag"`
	Following bool   `json:"following"`
}
