# glovebox-challenge-joe-edwards

This repository is to show a proof of concept for the GitHubProxy which implements PolicyProvider for the GitHub website.

## Dependencies
github.com/go-rod/rod v0.108.2

## Useage
The login method for the GitHubProvider takes a GitHub login (or email) and password for a valid GitHub account and performs the login action through a headless browser.  'login' and 'password' can be supplied as environment variables and loaded with the helper method LoadCredentialsFromEnvironment.

## Testing
run ```export login=<GitHub Login/Email> export password=<GitHub Password> && go test`` to execute all the unit tests associated with this project.  Supply the values in the '<...>' brackets with your account information.

## Usage
### Method 1:
- Initialize a GitHubProvider object
- Supply a valid github login/email and password  to the 'Login' method of your GitHubProvider object in your code.
### Method 2:
- Add environment variables 'login' and 'password' that correspond to a valid GitHub login/email address and password
- Use the function 'LoadCredentialsFromEnvironment' to extract login, and password values.  This function can also return an error if credentials have not been provided.
- Use the return values to supply to Method 1.

After running this function, you will see an artifact titled '<TIMESTAMP>_go_test.png' where 'TIMESTAMP' corresponds to the time the method was executed.  This contains a screenshot of the Login method after completion.  You should see the normal authenticated GitHub homepage containing data relating to the user credentials supplied.


