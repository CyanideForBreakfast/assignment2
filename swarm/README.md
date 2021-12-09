## Setup
- Start swarm, `docker swarm init --advertise-addr=<IP_ADDR>` (to leave use `docker swarm leave`)
- Create overlay network `docker network create --driver=overlay traefik-proxy`.
- Deploy go service `docker stack deploy -c go_server.yml go_server`.
- Deplay traefik service `docker stack deploy -c traefik.yml traefik`.
- To query use `curl --header 'Host:assignment-2-go-server' 'http://localhost:80'`
- Scale using `docker service scale <SERVICE_NAME>=<REPLICA_NUM>`
