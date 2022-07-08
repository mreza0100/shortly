package services

import (
	"context"
	"net/url"

	"github.com/mreza0100/shortly/internal/adapters/driven/kgs"
	"github.com/mreza0100/shortly/internal/ports/driven"
	"github.com/mreza0100/shortly/internal/ports/services"
	er "github.com/mreza0100/shortly/pkg/errors"
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

func (l *link) NewLink(ctx context.Context, destination, userEmail string) (string, error) {
	shortURL := l.KGS.GetKey()

	parsedURL, err := url.Parse(destination)
	if err != nil {
		return "", er.URLNotValid
	}
	if parsedURL.Scheme == "" {
		parsedURL.Scheme = "http"
	}

	err = l.cassandraWrite.SaveLink(ctx, shortURL, parsedURL.String(), userEmail)

	return shortURL, err
}

func (l *link) GetDestinationByLink(ctx context.Context, link string) (string, error) {
	destination, err := l.cassandraRead.GetDestinationByLink(ctx, link)

	return destination, err
}
