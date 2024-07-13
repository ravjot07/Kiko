# Use official Node.js image
FROM node:14

# Create app directory
WORKDIR /usr/src/app

# Create package.json to set module type
RUN echo '{ "type": "module" }' > package.json

# Install necessary dependencies
RUN npm install react react-dom

# Command to run the user's code
CMD ["node", "run_user_code.js"]
