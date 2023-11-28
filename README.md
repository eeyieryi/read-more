# read-more

An application to help me organize and archive texts that I will repeatedly be reading.

## Setup instructions

### Dev environment

Make sure to have go installed, used for building the fs.

0. Install the packages

    ```sh
    pnpm install
    ```

1. Copy and rename `.env.example` to `.env`

    ```sh
    cp .env.example .env
    ```

2. Initiate `dev` database

    ```sh
    pnpx prisma migrate dev --name init
    ```

3. Start the development server

    ```sh
    pnpm run dev
    ```
