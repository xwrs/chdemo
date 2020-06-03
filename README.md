# Demo TICK infrastructure 

## Execution

1. On the top level run `docker-compose up` command
2. Go to Grafana at http://localhost:3000
3. Login as admin. Skip changing password or change if you want
    - username: admin
    - password: admin
4. Go to Dashboards->Containers uptime to verify that dashboard is showing graph for a custom service
5. Verify that custom service is running:
    - running `curl http://localhost` should return `Hello World!`. If this is the case then custom service is up and running
6. Trigger shutdown of the service by calling `\api\down` endpoint:
    - `curl http://localhost/api/down`
7. Go to Grafana->Containers uptime dashboard to verify that panel is (or was) in Alerting state
8. Repeat step 5 to check if service has recovered
9. Repeat steps 6 and 8 to verify behavior multiple times