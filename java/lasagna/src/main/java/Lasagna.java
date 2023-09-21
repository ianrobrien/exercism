public class Lasagna {

  private final static int EXPECTED_MINUTES_IN_OVEN = 40;
  private final static int MINUTES_PER_LAYER = 2;

  public int expectedMinutesInOven() {
    return EXPECTED_MINUTES_IN_OVEN;
  }

  public int remainingMinutesInOven(int actualMinutes) {
    return expectedMinutesInOven() - actualMinutes;
  }

  public int preparationTimeInMinutes(int layers) {
    return layers * MINUTES_PER_LAYER;
  }

  public int totalTimeInMinutes(int layers, int actualMinutes) {
    return preparationTimeInMinutes(layers) + actualMinutes;
  }
}
