## Simple Golang bot for telegram
1. Setup you API_TOKEN from BotFather to .env.example file
   
2. Copy .env.example to .env
    ```shell
    cp .env.example .env
    ```

3. Build
   ```shell
   env GOOS=linux go build  -o bin/webhook main.go
    ```
   
4. Deploy with sls
    ```shell
       serverless deploy  -v
    ```
