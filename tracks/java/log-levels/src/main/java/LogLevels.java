public class LogLevels {

  public static String message(String logLine) {
    int startIndex = logLine.indexOf(':');
    return logLine.substring(startIndex + 1).trim();
  }

  public static String logLevel(String logLine) {
    int startIndex = logLine.indexOf('[');
    int stopIndex = logLine.indexOf(']');

    return logLine.substring(startIndex + 1, stopIndex).toLowerCase();
  }

  public static String reformat(String logLine) {
    return message(logLine) + " (" + logLevel(logLine) + ")";
  }
}
