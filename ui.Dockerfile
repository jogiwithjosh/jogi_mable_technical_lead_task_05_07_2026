FROM node:22-alpine AS builder

WORKDIR /workspace

COPY script ./script
COPY ecommerce ./ecommerce

WORKDIR /workspace/ecommerce

RUN npm install

RUN npm run build


FROM node:22-alpine

WORKDIR /app

ENV NODE_ENV=production
ENV PORT=5173

COPY --from=builder /workspace/ecommerce/package*.json ./

RUN npm install --omit=dev

COPY --from=builder /workspace/ecommerce/build ./build
COPY --from=builder /workspace/ecommerce/public ./public

EXPOSE 5173

CMD ["npm", "run", "start"]