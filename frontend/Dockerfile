FROM node:24 AS builder

WORKDIR /app

ARG VITE_FRONTEND_BUILD_MODE=public

COPY package*.json pnpm-lock.yaml ./
RUN npm i -g pnpm && pnpm install

COPY . .

RUN pnpm run build

FROM node:24-alpine AS runner

WORKDIR /app

COPY --from=builder --chown=node:node /app/build build/
COPY --from=builder --chown=node:node /app/package.json .

EXPOSE 3000

ENV NODE_ENV=production

USER node

CMD [ "node", "build/index.js" ]
