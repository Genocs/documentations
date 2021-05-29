https://docs.dapr.io/getting-started/get-started-component/

mddir my-components

dapr run --app-id myapp --dapr-http-port 3500 --components-path ./my-components

Invoke-RestMethod -Uri 'http://localhost:3500/v1.0/secrets/my-secret-store/esb-rabbitmq-connectionstring'