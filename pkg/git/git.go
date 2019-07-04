package git

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"

	billy "gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/memfs"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

// Git is a helper for git.
type Git struct {
	Filesystem billy.Filesystem
	repo       *git.Repository
}

func findDotGit(name string) (string, error) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return findDotGit(path.Join("..", name))
	}

	return filepath.Abs(name)
}

// NewGit instantiates and returns a Git struct.
func NewGit() (g *Git, err error) {
	p, err := findDotGit(".git")
	if err != nil {
		return
	}
	repo, err := git.PlainOpen(path.Dir(p))
	if err != nil {
		return
	}
	g = &Git{repo: repo}

	return g, err
}

// NewGitFromClone instantiates and returns a Git struct.
func NewGitFromClone(url string, ref plumbing.ReferenceName) (g *Git, err error) {
	fs := memfs.New()
	repo, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL:           url,
		Progress:      os.Stdout,
		ReferenceName: "refs/heads/" + ref,
		SingleBranch:  true,
		NoCheckout:    false,
	})

	if err != nil {
		return nil, err
	}

	g = &Git{Filesystem: fs, repo: repo}

	return g, err
}

// Branch returns the current git branch name.
func (g *Git) Branch() (branch string, isBranch bool, err error) {
	ref, err := g.repo.Head()
	if err != nil {
		return
	}
	if ref.Name().IsBranch() {
		isBranch = true
		branch = ref.Name().Short()
	}

	return branch, isBranch, err
}

// Ref returns the current git ref name.
func (g *Git) Ref() (ref string, err error) {
	r, err := g.repo.Head()
	if err != nil {
		return
	}

	ref = r.Name().String()

	return ref, err
}

// SHA returns the sha of the current commit.
func (g *Git) SHA() (sha string, err error) {
	ref, err := g.repo.Head()
	if err != nil {
		return
	}
	sha = ref.Hash().String()[0:7]

	return sha, err
}

// Tag returns the tag name if HEAD is a tag.
func (g *Git) Tag() (tag string, isTag bool, err error) {
	ref, err := g.repo.Head()
	if err != nil {
		return
	}
	tags, err := g.repo.Tags()
	if err != nil {
		return
	}
	err = tags.ForEach(func(_ref *plumbing.Reference) error {
		if _ref.Hash().String() == ref.Hash().String() {
			isTag = true
			tag = _ref.Name().Short()
			return nil
		}
		return nil
	})
	if err != nil {
		return
	}

	return tag, isTag, err
}

// Status returns the status of the working tree.
func (g *Git) Status() (status string, isClean bool, err error) {
	// temporary switch to calling out to git binary until issue with
	// go-git slowness on Worktree.Status is resolved
	// see: https://github.com/src-d/go-git/issues/844
	_, err = exec.Command("git", "diff-index", "--quiet", "HEAD", "--").CombinedOutput()
	if err == nil {
		isClean = true
		status = " nothing to commit, working tree clean"
	} else if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
		var porcelainStatus []byte
		porcelainStatus, err = exec.Command("git", "status", "--porcelain").Output()
		status = string(porcelainStatus)
	}

	return
}

// Message returns the commit message. In the case that a commit has multiple
// parents, the message of the last parent is returned.
func (g *Git) Message() (message string, err error) {
	ref, err := g.repo.Head()
	if err != nil {
		return
	}
	commit, err := g.repo.CommitObject(ref.Hash())
	if err != nil {
		return
	}
	if commit.NumParents() > 1 {
		parents := commit.Parents()
		for i := 1; i <= commit.NumParents(); i++ {
			var next *object.Commit
			next, err = parents.Next()
			if err != nil {
				return
			}
			if i == commit.NumParents() {
				message = next.Message
			}
		}
	} else {
		message = commit.Message
	}

	return message, err
}
