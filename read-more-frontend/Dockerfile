FROM node:lts-alpine3.18 AS base

WORKDIR /app

FROM base AS setup
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
COPY . /app
RUN pnpm install --frozen-lockfile

FROM setup AS build

ARG PUBLIC_BACKEND_BASE_URL
ENV PUBLIC_BACKEND_BASE_URL $PUBLIC_BACKEND_BASE_URL
ARG PUBLIC_BACKEND_API
ENV PUBLIC_BACKEND_API $PUBLIC_BACKEND_API

RUN pnpm run build

FROM base AS run

COPY --from=build /app/build /app/build
COPY --from=build /app/package.json /app/package.json

CMD node build/