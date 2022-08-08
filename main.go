package main

func main() {
	var gitHubProvider PolicyProvider = GitHubProvider{}
	login, password, _ := LoadCredentialsFromEnvironment()
	gitHubProvider.Login(login, password)
}
