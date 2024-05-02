# Use Nginx official image from Docker Hub
FROM nginx:alpine

# Copy static content from public folder to the container
COPY public /usr/share/nginx/html

# Expose port 80
EXPOSE 80

# Start Nginx and keep it running in the foreground
CMD ["nginx", "-g", "daemon off;"]
