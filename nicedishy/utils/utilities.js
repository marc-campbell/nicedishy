import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import utc from "dayjs/plugin/utc";
import queryString from "query-string";
import sortBy from "lodash/sortBy";
import cronstrue from "cronstrue";
import size from "lodash/size";
import each from "lodash/each";
import find from "lodash/find";
import * as jsdiff from "diff";

dayjs.extend(utc);
dayjs.extend(relativeTime);


/**
 * @param {String} - Returns the commit SHA of the current build
 */
export function getBuildVersion() {
  return window.env.BUILD_VERSION;
}

export function calculateTimeDifference(start, end) {
  const date1 = dayjs(start);
  const date2 = dayjs(end);
  const seconds = date2.diff(date1, "s");
  let formattedDiff;
  if (seconds >= 3600) {
    const hourDiff = date2.diff(date1, "h");
    formattedDiff = `${hourDiff} hour${hourDiff === 1 ? "" : "s"}`;
  } else if (seconds >= 60) {
    const minuteDiff = date2.diff(date1, "m");
    formattedDiff = `${minuteDiff} minute${minuteDiff === 1 ? "" : "s"}`;
  } else {
    formattedDiff = `${seconds} second${seconds === 1 ? "" : "s"}`;
  }

  return formattedDiff;
}

export function secondsAgo(time) {
  const date1 = dayjs(time);
  return dayjs().diff(date1, "s");
}

/**
 * Retrieves the type of application via a watched app's metadata
 *
 * @param {String} text The string you're checking to see if it needs resizing
 * @param {Int}    maxWidth The maximum width of the texts container
 * @param {String} defaultFontSize The default font-size of the string (ex 32px)
 * @param {Int} minFontSize The minimum font-size the string can be (ex 18)
 * @return {String} new font-size for text to fit one line (ex 28px)
 */
export function dynamicallyResizeText(text, maxWidth, defaultFontSize, minFontSize) {
  let size;
  let resizerElm = document.createElement("p");
  resizerElm.textContent = text;
  resizerElm.setAttribute("class", "u-fontWeight--bold");
  resizerElm.setAttribute("style", `visibility: hidden; z-index: -1; position: absolute; font-size: ${defaultFontSize}`);
  document.body.appendChild(resizerElm);

  if (resizerElm.getBoundingClientRect().width < maxWidth) {
    resizerElm.remove();
    return defaultFontSize;
  }

  while (resizerElm.getBoundingClientRect().width > maxWidth) {
    size = parseInt(resizerElm.style.fontSize, 10);
    resizerElm.style.fontSize = `${size - 1}px`;
  }

  resizerElm.remove();
  if (minFontSize && size < minFontSize) {
    return `${minFontSize}px`;
  } else {
    // Font size needs to be 1px smaller than the last calculated size to fully fit in the container
    return `${size - 1}px`;
  }
}

/**
 * @param {Number} numerator
 * @param {Number} denominator
 * @return {String} danger, warning or check
 */
export function getPercentageStatus(numerator, denominator) {
  if (!numerator || !denominator) {
    return "unknown";
  }
  const percentage = numerator / denominator;
  return percentage < 0.1 ? "danger" : percentage < 0.25 ? "warning" : "check";
}

export function rootPath(path) {
  if (path[0] !== "/") {
    return path = "/" + path;
  } else {
    return path;
  }
}

export function formatByteSize(bytes) {
  if (bytes < 1024) {
    return bytes + "b";
  } else if (bytes < 1048576) {
    return (bytes / 1024).toFixed(2) + "kb";
  } else if (bytes < 1073741824) {
    return (bytes / 1048576).toFixed(2) + "mb";
  } else {
    return (bytes / 1073741824).toFixed(2) + "gb";
  }
}

export const Utilities = {
  getToken() {
    if (this.localStorageEnabled()) {
      return window.localStorage.getItem("token");
    } else {
      return "";
    }
  },

  getClaim(name) {
    const token = this.getToken();
    var base64Url = token.split(".")[1];
    var base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
    var jsonPayload = decodeURIComponent(atob(base64).split("").map(function(c) {
        return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(""));
    const parsed = JSON.parse(jsonPayload);

    return parsed[name];
  },

  localStorageEnabled() {
    var test = "test";
    try {
      localStorage.setItem(test, test);
      localStorage.removeItem(test);
      return true;
    } catch (e) {
      return false;
    }
  },

  getCookieValue(a) {
    var b = document.cookie.match("(^|[^;]+)\\s*" + a + "\\s*=\\s*([^;]+)");
    return b ? b.pop() : "";
  },

  isLoggedIn() {
    const hasToken = this.getToken();
    return !!hasToken;
  },

  dateFormat(date, format, localize = true) {
    if (!localize) {
      return dayjs.utc(date).format(format);
    }

    return dayjs.utc(date).local().format(format);
  },

  dateFromNow(date) {
    return dayjs.utc(date).local().fromNow();
  },

  gqlUnauthorized(message) {
    return message === "GraphQL error: Unauthorized";
  },


  // Converts string to titlecase i.e. 'hello' -> 'Hello'
  // @returns {String}
  toTitleCase(word) {
    let i, j, str, lowers, uppers;
    const _word = typeof word === "string" ? word : this;
    str = _word.replace(/([^\W_]+[^\s-]*) */g, (txt) => {
      return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();
    });

    // Certain minor words should be left lowercase unless
    // they are the first or last words in the string
    lowers = ["A", "An", "The", "And", "But", "Or", "For", "Nor", "As", "At",
      "By", "For", "From", "In", "Into", "Near", "Of", "On", "Onto", "To", "With"];
    for (i = 0, j = lowers.length; i < j; i++) {
      str = str.replace(new RegExp("\\s" + lowers[i] + "\\s", "g"), (txt) => {
        return txt.toLowerCase();
      });
    }

    // Certain words such as initialisms or acronyms should be left uppercase
    uppers = ["Id", "Tv"];
    for (i = 0, j = uppers.length; i < j; i++) {
      str = str.replace(new RegExp("\\b" + uppers[i] + "\\b", "g"), uppers[i].toUpperCase());
    }

    return str;
  },

  logoutUser() {
    const token = this.getToken();
    // TODO: for now we just remove the token,
    if (token) {
      window.localStorage.removeItem("token");
    }

    if (window.location.pathname !== "/login") {
      window.location = "/login";
    }
  },

  isEmailValid(email) {
    const newEmail = email.trim();
    const exp = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return exp.test(newEmail);
  },


  bytesToSize(bytes) {
    const sizes = ["B", "KB", "MB", "GB", "TB"];
    if (bytes === 0) { return "0 B"; }
    let i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)));
    if (i === 0) { return bytes + " " + sizes[i]; }
    return (bytes / Math.pow(1024, i)).toFixed(1) + " " + sizes[i];
  }
};
