package metadata

type (
	Metadata struct {
		Domain   string    `json:"domain,omitempty"`
		Context  string    `json:"context,omitempty"`
		Policies []*Policy `json:"policies"`
		Roles    []*Role   `json:"roles"`
		CbURL    string    `json:"callback,omitempty"`
	}

	Role struct {
		Name       string   `json:"name"`
		Private    bool     `json:"private,omitempty"`
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

func NewRole(name string, private bool, policyRefs ...string) *Role {
	return &Role{name, private, policyRefs}
}

func NewMetadata(domain, context, callbackURL string, policies []*Policy, roles []*Role) *Metadata {
	return &Metadata{
		Domain:   domain,
		Context:  context,
		Policies: policies,
		Roles:    roles,
		CbURL:    callbackURL,
	}
}
