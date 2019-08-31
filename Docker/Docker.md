* docker container top ... process list in one container
* docker container inspect ... details of one container config
* docker container stats ... performance stats for all containers
* docker container run -it ... start new container interactively
* docker container exec -it ... run additional process in running container
** -d, --detach ... run conatiner in background and print container id
** -t, --tty ... pseudo-tty, simulates a real terminal, like what SSH does
** -i, --interactively ... interactive, Keep session open to recieve terminal input
* docker container start -ai CONTAINER ... start one or moer stopped containers
** -a, --attach ... attach STDOUT/STDERR and forward signals
* docker container run -p 80:80
** -p, --publish ... remember publishing ports is always in HOST:CONTAINER format
* docker container port
* docker network ls
* docker network inspect
* docker network create
