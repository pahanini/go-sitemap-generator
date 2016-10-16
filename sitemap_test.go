package sitemap

import (
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
	"io/ioutil"
)

const tmpDir = `testdata`

func init() {
	t, err := filepath.Abs(tmpDir)
	if err != nil {
		panic(err)
	}
	err = os.RemoveAll(t)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(t, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func TestGenerator(t *testing.T) {
	o := Options{
		MaxURLs:  2,
		Dir:      tmpDir,
		Filename: "a",
		BaseURL:  "http://example.com/",
	}

	g := New(o)
	require.NoError(t, g.Open())
	require.NoError(t, g.Add(URL{Loc: "test1"}))
	require.NoError(t, g.Add(URL{Loc: "test2"}))
	require.NoError(t, g.Close())

	o.Filename = "b"
	g = New(o)
	require.NoError(t, g.Open())
	require.NoError(t, g.Add(URL{Loc: "test1"}))
	require.NoError(t, g.Add(URL{Loc: "test2"}))
	require.NoError(t, g.Add(URL{Loc: "test3"}))
	require.NoError(t, g.Close())

	files, _ := ioutil.ReadDir(g.opt.Dir)
	require.Equal(t, 4, len(files))
	require.Equal(t, "a.xml", files[0].Name())
	require.Equal(t, "b-1.xml", files[1].Name())
	require.Equal(t, "b-2.xml", files[2].Name())
	require.Equal(t, "b.xml", files[3].Name())
}

func TestParamChecks(t *testing.T) {
	var (
		g   *Generator
		err error
	)
	opt := Options{
		MaxFileSize: 2,
		MaxURLs:     -1,
		BaseURL:     "/",
		Dir:         tmpDir,
	}
	g = New(opt)
	err = g.Open()
	require.Error(t, err)

	opt.MaxFileSize = len(header) + len(footer) + 10
	g = New(opt)
	err = g.Open()
	require.Error(t, err)

	opt.MaxURLs = 2
	g = New(opt)
	err = g.Open()
	require.NoError(t, err)
}

func TestInternals(t *testing.T) {
	var (
		g   *Generator
		err error
	)
	opt := Options{
		MaxFileSize: len(header) + len(footer) + 10,
		MaxURLs:     2,
		BaseURL:     "/",
		Dir:         tmpDir,
	}
	g = New(opt)
	err = g.Open()
	require.NoError(t, err)
	require.True(t, g.canFit(10))
	require.False(t, g.canFit(11))

	n1 := g.formatURLNode(URL{Loc: "test1"})
	require.Equal(t, `<url><loc>test1</loc></url>`, n1)

	n2 := g.formatURLNode(URL{Loc: "test2", ChangeFreq: ChangeFreqDaily})
	require.Equal(t, `<url><loc>test2</loc><changefreq>daily</changefreq></url>`, n2)
}
