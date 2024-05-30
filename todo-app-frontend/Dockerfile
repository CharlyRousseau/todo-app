FROM node:20

WORKDIR /app

COPY /todo-app-frontend/package*.json ./
RUN npm install

COPY /todo-app-frontend/ .
RUN npm run build

EXPOSE 3000
CMD ["npm", "run", "preview"]
