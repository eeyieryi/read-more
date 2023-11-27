# read-more

An application to help me organize and archive texts that I will repeatedly be reading.

## Setup instructions

0. Install the packages

    ```sh
    pnpm install
    ```

1. Copy and rename `.env.example` to `.env`

2. Initiate `dev` database

    ```sh
    pnpx prisma migrate dev --name init
    ```

3. Start the development server

    ```sh
    pnpm run dev
    ```

## Docker

0. Copy and rename `.env.example` to `.env.prod`

1. Edit the values in the `.env.prod` file

2. Edit the values in `compose.yaml` file

3. Run

    ```sh
    docker-compose up -d --build
    ```

4. Enjoy
