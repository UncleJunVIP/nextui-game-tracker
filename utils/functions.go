package utils

import (
	"path/filepath"
	"strings"
	"time"
)

func GetGameTrackerDBPath(isDev bool) string {
	if isDev {
		return "tracker.db"
	}

	return "/mnt/SDCARD/.userdata/shared/nextui-game-tracker/tracker.db"
}

func ParseRomName(path string) string {
	filename := filepath.Base(path)
	name := strings.ReplaceAll(filename, filepath.Ext(filename), "")
	return name
}

func Now() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func BoolToInt64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}
