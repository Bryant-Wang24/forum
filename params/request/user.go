package request

/*
{
	"user":{
	  "username": "Jacob",
	  "email": "jake@jake.jake",
	  "password": "jakejake"
	}
  }
*/

type UserRegistrationRequest struct {
	User UserRegistrationBody `json:"user"`
}

type UserRegistrationBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
{
  "user":{
    "email": "jake@jake.jake",
    "password": "jakejake"
  }
}
*/

type UserLoginRequest struct {
	User UserLoginBody `json:"user"`
}

type UserLoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
*

	{
	  "user": {
	    "image": "https://api.realworld.io/images/smiley-cyrus.jpeg",
	    "username": "xxxx123123",
	    "bio": "zzzz",
	    "email": "xxxx123123@gmail.com",
	    "password": "asdfasdf"
	  }
	}
*/
type EditUserRequest struct {
	EditUserBody EditUserBody `json:"user"`
}

type EditUserBody struct {
	Image    string `json:"image"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
