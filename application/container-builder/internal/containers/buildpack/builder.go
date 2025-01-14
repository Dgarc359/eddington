package image

import (
	"context"
	"fmt"
	"os"
	"strings"

	pack "github.com/buildpacks/pack/pkg/client"
	"github.com/null-channel/eddington/application/container-builder/models"
	"github.com/null-channel/eddington/proto/container"
	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
)

var log = zerolog.New(os.Stdout).With().Timestamp().Str("component", "builder").Logger()

type BuildOpt struct {
	BuildID   string
	RepoURL   string
	ImageName string
	BuildPack string
	Builder   string
}

// BuildPackInfo contains the buildpack and builder for a given language
type BuildPackInfo struct {
	BuildPack string
	Builder   string
}

type Builder struct {
	Client *pack.Client
	log    zerolog.Logger
	db     *bun.DB
	// used to store built images
	registry string
}

// NewBuilder returns a new builder
func NewBuilder(db *bun.DB, pack *pack.Client) (*Builder, error) {
	return &Builder{
		log:    log,
		db:     db,
		Client: pack,
	}, nil
}

// GetBuildPackInfo returns the buildpack and builder for a given language
func (b *Builder) GetBuildPackInfo(language string) (BuildPackInfo, error) {
	language = strings.ToLower(language)

	switch language {
	case "go":
		return BuildPackInfo{
			BuildPack: "paketo-buildpacks/go",
			Builder:   "paketobuildpacks/builder-jammy-full",
		}, nil
	case "nodejs":
		return BuildPackInfo{
			BuildPack: "paketo-buildpacks/nodejs",
			Builder:   "paketobuildpacks/builder-jammy-full",
		}, nil
	case "python":
		return BuildPackInfo{
			BuildPack: "paketo-buildpacks/python",
			Builder:   "paketobuildpacks/builder-jammy-full",
		}, nil

	default:
		return BuildPackInfo{}, fmt.Errorf("no buildpack found for %s", language)
	}
}
func (b *Builder) CreateImage(opt BuildOpt) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan error)

	go func() {
		err := b.Client.Build(ctx, pack.BuildOptions{
			AppPath:    opt.RepoURL,
			Builder:    opt.Builder,
			Image:      opt.ImageName,
			Buildpacks: []string{opt.BuildPack},
		})
		done <- err
	}()
	select {
	case <-ctx.Done():
		// The build was canceled
		b.log.Err(ctx.Err()).Msg("build canceled")
		return ctx.Err()
	case err := <-done:
		// The build has completed
		if err != nil {
			return err
		}

		// Build completed successfully
		b.log.Info().Msg("build completed successfully")
		// Update the build request
		err = b.UpdateBuildRequest(opt.BuildID, container.ContainerStatus_BUILT)
		if err != nil {
			b.log.Err(err).Msg("unable to update build request")
			return err
		}
		// TODO: Push the image to the registry

	}
	return nil
}

// NewBuildRequest creates a new build request in the database
func (b *Builder) NewBuild(customerID, repoName, buildID string) error {
	_, err := b.db.NewInsert().Model(&models.Build{
		CustomerID: customerID,
		RepoName:   repoName,
		BuildID:    buildID,
	}).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// UpdateBuildRequest updates the build request in the database
func (b *Builder) UpdateBuildRequest(buildID string, status container.ContainerStatus) error {
	_, err := b.db.NewUpdate().Model(&models.Build{
		BuildID: buildID,
		Status:  status,
	}).Where("build_id = ?", buildID).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// GetBuildRequest returns the build request from the database
func (b *Builder) GetBuild(buildID string) (*models.Build, error) {
	var req models.Build
	err := b.db.NewSelect().Model(&req).Where("build_id = ?", buildID).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return &req, nil
}
