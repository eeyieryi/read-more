FROM node:lts-alpine3.18 as base

FROM base AS setup
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
WORKDIR /app

FROM setup AS build

# Copy non ignored files to /app folder
COPY . /app
RUN mv /app/.env.prod /app/.env

# Clean up
RUN rm /app/nginx.conf

# Install packages
RUN pnpm install --frozen-lockfile

# Init prisma client
RUN pnpx prisma migrate dev --name init

# Build svelte kit project
RUN pnpm run build

FROM build AS run
CMD pnpm run preview --host
