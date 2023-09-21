public class CarsAssemble {

  private final static int CARS_PER_HOUR = 221;

  public double productionRatePerHour(int speed) {
    if (speed == 0) {
      return 0;
    }
    if (speed >= 1 && speed <= 4) {
      return CARS_PER_HOUR * speed;
    }
    if (speed >= 5 && speed <= 8) {
      return (CARS_PER_HOUR * speed) * 0.9;
    }
    if (speed == 9) {
      return (CARS_PER_HOUR * speed) * 0.8;
    }
    if (speed == 10) {
      return (CARS_PER_HOUR * speed) * 0.77;
    }
    throw new IllegalArgumentException("Speed must be between 0 and 10");
  }

  public int workingItemsPerMinute(int speed) {
    return (int) (productionRatePerHour(speed) / 60);
  }
}
