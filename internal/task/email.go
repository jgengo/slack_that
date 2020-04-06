package task

// GetIM opens an IM channel with a specified userEmail
func (c SlackClient) GetIM(userEmail string) (string, error) {

	var userID string

	user, err := c.Value.GetUserByEmail(userEmail)

	if user != nil && user.Profile.Email != "" {
		userID = user.ID
	} else {
		return "", err
	}

	_, _, channelID, err := c.Value.OpenIMChannel(userID)
	if err != nil {
		return "", err
	}
	return channelID, nil

}
