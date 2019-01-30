package commands

import (
	"strings"
	"testing"

	cmds "gx/ipfs/QmR77mMvvh8mJBBWQmBfQBu8oD38NUN4KE9SL2gDgAQNc6/go-ipfs-cmds"
)

func checkHelptextRecursive(t *testing.T, name []string, c *cmds.Command) {
	c.ProcessHelp()

	t.Run(strings.Join(name, "_"), func(t *testing.T) {
		if c.External {
			t.Skip("external")
		}

		t.Run("tagline", func(t *testing.T) {
			if c.Helptext.Tagline == "" {
				t.Error("no Tagline!")
			}
		})

		t.Run("longDescription", func(t *testing.T) {
			t.Skip("not everywhere yet")
			if c.Helptext.LongDescription == "" {
				t.Error("no LongDescription!")
			}
		})

		t.Run("shortDescription", func(t *testing.T) {
			t.Skip("not everywhere yet")
			if c.Helptext.ShortDescription == "" {
				t.Error("no ShortDescription!")
			}
		})

		t.Run("synopsis", func(t *testing.T) {
			t.Skip("autogenerated in go-ipfs-cmds")
			if c.Helptext.Synopsis == "" {
				t.Error("no Synopsis!")
			}
		})
	})

	for subname, sub := range c.Subcommands {
		checkHelptextRecursive(t, append(name, subname), sub)
	}
}

func TestHelptexts(t *testing.T) {
	Root.ProcessHelp()
	checkHelptextRecursive(t, []string{"ipfs"}, Root)
}
