package metadata

import (
	"time"

	"github.com/Masterminds/semver"
	"github.com/talos-systems/gitmeta/pkg/git"
)

// Metadata contains metadata.
type Metadata struct {
	git *git.Git

	Container *Container
	Git       *Git
	Version   *Version
	Built     string
}

// Container contains container specific metadata.
type Container struct {
	Image *Image
}

// Image contains information used to identity an image.
type Image struct {
	Name string
	Tag  string
}

// Git contains git specific metadata.
type Git struct {
	Branch   string
	Ref      string
	Message  string
	SHA      string
	Tag      string
	Status   string
	Describe string
	IsBranch bool
	IsClean  bool
	IsTag    bool
}

// Version contains version specific metadata.
type Version struct {
	Original     string
	Major        int64
	Minor        int64
	Patch        int64
	Prerelease   string
	IsPrerelease bool
}

// NewMetadata initializes and returns a Metadata struct.
func NewMetadata(git *git.Git) (m *Metadata, err error) {
	m = &Metadata{git: git}

	m.Built = time.Now().UTC().Format(time.RFC1123)
	if err := addMetadataForGit(m); err != nil {
		return nil, err
	}
	if err := addMetadataForContainer(m); err != nil {
		return nil, err
	}
	if err := addMetadataForVersion(m); err != nil {
		return nil, err
	}

	return m, nil
}

func addMetadataForVersion(m *Metadata) error {
	m.Version = &Version{}
	if m.Git.IsTag {
		var ver *semver.Version
		ver, err := semver.NewVersion(m.Git.Tag)
		if err != nil {
			return err
		}
		m.Version.Original = ver.Original()
		m.Version.Major = ver.Major()
		m.Version.Minor = ver.Minor()
		m.Version.Patch = ver.Patch()
		m.Version.Prerelease = ver.Prerelease()
		if ver.Prerelease() != "" {
			m.Version.IsPrerelease = true
		}
	}

	return nil
}

//nolint: unparam
func addMetadataForContainer(m *Metadata) error {
	tag := m.Git.Describe

	containerMetadata := &Container{
		Image: &Image{
			Name: "",
			Tag:  tag,
		},
	}

	m.Container = containerMetadata

	return nil
}

func addMetadataForGit(m *Metadata) (err error) {
	m.Git = &Git{}
	if err = addBranchMetadataForGit(m.git, m); err != nil {
		return err
	}
	if err = addRefMetadataForGit(m.git, m); err != nil {
		return err
	}
	if err = addMessageMetadataForGit(m.git, m); err != nil {
		return err
	}
	if err = addStatusMetadataForGit(m.git, m); err != nil {
		return err
	}
	if err = addSHAMetadataForGit(m.git, m); err != nil {
		return err
	}
	if err = addTagMetadataForGit(m.git, m); err != nil {
		return err
	}
	if err = addDescribeMetadataForGit(m.git, m); err != nil {
		return err
	}

	return nil
}

func addBranchMetadataForGit(g *git.Git, m *Metadata) error {
	branch, isBranch, err := g.Branch()
	if err != nil {
		return err
	}
	m.Git.Branch = branch
	m.Git.IsBranch = isBranch

	return nil
}

func addRefMetadataForGit(g *git.Git, m *Metadata) error {
	ref, err := g.Ref()
	if err != nil {
		return err
	}
	m.Git.Ref = ref

	return nil
}

func addMessageMetadataForGit(g *git.Git, m *Metadata) error {
	message, err := g.Message()
	if err != nil {
		return err
	}
	m.Git.Message = message

	return nil
}

func addSHAMetadataForGit(g *git.Git, m *Metadata) error {
	sha, err := g.SHA()
	if err != nil {
		return err
	}

	if !m.Git.IsClean {
		sha += "-dirty"
	}

	m.Git.SHA = sha

	return nil
}

func addStatusMetadataForGit(g *git.Git, m *Metadata) error {
	status, isClean, err := g.Status()
	if err != nil {
		return err
	}

	m.Git.Status = status
	m.Git.IsClean = isClean

	return nil
}

func addTagMetadataForGit(g *git.Git, m *Metadata) error {
	tag, isTag, err := g.Tag()
	if err != nil {
		return err
	}
	m.Git.IsTag = isTag
	if m.Git.IsTag {
		m.Git.Tag = tag
	} else {
		m.Git.Tag = "none"
	}

	return nil
}

func addDescribeMetadataForGit(g *git.Git, m *Metadata) error {
	describe, err := g.Describe()
	if err != nil {
		return err
	}
	m.Git.Describe = describe

	return nil
}
