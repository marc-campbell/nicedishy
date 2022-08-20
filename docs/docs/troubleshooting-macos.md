# Troubleshooting MacOS

## Authentication token
The MacOS version of the NiceDishy agent stores an authentication token in the local keychain. When you uninstall and reinstall the software, it's recommended to remove this token.
Before following the steps here, make sure that the NiceDishy software is not running.

#### Open Keychain Access
First, open KeyChain Access. You can do this with Cmd+Space and then type Keychain and press enter. A window similar to the one below will open:

![Keychain Access](/images/keychain-access.png)

#### Search for the token
In the search box on the top right, enter "com.nicedishy". This will filter the list of all stored tokens.

#### Select and delete 
There will be a token named "com.nicedishy.devicetoken". Highlight it, and press delete.

#### Reinstall
Now, open NiceDishy and select "Connect Device". This will repopulate this Keychain Access token with a new value and everything should work.

