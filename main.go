package main

func main() {
	var gitHubProvider PolicyProvider = GitHubProvider{}
	username, password, _ := LoadCredentialsFromEnvironment()
	gitHubProvider.Login(username, password)
}
