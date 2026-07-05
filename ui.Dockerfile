FROM node:22-alpine AS builder

WORKDIR /workspace

COPY script ./script
COPY ecommerce ./ecommerce

WORKDIR /workspace/ecommerce

RUN npm install

EXPOSE 5173

CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]