package media_unknown

import (
	"github.com/smartmediafiles/media/media/maps"
	"github.com/smartmediafiles/media/media/types"
)

// List of supported media.Unknown file types.
const (
	Unknown types.FileType = ""
)

// UnknownFileTypesExtensions is a map of media.Unknown file types and their extensions.
var UnknownFileTypesExtensions = maps.MapFileTypeExtensions{
	Unknown: {".*"},
}
