import java.util.Arrays;
import java.util.stream.Collectors;
import java.util.stream.Stream;

class Acronym {

  private final String phrase;

  Acronym(String phrase) {
    this.phrase = phrase;
  }

  String get() {
    return Arrays.stream(phrase.split("[ -]"))
        .flatMap(token -> Stream.of(token.split("(?=[a-zA-Z])")))
        .map(token -> token.substring(0, 1).toUpperCase())
        .collect(Collectors.joining());
  }
}
