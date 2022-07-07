package services

import (
	"context"

	"github.com/mreza0100/shortly/internal/adapters/driven/kgs"
	"github.com/mreza0100/shortly/internal/ports/driven"
	"github.com/mreza0100/shortly/internal/ports/services"
)

type LinkServiceOptions struct {
	CassandraRead  driven.CassandraReadPort
	CassandraWrite driven.CassandraWritePort
	KGS            kgs.KGS
	BaseURL        string
}

func NewLinkService(opt LinkServiceOptions) services.LinkServicePort {
	return &link{
		cassandraRead:  opt.CassandraRead,
		cassandraWrite: opt.CassandraWrite,
		KGS:            opt.KGS,
		baseURL:        opt.BaseURL,
	}
}

type link struct {
	cassandraRead  driven.CassandraReadPort
	cassandraWrite driven.CassandraWritePort
	KGS            kgs.KGS
	baseURL        string
}

func (l *link) NewLink(ctx context.Context, link, userEmail string) (string, error) {
	shortURL := l.KGS.GetKey()

	err := l.cassandraWrite.SaveLink(ctx, shortURL, link, userEmail)

	return shortURL, err
}
