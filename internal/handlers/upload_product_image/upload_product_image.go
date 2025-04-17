package uploadproductimage

import (
	"fmt"
	"net/http"
	"path/filepath"

	cloudinaryconfig "github.com/Alexander-s-Digital-Marketplace/core-service/internal/config/cloudinary"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadProductImage(c *gin.Context) (int, string) {

	cloudinaryconfig.CloudinaryConfig()

	err := c.Request.ParseMultipartForm(10 << 20) // Ограничение: 10MB
	if err != nil {
		return http.StatusBadRequest, "Error form parsing"
	}
	// Извлекаем файл из запроса
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		return http.StatusBadRequest, "Error file getting"
	}
	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename) // Получаем расширение файла
	newFileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)

	// Загрузка файла в Cloudinary
	uploadResult, err := cloudinaryconfig.CLD.Upload.Upload(c.Request.Context(), file, uploader.UploadParams{
		Folder:   "product_image", // Опционально: папка в Cloudinary
		PublicID: newFileName,
	})
	if err != nil {
		return http.StatusInternalServerError, "Error Cloudinary load"
	}

	return http.StatusOK, uploadResult.SecureURL
}
