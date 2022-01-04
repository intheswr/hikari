# hikari
accurate osu! name checker. 

> hikari is in a very early alpha stage. any bugs encountered can be sent to nina#4321

## instructions
first off, download the latest release, and put it inside it's own folder.

after you have downloaded hikari, you'll need to obtain both your osu! session token and your xsrf token.

you will find these two tokens by opening developer tools (F12 on your keyboard) and heading to the "storage" tab.
clicking on "Cookies" and opening the "osu.ppy.sh" tab will reveal both of these values. double click on the value section of each token, and paste them into their respective place in the "config.yaml" file.

you will then need to add your own word list under the file name "lists.txt"

## compiling from source (last resort)
using the latest version of both [git](https://git-scm.com/) and [go](https://golang.org/dl/):
```sh
git clone https://github.com/nina-x/hikari
cd hikari
go run .

```
