package actions

import (
	"io"
	"os"
	"path/filepath"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

// UploadHandler handles csv uploads to the /upload endpoint
func UploadHandler(c buffalo.Context) error {
	f, err := c.File("someFile")
	if err != nil {
		return errors.WithStack(err)
	}

	if !f.Valid() {
		return errors.WithStack(errors.New("invalid file upload"))
	}

	dir := filepath.Join(".", "uploads")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.WithStack(err)
	}
	osFile, err := os.Create(filepath.Join(dir, f.String()))
	if err != nil {
		return errors.WithStack(err)
	}
	defer osFile.Close()
	_, err = io.Copy(osFile, f)
	if err != nil {
		return errors.WithStack(err)
	}

	return c.Redirect(307, "customersPath()")
}
