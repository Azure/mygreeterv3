package logattrs

import (
	log "log/slog"
	// TODO: Internal AKS pkg, use generic pkg in future
	"go.goms.io/aks/rp/toolkit/aksbinversion"
)

var attrs []log.Attr

func init() {
	attrs = append(attrs,
		log.String("version", aksbinversion.GetVersion()),
		log.String("branch", aksbinversion.GetGitBranch()),
	)

}

func GetAttrs() []log.Attr {
	return attrs
}
