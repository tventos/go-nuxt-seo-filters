FROM node:20-alpine AS builder

WORKDIR /app

RUN npm install -g pnpm@8.15.7

COPY package*.json ./
COPY pnpm-lock.yaml* ./
RUN pnpm install --frozen-lockfile

COPY . .
RUN pnpm build

FROM node:20-alpine

WORKDIR /app

RUN npm install -g pnpm@8.15.7

COPY --from=builder /app/.output ./.output
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json

EXPOSE 3000

CMD ["node", ".output/server/index.mjs"]
