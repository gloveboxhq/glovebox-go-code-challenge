# carrierproxy-poc

This repository is used for creating a proof of concept around the `PolicyProvider` interface.

# About the project
Go-rod scraping library is used.

Login method from PolicyProvider interface has been implemented for facebook login. The process of login method is given below:

1) Create .env file and give appropriate username, password and showbrowser as given in .env_example file as shown below.

```
    FACEBOOK_USERNAME=your_username

    FACEBOOK_PASSWORD=your_password

    SHOW_BROWSER = true
```
  & run below command 

```bash 
    $ make serve
```

2) While login, There are chances of success or failure.
3) Goroutine and Channel are used to handle response from login function.


# Working of goroutine and channel

There are two goroutine functions.They are CheckLoginSuccess and CheckLoginFailure for handling success and failure. Goroutines communicate using channels.There are two channels.They are loginSuccessResponse and loginFailureResponse channels.

After providing username and password, We will press enter button .Then search for logout button, if it matches then it is a successful login.

In case of login failure, there are two possibilities either it fails due to incorrect email or incorrect password. Search for error text input that comes only after login failure. 

The select statement lets a goroutine wait on success or failure operations. 
A select blocks until one of its cases can run, then it executes that case. 

While running this program, both channels are ready for execution, so the select statement executes statement given inside that particular case.

The Time after function is used in select statement, which is executed if none of the channels are ready for execution.


