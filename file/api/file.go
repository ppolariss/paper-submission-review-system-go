package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	opcommon "github.com/opentreehole/go-common"
	"github.com/ppolariss/paper-submission-review-system-go/common"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	pdfSuffix = "pdf"
)

var (
	homeDir  string
	rootPath string
)

func init() {
	var err error
	homeDir, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	rootPath = path.Join(homeDir, "upload")

	// create rootPath if not exist
	if _, err = os.Stat(rootPath); os.IsNotExist(err) {
		err = os.Mkdir(rootPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	originName := file.Filename
	suffix := originName[strings.LastIndex(originName, ".")+1:]
	if suffix != pdfSuffix {
		return opcommon.BadRequest("请检查上传文件是否为PDF格式")
	}

	// create new uuid filename
	newUUID := uuid.New().String()
	newFilename := newUUID + "." + suffix
	err = c.SaveFile(file, path.Join(rootPath, newFilename))
	if err != nil {
		return err
	}

	return c.JSON(common.Response{Success: true, Data: newUUID})
}

func PreviewFile(c *fiber.Ctx) error {
	id := c.Params("essayId")
	file := path.Join(rootPath, id+".pdf")
	c.Set(fiber.HeaderContentType, "application/pdf")
	c.Set(fiber.HeaderContentDisposition, `inline; filename="`+filepath.Base(file)+`"`)
	return c.SendFile(file)
}

func DownloadFile(c *fiber.Ctx) error {
	id := c.Params("essayId")
	return c.Download(path.Join(rootPath, id+".pdf"))
}
