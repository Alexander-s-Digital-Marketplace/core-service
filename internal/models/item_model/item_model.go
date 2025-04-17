package itemmodel

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/Alexander-s-Digital-Marketplace/core-service/internal/database"
	"github.com/sirupsen/logrus"
)

type Item struct {
	Id      int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Content string `json:"content" gorm:"type:varchar(100)"`
}

func (item *Item) AddToTable() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.Create(&item).Error
	if err != nil {
		logrus.Errorln("Error add to table: ", err)
		return 503
	}
	return 200
}

func (item *Item) GetFromTable() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.First(&item).Error
	if err != nil {
		logrus.Errorln("Error get from table: ", err)
		return 503
	}
	return 200
}

// Метод для шифрования Content
func (item *Item) Encode(key []byte) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	ciphertext := make([]byte, aes.BlockSize+len(item.Content))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(item.Content))

	item.Content = hex.EncodeToString(ciphertext)
	return nil
}

// Метод для дешифрования Content
func (item *Item) Decode(key []byte) error {
	ciphertext, err := hex.DecodeString(item.Content)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	if len(ciphertext) < aes.BlockSize {
		return fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	item.Content = string(ciphertext)
	return nil
}
