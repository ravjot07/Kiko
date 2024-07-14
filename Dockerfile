# Use official Node.js image
FROM node:14

# Create app directory
WORKDIR /usr/src/app

# Create package.json to set module type
RUN echo '{ "type": "module" }' > package.json

# Install necessary dependencies
RUN npm install react react-dom @testing-library/react @testing-library/jest-dom @babel/preset-env @babel/preset-react jest babel-jest jest-environment-jsdom

# Copy Babel and Jest configurations
COPY .babelrc /usr/src/app/.babelrc
COPY jest.config.js /usr/src/app/jest.config.js

# Copy the user code and test files
COPY . .

# Command to run the user's code and tests
CMD ["npx", "jest", "--config", "/usr/src/app/jest.config.js"]
