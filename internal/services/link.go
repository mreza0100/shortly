package services

import (
	"context"
	"log"
	"net/url"
	"os"

	er "github.com/mreza0100/shortly/internal/pkg/errors"
	"github.com/mreza0100/shortly/internal/ports"
)

// Link Service Dependencies
type LinkServiceDep struct {
	StorageRead  ports.StorageReadPort
	StorageWrite ports.StorageWritePort
	KGS          ports.KGS
}

// Constructor of link service
func NewLinkService(opt *LinkServiceDep) ports.LinkServicePort {
	return &link{
		storageRead:  opt.StorageRead,
		storageWrite: opt.StorageWrite,
		kgs:          opt.KGS,
		errLogger:    log.New(os.Stderr, "UserService: ", log.LstdFlags),
	}
}

// Link service implementation
type link struct {
	storageRead  ports.StorageReadPort
	storageWrite ports.StorageWritePort
	kgs          ports.KGS
	errLogger    *log.Logger
}

// Service create new link and return short link
func (l *link) NewLink(ctx context.Context, destination, userEmail string) (string, error) {
	// Create new short link
	shortLink := l.kgs.GetKey()

	// Check if destination is a valid URL
	parsedURL, err := url.Parse(destination)
	if err != nil {
		return "", er.InvalidURL
	}
	// Check if destination have a scheme (http or https)
	if parsedURL.Scheme == "" {
		// If not, add http:// to destination
		parsedURL.Scheme = "http"
	}

	// Save new link to storage
	err = l.storageWrite.SaveLink(ctx, shortLink, parsedURL.String(), userEmail)
	if err != nil {
		l.errLogger.Printf("Error saving link: %e", err)
		return "", er.GeneralFailure
	}

	return shortLink, nil
}

// Service Get Destination By Link and return destination URL
func (l *link) GetDestinationByLink(ctx context.Context, short string) (string, error) {
	// Search for the link in storage
	link, err := l.storageRead.GetLinkByShort(ctx, short)
	if err != nil {
		// if link not found, just return Not Found error
		if err == er.NotFound {
			return "", err
		}
		l.errLogger.Printf("Error getting link by link: %e", err)
		return "", er.GeneralFailure
	}

	// Return destination URL
	return link.Destination, nil
}
