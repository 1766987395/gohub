package v1

type ChatglmController struct {
	BaseAPIController
}

// func (ctrl *ChatglmController) Index(c *gin.Context) {

// 	// request := requests.ChatglmRequest{}
// 	// if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
// 	// 	return
// 	// }

// 	userModel := auth.CurrentUser(c)

// 	keyParts := strings.Split(config.Get("chataikey.chatglmkey"), ".")
// 	if len(keyParts) != 2 {
// 		response.Error(c, nil, "invalid apiKey format")
// 		return
// 	}

// 	id, secret := keyParts[0], keyParts[1]

// 	payload := jwt.MapClaims{
// 		"api_key":   id,
// 		"exp":       time.Now().Add(time.Duration(request.ExpSeconds) * time.Second).Unix(),
// 		"timestamp": time.Now().Unix(),
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
// 	signedToken, err := token.SignedString([]byte(secret))
// 	if err != nil {
// 		response.Error(c, err)
// 		return
// 	}

// 	response.JSON(c, gin.H{
// 		"token": signedToken,
// 	})
// }
