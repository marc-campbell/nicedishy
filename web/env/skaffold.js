module.exports = {
  ENVIRONMENT: "development",
  API_ENDPOINT: "http://localhost:30065/api/v1",
  BUILD_VERSION: (function () {
    return String(Date.now());
  }()),
};
