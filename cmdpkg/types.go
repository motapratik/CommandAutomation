package cmdpkg

//Hostlist is list of all hosts
type Hostlist struct {
	Hosts []Host `json:"hosts"`
}

// Host Detail
type Host struct {
	Hostname   string   `json:"hostname"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	Sshkeyfile string   `json:"sshkeyfile"`
	Ostype     string   `json:"ostype"`
	Conntype   string   `json:"conntype"`
	Commands   []string `json:"commands"`
}
