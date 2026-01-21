FROM node:24 AS builder

WORKDIR /app

COPY package*.json .
COPY pnpm-lock.yaml .

RUN npm i -g pnpm ; pnpm install

COPY . .

RUN pnpm run build ; pnpm prune --prod

RUN npm prune --production

FROM node:24-alpine AS runner

WORKDIR /app

COPY --from=builder /app/build build/
COPY --from=builder /app/package.json .

EXPOSE 3000

ENV NODE_ENV=production

CMD [ "node", "build" ]
