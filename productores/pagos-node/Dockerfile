FROM node:18-alpine
WORKDIR /app
COPY package*.json ./
RUN npm install --omit=dev
COPY productor.js .
CMD ["node", "productor.js"]
