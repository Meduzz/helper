package metadata

import "fmt"

const (
	IDServer = "https://id.chimps.se"
)

type (
	Metadata struct {
		IDRef          string    `json:"idref"`
		Domain         string    `json:"domain,omitempty"`
		Context        string    `json:"context,omitempty"`
		SubscribeURL   string    `json:"subscribe,omitempty"`
		UnsubscribeURL string    `json:"unsubscribe,omitempty"`
		Policies       []*Policy `json:"policies"`
		Roles          []*Role   `json:"roles"`
	}

	Role struct {
		Name       string   `json:"name"`
		PolicyRefs []string `json:"policies"`
	}

	Policy struct {
		Name  string  `json:"name"`
		Paths []*Path `json:"paths"`
	}

	Path struct {
		Method string `json:"method"`
		URL    string `json:"url"`
	}
)

func IDRef(id string) string {
	return fmt.Sprintf("%s/id/%s", IDServer, id)
}

func NewPath(method, url string) *Path {
	return &Path{method, url}
}

func GET(url string) *Path {
	return NewPath("GET", url)
}

func POST(url string) *Path {
	return NewPath("POST", url)
}

func PUT(url string) *Path {
	return NewPath("PUT", url)
}

func DELTETE(url string) *Path {
	return NewPath("DELETE", url)
}

func PATCH(url string) *Path {
	return NewPath("PATCH", url)
}

func Policies(policies ...*Policy) []*Policy {
	return policies
}

func NewPolicy(name string, paths ...*Path) *Policy {
	return &Policy{name, paths}
}

func Roles(roles ...*Role) []*Role {
	return roles
}

func NewRole(name string, roleRefs ...string) *Role {
	return &Role{name, roleRefs}
}

func NewMetadata(idref, domain, context, subscribe, unsubscribe string, policies []*Policy, roles []*Role) *Metadata {
	return &Metadata{
		IDRef:          idref,
		Domain:         domain,
		Context:        context,
		SubscribeURL:   subscribe,
		UnsubscribeURL: unsubscribe,
		Policies:       policies,
		Roles:          roles,
	}
}
