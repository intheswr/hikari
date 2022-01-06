# hikari
accurate osu! name checker. 

> hikari is in a very early alpha stage. any bugs encountered can be sent to nina#4321

## instructions
first off, download the latest release, and put it inside it's own folder.

after you have downloaded hikari, you'll need to obtain both your osu! session token and your xsrf token. you will need to logged into, and on the osu! website.

### chromium based browsers
navigate to the search bar and click the "lock" icon
click cookies --> ppy.sh --> cookies
click xsrf_token then copy the content and paste it into the respective value in the config.yaml 
repeat the same for osu! session

### firefox based browsers
you will find these two tokens by opening developer tools (F12 on your keyboard) and heading to the "storage" tab
clicking on "Cookies" and opening the "osu.ppy.sh" tab will reveal both of these values
double click on the value section of each token, and paste them into their respective place in the "config.yaml" file

you will then need to add your own word list under the file name "lists.txt"

## config explaination
the user section is where you place both your osu!session token as well as your xsrf token, both of which you should have gotten in the previous steps.
```yaml
user: 
  session: ~
  xsrf: ~
```

the discord section is where you will enable discord webhooks for when usernames are found. active will toggle this feature, and webhook is where you place the webhook url (discord intergration is not complete yet.)
```yaml
discord:
  active: false
  webhook: ~
 ```
 
 the config section decides the specifics of what hikari will do. requests defines the amount of names that will be checked before hikari pauses to prevent ratelimiting, while the cooldown is the time in seconds of how long this pause will be. verbose will toggle the showing of unavailable names, as well as a few other things that may help with debugging. the dropping option will dictate whether or not names that will soon be available are classed as available or not. dropping names will have the date appended to the end of the username.
```yaml
config:
  requests: 60
  cooldown: 60
  verbose: true
  dropping: false
```

## compiling from source (last resort)
using the latest version of both [git](https://git-scm.com/) and [go](https://golang.org/dl/):
```sh
git clone https://github.com/nina-x/hikari
cd hikari
go run .

```
