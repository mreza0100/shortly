package services

import (
	"context"
	"log"
	"net/url"
	"os"

	er "github.com/mreza0100/shortly/internal/pkg/errors"
	"github.com/mreza0100/shortly/internal/ports"
)

type LinkServiceOptions struct {
	CassandraRead  ports.StorageReadPort
	CassandraWrite ports.StorageWritePort
	KGS            ports.KGS
}

func NewLinkService(opt *LinkServiceOptions) ports.LinkServicePort {
	return &link{
		cassandraRead:  opt.CassandraRead,
		cassandraWrite: opt.CassandraWrite,
		KGS:            opt.KGS,
		errLogger:      log.New(os.Stderr, "UserService: ", log.LstdFlags),
	}
}

type link struct {
	cassandraRead  ports.StorageReadPort
	cassandraWrite ports.StorageWritePort
	KGS            ports.KGS
	errLogger      *log.Logger
}

func (l *link) NewLink(ctx context.Context, destination, userEmail string) (string, error) {
	shortURL := l.KGS.GetKey()

	parsedURL, err := url.Parse(destination)
	if err != nil {
		return "", er.InvalidURL
	}
	if parsedURL.Scheme == "" {
		parsedURL.Scheme = "http"
	}

	err = l.cassandraWrite.SaveLink(ctx, shortURL, parsedURL.String(), userEmail)
	if err != nil {
		l.errLogger.Printf("Error saving link: %e", err)
		return "", er.GeneralFailure
	}

	return shortURL, nil
}

func (l *link) GetDestinationByLink(ctx context.Context, short string) (string, error) {
	link, err := l.cassandraRead.GetLinkByShort(ctx, short)
	if err != nil {
		if err == er.NotFound {
			return "", err
		}
		l.errLogger.Printf("Error getting link by link: %e", err)
		return "", er.GeneralFailure
	}

	return link.Destination, nil
}
