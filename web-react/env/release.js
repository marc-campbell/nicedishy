module.exports = {
  ENVIRONMENT: "production",
  API_ENDPOINT: "https://api.nicedishy.com/api/v1",
  BUILD_VERSION: (function () {
    return process.env.BUILD_VERSION;
  }()),
};

