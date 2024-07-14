module.exports = {
  transform: {
    "^.+\\.jsx?$": "babel-jest"
  },
  testEnvironment: "jest-environment-jsdom",
  transformIgnorePatterns: ["/node_modules/"]
};
