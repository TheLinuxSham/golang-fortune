# Welcome

This is Fortune, a golang program. It spins up a server at port 3030. If you send a GET request at 127.0.0.1:3030 you will get a json response back. The response will contain an animal and a "fortune".

The program is inspired by cowsays.

### Is there an option for a docker image?
Yes! You can build the docker image with
```docker buildx build -t golang-fortune:1.0 .``` After succesfully building it, you might run it with ```docker run -p 3030:3030 golang-fortune:1.0```

### Do I need a docker image?
No,  you can execute ```make run``` to run the program or ```make build``` to compile the program. Both ways are totally fine. The compiled program contains everything it needs and is not dependend on anything else.

### How does it work?
You start the program. Then you can, for example, ```curl 127.0.0.1:3030``` or open ```127.0.0.1:3030``` in your web browser to get a fortune. (Use localhost instead of 127.0.0.1, if you're on windows.)

### Where are the fortunes stored? I want to modify them
The file is /internal/embedding/fortunes.json. Please note the structure and don't change it. You can add animals and quotes as you like though.

### I use port 3030 already and need to change it...
Head over to /internal/server/server.go and change Addr.

# What's next?
- implement animal selection
- write more tests
- add option to return html
- ascii arts, maybe?