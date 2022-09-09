package marketplace

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/reearth/reearth/server/pkg/id"
	"github.com/reearth/reearth/server/pkg/log"
	"github.com/reearth/reearth/server/pkg/plugin/pluginpack"
	"github.com/reearth/reearthx/rerror"
	"golang.org/x/oauth2/clientcredentials"
)

var pluginPackageSizeLimit int64 = 10 * 1024 * 1024 // 10MB

type Marketplace struct {
	endpoint string
	conf     *clientcredentials.Config
}

func New(endpoint string, conf *clientcredentials.Config) *Marketplace {
	return &Marketplace{
		endpoint: strings.TrimSuffix(endpoint, "/"),
		conf:     conf,
	}
}

func (m *Marketplace) FetchPluginPackage(ctx context.Context, pid id.PluginID) (*pluginpack.Package, error) {
	purl, err := m.getPluginURL(pid)
	if err != nil {
		return nil, err
	}
	return m.downloadPluginPackage(ctx, purl)
}

func (m *Marketplace) getPluginURL(pid id.PluginID) (string, error) {
	return strings.TrimSpace(fmt.Sprintf("%s/api/plugins/%s/%s.zip", m.endpoint, pid.Name(), pid.Version().String())), nil
}

/*
func (m *Marketplace) getPluginURL(ctx context.Context, pid id.PluginID) (string, error) {
	body := strings.NewReader(fmt.Sprintf(
		`{"query":"query { node(id:"%s" type:PLUGIN) { ...Plugin { url } } }"}`,
		pid.Name(),
	))
	req, err := http.NewRequestWithContext(ctx, "POST", m.endpoint+"/graphql", body)
	if err != nil {
		return "", rerror.ErrInternalBy(err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := m.client.Do(req)
	if err != nil {
		return "", rerror.ErrInternalBy(err)
	}
	if res.StatusCode != http.StatusOK {
		return "", rerror.ErrNotFound
	}
	defer func() {
		_ = res.Body.Close()
	}()
	var pluginRes response
	if err := json.NewDecoder(res.Body).Decode(&pluginRes); err != nil {
		return "", rerror.ErrInternalBy(err)
	}
	if pluginRes.Errors != nil {
		return "", rerror.ErrInternalBy(fmt.Errorf("gql returns errors: %v", pluginRes.Errors))
	}

	purl := pluginRes.PluginURL()
	if purl == "" {
		return "", rerror.ErrNotFound
	}
	return purl, nil
}

type response struct {
	Data   pluginNodeQueryData `json:"data"`
	Errors any                 `json:"errors"`
}

func (r response) PluginURL() string {
	return r.Data.Node.URL
}

type pluginNodeQueryData struct {
	Node plugin
}

type plugin struct {
	URL string `json:"url"`
}
*/

func (m *Marketplace) downloadPluginPackage(ctx context.Context, url string) (*pluginpack.Package, error) {
	var client *http.Client
	if m.conf != nil && m.conf.ClientID != "" && m.conf.ClientSecret != "" && m.conf.TokenURL != "" {
		client = m.conf.Client(ctx)
	}
	if client == nil {
		client = http.DefaultClient
	}

	log.Infof("marketplace: downloading plugin package from \"%s\"", url)

	res, err := client.Get(url)
	if err != nil {
		return nil, rerror.ErrInternalBy(err)
	}
	defer func() {
		_ = res.Body.Close()
	}()
	if res.StatusCode == http.StatusNotFound {
		return nil, rerror.ErrNotFound
	}
	if res.StatusCode != http.StatusOK {
		return nil, rerror.ErrInternalBy(fmt.Errorf("status code is %d", res.StatusCode))
	}
	return pluginpack.PackageFromZip(res.Body, nil, pluginPackageSizeLimit)
}