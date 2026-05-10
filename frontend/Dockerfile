FROM node:lts AS builder

WORKDIR /app

ARG VITE_FRONTEND_BUILD_MODE=public

COPY package*.json pnpm-lock.yaml pnpm-workspace.yaml ./
RUN corepack enable
RUN pnpm install

COPY . .

RUN pnpm run build

FROM node:lts-alpine AS runner

WORKDIR /app

COPY --from=builder --chown=node:node /app/build build/
COPY --from=builder --chown=node:node /app/package.json .

EXPOSE 3000

ENV NODE_ENV=production

USER node

CMD [ "node", "build/index.js" ]
