# Discord Username
Starting in **May**, Discord users will have their own unique **@username**. Usernames like **example#0000** will no longer exist.
This simple script checks whether a given **username** is available or not.

# Response Code
* 40001 - Username available or probably banned.
* 50035 - Username unavailable. To check the request in more detail, go to **line 46** and add **fmt.Println(string(body))**.