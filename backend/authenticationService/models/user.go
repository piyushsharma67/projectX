package models

type User struct{
    
    Name  string `json:"name"`
    Email string `json:"email"`
    Mobile   string `json:"mobile"`
    Password string `json:"password"`
    HashedPassword string 

}

type UserLoginDetails struct{
	User *User
	Token string
}