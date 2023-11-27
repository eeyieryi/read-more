# read-more

An application to help me organize and archive texts that I will repeatedly be reading.

## Setup instructions

### Dev environment

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

## Docker

### Build

0. Copy and rename `.env.example` to `.env.prod`

```sh
    cp .env.example .env.prod
```

1. (Optional) Edit the values in the `.env.prod` file

```sh
    vi .env.prod
```

2. (Optional) Edit the values in `compose.yaml` file

```sh
    vi compose.yaml
```

3. Run

```sh
    docker-compose up -d --build
```

4. Go to [http://localhost:5789/](http://localhost:5789)

5. Enjoy!

### Shutdown

0. To shutdown the app, just stop the container with:

```sh
    docker-compose down
```
