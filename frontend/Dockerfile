# Use a Node.js base image
FROM node:23

# Set the working directory
WORKDIR /app

# Copy Yarn cache into the container
COPY yarn_cache /usr/local/share/.cache/yarn

# Copy package.json and yarn.lock to leverage Docker cache
COPY package.json yarn.lock ./

# Install dependencies using Yarn in offline mode
RUN yarn install --frozen-lockfile

# Copy the rest of the application
COPY . .

# Expose the development server port
EXPOSE 5173

# Start the application
CMD ["yarn", "dev", "--host"]
