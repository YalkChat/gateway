package database

import (
	"yalk/chat/models/db"
	"yalk/chat/models/events"

	"gorm.io/gorm"
)

type DatabaseImpl struct {
	conn *gorm.DB
}

func NewDatabase(conn *gorm.DB) *DatabaseImpl {
	return &DatabaseImpl{
		conn: conn,
	}
}

func (dbi *DatabaseImpl) SaveMessage(newMessage *events.Message) (*db.Message, error) {
	dbMessage := &db.Message{
		ChatID:  newMessage.ChatID,
		UserID:  newMessage.UserID,
		Content: newMessage.Content,
	}
	if err := dbi.conn.Create(dbMessage).Error; err != nil {
		return nil, err
	}
	return dbMessage, nil

}

func (dbi *DatabaseImpl) GetMessage(messageID uint) (*db.Message, error) {
	var message *db.Message
	if err := dbi.conn.Where("id = ?", messageID).First(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

func (dbi *DatabaseImpl) GetClients(chatID uint) ([]uint, error) {
	var chat *db.Chat
	result := dbi.conn.Preload("Users").Find(&chat, "id = ?", chatID)
	if result.Error != nil {
		return nil, result.Error
	}

	var userIDs []uint
	for _, user := range chat.Users {
		userIDs = append(userIDs, user.ID)
	}

	return userIDs, nil
}

// TODO: Decide what this function should reeturn
func (dbi *DatabaseImpl) NewUserWithPassword(newUser *events.UserCreationEvent) (*db.User, error) {
	dbNewUser := &db.User{Email: newUser.Email, Password: newUser.Password}
	if err := dbi.conn.Create(newUser).Error; err != nil {
		return nil, err
	}
	return dbNewUser, nil
}

// GetUsersByChatId retrieves the users associated with a specific chat ID
func (dbi *DatabaseImpl) GetUsersByChatId(chatID uint) ([]*db.User, error) {
	var chat db.Chat
	result := dbi.conn.Preload("Users").Find(&chat, "id = ?", chatID)
	if result.Error != nil {
		return nil, result.Error
	}

	return chat.Users, nil
}

// NewUser creates a new user in the database
func (dbi *DatabaseImpl) NewUser(newUser *db.User) (*db.User, error) {
	dbNewUser := &db.User{
		Email: newUser.Email,
		// Add other fields from events.User to db.User here
	}
	if err := dbi.conn.Create(dbNewUser).Error; err != nil {
		return nil, err
	}
	return dbNewUser, nil
}

// NewChat creates a new chat in the database
func (dbi *DatabaseImpl) NewChat(newChat *events.Chat) (*db.Chat, error) {
	dbNewChat := &db.Chat{
		Name: newChat.Name,
		// Add other fields from events.Chat to db.Chat here
	}
	if err := dbi.conn.Create(dbNewChat).Error; err != nil {
		return nil, err
	}
	return dbNewChat, nil
}

// NewChatType creates a new chat type in the database
func (dbi *DatabaseImpl) NewChatType(newChatType *events.ChatType) (*db.ChatType, error) {
	dbNewChatType := &db.ChatType{
		Name: newChatType.Name,
		// Add other fields from events.ChatType to db.ChatType here
	}
	if err := dbi.conn.Create(dbNewChatType).Error; err != nil {
		return nil, err
	}
	return dbNewChatType, nil
}

// IsServerInitialized checks if the server is initialized
func (dbi *DatabaseImpl) IsServerInitialized() (bool, error) {
	// Your logic here to check if the server is initialized
	return true, nil // Placeholder
}

// SaveServerSettings saves server settings to the database
func (dbi *DatabaseImpl) SaveServerSettings(newSettings *events.ServerSettings) error {
	dbSettings := &db.ServerSettings{
		// Map fields from events.ServerSettings to db.ServerSettings here
	}
	if err := dbi.conn.Create(dbSettings).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID retrieves the user by their ID
func (dbi *DatabaseImpl) GetUserByID(userID uint) (*db.User, error) {
	var user *db.User
	result := dbi.conn.Preload("Chats").Preload("Chats.ChatType").Find(&user, "id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (dbi *DatabaseImpl) GetUserByUsername(username string) (*db.User, error) {
	var user *db.User
	result := dbi.conn.Where("username = ?", username).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
